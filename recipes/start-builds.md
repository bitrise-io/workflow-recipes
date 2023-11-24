# Start (parallel) builds from the Workflow

## Description

Start one or more builds with specified Workflows from the parent Workflow and optionally wait for their completion.

## Prerequisites

1. Make sure you have a valid Bitrise API key in your Secrets (`$BITRISE_API_KEY`). See [Personal access](https://devcenter.bitrise.io/en/accounts/personal-access-tokens.html##) tokens for more details.
2. Have Workflow(s) you would like to run in parallel (`workflow-1` and `workflow-2` in the example).

## Instructions

1. Add a [Bitrise Start Build](https://www.bitrise.io/integrations/steps/build-router-start) Step. Set the input variables:
    - **workflows**: The Workflow(s) to start. One Workflow per line.
    - **Bitrise Access Token**: `$BITRISE_API_KEY`.
2. (Optional) Add any Step you would like to run in parallel in the parent Workflow while the triggered Workflows are running.
3. (Optional) Add a [Bitrise Wait for Build](https://www.bitrise.io/integrations/steps/build-router-wait) Step. Set the input variables:
    - **Bitrise Access Token**: `$BITRISE_API_KEY`

With the Steps above you can only start a build for the same app. If you would like to start a build for an other app, you can use the [Trigger Bitrise workflow](https://www.bitrise.io/integrations/steps/trigger-bitrise-workflow) Step.

## bitrise.yml

```yaml
parent-workflow:
  steps:
  - build-router-start@0:
      inputs:
      - workflows: |-
          workflow-1
          workflow-2
      - access_token: $BITRISE_API_KEY
  - script@1:
      inputs:
      - content: echo "Doing something else..."
  - build-router-wait@0:
      inputs:
      - access_token: $BITRISE_API_KEY
```
