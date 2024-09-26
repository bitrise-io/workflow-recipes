# (Android) Cache Gradle dependencies (Beta)

## Description

Cache project dependencies that Gradle downloads with the new key-based caching Steps, **Save Gradle Cache** and **Restore Gradle Cache**.

If you want to cache not only the dependencies but build tasks as well, check out [this workflow recipe](android-key-cache-build-tasks.md).

## Instructions

1. Add the [Restore Gradle Cache](https://bitrise.io/integrations/steps/restore-gradle-cache) Step to the Workflow.
1. Add the usual Android Steps, such as [Android Build](https://bitrise.io/integrations/steps/android-build).
1. Add the [Save Gradle Cache](https://bitrise.io/integrations/steps/save-gradle-cache) Step.
1. If you want to cache not only the dependencies but build tasks  as well, check out [this workflow recipe](android-key-cache-build-tasks.md).

### Fine tune cache behaviour

The Gradle specific cache steps use optimal cache key and path configurations maintained by Bitrise. If you want full control over what should be cached then please check out the generic [Restore Cache](https://bitrise.io/integrations/steps/restore-cache) and [Save Cache](https://bitrise.io/integrations/steps/save-cache) Steps.

You can always check out what key and path settings the Gradle cache step uses:
[Github code snippet](https://github.com/bitrise-steplib/bitrise-step-save-gradle-cache/blob/main/step/step.go#L14-L53)

## bitrise.yml

```yaml
- restore-gradle-cache@2: {}
- android-build@1:
    inputs:
    - variant: debug
    - build_type: apk
- save-gradle-cache@1: {}
```
