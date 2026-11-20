# (Android) Run UI / instrumentation tests on virtual devices

## Description

Run UI / instrumentation (for example, Espresso) or robo/gameloop tests on virtual devices. [Our device testing solution](https://docs.bitrise.io/en/bitrise-ci/testing/device-testing-with-firebase/device-testing-for-android) is based on Firebase Test Lab. You can find the test results, logs, videos and screenshots on the **Tests** tab of the build page.

## Instructions

1. Add an [Android Build for UI Testing](https://bitrise.io/integrations/steps/android-build-for-ui-testing) Step. Set the input variables:
    - **Project Location**: Use the default `$BITRISE_SOURCE_DIR` or `$PROJECT_LOCATION`. You can set a specific path but the automatically exposed Environment Variables are usually the best option.
    - **Variant**: Use the `$VARIANT` Environment Variable, or specify a variant manually.
    - **Module**: Use the `$MODULE` Environment Variable, or specify a module manually.
2. Add a [Virtual Device Testing for Android](https://bitrise.io/integrations/steps/virtual-device-testing-for-android) Step. Set the input variables:
    - **Test type**: `instrumentation` (or `robo` or `gameloop`)
    - (Optional) **Test devices** (default: `MediumPhone.arm,33,en,portrait`)
    - (Optional) **Number of times a test execution is reattempted**: reruns the whole test suite if a test fails. A test that passes on a rerun is marked as flaky: see [Detect and quarantine flaky tests](./flaky-tests.md).
    - (Optional) **Download files**: set it to `true` to download the test outputs. Flaky tests are only listed in `$BITRISE_FLAKY_TEST_CASES` if this is enabled.
3. Add a [Deploy to Bitrise.io](https://bitrise.io/integrations/steps/deploy-to-bitrise-io) Step that makes the test results available on the **Tests** tab.

## bitrise.yml

```yaml
- android-build-for-ui-testing@1:
    inputs:
    - variant: $VARIANT
    - module: $MODULE
- virtual-device-testing-for-android@1:
    inputs:
    - test_type: instrumentation
    - num_flaky_test_attempts: 1
    - download_test_results: "true"
- deploy-to-bitrise-io@2: {}
```

## Links

* https://docs.bitrise.io/en/bitrise-ci/testing/device-testing-with-firebase/device-testing-for-android
