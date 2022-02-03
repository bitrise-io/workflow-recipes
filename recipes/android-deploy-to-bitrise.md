# (Android) Deploy to Bitrise.io

## Description

Build and distribute your app to testers via the [Bitrise.io Ship add-on](https://devcenter.bitrise.io/en/deploying/deploying-with-ship/getting-started-with-ship.html).

## Prerequisites

* If you want to deploy a release build, don't forget to [set up code signing on Bitrise](https://devcenter.bitrise.io/en/code-signing/android-code-signing.html) to build and sign the APK with your release key.

## Instructions

1. Add the [Android Build](https://github.com/bitrise-steplib/bitrise-step-android-build) Step and set the following inputs:
    - **Build type**: Set this to `apk`.
    - **Variant**: Use `release`, `debug`, or one of your custom variants if you have any.
2. If you build a release variant, add the [Android Sign](https://github.com/bitrise-steplib/steps-sign-apk) Step. You can skip this if you plan to deploy an unsigned debug variant.
3. Add a [Deploy to Bitrise.io - Apps, Logs, Artifacts](https://www.bitrise.io/integrations/steps/deploy-to-bitrise-io) Step.

## bitrise.yml

```yaml
- android-build@1:
    inputs:
    - variant: release
    - build_type: apk
- sign-apk@1: {}
- deploy-to-bitrise-io@2: {}
```
