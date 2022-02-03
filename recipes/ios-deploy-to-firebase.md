# (iOS) Deploy to Firebase App Distribution

## Description

Build and distribute your app to testers via Firebase App Distribution.

## Prerequisites

* An existing Firebase project where your exact bundle ID is registered. Follow the [Firebase documentation](https://firebase.google.com/docs/app-distribution/ios/distribute-console) for details.
* Obtain a token from Firebase by running `firebase login:ci` locally. See the [Firebase CLI](https://firebase.google.com/docs/cli#sign-in-test-cli) docs for more details.
* Add this token as a secret your Bitrise project with the name `FIREBASE_TOKEN`.
* Get your Firebase App ID from your project's **General Settings** page and pass this value as an input to the `firebase-app-distribution` Step.
* Settings up code signing on Bitrise is not part of this guide, please follow our [code signing docs](https://devcenter.bitrise.io/en/code-signing/ios-code-signing.html#ios-code-signing-53933) for instructions.

## Instructions

1. Add the [Xcode Archive](https://github.com/bitrise-steplib/steps-xcode-archive) Step and set the required inputs, such as scheme, distribution method and the desired code signing method.
2. Add the [Firebase App Distribution](https://github.com/guness/bitrise-step-firebase-app-distribution) Step and set the following inputs:
    * **Firebase token**: use the secret env var previously defined: `$FIREBASE_TOKEN`.
    * **Firebase App ID**: see the Prerequisites section above for details.
    * Optionally, you can define test groups or individual testers in the Step inputs.

## bitrise.yml

```yaml
- xcode-archive@6:
    inputs:
    - distribution_method: development
    - scheme: # your scheme goes here
    - automatic_code_signing: api-key
- firebase-app-distribution@0:
    inputs:
    - firebase_token: $FIREBASE_TOKEN
    - app: # your app ID from Firebase
    - testers: email@company.com # optional
    - groups: qa-team #optional
```
