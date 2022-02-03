# (Android) Run unit tests

## Description

Run unit tests (for example, `testDebugUnitTest`).

## Instructions

1. Add an [Android Unit Test](https://www.bitrise.io/integrations/steps/android-unit-test) Step. Input variables you might set:
    - **Project Location**: Use the default `$BITRISE_SOURCE_DIR` or `$PROJECT_LOCATION`. You can set a specific path but the automatically exposed Environment Variables are usually the best option.
    - **Variant**: Use the `$VARIANT` Enviromment Variable, or specify a variant manually.
    - **Module**: Specify one or leave it blank to run tests in all of the modules.
2. Add a [Deploy to Bitrise.io - Apps, Logs, Artifacts](https://www.bitrise.io/integrations/steps/deploy-to-bitrise-io) Step that makes the test results available in the [Test Reports add-on](https://devcenter.bitrise.io/en/testing/test-reports.html).

## bitrise.yml

```yaml
- android-unit-test@1:
    inputs:
    - project_location: $PROJECT_LOCATION
    - variant: $VARIANT
- deploy-to-bitrise-io@2: {}
```

## Related Links

* https://devcenter.bitrise.io/en/testing/android-unit-tests.html
