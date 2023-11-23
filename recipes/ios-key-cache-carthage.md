# (iOS) Cache Carthage dependencies (Beta)

## Description

Cache the contents of the `Carthage` folder with the new key-based caching Steps, **Save Cache** and **Restore Cache**.

## Instructions

1. Add the [Restore Carthage Cache](https://bitrise.io/integrations/steps/restore-carthage-cache) Step to the Workflow.
1. Add the [Carthage](https://bitrise.io/integrations/steps/carthage) Step.
1. Add the [Save Carthage Cache](https://bitrise.io/integrations/steps/save-carthage-cache) Step.

### Fine tune cache behaviour

The Carthage specific cache steps use optimal cache key and path configurations maintained by Bitrise. If you want full control over what should be cached then please check out the generic [Restore Cache](https://bitrise.io/integrations/steps/restore-cache) and [Save Cache](https://bitrise.io/integrations/steps/save-cache) Steps.

For example, to setup the cache key value based on your need you simply need to enter this to the **Restore Cache** and **Save Cache** step **Cache keys** input:
```
carthage-cache-{{ checksum "Cartfile.resolved" }}
carthage-cache-
```
The first key will result in a unique string based on the exact dependencies defined in `Cartfile.resolved` (make sure to commit the file!). If there is no cache to restore with that key, the Step will move on to the second key and will restore a cache with a key that starts with `carthage-cache-`. This will result in a cache that was saved for slightly different dependencies, but it's still better than not restoring any cache.

And if you need to fine tune what gets saved then you need to set the **Paths to cache** input to the `Carthage` folder within the project.

## bitrise.yml

```yaml
- restore-carthage-cache@1: {}
- carthage@1: {}
- save-carthage-cache@1: {}
```
