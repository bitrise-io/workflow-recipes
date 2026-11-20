# (Android) Show screenshots of UI tests on the Tests tab

## Description

Show the screenshots, recordings and logs of your UI tests under their test cases on the **Tests** tab of the build page.

The Steps that produce an `.xcresult` file, such as Xcode Test for iOS, attach these files automatically. For other test tools, name each file after the test case that produced it, and the Export test results to Test Reports Step finds it next to the JUnit XML results.

## Prerequisites

1. A Linux stack: Android emulators don't run on macOS stacks.
1. Your tests save their files under the attachment name `<classname>__<name>__<label>.<ext>`, for example `com.example.LoginTest__emptyState__failure.png`. The [naming guide](https://docs.bitrise.io/en/bitrise-ci/testing/test-reports/test-attachments/naming-attachments-after-test-cases#setting-up-your-test-tool) has ready-made setups for Espresso, Paparazzi, Roborazzi, Detox and Playwright. The Espresso one takes a screenshot when a check fails and saves it with the test storage service.

## Instructions

1. Add an [AVD Manager](https://bitrise.io/integrations/steps/avd-manager) Step and set the emulator in the **Android API Level**, **OS Tag** and **ABI** inputs.
1. Add a [Gradle Runner](https://bitrise.io/integrations/steps/gradle-runner) Step that runs your UI tests, for example with the `connectedDebugAndroidTest` task.
1. Add an [Export test results to Test Reports](https://bitrise.io/integrations/steps/custom-test-results-export) Step (version 1.3.0 or newer) with the following inputs. The Step runs even if the tests failed.
    - Test result base path: a folder that contains both the JUnit XML results and the attachments, for example `$BITRISE_SOURCE_DIR/app/build/outputs`. The Step looks for attachments only under this folder.
    - Test result search pattern: `*/androidTest-results/*.xml`, so that only the JUnit XML results match.
    - The name of the test: `UI tests`.
1. Add a [Deploy to Bitrise.io](https://bitrise.io/integrations/steps/deploy-to-bitrise-io) Step (version 2.26.0 or newer) that deploys the test results with their attachments to the **Tests** tab.

If an attachment doesn't show up, the Export test results to Test Reports Step logs a warning for names that match more than one test case, and the Deploy to Bitrise.io Step logs a warning for files it can't link to a test case.

## bitrise.yml

```yaml
- avd-manager@2:
    inputs:
    - api_level: 34
    - tag: google_apis
    - abi: x86_64
- gradle-runner@6:
    inputs:
    - gradlew_path: ./gradlew
    - gradle_task: connectedDebugAndroidTest
- custom-test-results-export@2:
    inputs:
    - base_path: $BITRISE_SOURCE_DIR/app/build/outputs
    - search_pattern: "*/androidTest-results/*.xml"
    - test_name: UI tests
- deploy-to-bitrise-io@2: {}
```

## Relevant Links

* https://docs.bitrise.io/en/bitrise-ci/testing/test-reports/test-attachments
* https://docs.bitrise.io/en/bitrise-ci/testing/test-reports/test-attachments/naming-attachments-after-test-cases
* https://docs.bitrise.io/en/bitrise-ci/testing/test-reports/test-attachments/referencing-attachments-in-junit-xml
