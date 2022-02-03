# (Android) Deploy to Firebase App Distribution

## Description

Build and distribute your app to testers via Firebase App Distribution.

This example builds and deploys an APK, but the workflow can be tweaked to distribute AAB instead.

## Prerequisites

* An existing Firebase project where your exact package name is registered. See the [Firebase documentation](https://firebase.google.com/docs/app-distribution/android/distribute-console?apptype=apk) for details.
* Obtain a token from Firebase by running `firebase login:ci` locally. See the [Firebase CLI](https://firebase.google.com/docs/cli#sign-in-test-cli) docs for details.
* Add this token as a [Secret](https://devcenter.bitrise.io/en/builds/secrets.html) to your Bitrise app with the name `FIREBASE_TOKEN`.
* Get your Firebase App ID from your project's General Settings page and pass this value as an input to the `firebase-app-distribution` step.
* If you want to deploy a release build, don't forget to [set up code signing on Bitrise](https://devcenter.bitrise.io/en/code-signing/android-code-signing.html) to build and sign the APK with your release key.

## Instructions

1. Add the [Android Build](https://github.com/bitrise-steplib/bitrise-step-android-build) Step and set the following inputs:
    - **Build type**: Set this to `apk`.
    - **Variant**: Use `release`, `debug`, or one of your custom variants if you have any.
2. If you build a release variant, add the [Android Sign](https://github.com/bitrise-steplib/steps-sign-apk) Step. You can skip this if you plan to deploy an unsigned debug variant.
3. Add the [Firebase App Distribution](https://github.com/guness/bitrise-step-firebase-app-distribution) step and set the following inputs:
    * Firebase token: use the secret env var previously defined: `$FIREBASE_TOKEN`
    * App path: this should point to the APK that the previous steps have built and signed. By default, it's located at `$BITRISE_DEPLOY_DIR/app-release-bitrise-signed.apk`, but the exact file name might be different based on your project config.
    * Firebase App ID: see the Prerequisites section above for details
    * Optional: you can define test groups or individual testers in the step inputs

## bitrise.yml

```yaml
- android-build@1:
    inputs:
    - variant: release
    - build_type: apk
- sign-apk@1: {}
- firebase-app-distribution@0:
    inputs:
    - firebase_token: $FIREBASE_TOKEN
    - app_path: $BITRISE_DEPLOY_DIR/app-release-bitrise-signed.apk
    - app: your_app_id_from_firebase
    - testers: email@company.com # optional
    - groups: qa-team #optional
```
