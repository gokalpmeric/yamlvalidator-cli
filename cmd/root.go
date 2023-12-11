package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var filename string

var rootCmd = &cobra.Command{
	Use:   "yamlvalidator",
	Short: "YAML Validator is a tool for validating YAML files",
	Long:  `YAML Validator is a CLI tool to check the validity of YAML files.`,
	Run: func(cmd *cobra.Command, args []string) {
		if filename != "" {
			validateYAML(filename) // Call a separate function to validate the YAML
		} else {
			cmd.Help()
		}
	},
}

func validateYAML(filename string) {
	// Read the YAML file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", filename, err)
		os.Exit(1)
	}

	// Parse and validate the YAML file
	var content interface{}
	err = yaml.Unmarshal(data, &content)
	if err != nil {
		fmt.Printf("Error in %s: %v\n", filename, err)
		os.Exit(1)
	}

	// If no error, the YAML is valid
	fmt.Printf("%s is a valid YAML file.\n", filename)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&filename, "validate", "v", "", "YAML file to validate")
}
