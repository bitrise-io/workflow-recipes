# (iOS) Cache CocoaPods dependencies (Beta)

## Description

Cache the contents of the `Pods` folder with the new key-based caching Steps, **Save Cache** and **Restore Cache**.

## Instructions

1. Add the [Restore Cocoapods Cache](https://bitrise.io/integrations/steps/restore-cocoapods-cache) Step to the Workflow.
1. Add the [Run CocoaPods install](https://bitrise.io/integrations/steps/cocoapods-install) Step.
1. Add the [Save Cocoapods Cache](https://bitrise.io/integrations/steps/save-cocoapods-cache) Step.

### Fine tune cache behaviour

The Cocoapods specific cache steps use optimal cache key and path configurations maintained by Bitrise. If you want full control over what should be cached then please check out the generic [Restore Cache](https://bitrise.io/integrations/steps/restore-cache) and [Save Cache](https://bitrise.io/integrations/steps/save-cache) Steps.

For example, to setup the cache key value based on your need you simply need to enter this to the **Restore Cache** and **Save Cache** step **Cache keys** input:
```
cocoapods-cache-{{ checksum "Podfile.lock" }}
cocoapods-cache-
```
The first key will result in a unique string based on the exact dependencies defined in `Podfile.lock` (make sure to commit the file!). If there is no cache to restore with that key, the Step will move on to the second key and will restore a cache with a key that starts with `cocoapods-cache-`. This will result in a cache that was saved for slightly different dependencies, but it's still better than not restoring any cache.

And if you need to fine tune what gets saved then you need to set the **Paths to cache** input to the `Pods` folder.

## bitrise.yml

```yaml
- restorecocopods-cache@1: {}
- cocoapods-install@2: {}
- save-cocopods-cache@1: {}
```
