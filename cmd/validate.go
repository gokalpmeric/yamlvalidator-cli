package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate [file]",
	Short: "Validate a YAML file",
	Long:  `Validate a YAML file for proper format and syntax.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]

		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", filename, err)
			os.Exit(1)
		}

		var content interface{}
		err = yaml.Unmarshal(data, &content)
		if err != nil {
			fmt.Printf("Error in %s: %v\n", filename, err)
			os.Exit(1)
		}

		fmt.Printf("%s is a valid YAML file.\n", filename)
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
