# (iOS) Run tests on a simulator

## Description

Run unit or UI tests of an iOS app on a simulator.

## Instructions

1. Add an [Xcode Test for iOS](https://github.com/bitrise-steplib/steps-xcode-test) Step. Override any of the following inputs if needed:
    - **Project path**: The default value is `$BITRISE_PROJECT_PATH` and in most cases you don't have to change it.
    - **Scheme**: The default value is `$BITRISE_SCHEME`, this variable stores the scheme that you set when adding the app on Bitrise. You can specify a different scheme if you want but it must be a shared scheme.
    - **Device destination specifier** (default: `platform=iOS Simulator,name=iPhone 8 Plus,OS=latest`).
2. Add a [Deploy to Bitrise.io - Apps, Logs, Artifacts](https://www.bitrise.io/integrations/steps/deploy-to-bitrise-io) Step that makes the test results available in the [Test Reports add-on](https://devcenter.bitrise.io/en/testing/test-reports.html).

## bitrise.yml

```yaml
- xcode-test@4: {}
- deploy-to-bitrise-io@2: {}
```
