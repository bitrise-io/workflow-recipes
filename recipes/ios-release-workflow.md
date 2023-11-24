# (iOS) Example Release Workflow

## Description

Example Workflow for uploading a release draft of an iOS app to the App Store. The Worklow contains:

1. Installing [Cocoapods](/recipes/ios-key-cache-cocoapods.md) and [Carthage](/recipes/ios-install-carthage-dependencies.md) dependecies.
2. [Setting the version number](https://www.bitrise.io/integrations/steps/set-ios-version) based on [env vars passed to build](https://devcenter.bitrise.io/en/builds/environment-variables.html#setting-a-custom-env-var-when-starting-a-build) (`$VERSION_NUMBER`).
3. [Building a release build and uploading to App Store](/recipes/ios-deploy-to-appstore.md).

## bitrise.yml

```yaml
---
format_version: '13'
default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
project_type: ios

meta:
  bitrise.io:
    stack: osx-xcode-15.0.x
    machine_type_id: g2-m1.4core

workflows:
  release:
    steps:
    - activate-ssh-key@4:
        run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
    - git-clone@8: {}
    - restore-cocoapods-cache@1: {}
    - carthage@3:
        inputs:
        - carthage_options: "--use-xcframeworks --platform iOS"
    - set-xcode-build-number@1:
        inputs:
        - build_short_version_string: $VERSION_NUMBER
        - build_version: $BITRISE_BUILD_NUMBER
        - plist_path: BitriseTest/Info.plist
    - recreate-user-schemes@1:
        inputs:
        - project_path: $BITRISE_PROJECT_PATH
    - xcode-archive@5:
        inputs:
        - project_path: $BITRISE_PROJECT_PATH
        - scheme: $BITRISE_SCHEME
        - automatic_code_signing: apple-id
        - distribution_method: app-store
    - deploy-to-itunesconnect-application-loader@1:
        inputs:
        - connection: apple_id

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

* https://devcenter.bitrise.io/en/builds/environment-variables.html#setting-a-custom-env-var-when-starting-a-build
