# (iOS) Merging test results and deploying to the Test Reports add-on

## Description

Test Reports add-on is tied to Bitrise builds. To make all the test reports generated in different builds appear on a single page in the add-on, the reports need to be merged and deployed in an additional build.

This example uses the [sample-swift-project-with-parallel-ui-test](https://github.com/bitrise-io/sample-swift-project-with-parallel-ui-test) iOS Open Source sample app and extends the ‘Run iOS test groups in parallel’ example Pipeline config with merging and deploying test results.

`run_ui_tests` and `run_unit_tests` Workflows are extended with a `deploy-to-bitrise-io` Step to make the generated test results available for the next Stage.

`run_tests_groups` Pipeline is extended with a new Stage: `deploy_test_results`.

This Stage runs the `deploy_test_results` Workflow:
1. `artifact-pull` Step downloads all the previous stage (`run_tests_groups`) generated zipped test results.
1. `script` Step unzips each test result into a new test run directory within the Test Report add-on deploy dir and creates the related `test-info.json` file.
1. `deploy-to-bitrise-io` Step deploys the merged test results.

![A screenshot of the example Pipeline in Bitrise's web UI](./ios-merging-test-results-and-deploying-to-the-test-reports-add-on.png)

## Instructions

1. Visit the [Create New App page](https://app.bitrise.io/apps/add) to create a new App.
1. When prompted to select a git repository, choose **Other/Manual** and paste the sample project repository URL (`https://github.com/bitrise-io/sample-swift-project-with-parallel-ui-test`) in the **Git repository (clone) URL** field.
1. Confirm that this is a public repository in the resulting pop-up.
1. Select the `master` branch to scan.
1. Wait for the project scanner to complete.
1. Select any of the offered Distribution methods (for example development, it does not really matter as now we are focusing on testing).
1. Confirm the offered stack, skip choosing the app icon and the webhook registration and kick off the first build.
1. Open the new Bitrise project’s Workflow Editor. 
1. Go to the **bitrise.yml** tab and replace the existing `bitrise.yml` with the contents of the example bitrise.yml below.
1. Click the **Start/Schedule a Build** button, and select the `run_tests_groups` option in the **Workflow, Pipeline** dropdown menu at the bottom of the popup.
1. Open the Pipeline’s build page.
1. Select the `deploy_test_results` build.
1. Click on **Details & Add-ons** on the build details page and select the Test Reports add-on to view the merged test reports.

## bitrise.yml

```yaml
---
format_version: '11'
default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
project_type: ios

app:
  envs:
  - BITRISE_PROJECT_PATH: BullsEye.xcworkspace
  - BITRISE_SCHEME: BullsEye

meta:
  bitrise.io:
    stack: osx-xcode-13.2.x

pipelines:
  run_tests_groups:
    stages:
    - build_tests: {}
    - run_tests_groups: {}
    - deploy_test_results: {}

stages:
  build_tests:
    workflows:
    - build_tests: {}

  run_tests_groups:
    workflows:
    - run_ui_tests: {}
    - run_unit_tests: {}
    
  deploy_test_results:
    workflows:
    - deploy_test_results: {}

workflows:
  build_tests:
    steps:
    - git-clone: {}
    - xcode-build-for-test:
        inputs:
        - destination: generic/platform=iOS Simulator
    - deploy-to-bitrise-io: {}

  run_ui_tests:
    before_run:
    - _pull_test_bundle
    steps:
    - xcode-test-without-building:
        inputs:
        - xctestrun: "$BITRISE_TEST_BUNDLE_PATH/BullsEye_UITests_iphonesimulator15.2-arm64-x86_64.xctestrun"
        - destination: platform=iOS Simulator,name=iPhone 12 Pro Max
    - deploy-to-bitrise-io: {}

  run_unit_tests:
    before_run:
    - _pull_test_bundle
    steps:
    - xcode-test-without-building:
        inputs:
        - xctestrun: "$BITRISE_TEST_BUNDLE_PATH/BullsEye_UnitTests_iphonesimulator15.2-arm64-x86_64.xctestrun"
        - destination: platform=iOS Simulator,name=iPhone 12 Pro Max
    - deploy-to-bitrise-io: {}
        
  deploy_test_results:
    steps:
    - git::https://github.com/bitrise-steplib/bitrise-step-artifact-pull.git@main:
        inputs:
        - artifact_sources: run_tests_groups.*
    - script:
        inputs:
        - content: |
            #!/usr/bin/env bash
            set -eo pipefail

            echo "Pulled Test Results: $BITRISE_ARTIFACT_PATHS"

            i=1
            IFS='|'
            read -ra FILES <<< "$BITRISE_ARTIFACT_PATHS"
            for f in ${FILES[@]}
            do
              if [ "${f: -4}" = ".zip" ]; then
                testing_addon_test_dir="${BITRISE_TEST_RESULT_DIR}/test_${i}"

                echo "Unzipping $f"

                dir="$(dirname $f)"
                ( cd $dir ; unzip "$f" -d $testing_addon_test_dir &> /dev/null )

                test_info="${testing_addon_test_dir}/test-info.json"
                echo "Generating test info at: $test_info"
                echo "{ \"test-name\": \"Tests ${i}\" }" > "$test_info"

                ((i++))
              fi
            done
    - deploy-to-bitrise-io: {}
        
  _pull_test_bundle:
    steps:
    - git::https://github.com/bitrise-steplib/bitrise-step-artifact-pull.git@main:
        inputs:
        - export_map: 'BITRISE_TEST_BUNDLE_ZIP_PATH: .*\.zip'
        - artifact_sources: build_tests.build_tests.*
    - script:
        inputs:
        - content: |-
            #!/usr/bin/env bash
            set -e
            set -o pipefail

            unzip "$BITRISE_TEST_BUNDLE_ZIP_PATH" -d "./test_bundle"
            envman add --key "BITRISE_TEST_BUNDLE_PATH" --value "./test_bundle"
```
