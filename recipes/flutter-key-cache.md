# (Flutter) Cache Dart dependencies (Beta)

## Description

Cache the contents of the [Dart pub system cache](https://dart.dev/tools/pub/glossary#system-cache) folder with the new key-based caching Steps, **Save Dart Cache** and **Restore Dart Cache**.

## Instructions

1. Add the [Restore Dart Cache](https://bitrise.io/integrations/steps/restore-dart-cache) Step to the Workflow.
1. Add one of Flutter Steps to the workflow, such as [Flutter Build](https://www.bitrise.io/integrations/steps/flutter-build)
1. Add the [Save Dart Cache](https://bitrise.io/integrations/steps/save-dart-cache) Step.

### Fine tune cache behaviour

The Dart specific cache steps use optimal cache key and path configurations maintained by Bitrise. If you want full control over what should be cached then please check out the generic [Restore Cache](https://bitrise.io/integrations/steps/restore-cache) and [Save Cache](https://bitrise.io/integrations/steps/save-cache) Steps.

For example, to setup the cache key value based on your need you simply need to enter this to the **Restore Cache** and **Save Cache** step **Cache keys** input:
```
dart-cache-{{ checksum "pubspec.lock" }}
dart-cache-
```
The first key will result in a unique string based on the exact dependencies defined in `pubspec.lock` (make sure to commit the file!). If there is no cache to restore with that key, the Step will move on to the second key and will restore a cache with a key that starts with `dart-cache-`. This will result in a cache that was saved for slightly different dependencies, but it's still better than not restoring any cache.

And if you need to fine tune what gets saved then you need to enter this to the **Save Cache** step **Paths to cache** input:
```
~/.pub-cache
```

## bitrise.yml

```yaml
- restore-dart-cache@1: {}
- flutter-build@0: {}
- save-dart-cache@1: {}
```
