# (iOS) Cache Swift Package Manager dependencies (Beta)

## Description

Cache the resolved Swift package dependencies with the new key-based caching Steps, **Save Cache** and **Restore Cache**.

## Instructions

1. Add the [Restore SPM Cache](https://bitrise.io/integrations/steps/restore-spm-cache) Step to the Workflow.
1. Add one of the usual iOS build steps, such as [Xcode Test for iOS](https://www.bitrise.io/integrations/steps/xcode-test).
1. Add the [Save SPM Cache](https://bitrise.io/integrations/steps/save-spm-cache) Step.

### Fine tune cache behaviour

The SPM specific cache steps use optimal cache key and path configurations maintained by Bitrise. If you want full control over what should be cached then please check out the generic [Restore Cache](https://bitrise.io/integrations/steps/restore-cache) and [Save Cache](https://bitrise.io/integrations/steps/save-cache) Steps.

You can always check out what key and path settings the SPM cache step uses:
[Github code snippet](https://github.com/bitrise-steplib/bitrise-step-save-spm-cache/blob/main/step/step.go#L13-L26)

## bitrise.yml

```yaml
- restore-spm-cache@1: {}
- xcode-test@5: {}
- save-spm-cache@1: {}
```
