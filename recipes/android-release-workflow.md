# (Android) Example Release Workflow

## Description

Example workflow for uploading a release draft of an app to Google Play. The workflow contains:

1. [Setting the verion name](https://www.bitrise.io/integrations/steps/change-android-versioncode-and-versionname) based on [Env Vars passed to the build](https://devcenter.bitrise.io/en/builds/environment-variables.html#setting-a-custom-env-var-when-starting-a-build) (`$VERSION_NAME`).
2. [Creating a release Android App Bundle and uploading it to Google Play](/recipes/android-deploy-to-google-play.md).

## bitrise.yml

```yaml
---
format_version: '11'
default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
project_type: android
workflows:
  release:
    steps:
    - activate-ssh-key@4:
        run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
    - git-clone@6: {}
    - change-android-versioncode-and-versionname@1:
        inputs:
        - new_version_name: "$VERSION_NAME"
        - new_version_code: "$BITRISE_BUILD_NUMBER"
        - build_gradle_path: "$PROJECT_LOCATION/$MODULE/build.gradle"
    - android-build@1:
        inputs:
        - project_location: "$PROJECT_LOCATION"
        - module: "$MODULE"
        - build_type: aab
        - variant: release
    - sign-apk@1: {}
    - google-play-deploy@3:
        inputs:
        - service_account_json_key_path: "$BITRISEIO_SERVICE_ACCOUNT_JSON_KEY_URL"
        - package_name: io.bitrise.sample.android
        - status: draft
        - track: production
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
```

## Relevant Links

* https://devcenter.bitrise.io/en/builds/environment-variables.html#setting-a-custom-env-var-when-starting-a-build
