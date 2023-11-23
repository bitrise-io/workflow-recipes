# (iOS) Deploy to bitrise.io

## Description

Build and distribute your app to testers via [Bitrise.io Ship](https://devcenter.bitrise.io/en/deploying/deploying-with-ship.html).

## Prerequisites

1. You have code signing set up. See [iOS Code Signing](https://devcenter.bitrise.io/en/code-signing/ios-code-signing.html) for more details.

## Instructions

1. Add the [Xcode Archive & Export for iOS](https://bitrise.io/integrations/steps/xcode-archive) step. Set the input variables:
    - **Project path**: by default, `$BITRISE_PROJECT_PATH`.
    - **Scheme**: by default, `$BITRISE_SCHEME`.
    - **Distribution method**: `development`, `ad-hoc` or `enterprise`.
2. Add the [Deploy to Bitrise.io - Apps, Logs, Artifacts](https://www.bitrise.io/integrations/steps/deploy-to-bitrise-io) Step.

## bitrise.yml

```yaml
- xcode-archive@5:
    inputs:
    - project_path: "$BITRISE_PROJECT_PATH"
    - scheme: "$BITRISE_SCHEME"
    - automatic_code_signing: apple-id
    - distribution_method: development
- deploy-to-bitrise-io@2: {}
```
