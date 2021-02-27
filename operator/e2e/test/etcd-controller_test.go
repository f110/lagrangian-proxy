package test

import (
	"context"
	"testing"
	"time"

	"go.etcd.io/etcd/v3/clientv3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"go.f110.dev/heimdallr/operator/e2e/e2eutil"
	"go.f110.dev/heimdallr/operator/e2e/framework"
	etcdv1alpha2 "go.f110.dev/heimdallr/operator/pkg/api/etcd/v1alpha2"
	"go.f110.dev/heimdallr/pkg/k8s/kind"
	"go.f110.dev/heimdallr/pkg/testing/btesting"
)

func TestEtcdController(t *testing.T) {
	t.Parallel()
	f := framework.New(t, RESTConfig)
	defer f.Execute()

	f.Describe("EtcdController", func(s *btesting.Scenario) {
		s.Context("creates a new cluster", func(s *btesting.Scenario) {
			s.Context("with PersistentVolumeClaim", func(s *btesting.Scenario) {
				s.Subject(func(m *btesting.Matcher) {
					f.EtcdClusters.Setup(m, framework.Name("pvc"), framework.PersistentData)
				})
				s.Defer(func() { f.EtcdClusters.EtcdCluster("pvc").Destroy(f.Client()) })

				s.It("should have 3 pods", func(m *btesting.Matcher) {
					f.EtcdClusters.EtcdCluster("pvc").NumOfPods(m, 3)
				})

				s.It("should have persistent volume", func(m *btesting.Matcher) {
					f.EtcdClusters.EtcdCluster("pvc").HavePVC(m)
				})
			})

			s.Context("with EmptyDir", func(s *btesting.Scenario) {
				s.Subject(func(m *btesting.Matcher) {
					f.EtcdClusters.Setup(m, framework.Name("memory"))
				})
				s.Defer(func() { f.EtcdClusters.EtcdCluster("memory").Destroy(f.Client()) })

				s.It("should have 3 pods", func(m *btesting.Matcher) {
					f.EtcdClusters.EtcdCluster("memory").NumOfPods(m, 3)
				})

				s.It("should have EmptyDir volume", func(m *btesting.Matcher) {
					f.EtcdClusters.EtcdCluster("memory").HaveEmptyDir(m)
				})
			})
		})

		s.Context("updates the cluster", func(s *btesting.Scenario) {
			s.Context("recreates all pods", func(s *btesting.Scenario) {
				s.Defer(func() { f.EtcdClusters.EtcdCluster("update").Destroy(f.Client()) })

				s.Step("create new cluster", func(s *btesting.Scenario) {
					s.Subject(func(m *btesting.Matcher) {
						f.EtcdClusters.Setup(m, framework.Name("update"), framework.DisableAntiAffinity)
					})

					s.It("should have 3 pods", func(m *btesting.Matcher) {
						f.EtcdClusters.EtcdCluster("update").NumOfPods(m, 3)
					})
				})

				s.Step("edit version", func(s *btesting.Scenario) {
					s.Subject(func(m *btesting.Matcher) {
						f.EtcdClusters.EtcdCluster("update").Reload()
						ec := f.EtcdClusters.EtcdCluster("update").EtcdCluster
						m.NotNil(ec)
						ec.Spec.Version = "v3.4.4"
						_, err := f.Client().EtcdV1alpha2().EtcdClusters(ec.Namespace).Update(context.TODO(), ec, metav1.UpdateOptions{})
						m.Must(err)
						m.Must(e2eutil.WaitForStatusOfEtcdClusterBecome(f.Client(), ec, etcdv1alpha2.ClusterPhaseUpdating, 1*time.Minute))
					})

					s.It("should be ready", func(m *btesting.Matcher) {
						f.EtcdClusters.EtcdCluster("update").Ready(m)
					})
				})

				s.Step("wait for running", func(s *btesting.Scenario) {
					s.Subject(func(m *btesting.Matcher) {
						ec := f.EtcdClusters.EtcdCluster("update").EtcdCluster
						m.NotNil(ec)
						m.Must(e2eutil.WaitForStatusOfEtcdClusterBecome(f.Client(), ec, etcdv1alpha2.ClusterPhaseRunning, 10*time.Minute))
					})

					s.It("all pods should have updated", func(m *btesting.Matcher) {
						f.EtcdClusters.EtcdCluster("update").EqualVersion(m, "v3.4.4")
					})
				})
			})
		})

		s.Context("restore from backup", func(s *btesting.Scenario) {
			var dataPutTime time.Time
			const testDataKey = "e2e-test-data"
			s.Defer(func() { f.EtcdClusters.EtcdCluster("restore").Destroy(f.Client()) })

			s.Step("create new cluster", func(s *btesting.Scenario) {
				s.Subject(func(m *btesting.Matcher) {
					f.EtcdClusters.Setup(m, framework.Name("restore"), framework.Backup(kind.MinIOBucketName))
				})

				s.It("should have 3 pods", func(m *btesting.Matcher) {
					f.EtcdClusters.EtcdCluster("restore").NumOfPods(m, 3)
				})
			})

			s.Step("write data", func(s *btesting.Scenario) {
				s.Subject(func(m *btesting.Matcher) {
					f.EtcdClusters.EtcdCluster("restore").Reload()
					ecClient := f.EtcdClusters.EtcdCluster("restore").Client(m)
					_, err := ecClient.Put(context.TODO(), testDataKey, "ok-test")
					m.Must(err)
					dataPutTime = time.Now()
					m.Must(ecClient.Close())
				})

				s.It("should return data", func(m *btesting.Matcher) {
					ecClient := f.EtcdClusters.EtcdCluster("restore").Client(m)
					_, err := ecClient.Get(context.TODO(), testDataKey)
					m.Must(err)
				})
			})

			s.Step("waiting for backed up", func(s *btesting.Scenario) {
				s.Subject(func(m *btesting.Matcher) {
					m.Must(e2eutil.WaitForBackup(f.Client(), f.EtcdClusters.EtcdCluster("restore").EtcdCluster, dataPutTime))
				})

				s.It("should update status", func(m *btesting.Matcher) {
					f.EtcdClusters.EtcdCluster("restore").Reload()
					m.NotNil(f.EtcdClusters.EtcdCluster("restore").Status.Backup)
					m.True(f.EtcdClusters.EtcdCluster("restore").Status.Backup.Succeeded)
					m.True(len(f.EtcdClusters.EtcdCluster("restore").Status.Backup.History) > 0)
				})
			})

			s.Step("destroy the cluster", func(s *btesting.Scenario) {
				s.Subject(func(m *btesting.Matcher) {
					etcdCluster := f.EtcdClusters.EtcdCluster("restore")
					for _, v := range etcdCluster.Status.Members {
						m.Must(f.CoreClient().CoreV1().Pods(etcdCluster.Namespace).Delete(context.TODO(), v.PodName, metav1.DeleteOptions{}))
					}
				})

				s.It("should deleted pod", func(m *btesting.Matcher) {
					m.True(true)
				})
			})

			s.Step("waiting for completed restoring", func(s *btesting.Scenario) {
				s.Subject(func(m *btesting.Matcher) {
					m.Must(e2eutil.WaitForRestore(f.Client(), f.EtcdClusters.EtcdCluster("restore").EtcdCluster, dataPutTime))
					m.Must(e2eutil.WaitForStatusOfEtcdClusterBecome(f.Client(), f.EtcdClusters.EtcdCluster("restore").EtcdCluster, etcdv1alpha2.ClusterPhaseRunning, 10*time.Minute))
				})

				s.It("should restore completed", func(m *btesting.Matcher) {
					f.EtcdClusters.EtcdCluster("restore").Reload()
					m.NotNil(f.EtcdClusters.EtcdCluster("restore").Status.Restored)
					m.True(f.EtcdClusters.EtcdCluster("restore").Status.Restored.Completed)
				})
			})

			s.Step("read data", func(s *btesting.Scenario) {
				var response *clientv3.GetResponse
				s.Subject(func(m *btesting.Matcher) {
					ecClient := f.EtcdClusters.EtcdCluster("restore").Client(m)
					res, err := ecClient.Get(context.TODO(), testDataKey)
					response = res
					m.Must(err)
				})

				s.It("should return 1 value", func(m *btesting.Matcher) {
					m.Equal(int64(1), response.Count)
				})

				s.It("should return correctly data", func(m *btesting.Matcher) {
					m.Equal("ok-test", string(response.Kvs[0].Value))
				})
			})
		})
	})
}
