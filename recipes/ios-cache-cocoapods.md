# (iOS) Cache CocoaPods dependencies

## Description

Cache the content of the `Pods` folder.

## Instructions

1. Add the [Bitrise.io Cache:Pull](https://www.bitrise.io/integrations/steps/cache-pull) Step.
2. Add the [Run CocoaPods install](https://github.com/bitrise-steplib/steps-cocoapods-install) Step.
3. Add the [Bitrise.io Cache:Push](https://www.bitrise.io/integrations/steps/cache-push) Step.
    - Optionally you can set **Compress Archive** to `true`. This is useful if your cached folders are bigger.

## bitrise.yml

```yaml
- cache-pull@2: {}
- cocoapods-install@2: {}
- cache-push@2: {}
```
