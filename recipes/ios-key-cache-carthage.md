# (iOS) Cache Carthage dependencies (Beta)

## Description

Cache the contents of the `Carthage` folder with the new key-based caching Steps, **Save Cache** and **Restore Cache**.

## Instructions

1. Add the [Restore Carthage Cache](https://bitrise.io/integrations/steps/restore-carthage-cache) Step to the Workflow.
1. Add the [Carthage](https://bitrise.io/integrations/steps/carthage) Step.
1. Add the [Save Carthage Cache](https://bitrise.io/integrations/steps/save-carthage-cache) Step.

### Fine tune cache behaviour

The Carthage specific cache steps use optimal cache key and path configurations maintained by Bitrise. If you want full control over what should be cached then please check out the generic [Restore Cache](https://bitrise.io/integrations/steps/restore-cache) and [Save Cache](https://bitrise.io/integrations/steps/save-cache) Steps.

You can always check out what key and path settings the Carthage cache step uses:
[Github code snippet](https://github.com/bitrise-steplib/bitrise-step-save-carthage-cache/blob/main/step/step.go#L14-L34)

## bitrise.yml

```yaml
- restore-carthage-cache@2: {}
- carthage@3: {}
- save-carthage-cache@1: {}
```
