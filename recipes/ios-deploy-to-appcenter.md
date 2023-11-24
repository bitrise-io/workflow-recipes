# (iOS) Deploy to Visual Studio App Center

## Description

Build and distribute your app to testers via AppCenter.

## Prerequisites

1. An existing [Visual Studio App Center](https://docs.microsoft.com/en-us/appcenter/dashboard/) project where your app is registered.
2. Adding the API token as a [Secret](https://devcenter.bitrise.io/en/builds/secrets.html) your Bitrise project with the name `APPCENTER_API_TOKEN`.
3. You have code signing set up. See [iOS Code Signing](https://devcenter.bitrise.io/en/code-signing/ios-code-signing.html) for more details.

## Instructions

1. Add the [Xcode Archive & Export for iOS](https://bitrise.io/integrations/steps/xcode-archive) step. Set the input variables:
    - **Project path**: by default `$BITRISE_PROJECT_PATH`.
    - **Scheme**: by default `$BITRISE_SCHEME`.
    - **Distribution method**: `development`, `ad-hoc` or `enterprise`.
3. Add the [AppCenter iOS Deploy](https://www.bitrise.io/integrations/steps/appcenter-deploy-ios) step and set the following inputs:
    - **API Token**: `$APPCENTER_API_TOKEN`.
    - **Owner name**: for example, `my-company`.
    - **App name**: for example, `my-app` Use the [App Center CLI](https://github.com/Microsoft/appcenter-cli) to get the app name since it might not be the same as the one you can see on the Visual Studio App Center website.
    - Check out other options in the Step documentation or in the Workflow Editor.

## bitrise.yml

```yaml
- xcode-archive@5:
    inputs:
    - project_path: $BITRISE_PROJECT_PATH
    - scheme: $BITRISE_SCHEME
    - automatic_code_signing: apple-id
    - distribution_method: development
- appcenter-deploy-ios@2:
    inputs:
    - owner_name: my-company
    - app_name: my-app
    - api_token: $APPCENTER_API_TOKEN
```
