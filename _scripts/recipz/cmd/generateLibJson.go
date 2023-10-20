/*
Copyright © 2023 Bitrise
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bitrise-io/go-utils/fileutil"
	"github.com/bitrise-io/workflow-recipes/_scripts/recipz/lib/recipemdtojson"
	"github.com/spf13/cobra"
)

type RecipesLib struct {
	Recipes []recipemdtojson.RecipeJSON `json:"recipes"`
}

var generateLibJsonCmdConfigs = struct {
	RecipesDirPath string
}{}

// generateLibJsonCmd represents the generateLibJson command
var generateLibJsonCmd = &cobra.Command{
	Use:   "generateLibJson",
	Short: "Generate a Bitrise Recipes JSON",
	Long: `Generate a Bitrise Recipes JSON
which includes all the recipes in from the specified directory,
compiled into a processed JSON representation.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// IMPLEMENT HERE
		// iterate through the speficied directory
		files, err := os.ReadDir(generateLibJsonCmdConfigs.RecipesDirPath)
		if err != nil {
			log.Fatal(err)
		}

		recipesLib := RecipesLib{
			Recipes: []recipemdtojson.RecipeJSON{},
		}
		// from that directory read all `.md` recipe files
		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".md") {
				recipeFilePath := filepath.Join(generateLibJsonCmdConfigs.RecipesDirPath, file.Name())
				log.Println("recipeFilePath:", recipeFilePath)
				recipeMDContent, err := fileutil.ReadStringFromFile(recipeFilePath)
				if err != nil {
					return fmt.Errorf("failed to read recipe file at path (%s), err: %w", recipeFilePath, err)
				}

				recipeJSON, err := recipemdtojson.RecipeMDToJSON(recipeMDContent)
				if err != nil {
					return fmt.Errorf("failed to parse recipe (path: %s), err: %w", recipeFilePath, err)
				}
				recipesLib.Recipes = append(recipesLib.Recipes, recipeJSON)
			}
		}

		recipesLibJSON, err := json.Marshal(recipesLib.Recipes)
		if err != nil {
			return fmt.Errorf("failed to generate Recipes Lib JSON, err: %w", err)
		}

		fmt.Printf("%s\n", recipesLibJSON)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateLibJsonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateLibJsonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateLibJsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	generateLibJsonCmd.Flags().StringVar(&generateLibJsonCmdConfigs.RecipesDirPath, "recipes-dir", "", "Recipes dir path - the path of the directory which includes the recipe md files")
	generateLibJsonCmd.MarkFlagRequired("recipes-dir")
}
