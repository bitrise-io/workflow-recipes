# (React Native) Run tests

## Description

Run tests (for example, Jest) and view the results on the **Tests** tab of the build page.

## Prerequisites

1. Jest writes a JUnit XML report. For example, add [jest-junit](https://www.npmjs.com/package/jest-junit) to your dev dependencies and register it as a reporter in your Jest config:

    ```json
    "jest": {
      "preset": "react-native",
      "reporters": [
        "default",
        ["jest-junit", { "outputDirectory": "<rootDir>/reports" }]
      ]
    }
    ```

## Instructions

1. Add either the [Run yarn command](https://bitrise.io/integrations/steps/yarn) or the [Run npm command](https://bitrise.io/integrations/steps/npm) Step based on your project setup.
2. Set the **The yarn command to run** or **The npm command with arguments to run** input to `test`.
3. Add an [Export test results to Test Reports](https://bitrise.io/integrations/steps/custom-test-results-export) Step. The Step runs even if the tests failed.
    - The name of the test: for example, `Jest`.
    - Test result search pattern: `*/reports/*.xml`, matching the `outputDirectory` of `jest-junit`.
4. Add a [Deploy to Bitrise.io](https://bitrise.io/integrations/steps/deploy-to-bitrise-io) Step that makes the test results available on the **Tests** tab.

## bitrise.yml

Using `yarn`:

```yaml
- yarn@2:
    inputs:
    - command: test
- custom-test-results-export@2:
    inputs:
    - test_name: Jest
    - search_pattern: "*/reports/*.xml"
- deploy-to-bitrise-io@2: {}
```

Using `npm`:

```yaml
- npm@3:
    inputs:
    - command: test
- custom-test-results-export@2:
    inputs:
    - test_name: Jest
    - search_pattern: "*/reports/*.xml"
- deploy-to-bitrise-io@2: {}
```

## Relevant Links

* https://docs.bitrise.io/en/bitrise-ci/testing/testing-react-native-apps/running-unit-and-ui-tests-for-react-native-apps
* https://docs.bitrise.io/en/bitrise-ci/testing/test-reports/deploying-test-results
