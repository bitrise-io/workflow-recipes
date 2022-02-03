# (Flutter) Run Dart Analyzer

## Description
Runs the Dart Analyzer.

## Instructions

1. Add the [Flutter Analyze](https://www.bitrise.io/integrations/steps/flutter-analyze) Step to your Workflow.

## bitrise.yml

```yaml
- flutter-analyze@0:
    inputs:
    - project_location: $BITRISE_FLUTTER_PROJECT_LOCATION
```
