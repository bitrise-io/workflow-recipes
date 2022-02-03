# (iOS) Install CocoaPods Dependencies

## Description

Installing CocoaPods dependecies. **Make sure that you are using the workspace and not the project file in your steps. Check the value of `$BITRISE_PROJECT_PATH` env var.**

## Instructions

1. Add the [Run CocoaPods install](https://github.com/bitrise-steplib/steps-cocoapods-install) Step.
2. (Optional) If your Podfile is not in the root, you can set the **Podfile path** input.  

## bitrise.yml

```yaml
- cocoapods-install@2: {}
```
