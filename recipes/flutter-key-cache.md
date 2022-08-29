# (Flutter) Cache Dart dependencies (Beta)

## Description

Cache the contents of the [Dart pub system cache](https://dart.dev/tools/pub/glossary#system-cache) folder with the new key-based caching steps, **Save Cache** and **Restore Cache**.

## Instructions

1. Add the [Restore cache](https://github.com/bitrise-steplib/bitrise-step-restore-cache) Step to the workflow.
1. For the **Cache keys** input, add the following keys:
    ```
    dart-cache-{{ checksum "pubspec.lock" }}
    dart-cache-
    ```
    The first key will result in a unique string based on the exact dependencies defined in `pubspec.lock` (make sure to commit the file!). If there is no cache to restore with that key, the Step will use the second key and will restore a cache with a key that starts with `dart-cache-`. This will result in a cache that was saved for slightly different dependencies, but it's still better than not restoring any cache.
1. Add one of Flutter Steps to the workflow, such as [Flutter Build](https://www.bitrise.io/integrations/steps/flutter-build)
1. Add the [Save cache](https://github.com/bitrise-steplib/bitrise-step-restore-cache) Step.
    - For its **Cache key** input, set `dart-cache-{{ checksum "pubspec.lock" }}`. The checksum at the end guarantees a new cache archive when dependencies change.
    - Set the **Paths to cache** input to the following (or adjust it if your project has a different folder structure):
    ```
    ~/.pub-cache
    .dart_tool
    ```

## bitrise.yml

```yaml
- restore-cache@1:
    inputs:
    - key: |-
        dart-cache-{{ checksum "pubspec.lock" }}
        dart-cache-
- cocoapods-install@2: {}
- save-cache@1:
    inputs:
    - key: dart-cache-{{ checksum "pubspec.lock" }}
    - paths: |-
        ~/.pub-cache
        .dart_tool
```
