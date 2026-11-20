# (Android) Run unit tests

## Description

Run unit tests (for example, `testDebugUnitTest`) and view the results on the **Tests** tab of the build page.

## Instructions

1. Add an [Android Unit Test](https://bitrise.io/integrations/steps/android-unit-test) Step. Input variables you might set:
    - **Project Location**: Use the default `$BITRISE_SOURCE_DIR` or `$PROJECT_LOCATION`. You can set a specific path but the automatically exposed Environment Variables are usually the best option.
    - **Variant**: Use the `$VARIANT` Environment Variable, or specify a variant manually.
    - **Module**: Specify one or leave it blank to run tests in all of the modules.
2. Add a [Deploy to Bitrise.io](https://bitrise.io/integrations/steps/deploy-to-bitrise-io) Step. It makes the test results available on the **Tests** tab, and it runs even if the tests failed.

To rerun failed tests and mark the ones that pass on a rerun as flaky, apply the [Gradle Test Retry plugin](https://github.com/gradle/test-retry-gradle-plugin) in your project. See [Detecting and quarantining flaky tests](/recipes/flaky-tests.md).

## bitrise.yml

```yaml
- android-unit-test@1:
    inputs:
    - project_location: $PROJECT_LOCATION
    - variant: $VARIANT
- deploy-to-bitrise-io@2: {}
```

## Related Links

* https://docs.bitrise.io/en/bitrise-ci/testing/testing-android-apps/android-unit-tests
