# (Expo) Build using Turtle CLI

## Description

Publish an app to Expo's servers and build an iOS App Store .ipa and Android .aab files from your Expo project using Turtle CLI.

## Prerequisites

1. Generate an iOS Distribution Certificate and an App Store Provisioning Profile based on the [Generating iOS code signing files](https://devcenter.bitrise.io/en/code-signing/ios-code-signing/generating-ios-code-signing-files.html) guide.
2. Generate an Android Keystore by following the [Android code signing with Android Studio](https://devcenter.bitrise.io/en/code-signing/android-code-signing/android-code-signing-with-android-studio.html) guide.
3. Make sure you can [Publish your Expo project](https://docs.expo.dev/classic/turtle-cli/#publish-your-project) locally.

## Instructions

1. Upload the project's iOS Distribution Certificate and App Store Provisioning Profile on the Bitrise project's Workflow Editor / Code signing tab.
2. Upload the project's Android Keystore on the Bitrise project's Workflow Editor / Code signing tab.
3. Create a new Secret (`IOS_DEVELOPMENT_TEAM`) with the ID of the iOS Development Team, issued the project's Certificate and Provisioning Profile.
4. Store the Expo account, used for publishing the Expo app and fetching the app manifest, in `EXPO_USERNAME` and `EXPO_PASSWORD` secrets.
5. Copy-paste `envs` to your workflow.
6. Copy-paste `steps` to your workflow.
    - The built .ipa and .aab files are exposed via `BITRISE_IPA_PATH` and `BITRISE_AAB_PATH` env vars.

## bitrise.yml

```yaml
  turtle_build:
    envs:
    - KEYSTORE_PATH: /tmp/keystore.jks
    - KEYSTORE_ALIAS: $BITRISEIO_ANDROID_KEYSTORE_ALIAS
    - EXPO_ANDROID_KEYSTORE_PASSWORD: $BITRISEIO_ANDROID_KEYSTORE_PASSWORD
    - EXPO_ANDROID_KEY_PASSWORD: $BITRISEIO_ANDROID_KEYSTORE_PRIVATE_KEY_PASSWORD
    - PROFILE_PATH: /tmp/profile.mobileprovision
    - CERTIFICATE_PATH: /tmp/certificate.p12
    - EXPO_IOS_DIST_P12_PASSWORD: $BITRISE_CERTIFICATE_PASSPHRASE
    - IOS_DEVELOPMENT_TEAM: $IOS_DEVELOPMENT_TEAM
    - EXPO_USERNAME: $EXPO_USERNAME
    - EXPO_PASSWORD: $EXPO_PASSWORD
    steps:
    - script@1:
        title: Install dependencies
        inputs:
        - content: |-
            #!/usr/bin/env bash
            set -ex

            node --version
            fastlane --version

            npm install -g turtle-cli
            turtle --version

            npm install -g expo-cli
            expo --version
    - file-downloader@1:
        title: Download Android Keystore
        inputs:
        - destination: $KEYSTORE_PATH
        - source: $BITRISEIO_ANDROID_KEYSTORE_URL
    - file-downloader@1:
        title: Download iOS Certificate
        inputs:
        - destination: $CERTIFICATE_PATH
        - source: $BITRISE_CERTIFICATE_URL
    - file-downloader@1:
        title: Download iOS Provisioning Profile
        inputs:
        - destination: $PROFILE_PATH
        - source: $BITRISE_PROVISION_URL
    - npm@1:
        title: Install project dependencies
        inputs:
        - command: install
    - set-java-version@1:
        title: Set Java version to Java 8
        inputs:
        - set_java_version: "8"
    - script@1:
        title: Run Expo publish
        inputs:
        - content: |-
            #!/usr/bin/env bash
            set -ex

            expo login -u $EXPO_USERNAME -p $EXPO_PASSWORD --non-interactive
            expo publish
    - script@1:
        title: Run Turtle build
        inputs:
        - content: |-
            #!/usr/bin/env bash
            set -ex

            turtle setup:android
            aab_path=$BITRISE_DEPLOY_DIR/expo-project.aab
            turtle build:android --type app-bundle --keystore-path $KEYSTORE_PATH --keystore-alias $KEYSTORE_ALIAS -o $aab_path
            envman add --key BITRISE_AAB_PATH --value $aab_path

            turtle setup:ios
            ipa_path=$BITRISE_DEPLOY_DIR/expo-project.ipa
            turtle build:ios --team-id $IOS_DEVELOPMENT_TEAM --dist-p12-path $CERTIFICATE_PATH --provisioning-profile-path $PROFILE_PATH -o $ipa_path
            envman add --key BITRISE_IPA_PATH --value $ipa_path
```
