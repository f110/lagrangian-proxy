package release

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
	"gopkg.in/yaml.v2"
)

type basic struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
}

func finalizer(in io.Reader, out io.Writer, version string) error {
	d := yaml.NewDecoder(in)
	e := yaml.NewEncoder(out)
	for {
		v := make(map[interface{}]interface{})
		err := d.Decode(v)
		if err == io.EOF {
			break
		}
		if err != nil {
			return xerrors.Errorf(": %w", err)
		}

		kind, err := detectKind(v)
		if err != nil {
			return xerrors.Errorf(": %w", err)
		}
		switch kind {
		case "CustomResourceDefinition":
			editCustomResourceDefinition(v)
		case "Deployment":
			editDeployment(v, version)
		}

		if err := e.Encode(v); err != nil {
			return xerrors.Errorf(": %w", err)
		}
	}
	if err := e.Close(); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func detectKind(v interface{}) (string, error) {
	b, err := yaml.Marshal(v)
	if err != nil {
		return "", xerrors.Errorf(": %w", err)
	}
	bc := &basic{}
	if err := yaml.Unmarshal(b, bc); err != nil {
		return "", xerrors.Errorf(": %w", err)
	}

	if bc.Kind != "" {
		return bc.Kind, nil
	}

	return "", errors.New("failed parse document")
}

func editCustomResourceDefinition(v map[interface{}]interface{}) {
	delete(v, "status")

	m := v["metadata"].(map[interface{}]interface{})
	delete(m, "creationTimestamp")
	delete(m, "annotations")
}

func editDeployment(v map[interface{}]interface{}, version string) {
	containers := v["spec"].(map[interface{}]interface{})["template"].(map[interface{}]interface{})["spec"].(map[interface{}]interface{})["containers"].([]interface{})
	for _, c := range containers {
		v := c.(map[interface{}]interface{})
		if i, ok := v["image"]; ok {
			image := i.(string)
			if strings.Contains(image, "heimdallr-operator") {
				s := strings.Split(image, ":")
				v["image"] = s[0] + ":" + version
			}
		}
	}
}

func manifestCleaner(input, output, version string) error {
	reader, err := os.Open(input)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	writer, err := os.Create(output)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return finalizer(reader, writer, version)
}

func ManifestCleaner(rootCmd *cobra.Command) {
	input := ""
	output := ""
	version := ""
	cleaner := &cobra.Command{
		Use: "manifest-cleaner",
		RunE: func(_ *cobra.Command, _ []string) error {
			return manifestCleaner(input, output, version)
		},
	}
	cleaner.Flags().StringVar(&input, "input", "", "Input file")
	cleaner.Flags().StringVar(&output, "output", "", "Output path")
	cleaner.Flags().StringVar(&version, "version", "", "Version string")

	rootCmd.AddCommand(cleaner)
}