# (Android) Example Nightly Workflow

## Description

Example workflow for Android nightly builds. The workflow contains:

1. [Setting the version code and version name](https://www.bitrise.io/integrations/steps/change-android-versioncode-and-versionname). By default the app will get the build number (`$BITRISE_BUILD_NUMBER`) as the version code.
2. [Building a release Android App Bundle and uploading to Google Play internal testing](/recipes/android-deploy-to-google-play.md).
3. [Building a test app and uploading to bitrise.io](/recipes/android-deploy-to-bitrise.md).
4. [Sending the QR code of the test build to Slack](/recipes/slack-send-qr-code.md).

Check out the [guide](https://devcenter.bitrise.io/en/builds/starting-builds/scheduling-builds.html) to run scheduled builds.

## bitrise.yml

```yaml
---
format_version: '11'
default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
project_type: android
workflows:
  nightly:
    steps:
    - activate-ssh-key@4:
        run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
    - git-clone@6: {}
    - cache-pull@2: {}
    - change-android-versioncode-and-versionname@1:
        inputs:
        - new_version_name: 1.0.0
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
        - status: completed
        - track: internal
    - android-build@1:
        inputs:
        - project_location: "$PROJECT_LOCATION"
        - module: "$MODULE"
        - variant: "$VARIANT"
    - deploy-to-bitrise-io@2: {}
    - create-install-page-qr-code@1: {}
    - slack@3:
        inputs:
        - channel: "#build-notifications"
        - thumb_url: "$BITRISE_PUBLIC_INSTALL_PAGE_QR_CODE_IMAGE_URL"
        - webhook_url: "$SLACK_WEBHOOK"
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
```

## Relevant Links

* https://devcenter.bitrise.io/en/builds/starting-builds/scheduling-builds.html
