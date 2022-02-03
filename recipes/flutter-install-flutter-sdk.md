# (Flutter) Install Flutter SDK

## Description

Install the latest stable/beta or a specific version of Flutter.

## Instructions

1. Add the [Flutter Install](https://www.bitrise.io/integrations/steps/flutter-installer) step. **Use this step before the Cache Pull step to make sure caching works correctly.**
2. You can install either the latest stable/beta versions or a specific version:
    - By default, the Step installs the latest stable version.
    - To install the latest beta, set the **Flutter SDK git repository version** input to `beta`.
    - **Recommended:** To install a specific version, set the **Flutter SDK installation bundle URL** input. You can find the list of Flutter installation bundles here: https://flutter.dev/docs/development/tools/sdk/releases. Make sure you set the bundle based on the stack (MacOS or Linux).

## bitrise.yml

Specific version (recommended):

```yaml
- flutter-installer@0:
    inputs:
    - installation_bundle_url: https://storage.googleapis.com/flutter_infra_release/releases/stable/macos/flutter_macos_2.5.3-stable.zip
```

Latest stable:

```yaml
- flutter-installer@0: {}
```

Latest beta:

```yaml
- flutter-installer@0:
    inputs:
    - version: beta
```
