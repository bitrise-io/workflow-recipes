package recipemdtojson

import (
	"errors"
	"strings"
)

type RecipeJSON struct {
	Markdown      string   `json:"markdown"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Prerequisites string   `json:"prerequisites"`
	Instructions  string   `json:"instructions"`
	BitriseYML    string   `json:"bitrise_yml"`
	Tags          []string `json:"tags"`
}

func RecipeMDToJSON(recipeMD string) (RecipeJSON, error) {
	sections := strings.Split(recipeMD, "\n\n## ")
	if len(sections) < 2 {
		return RecipeJSON{}, errors.New("invalid recipe markdown")
	}

	recipe := RecipeJSON{
		Markdown: recipeMD,
		Tags:     []string{},
	}
	for _, section := range sections {
		lines := strings.Split(section, "\n")
		if strings.HasPrefix(lines[0], "# ") {
			recipe.Title = strings.TrimPrefix(lines[0], "# ")
		} else {
			switch lines[0] {
			case "Description":
				recipe.Description = strings.TrimSpace(strings.Join(lines[1:], "\n"))
			case "Prerequisites":
				recipe.Prerequisites = strings.TrimSpace(strings.Join(lines[1:], "\n"))
			case "Instructions":
				recipe.Instructions = strings.TrimSpace(strings.Join(lines[1:], "\n"))
			case "bitrise.yml":
				recipe.BitriseYML = strings.TrimSpace(strings.Join(lines[1:], "\n"))
			case "Recipe Tags":
				recipe.Tags = strings.Split(strings.ReplaceAll(strings.TrimSpace(strings.Join(lines[1:], "\n")), " ", ""), ",")
			}
		}
	}

	if len(recipe.Title) < 1 {
		return RecipeJSON{}, errors.New("recipe title can not be empty")
	}
	if len(recipe.Description) < 1 {
		return RecipeJSON{}, errors.New("recipe description section can not be empty")
	}
	if len(recipe.BitriseYML) < 1 {
		return RecipeJSON{}, errors.New("recipe bitrise.yml section can not be empty")
	}

	return recipe, nil
}
