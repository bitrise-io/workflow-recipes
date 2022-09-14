# (Android) Cache Gradle dependencies (Beta)

## Description

Cache project dependencies that Gradle downloads with the new key-based caching Steps, **Save Cache** and **Restore Cache**.

If you want to cache not only the dependencies but build tasks as well, check out [this workflow recipe](android-key-cache-build-tasks.md).

## Instructions

1. Add the [Restore cache](https://github.com/bitrise-steplib/bitrise-step-restore-cache) Step to the Workflow.
1. Add the following keys to the **Cache keys** input:
    ```
    gradle-cache-{{ checksum "**/*.gradle*" "**/gradle-wrapper.properties" "gradle.properties" }}
    gradle-cache-
    ```
    The first key will result in a unique string based on the exact dependencies defined in your Gradle config. If there is no cache to restore with that key, the Step will move on to the second key and will restore a cache with a key that starts with `gradle-cache-`. This will result in a cache that was saved for slightly different dependencies, but it's still better than not restoring any cache.
1. Add the usual Android Steps, such as [Android Build](https://github.com/bitrise-steplib/bitrise-step-android-build).
1. Add the [Save cache](https://github.com/bitrise-steplib/bitrise-step-restore-cache) Step.
    - Add `gradle-cache-{{ checksum "**/*.gradle*" "**/gradle-wrapper.properties" "gradle.properties" }}` to the **Cache key** input. The checksum at the end guarantees a new cache archive when dependencies change.
    - Set the **Paths to cache** input to the following:

    ```
    ~/.gradle/caches
    ~/.gradle/wrapper
    .gradle/configuration-cache
    ```
1. If you want to cache not only the dependencies but build tasks  as well, check out [this workflow recipe](android-key-cache-build-tasks.md).


## bitrise.yml

```yaml
- restore-cache@1:
    inputs:
    - key: |
        gradle-cache-{{ checksum "**/*.gradle*" "**/gradle-wrapper.properties" "gradle.properties" }}
        gradle-cache-
- android-build@1:
    inputs:
    - variant: debug
    - build_type: apk
- save-cache@1:
    inputs:
    - key: gradle-cache-{{ checksum "**/*.gradle*" "**/gradle-wrapper.properties" "gradle.properties" }}
    - paths: |-
        ~/.gradle/caches
        ~/.gradle/wrapper
        .gradle/configuration-cache
```
