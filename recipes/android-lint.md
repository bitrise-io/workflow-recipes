# (Android) Run Lint

## Description
Runs Lint on your Android project and generates a report with the results.

## Instructions

1. Add the [Android Lint](https://www.bitrise.io/integrations/steps/android-lint) step. Set the input variables:
    - **Project Location**: Use the default `$BITRISE_SOURCE_DIR` or `$PROJECT_LOCATION`. You can set a specific path but the automatically exposed Environment Variables are usually the best option.
    - **Variant**: Use the `$VARIANT` Enviromment Variable, or specify a variant manually.
    - **Module**: Specify one or leave it blank to run lint in all of the modules.
2. Add a [Deploy to Bitrise.io](https://www.bitrise.io/integrations/steps/deploy-to-bitrise-io) Step. This Step uploads the lint report as a [build artifact](https://devcenter.bitrise.io/en/builds/managing-build-files/build-artifacts-online.html).

## bitrise.yml

```yaml
- android-lint@0:
    inputs:
    - variant: $VARIANT
- deploy-to-bitrise-io@2: {}
```
