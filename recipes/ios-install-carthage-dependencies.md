# (iOS) Install Carthage Dependencies

## Description

Installing Carthage dependecies.

## Instructions

1. Add the [Carthage](https://bitrise.io/integrations/steps/carthage) Step. Set the input variables:
    - **Github Personal Access Token**: We recommend adding a GitHub access token to your Secrets (`$GITHUB_ACCESS_TOKEN`). We need this token to avoid GitHub rate limit issue. See the GitHub guide: [Creating an access token for command-line use](https://help.github.com/articles/creating-an-access-token-for-command-line-use/) on how to create Personal Access Token. Uncheck every scope box when creating this token. There is no reason this token needs access to private information.
    - (Optional) **Additional options for carthage command**: See the [Carthage docs](https://github.com/Carthage/Carthage) for the available options, for example, `--use-xcframeworks --platform iOS`.

## bitrise.yml

```
- carthage@3:
    inputs:
    - carthage_options: "--use-xcframeworks --platform iOS"
```
