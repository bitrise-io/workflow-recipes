# (Flutter) Run tests

## Description
Performs any test in a Flutter project and makes the results available on the **Tests** tab of the build page.

## Instructions

1. Add the [Flutter Test](https://bitrise.io/integrations/steps/flutter-test) Step to your Workflow. Set the input variables:
    - **Project Location**: For example, `$BITRISE_FLUTTER_PROJECT_LOCATION`.
    - Check out optional inputs in the Workflow Editor or in the Step description.
2. Add a [Deploy to Bitrise.io](https://bitrise.io/integrations/steps/deploy-to-bitrise-io) step. It makes the test results available on the **Tests** tab, and it runs even if the tests failed.

## bitrise.yml

```yaml
- flutter-test@1:
    inputs:
    - project_location: $BITRISE_FLUTTER_PROJECT_LOCATION
- deploy-to-bitrise-io@2: {}
```

## Relevant Links

* https://docs.bitrise.io/en/bitrise-ci/testing/testing-flutter-apps/running-unit-and-ui-tests-on-flutter-apps
