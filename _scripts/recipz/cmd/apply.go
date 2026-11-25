/*
Copyright © 2023 Bitrise
*/
package cmd

import (
	"fmt"

	"github.com/bitrise-io/go-utils/fileutil"
	"github.com/bitrise-io/workflow-recipes/_scripts/recipz/lib/recipemdtojson"
	"github.com/spf13/cobra"
)

var applyConfigs = struct {
	RecipePath        string
	ExampleBitriseYML string
	OpenAIAPIKey      string
}{}

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("apply called", applyConfigs)

		recipeMDContent, err := fileutil.ReadStringFromFile(applyConfigs.RecipePath)
		if err != nil {
			return fmt.Errorf("failed to read recipe file at path (%s), err: %w", applyConfigs.RecipePath, err)
		}
		bitriseYMLContent, err := fileutil.ReadStringFromFile(applyConfigs.ExampleBitriseYML)
		if err != nil {
			return fmt.Errorf("failed to read bitrise.yml file at path (%s), err: %w", applyConfigs.ExampleBitriseYML, err)
		}

		recipeJSON, err := recipemdtojson.RecipeMDToJSON(recipeMDContent)
		if err != nil {
			return fmt.Errorf("failed to parse recipe (path: %s), err: %w", applyConfigs.RecipePath, err)
		}

		const prompt = "I have the following bitrise.yml:\n\n```\n%s\n```\n\n" +
			"I'd like to modify this `bitrise.yml` to achieve the following:\n\n```\n%s\n```\n\n" +
			"This is an example YML snippet:\n\n%s\n\n" +
			"Please always respond with the modified `bitrise.yml`."
		fullPrompt := fmt.Sprintf(prompt, bitriseYMLContent, recipeJSON.Description, recipeJSON.BitriseYML)
		fmt.Println("fullPrompt:", fullPrompt)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	applyCmd.Flags().StringVar(&applyConfigs.RecipePath, "recipe", "", "Recipe path")
	applyCmd.MarkFlagRequired("recipe")

	applyCmd.Flags().StringVar(&applyConfigs.ExampleBitriseYML, "bitrise-yml", "", "Bitrise YML path")
	applyCmd.MarkFlagRequired("bitrise-yml")

	applyCmd.Flags().StringVar(&applyConfigs.OpenAIAPIKey, "openai-api-key", "", "OpenAI API Key")
	applyCmd.MarkFlagRequired("openai-api-key")
}
