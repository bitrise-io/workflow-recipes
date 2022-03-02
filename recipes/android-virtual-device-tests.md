# (Android) Run UI / instrumentation tests on virtual devices

## Description

Run UI / instrumentation (for example, Espresso) or robo/gameloop tests on virtual devices. [Our device testing solution](https://devcenter.bitrise.io/en/testing/device-testing-for-android.html).
) is based on Firebase Test Lab. You can find the resulting logs, videos and screenshots on Bitrise.

## Instructions

1. Add an [Android Build for UI Testing](https://github.com/bitrise-steplib/bitrise-step-android-build-for-ui-testing) Step. Set the input variables:
    - **Project Location**: Use the default `$BITRISE_SOURCE_DIR` or `$PROJECT_LOCATION`. You can set a specific path but the automatically exposed Environment Variables are usually the best option.
    - **Variant**: Use the `$VARIANT` Enviromment Variable, or specify a variant manually.
    - **Module**: Specify one or leave it blank to run tests in all of the modules.
2. Add a [[BETA] Virtual Device Testing for Android](https://www.bitrise.io/integrations/steps/virtual-device-testing-for-android) Step. Set the input variables:
    - **Test type**: `instrumentation` (or `robo` or `gameloop`)
    - (Optional) **Test devices** (default: `NexusLowRes,24,en,portrait`)
3. Add a [Deploy to Bitrise.io - Apps, Logs, Artifacts](https://www.bitrise.io/integrations/steps/deploy-to-bitrise-io) Step that makes the test results available in the [Test Reports add-on](https://devcenter.bitrise.io/en/testing/test-reports.html).

## bitrise.yml

```yaml
- android-build-for-ui-testing@0:
    inputs:
    - variant: $VARIANT
    - module: $MODULE
- virtual-device-testing-for-android@1:
    inputs:
    - test_type: instrumentation
- deploy-to-bitrise-io@2: {}
```

## Links

* https://devcenter.bitrise.io/en/testing/device-testing-for-android.html
