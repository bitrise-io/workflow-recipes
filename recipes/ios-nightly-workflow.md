# (iOS) Example Nightly Workflow

## Description

Example Workflow for nightly builds for iOS apps. The Workflow contains:

1. Installing [Cocoapods](/recipes/ios-key-cache-cocoapods.md) and [Carthage](/recipes/ios-install-carthage-dependencies.md) dependecies.
2. [Setting the version and build number](https://www.bitrise.io/integrations/steps/set-ios-version). By default, the app will get the build number (`$BITRISE_BUILD_NUMBER`) as the version code.
3. [Building a release build and uploading to TestFlight](/recipes/ios-deploy-to-appstore.md).
4. [Building a test app and uploading to bitrise.io](/recipes/ios-deploy-to-bitrise.md).
5. [Sending the QR code of the test build to Slack](/recipes/slack-send-qr-code.md).

Check out the [guide](https://devcenter.bitrise.io/en/builds/starting-builds/scheduling-builds.html) to run scheduled builds.

## bitrise.yml

```yaml
---
format_version: '17'
default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
project_type: ios

meta:
  bitrise.io:
    stack: osx-xcode-16.0.x
    machine_type_id: g2.mac.medium

workflows:
  nightly:
    steps:
    - activate-ssh-key@4:
        run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
    - git-clone@8: {}
    - cocoapods-install@2: {}
    - carthage@3:
        inputs:
        - carthage_options: "--use-xcframeworks --platform iOS"
    - set-xcode-build-number@2:
        inputs:
        - build_short_version_string: '1.0'
    - xcode-archive@5:
        inputs:
        - project_path: $BITRISE_PROJECT_PATH
        - scheme: $BITRISE_SCHEME
        - automatic_code_signing: apple-id
        - distribution_method: app-store
    - deploy-to-itunesconnect-application-loader@1:
        inputs:
        - connection: apple_id
    - xcode-archive@5:
        inputs:
        - project_path: $BITRISE_PROJECT_PATH
        - scheme: $BITRISE_SCHEME
        - automatic_code_signing: apple-id
        - distribution_method: development
        - deploy-to-bitrise-io@2: {}
    - deploy-to-bitrise-io@2: {}
    - create-install-page-qr-code@1: {}
    - slack@4:
        inputs:
        - channel: "#build-notifications"
        - thumb_url: $BITRISE_PUBLIC_INSTALL_PAGE_QR_CODE_IMAGE_URL
        - webhook_url: $SLACK_WEBHOOK

app:
  envs:
  - BITRISE_PROJECT_PATH: BitriseTest.xcworkspace
    opts:
      is_expand: false
  - BITRISE_SCHEME: BitriseTest
    opts:
      is_expand: false
  - BITRISE_DISTRIBUTION_METHOD: development
    opts:
      is_expand: false
```

## Relevant Links

* https://devcenter.bitrise.io/en/builds/starting-builds/scheduling-builds.html
