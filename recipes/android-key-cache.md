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

For example, to setup the cache key value based on your need you simply need to enter this to the **Restore Cache** and **Save Cache** step **Cache keys** input:
```
gradle-cache-{{ checksum "**/*.gradle*" "**/gradle-wrapper.properties" "gradle.properties" }}
gradle-cache-
```
The first key will result in a unique string based on the exact dependencies defined in your Gradle config. If there is no cache to restore with that key, the Step will move on to the second key and will restore a cache with a key that starts with `gradle-cache-`. This will result in a cache that was saved for slightly different dependencies, but it's still better than not restoring any cache.

And if you need to fine tune what gets saved then you need to enter this to the **Save Cache** step **Paths to cache** input:
```
~/.gradle/caches
~/.gradle/wrapper
.gradle/configuration-cache
```

## bitrise.yml

```yaml
- restore-gradle-cache@1: {}
- android-build@1:
    inputs:
    - variant: debug
    - build_type: apk
- save-gradle-cache@1: {}
```
