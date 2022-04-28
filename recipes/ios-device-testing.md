# (iOS) Run Tests on a Physical Device

## Description

Run unit or UI tests on a physical device. Our device testing solution is based on Firebase Test Lab. You can find the resulting logs, videos and screenshots on Bitrise.

## Prerequisites

1. The source code is cloned and the dependencies (for example, Cocoapods, Carthage) are installed.
2. You have code signing set up. See [iOS Code Signing](https://devcenter.bitrise.io/en/code-signing/ios-code-signing.html) for more details.

## Instructions

1. Add an [Xcode Build for testing for iOS](https://github.com/bitrise-steplib/steps-xcode-build-for-test) Step.
2. Add a [[BETA] iOS Device Testing](https://www.bitrise.io/integrations/steps/virtual-device-testing-for-ios) Step.
    - Setup code signing for the Step.
3. Add a [Deploy to Bitrise.io - Apps, Logs, Artifacts](https://www.bitrise.io/integrations/steps/deploy-to-bitrise-io) Step that makes the test results available in the [Test Reports add-on](https://devcenter.bitrise.io/en/testing/test-reports.html).

## bitrise.yml

```yaml
- xcode-build-for-test@1:
    inputs:
    - automatic_code_signing: api_key
- virtual-device-testing-for-ios@1: {}
- deploy-to-bitrise-io@2: {}
```

## Relevant Links

* https://devcenter.bitrise.io/en/testing/device-testing-for-ios.html
