# (iOS) Example CI Worklow

## Description

Example Workflow for commits on the main branch of an iOS app. The Workflow contains:

1. Installing [Cocoapods](/recipes/ios-cache-cocoapods.md) and [Carthage](/recipes/ios-install-carthage-dependencies.md) dependecies.
2. [Running all unit and UI tests in simulator](/recipes/ios-simulator-test.md)
3. [Building a test app and uploading to bitrise.io](/recipes/ios-deploy-to-bitrise.md)
4. [Sending a Slack notification with the build status](/recipes/slack-send-build-status.md)
5. [Filling the cache for upcoming pull request builds](/recipes/pull-request-build-caching.md)

## bitrise.yml

```yaml
---
format_version: '11'
default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
project_type: ios
workflows:
  ci:
    steps:
    - activate-ssh-key@4:
        run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
    - git-clone@6: {}
    - cache-pull@2: {}
    - cocoapods-install@2: {}
    - carthage@3:
        inputs:
        - carthage_options: "--use-xcframeworks --platform iOS"
    - recreate-user-schemes@1:
        inputs:
        - project_path: "$BITRISE_PROJECT_PATH"
    - xcode-test@4:
        inputs:
        - log_formatter: xcodebuild
        - xcodebuild_options: "-enableCodeCoverage YES"
    - xcode-archive@4:
        inputs:
        - project_path: "$BITRISE_PROJECT_PATH"
        - scheme: "$BITRISE_SCHEME"
        - automatic_code_signing: apple-id
        - distribution_method: development
    - deploy-to-bitrise-io@2: {}
    - slack@3:
        inputs:
        - channel: "#build-notifications"
        - webhook_url: "$SLACK_WEBHOOK"
    - cache-push@2: {}
app:
  envs:
  - opts:
      is_expand: false
    BITRISE_PROJECT_PATH: BitriseTest.xcworkspace
  - opts:
      is_expand: false
    BITRISE_SCHEME: BitriseTest
  - opts:
      is_expand: false
    BITRISE_DISTRIBUTION_METHOD: development
trigger_map:
- push_branch: main
  workflow: ci
```
