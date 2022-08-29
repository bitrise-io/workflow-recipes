# (iOS) Cache Carthage dependencies (Beta)

## Description

Cache the contents of the `Carthage` folder with the new key-based caching steps, **Save Cache** and **Restore Cache**.

## Instructions

1. Add the [Restore cache](https://github.com/bitrise-steplib/bitrise-step-restore-cache) Step to the workflow.
1. For the **Cache keys** input, add the following keys:
    ```
    carthage-cache-{{ checksum "Cartfile.resolved" }}
    carthage-cache-
    ```
    The first key will result in a unique string based on the exact dependencies defined in `Cartfile.resolved` (make sure to commit the file!). If there is no cache to restore with that key, the Step will use the second key and will restore a cache with a key that starts with `carthage-cache-`. This will result in a cache that was saved for slightly different dependencies, but it's still better than not restoring any cache.
1. Add the [Carthage](https://github.com/bitrise-steplib/steps-carthage) Step.
1. Add the [Save cache](https://github.com/bitrise-steplib/bitrise-step-restore-cache) Step.
    - For its **Cache key** input, set `carthage-cache-{{ checksum "Cartfile.resolved" }}`. The checksum at the end guarantees a new cache archive when dependencies change.
    - Set the **Paths to cache** input to the `Carthage` folder within the project (or adjust it if your project has a different folder structure)

## bitrise.yml

```yaml
- restore-cache@1:
    inputs:
    - key: |
        carthage-cache-{{ checksum "Cartfile.resolved" }}
        carthage-cache-
- carthage@1: {}
- save-cache@3:
    inputs:
    - key: carthage-cache-{{ checksum "Cartfile.resolved" }}
    - paths: Carthage
```
