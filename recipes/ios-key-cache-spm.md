# (iOS) Cache Swift Package Manager dependencies (Beta)

## Description

Cache the resolved Swift package dependencies with the new key-based caching Steps, **Save Cache** and **Restore Cache**.

## Instructions

1. Add the [Restore cache](https://github.com/bitrise-steplib/bitrise-step-restore-cache) Step to the Workflow.
1. Add the following keys to the **Cache keys** input:
    ```
    spm-cache-{{ checksum "**/Package.resolved" }}
    spm-cache-
    ```
    The first key will result in a unique string based on the exact dependencies defined in `Package.resolved` within `.xcodeproj` (make sure to commit the file!). If there is no cache to restore with that key, the Step will move on to the second key and will restore a cache with a key that starts with `spm-cache-`. This will result in a cache that was saved for slightly different dependencies, but it's still better than not restoring any cache.
1. Add one of the usual iOS build steps, such as [Xcode Test for iOS](https://www.bitrise.io/integrations/steps/xcode-test).
    - Add `spm-cache-{{ checksum "**/Package.resolved" }}` to the **Cache key** input. The checksum at the end guarantees a new cache archive when dependencies change.
    - Set the **Paths to cache** input to `~/Library/Developer/Xcode/DerivedData/**/SourcePackages`

## bitrise.yml

```yaml
- restore-cache@1:
    inputs:
    - key: |
        spm-cache-{{ checksum "**/Package.resolved" }}
        spm-cache-
- xcode-test@4: {}
- save-cache@3:
    inputs:
    - key: spm-cache-{{ checksum "**/Package.resolved" }}
    - paths: ~/Library/Developer/Xcode/DerivedData/**/SourcePackages
```
