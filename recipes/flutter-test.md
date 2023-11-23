# (Flutter) Run tests

## Description
Performs any test in a Flutter project.

## Instructions

1. Add the [Flutter Test](https://bitrise.io/integrations/steps/flutter-test) Step to your Workflow. Set the input variables:
    - **Project Location**: For example, `$BITRISE_FLUTTER_PROJECT_LOCATION`.
    - Check out optional inputs in the Workflow Editor or in the Step description.
2. Add a [Deploy to Bitrise.io - Apps, Logs, Artifacts](https://www.bitrise.io/integrations/steps/deploy-to-bitrise-io) step that makes the test results available in the [Test Reports add-on](https://devcenter.bitrise.io/en/testing/test-reports.html). The failed tests will be also available under the `Test Results` tab on the build details page.

## bitrise.yml

```yaml
- flutter-test@1:
    inputs:
    - project_location: "$BITRISE_FLUTTER_PROJECT_LOCATION"
- deploy-to-bitrise-io@2: {}
```
