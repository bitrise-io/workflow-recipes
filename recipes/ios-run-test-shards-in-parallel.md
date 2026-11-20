# (iOS) Run tests in parallel shards

## Description

Split the tests of a Test Plan into shards and run each shard on its own simulator, in parallel copies of the same Workflow in a Pipeline.

This example uses the [sample-swift-project-with-parallel-ui-test](https://github.com/bitrise-io/sample-swift-project-with-parallel-ui-test) iOS Open Source sample app.

The `build_and_run_test_shards` Pipeline has two Workflows:

1. `build_for_testing`: Builds the app and the tests of the `FullTests` Test Plan with the `xcode-build-for-test` Step. The `xcode-test-shard-calculation` Step lists the built tests and splits them alphabetically into `$TEST_SHARD_COUNT` shards. The `deploy-to-bitrise-io` Step shares the test bundle and the shards with the next Workflow.
1. `run_test_shard`: The Pipeline runs `$TEST_SHARD_COUNT` copies of this Workflow. Each copy pulls the test bundle, finds its `.xctestrun` file (its name contains the SDK version, which changes with the Xcode version of the stack), and runs only the tests of its own shard with the `xcode-test-without-building` Step. The shard is selected by `$BITRISE_IO_PARALLEL_INDEX`, the zero-based index of the copy. The `deploy-to-bitrise-io` Step makes the test results available on the **Tests** tab of the build page.

## Instructions

1. Visit the [Create New App page](https://app.bitrise.io/apps/add) to create a new App.
1. When prompted to select a git repository, choose **Other/Manual** and paste the sample project repository URL (`https://github.com/bitrise-io/sample-swift-project-with-parallel-ui-test`) in the **Git repository (clone) URL** field.
1. Confirm that this is a public repository in the resulting pop-up.
1. Select the `master` branch to scan.
1. Wait for the project scanner to complete.
1. Select any of the offered Distribution methods (for example development, it does not really matter as now we are focusing on testing).
1. Confirm the offered stack, skip choosing the app icon and the webhook registration and kick off the first build.
1. Open the new Bitrise project's Workflow Editor.
1. Go to the **bitrise.yml** tab and replace the existing `bitrise.yml` with the contents of the example bitrise.yml below.
1. Click the **Start/Schedule a Build** button, and select the `build_and_run_test_shards` option in the **Workflow, Pipeline** dropdown menu at the bottom of the popup.

In your own project, set the number of shards in the `TEST_SHARD_COUNT` Environment Variable, and the Test Plan in the `xcode-build-for-test` Step and in the script that finds the `.xctestrun` file.

## bitrise.yml

```yaml
---
format_version: '26'
default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
project_type: ios

app:
  envs:
  - BITRISE_PROJECT_PATH: BullsEye.xcworkspace
  - BITRISE_SCHEME: BullsEye
  - TEST_SHARD_COUNT: 3

meta:
  bitrise.io:
    stack: osx-xcode-latest-stable
    machine_type_id: g2.mac.medium

pipelines:
  build_and_run_test_shards:
    workflows:
      build_for_testing: {}
      run_test_shard:
        depends_on: [build_for_testing]
        parallel: $TEST_SHARD_COUNT

workflows:
  build_for_testing:
    steps:
    - git-clone@8: {}
    - xcode-build-for-test@3:
        inputs:
        - destination: generic/platform=iOS Simulator
        - test_plan: FullTests
    - xcode-test-shard-calculation@0:
        inputs:
        - product_path: $BITRISE_XCTESTRUN_FILE_PATH
        - shard_count: $TEST_SHARD_COUNT
    - deploy-to-bitrise-io@2:
        inputs:
        - pipeline_intermediate_files: |-
            $BITRISE_TEST_BUNDLE_PATH:BITRISE_TEST_BUNDLE_PATH
            $BITRISE_TEST_SHARDS_PATH:BITRISE_TEST_SHARDS_PATH

  run_test_shard:
    steps:
    - pull-intermediate-files@1:
        inputs:
        - artifact_sources: build_for_testing
    - script@1:
        title: Find the xctestrun file
        inputs:
        - content: |-
            set -eo pipefail
            xctestrun=$(ls "$BITRISE_TEST_BUNDLE_PATH"/BullsEye_FullTests_*.xctestrun)
            envman add --key XCTESTRUN_PATH --value "$xctestrun"
    - xcode-test-without-building@1:
        inputs:
        - xctestrun: $XCTESTRUN_PATH
        - destination: platform=iOS Simulator,name=Bitrise iOS default,OS=latest
        - only_testing: $BITRISE_TEST_SHARDS_PATH/$BITRISE_IO_PARALLEL_INDEX
    - deploy-to-bitrise-io@2: {}
```

## Relevant Links

* https://docs.bitrise.io/en/bitrise-ci/testing/test-sharding
* https://docs.bitrise.io/en/bitrise-ci/workflows-and-pipelines/build-pipelines/configuring-a-bitrise-pipeline#running-variations-of-the-same-workflow
