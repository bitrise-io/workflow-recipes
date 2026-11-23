# (Android) Run UI / instrumentation tests on local emulator

## Description

Run UI / instrumentation tests on a local emulator instance with Gradle, and view the results on the **Tests** tab of the build page.

## Prerequisites

1. A Linux stack: Android emulators don't run on macOS stacks.

## Instructions

1. Add an [AVD Manager](https://bitrise.io/integrations/steps/avd-manager) Step and set the emulator in the **Android API Level**, **OS Tag** and **ABI** inputs. The ABI must match the CPU architecture of the stack. The Step waits for the emulator to boot. To customize the emulator further, see the [step configuration](https://github.com/bitrise-steplib/steps-avd-manager#%EF%B8%8F-configuration).
2. Add a [Gradle Runner](https://bitrise.io/integrations/steps/gradle-runner) step. Set the input variables:
    - **gradlew file path**: for example, `./gradlew`.
    - **Gradle task to run**: for example, `:app:connectedDebugAndroidTest`. Without the `:app:` prefix, Gradle runs the UI tests of every module, and the results of the other modules don't reach the **Tests** tab unless the base path below covers them.
3. Add an [Export test results to Test Reports](https://bitrise.io/integrations/steps/custom-test-results-export) Step with the following inputs. The Step runs even if the tests failed.
    - The name of the test: `Emulator tests`.
    - Test result base path: `$BITRISE_SOURCE_DIR/app/build/outputs/androidTest-results`. You might want to adjust the path based on the module name(s) in your project.
    - Test result search pattern: `*.xml`. The pattern also matches files in subfolders, and the Step merges every JUnit XML it finds into one test report.
4. Add a [Deploy to Bitrise.io](https://bitrise.io/integrations/steps/deploy-to-bitrise-io) Step that makes the test results available on the **Tests** tab.

The [Android Instrumented Test](https://bitrise.io/integrations/steps/android-instrumented-test) Step runs pre-built test APKs on the emulator instead, but it doesn't send its results to the **Tests** tab.

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
    - gradle_task: :app:connectedDebugAndroidTest
- custom-test-results-export@2:
    inputs:
    - search_pattern: "*.xml"
    - base_path: $BITRISE_SOURCE_DIR/app/build/outputs/androidTest-results
    - test_name: Emulator tests
- deploy-to-bitrise-io@2: {}
```

## Relevant Links

* https://docs.bitrise.io/en/bitrise-ci/testing/testing-android-apps/running-instrumented-tests-for-android-apps
* https://docs.bitrise.io/en/bitrise-ci/testing/test-reports/deploying-test-results
