# (Android) Run tests using the emulator

## Description

Run any kind of tests (unit, instrumentation) on a local emulator instance.

## Instructions

1. Add an [AVD Manager](https://github.com/bitrise-steplib/steps-avd-manager) Step. To customize the emulator, see the [step configuration](https://github.com/bitrise-steplib/steps-avd-manager#%EF%B8%8F-configuration).
2. Add a [Wait for Android emulator](https://github.com/bitrise-steplib/steps-wait-for-android-emulator) step.
3. Add a [Gradle Runner](https://github.com/bitrise-steplib/steps-gradle-runner) step. Set the input variables:
    - **gradlew file path**: for example, `./gradlew`.
    - **Gradle task to run**: for example, `connectedDebugAndroidTest`.
4. Add a [Export test results to Test Reports add-on](https://github.com/bitrise-steplib/step-custom-test-results-export) Step with the following inputs:
    - The name of the test: `Emulator tests`.
    - Test result base path: `$BITRISE_SOURCE_DIR/app/build/outputs/androidTest-results`. You might want to adjust the path based on the module name(s) in your project.
    - Test result search pattern: `*.xml`.
5. Add a [Deploy to Bitrise.io - Apps, Logs, Artifacts](https://www.bitrise.io/integrations/steps/deploy-to-bitrise-io) Step that makes the test results available in the [Test Reports add-on](https://devcenter.bitrise.io/en/testing/test-reports.html).

## bitrise.yml

```yaml
- avd-manager@1: {}
- wait-for-android-emulator@1:
- gradle-runner@2:
    inputs:
    - gradlew_path: ./gradlew
    - gradle_task: connectedDebugAndroidTest
- custom-test-results-export@0:
    inputs:
    - search_pattern: "*.xml"
    - base_path: $BITRISE_SOURCE_DIR/app/build/outputs/androidTest-results
    - test_name: Emulator tests
- deploy-to-bitrise-io@2:
```
