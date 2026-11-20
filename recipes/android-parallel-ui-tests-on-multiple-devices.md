# (Android) Run UI tests in parallel on multiple devices

## Description

Running the UI or instrumented tests of a single module on different emulated devices, in parallel Workflows utilizing Pipelines.

The Pipeline runs three Workflows in parallel: `ui_test_on_phone`, `ui_test_on_tablet` and `ui_test_on_foldable`. They use the `run_ui_tests` [Step bundle](https://docs.bitrise.io/en/bitrise-ci/workflows-and-pipelines/steps/step-bundles) with a different `emulator_profile` input. The bundle:

1. Starts an emulator with the given device profile and waits for it to boot using the `avd-manager` Step.
1. Runs the UI tests of the given module and variant with Gradle using the `gradle-runner` Step.
1. Collects the JUnit XML results using the `custom-test-results-export` Step, named after the device profile.
1. Runs the `deploy-to-bitrise-io` Step, which makes the test results available on the **Tests** tab of the build page.

To split the tests of one device into shards instead, see [UI tests in parallel shards](./android-parallel-ui-test-shards.md).

## Prerequisites

1. A Linux stack: Android emulators don't run on macOS stacks.

## Instructions

To test this configuration in a new Bitrise example project, do the following:

1. Visit the [Create New App page](https://app.bitrise.io/apps/add) to create a new App.
1. When prompted to select a git repository, choose **Other/Manual** and paste the sample project repository URL (`https://github.com/bitrise-io/Bitrise-Android-Modules-Sample.git`) in the **Git repository (clone) URL** field.
1. Confirm that this is a public repository in the resulting pop-up.
1. Select the `main` branch to scan.
1. Wait for the project scanner to complete.
1. Enter `app` as the specified module.
1. Enter `debug` as the specified variant.
1. Continue through the prompts as normal, no changes are needed.
1. Open the new Bitrise project's Workflow Editor.
1. Go to the **bitrise.yml** tab, and replace the existing yaml contents with the contents of the example `bitrise.yml` below.
1. Click the **Start/Schedule a Build** button, and select the `ui_test_on_multiple_devices` option in the **Workflow, Pipeline** dropdown menu at the bottom of the popup.

To list the device profiles you can use, run `avdmanager list device` and use the `id` values.

## bitrise.yml

```yaml
format_version: "26"
default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
project_type: android

meta:
  bitrise.io:
    stack: ubuntu-latest-stable
    machine_type_id: standard

pipelines:
  ui_test_on_multiple_devices:
    workflows:
      ui_test_on_phone: {}
      ui_test_on_tablet: {}
      ui_test_on_foldable: {}

step_bundles:
  run_ui_tests:
    inputs:
    - emulator_profile: pixel_5
    steps:
    - git-clone@8: {}
    - avd-manager@2:
        inputs:
        - profile: $emulator_profile
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
        - test_name: UI tests on $emulator_profile
    - deploy-to-bitrise-io@2: {}

workflows:
  ui_test_on_phone:
    steps:
    - bundle::run_ui_tests: {}

  ui_test_on_tablet:
    steps:
    - bundle::run_ui_tests:
        inputs:
        - emulator_profile: 10.1in WXGA (Tablet)

  ui_test_on_foldable:
    steps:
    - bundle::run_ui_tests:
        inputs:
        - emulator_profile: 8in Foldable
```
