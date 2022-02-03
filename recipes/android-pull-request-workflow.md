# (Android) Example Pull Request Worklow

## Description

Example workflow for Android Pull Request validation. The workflow contains:

1. [Running unit tests](/recipes/android-unit-test.md)
2. [Running UI tests on a virtual device](/recipes/android-virtual-device-tests.md)
3. [Running lint](/recipes/android-lint.md)
4. [Building a test app and uploading to bitrise.io](/recipes/android-deploy-to-bitrise.md)
5. [Sending the QR code of the test build to the Pull Request](/recipes/github-pull-request-build-qr-code.md)
6. Triggering the Workflow for pull requests.

## bitrise.yml

```yaml
---
format_version: '11'
default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
project_type: android
workflows:
  pull-request:
    steps:
    - activate-ssh-key@4:
        run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
    - git-clone@6: {}
    - cache-pull@2: {}
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
        - test_type: instrumentation
    - android-lint@0:
        inputs:
        - variant: "$VARIANT"
    - android-build@1:
        inputs:
        - project_location: "$PROJECT_LOCATION"
        - module: "$MODULE"
        - variant: "$VARIANT"
    - deploy-to-bitrise-io@2: {}
    - create-install-page-qr-code@1: {}
    - comment-on-github-pull-request@0:
        inputs:
        - body: |-
            ![QR code]($BITRISE_PUBLIC_INSTALL_PAGE_QR_CODE_IMAGE_URL)

            $BITRISE_PUBLIC_INSTALL_PAGE_URL
        - personal_access_token: "$GITHUB_ACCESS_TOKEN"
    - cache-push@2: {}
app:
  envs:
  - opts:
      is_expand: false
    PROJECT_LOCATION: "."
  - opts:
      is_expand: false
    MODULE: app
  - VARIANT: debug
    opts:
      is_expand: false
trigger_map:
- pull_request_source_branch: "*"
  workflow: pull-request
```
