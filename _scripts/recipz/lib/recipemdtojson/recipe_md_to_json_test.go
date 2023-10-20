package recipemdtojson

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecipeMDToJSON(t *testing.T) {
	t.Log("Template example")
	{
		recipeTemplate := `# [Name of the recipe]

## Description

[Description]

## Prerequisites

1. ...

## Instructions

1. ...
    - ...
    - ...
    - ...

## bitrise.yml

code

`

		recipeJSON, err := RecipeMDToJSON(recipeTemplate)
		require.NoError(t, err)
		require.Equal(t, RecipeJSON{
			Markdown:      recipeTemplate,
			Title:         "[Name of the recipe]",
			Description:   "[Description]",
			Prerequisites: "1. ...",
			Instructions: `1. ...
    - ...
    - ...
    - ...`,
			BitriseYML: "code",
			Tags:       []string{},
		}, recipeJSON)
	}

	t.Log("Missing Prerequisites")
	{
		recipeTemplate := `# [Name of the recipe]

## Description

This is a

Multi section description

## Instructions

1. ...
    - ...
    - ...
    - ...

## bitrise.yml

code

`

		recipeJSON, err := RecipeMDToJSON(recipeTemplate)
		require.NoError(t, err)
		require.Equal(t, RecipeJSON{
			Markdown:      recipeTemplate,
			Title:         "[Name of the recipe]",
			Description:   "This is a\n\nMulti section description",
			Prerequisites: "",
			Instructions: `1. ...
    - ...
    - ...
    - ...`,
			BitriseYML: "code",
			Tags:       []string{},
		}, recipeJSON)
	}

	t.Log("Extra Links section at the end")
	{
		recipeTemplate := `# [Name of the recipe]

## Description

[Description]

## Prerequisites

1. ...

## Instructions

1. ...
    - ...
    - ...
    - ...

## bitrise.yml

code

## Links

* https://devcenter.bitrise.io/en/testing/device-testing-for-android.html
`

		recipeJSON, err := RecipeMDToJSON(recipeTemplate)
		require.NoError(t, err)
		require.Equal(t, RecipeJSON{
			Markdown:      recipeTemplate,
			Title:         "[Name of the recipe]",
			Description:   "[Description]",
			Prerequisites: "1. ...",
			Instructions: `1. ...
    - ...
    - ...
    - ...`,
			BitriseYML: "code",
			Tags:       []string{},
		}, recipeJSON)
	}
}

func TestRecipeMDToJSON_missing_required_fields(t *testing.T) {
	t.Log("Missing title")
	{
		recipeTemplate := `
## Description

[Description]

## Prerequisites

1. ...

## Instructions

1. ...
    - ...
    - ...
    - ...

## bitrise.yml

code

`

		recipeJSON, err := RecipeMDToJSON(recipeTemplate)
		require.EqualError(t, err, "recipe title can not be empty")
		require.Equal(t, RecipeJSON{}, recipeJSON)
	}

	t.Log("Missing description")
	{
		recipeTemplate := `# [Name of the recipe]

## NoDescription

No description section provided

## Prerequisites

1. ...

## Instructions

1. ...
	- ...
	- ...
	- ...

## bitrise.yml

code

`

		recipeJSON, err := RecipeMDToJSON(recipeTemplate)
		require.EqualError(t, err, "recipe description section can not be empty")
		require.Equal(t, RecipeJSON{}, recipeJSON)
	}

	t.Log("Missing bitrise.yml section")
	{
		recipeTemplate := `# [Name of the recipe]

## Description

[Description]

## Prerequisites

1. ...

## Instructions

1. ...
	- ...
	- ...
	- ...

## Links

- Link1

`

		recipeJSON, err := RecipeMDToJSON(recipeTemplate)
		require.EqualError(t, err, "recipe bitrise.yml section can not be empty")
		require.Equal(t, RecipeJSON{}, recipeJSON)
	}
}

func TestRecipeMDToJSON_tags(t *testing.T) {
	t.Log("Recipe Tags listed")
	{
		recipeTemplate := `# [Name of the recipe]

## Description

[Description]

## Prerequisites

1. ...

## Instructions

1. ...
    - ...
    - ...
    - ...

## bitrise.yml

code

## Recipe Tags

no-auto-apply, no-single-bitrise-yml

`

		recipeJSON, err := RecipeMDToJSON(recipeTemplate)
		require.NoError(t, err)
		require.Equal(t, RecipeJSON{
			Markdown:      recipeTemplate,
			Title:         "[Name of the recipe]",
			Description:   "[Description]",
			Prerequisites: "1. ...",
			Instructions: `1. ...
    - ...
    - ...
    - ...`,
			BitriseYML: "code",
			Tags:       []string{"no-auto-apply", "no-single-bitrise-yml"},
		}, recipeJSON)
	}
}
