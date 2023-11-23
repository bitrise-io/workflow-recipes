# (iOS) Cache Swift Package Manager dependencies (Beta)

## Description

Cache the resolved Swift package dependencies with the new key-based caching Steps, **Save Cache** and **Restore Cache**.

## Instructions

1. Add the [Restore SPM Cache](https://bitrise.io/integrations/steps/restore-spm-cache) Step to the Workflow.
1. Add one of the usual iOS build steps, such as [Xcode Test for iOS](https://www.bitrise.io/integrations/steps/xcode-test).
1. Add the [Save SPM Cache](https://bitrise.io/integrations/steps/save-spm-cache) Step.

### Fine tune cache behaviour

The SPM specific cache steps use optimal cache key and path configurations maintained by Bitrise. If you want full control over what should be cached then please check out the generic [Restore Cache](https://bitrise.io/integrations/steps/restore-cache) and [Save Cache](https://bitrise.io/integrations/steps/save-cache) Steps.

For example, to setup the cache key value based on your need you simply need to enter this to the **Restore Cache** and **Save Cache** step **Cache keys** input:
```
spm-cache-{{ checksum "**/Package.resolved" }}
spm-cache-
```
The first key will result in a unique string based on the exact dependencies defined in `Package.resolved` within `.xcodeproj` (make sure to commit the file!). If there is no cache to restore with that key, the Step will move on to the second key and will restore a cache with a key that starts with `spm-cache-`. This will result in a cache that was saved for slightly different dependencies, but it's still better than not restoring any cache.

And if you need to fine tune what gets saved then you need to set the **Paths to cache** input to `~/Library/Developer/Xcode/DerivedData/**/SourcePackages`.

## bitrise.yml

```yaml
- restore-spm-cache@1: {}
- xcode-test@5: {}
- save-spm-cache@1: {}
```
