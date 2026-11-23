# (iOS) Run tests on a simulator

## Description

Run unit or UI tests of an iOS app on a simulator and view the results on the **Tests** tab of the build page.

## Instructions

1. Add an [Xcode Test for iOS](https://bitrise.io/integrations/steps/xcode-test) Step. Override any of the following inputs if needed:
    - **Project path**: The default value is `$BITRISE_PROJECT_PATH` and in most cases you don't have to change it.
    - **Scheme**: The default value is `$BITRISE_SCHEME`, this variable stores the scheme that you set when adding the app on Bitrise. You can specify a different scheme if you want but it must be a shared scheme.
    - **Device destination specifier** (default: `platform=iOS Simulator,name=Bitrise iOS default,OS=latest`). `Bitrise iOS default` is a pre-warmed simulator that is already created on the build machine.
    - **Test Repetition Mode** (default: `retry_on_failure`) and **Maximum Test Repetitions** (default: `3`): failed tests are rerun, and a test that passes on a rerun is marked as flaky instead of failing the build. See [Detecting and quarantining flaky tests](/recipes/flaky-tests.md).
2. Add a [Deploy to Bitrise.io](https://bitrise.io/integrations/steps/deploy-to-bitrise-io) Step. It makes the test results available on the **Tests** tab, and it runs even if the tests failed.

## bitrise.yml

```yaml
- xcode-test@6: {}
- deploy-to-bitrise-io@2: {}
```

## Relevant Links

* https://docs.bitrise.io/en/bitrise-ci/testing/testing-ios-apps/running-unit-and-ui-tests-for-ios-apps
* https://docs.bitrise.io/en/bitrise-ci/testing/test-reports/viewing-test-results
