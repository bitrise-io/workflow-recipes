# (iOS) Cache CocoaPods dependencies (Beta)

## Description

Cache the contents of the `Pods` folder with the new key-based caching Steps, **Save Cache** and **Restore Cache**.

## Instructions

1. Add the [Restore cache](https://github.com/bitrise-steplib/bitrise-step-restore-cache) Step to the Workflow.
1. Add the following keys to the **Cache keys** input:
    ```
    cocoapods-cache-{{ checksum "Podfile.lock" }}
    cocoapods-cache-
    ```
    The first key will result in a unique string based on the exact dependencies defined in `Podfile.lock` (make sure to commit the file!). If there is no cache to restore with that key, the Step will move on to the second key and will restore a cache with a key that starts with `cocoapods-cache-`. This will result in a cache that was saved for slightly different dependencies, but it's still better than not restoring any cache.
1. Add the [Run CocoaPods install](https://github.com/bitrise-steplib/steps-cocoapods-install) Step.
1. Add the [Save cache](https://github.com/bitrise-steplib/bitrise-step-restore-cache) Step.
    - Add `cocoapods-cache-{{ checksum "Podfile.lock" }}` to the **Cache key** input. The checksum at the end guarantees a new cache archive when dependencies change.
    - Set the **Paths to cache** input to `Pods` (or adjust it if your project has a different folder structure)

## bitrise.yml

```yaml
- restore-cache@1:
    inputs:
    - key: |
        cocoapods-cache-{{ checksum "Podfile.lock" }}
        cocoapods-cache-
- cocoapods-install@2: {}
- save-cache@1:
    inputs:
    - key: cocoapods-cache-{{ checksum "Podfile.lock" }}
    - paths: Pods
```
