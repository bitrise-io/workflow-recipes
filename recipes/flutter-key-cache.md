# (Flutter) Cache Dart dependencies (Beta)

## Description

Cache the contents of the [Dart pub system cache](https://dart.dev/tools/pub/glossary#system-cache) folder with the new key-based caching Steps, **Save Dart Cache** and **Restore Dart Cache**.

## Instructions

1. Add the [Restore Dart Cache](https://bitrise.io/integrations/steps/restore-dart-cache) Step to the Workflow.
1. Add one of Flutter Steps to the workflow, such as [Flutter Build](https://www.bitrise.io/integrations/steps/flutter-build)
1. Add the [Save Dart Cache](https://bitrise.io/integrations/steps/save-dart-cache) Step.

### Fine tune cache behaviour

The Dart specific cache steps use optimal cache key and path configurations maintained by Bitrise. If you want full control over what should be cached then please check out the generic [Restore Cache](https://bitrise.io/integrations/steps/restore-cache) and [Save Cache](https://bitrise.io/integrations/steps/save-cache) Steps.

You can always check out what key and path settings the Dart cache step uses:
[Github code snippet](https://github.com/bitrise-steplib/bitrise-step-save-dart-cache/blob/main/step/step.go#L13-L22)

## bitrise.yml

```yaml
- restore-dart-cache@1: {}
- flutter-build@0: {}
- save-dart-cache@1: {}
```
