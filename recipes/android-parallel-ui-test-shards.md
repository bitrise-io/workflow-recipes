# (Android) Run UI tests in parallel shards

## Description

Split the UI / instrumented tests of a module into shards and run each shard on its own emulator, in parallel copies of the same Workflow in a Pipeline.

The `ui_test_shards` Pipeline runs `$TEST_SHARD_COUNT` copies of the `ui_tests` Workflow. Each copy gets its zero-based index in `$BITRISE_IO_PARALLEL_INDEX` and the number of copies in `$BITRISE_IO_PARALLEL_TOTAL`, and passes them to the Android test runner, which runs only the tests of that shard.

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
1. Click the **Start/Schedule a Build** button, and select the `ui_test_shards` option in the **Workflow, Pipeline** dropdown menu at the bottom of the popup.

In your own project, set the number of shards in the `TEST_SHARD_COUNT` Environment Variable, and the module in the Gradle task and in the test result base path.

## bitrise.yml

```yaml
format_version: "26"
default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
project_type: android

app:
  envs:
  - TEST_SHARD_COUNT: 2

meta:
  bitrise.io:
    stack: ubuntu-latest-stable
    machine_type_id: standard

pipelines:
  ui_test_shards:
    workflows:
      ui_tests:
        parallel: $TEST_SHARD_COUNT

workflows:
  ui_tests:
    steps:
    - git-clone@8: {}
    - avd-manager@2:
        inputs:
        - api_level: 34
        - tag: google_apis
        - abi: x86_64
    - gradle-runner@6:
        inputs:
        - gradlew_path: ./gradlew
        - gradle_task: |-
            :app:connectedDebugAndroidTest \
              -Pandroid.testInstrumentationRunnerArguments.numShards=$BITRISE_IO_PARALLEL_TOTAL \
              -Pandroid.testInstrumentationRunnerArguments.shardIndex=$BITRISE_IO_PARALLEL_INDEX
    - custom-test-results-export@2:
        inputs:
        - search_pattern: "*.xml"
        - base_path: $BITRISE_SOURCE_DIR/app/build/outputs/androidTest-results
        - test_name: UI tests shard $BITRISE_IO_PARALLEL_INDEX
    - deploy-to-bitrise-io@2: {}
```

## Relevant Links

* https://docs.bitrise.io/en/bitrise-ci/testing/test-sharding
* https://docs.bitrise.io/en/bitrise-ci/workflows-and-pipelines/build-pipelines/configuring-a-bitrise-pipeline#running-variations-of-the-same-workflow
