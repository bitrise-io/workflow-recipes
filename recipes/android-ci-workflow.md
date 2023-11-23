# (Android) Example CI Workflow

## Description

Example Workflow for commits on the main branch of an Android app. The workflow contains:

1. [Running unit tests](/recipes/android-unit-test.md)
2. [Running UI tests on a virtual device](/recipes/android-virtual-device-tests.md)
3. [Running lint](/recipes/android-lint.md)
4. Building a test app
5. [Sending a Slack notification with the build status](/recipes/slack-send-build-status.md)
6. [Filling the cache for upcoming Pull Request builds](/recipes/pull-request-build-caching.md)

## Instructions

Use the yaml below and change the following env var values to match your project settings:
- `$PROJECT_LOCATION`
- `$MODULE`
- `$VARIANT`
- `$SLACK_WEBHOOK`

## bitrise.yml

```yaml
---
format_version: '13'
default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
project_type: android

meta:
  bitrise.io:
    stack: linux-docker-android-20.04
    machine_type_id: standard

workflows:
  ci:
    steps:
    - activate-ssh-key@4:
        run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
    - git-clone@8: {}
    - restore-gradle-cache@1: {}
    - android-unit-test@1:
        inputs:
        - project_location: $PROJECT_LOCATION
        - variant: $VARIANT
    - android-build-for-ui-testing@0:
        inputs:
        - variant: $VARIANT
        - module: $MODULE
    - virtual-device-testing-for-android@1:
        inputs:
        - test_devices: NexusLowRes,30,en,portrait
        - test_type: instrumentation
    - android-lint@0:
        inputs:
        - variant: $VARIANT
    - android-build@1:
        inputs:
        - project_location: $PROJECT_LOCATION
        - module: $MODULE
        - variant: $VARIANT
    - deploy-to-bitrise-io@2: {}
    - slack@4:
        inputs:
        - channel: "#build-notifications"
        - webhook_url: $SLACK_WEBHOOK
    - save-gradle-cache@1: {}

app:
  envs:
  - PROJECT_LOCATION: "."
    opts:
      is_expand: false
  - MODULE: app
    opts:
      is_expand: false
  - VARIANT: debug
    opts:
      is_expand: false

trigger_map:
- push_branch: main
  workflow: ci
```
