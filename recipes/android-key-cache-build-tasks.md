# (Android) Cache Gradle build tasks (Beta)

## Description

Cache Gradle tasks with the new key-based caching Steps, **Save Gradle Cache** and **Restore Gradle Cache**.

## Prerequisites

Make sure to read how to [cache Gradle dependencies](android-key-cache.md) and set up the Workflow according to the guide. Caching build tasks is an opt-in feature that builds on caching Gradle dependencies.

## Instructions

[Gradle build cache](https://docs.gradle.org/current/userguide/build_cache.html) is a feature that enables the storage of the task outputs in the shared Gradle cache folder. Caching this folder in CI builds means that Gradle can reuse the task outputs from previous builds and can skip running the tasks when the inputs are unchanged.

This is an opt-in feature. There are two ways to enable the build cache in a Gradle project:

- add `org.gradle.caching = true` to the `gradle.properties` file in the project
- pass the `--build-cache` CLI flag to each Gradle execution

If you choose the second option and use Bitrise Android Steps, there is a Step input for additional Gradle arguments where you can define `--build-cache`.


## bitrise.yml

```yaml
- restore-gradle-cache@1: {}
- android-build@1:
    inputs:
    - variant: debug
    - build_type: apk
    - arguments: --build-cache
- save-gradle-cache@1: {}
```
