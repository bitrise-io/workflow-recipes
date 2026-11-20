# (iOS) Run Tests on a Physical Device

## Description

Run unit or UI tests on a physical device. Our device testing solution is based on Firebase Test Lab. You can find the test results, logs, videos and screenshots on the **Tests** tab of the build page.

## Prerequisites

1. The source code is cloned and the dependencies (for example, Cocoapods, Carthage) are installed.
2. You have code signing set up. See [iOS Code Signing](https://docs.bitrise.io/en/bitrise-ci/code-signing/ios-code-signing/ios-code-signing) for more details.

## Instructions

1. Add an [Xcode Build for testing for iOS](https://bitrise.io/integrations/steps/xcode-build-for-test) Step.
    - Set up code signing for the Step.
2. Add an [iOS Device Testing](https://bitrise.io/integrations/steps/virtual-device-testing-for-ios) Step. Optional inputs:
    - **Number of times a test execution is reattempted**: reruns the whole test suite if a test fails. A test that passes on a rerun is marked as flaky.
    - **Download files**: set it to `true` to download the test outputs. Flaky tests are only listed in `$BITRISE_FLAKY_TEST_CASES` if this is enabled.
3. Add a [Deploy to Bitrise.io](https://bitrise.io/integrations/steps/deploy-to-bitrise-io) Step that makes the test results available on the **Tests** tab.

## bitrise.yml

```yaml
- xcode-build-for-test@3:
    inputs:
    - automatic_code_signing: api-key
- virtual-device-testing-for-ios@1:
    inputs:
    - num_flaky_test_attempts: 1
    - download_test_results: "true"
- deploy-to-bitrise-io@2: {}
```

## Relevant Links

* https://docs.bitrise.io/en/bitrise-ci/testing/device-testing-with-firebase/device-testing-for-ios
