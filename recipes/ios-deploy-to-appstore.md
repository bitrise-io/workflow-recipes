# (iOS) Deploy to App Store Connect / TestFlight 

## Description
Archiving the app and uploading to App Store Connect to either release it to App Store or to TestFlight.

## Prerequisites

1. The source code is cloned and the dependencies (for example, Cocoapods, Carthage) are installed.
2. You have code signing set up. See [iOS Code Signing](https://devcenter.bitrise.io/en/code-signing/ios-code-signing.html) for more details.
3. You have Apple Developer connection set up. See [Apple services connection](https://devcenter.bitrise.io/en/accounts/connecting-to-services/apple-services-connection.html) for more details.

## Instructions

1. (Optional) Add the [Set Xcode Project Build Number](https://www.bitrise.io/integrations/steps/set-xcode-build-number) Step. Set the input variables:
    - **Info.plist file path**: for example, `MyApp/Info.plist`.
    - **Build Number**: for example, `42`.
    - **Version Number**: for example, `1.1`.
2. Add the [Xcode Archive & Export for iOS](https://github.com/bitrise-steplib/steps-xcode-archive) step. Set the input variables:
    - **Project path**: by default `$BITRISE_PROJECT_PATH`. Normally, you don't have to change this. 
    - **Scheme**: by default `$BITRISE_SCHEME`. This Environment Variable stores the scheme that you set when adding the app. The scheme always must be a shared scheme.
    - **Distribution method**: it must be set to `app-store`. 
3. Add the [Deploy to App Store Connect - Application Loader (formerly iTunes Connect)](https://github.com/bitrise-steplib/steps-deploy-to-itunesconnect-application-loader) Step. Set the input variables:
    - **Bitrise Apple Developer Connection**: for example, `api_key`

Alternatively you can use the [Deploy to App Store Connect with Deliver (formerly iTunes Connect)](https://github.com/bitrise-steplib/steps-deploy-to-itunesconnect-deliver) step as well, which gives you more options.

## bitrise.yml

```yaml
- set-xcode-build-number@1:
    inputs:
    - build_short_version_string: '1.0'
    - plist_path: BitriseTest/Info.plist
- xcode-archive@4:
    inputs:
    - project_path: "$BITRISE_PROJECT_PATH"
    - scheme: "$BITRISE_SCHEME"
    - automatic_code_signing: api_key
    - distribution_method: app-store
- deploy-to-itunesconnect-application-loader@1:
    inputs:
    - connection: api_key
```
