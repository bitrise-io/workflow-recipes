# (iOS) Cache CocoaPods dependencies (Beta)

## Description

Cache the contents of the `Pods` folder with the new key-based caching Steps, **Save Cache** and **Restore Cache**.

## Instructions

1. Add the [Restore Cocoapods Cache](https://bitrise.io/integrations/steps/restore-cocoapods-cache) Step to the Workflow.
1. Add the [Run CocoaPods install](https://bitrise.io/integrations/steps/cocoapods-install) Step.
1. Add the [Save Cocoapods Cache](https://bitrise.io/integrations/steps/save-cocoapods-cache) Step.

### Fine tune cache behaviour

The Cocoapods specific cache steps use optimal cache key and path configurations maintained by Bitrise. If you want full control over what should be cached then please check out the generic [Restore Cache](https://bitrise.io/integrations/steps/restore-cache) and [Save Cache](https://bitrise.io/integrations/steps/save-cache) Steps.

You can always check out what key and path settings the Cocapods cache step uses:
[Github code snippet](https://github.com/bitrise-steplib/bitrise-step-save-cocoapods-cache/blob/main/step/step.go#L13-L23)

## bitrise.yml

```yaml
- restorecocopods-cache@1: {}
- cocoapods-install@2: {}
- save-cocopods-cache@1: {}
```
