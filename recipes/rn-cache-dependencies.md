# (React Native) Cache dependencies (node_modules)

## Description

Set up caching for dependecies (`node_modules` folder) of a React Native app.

## Instructions

1. Add the [Bitrise.io Cache:Pull](https://www.bitrise.io/integrations/steps/cache-pull) Step.
2. Add either the [Run yarn command](https://www.bitrise.io/integrations/steps/yarn) or the [Run npm command](https://github.com/bitrise-steplib/steps-npm) Step based on your project setup. Set the input variables:
    - Set the **The yarn command to run** or **The npm command with arguments to run** input to `install`.
    - Set **Cache node_modules** to `yes`.
3. Add the [Bitrise.io Cache:Push](https://www.bitrise.io/integrations/steps/cache-push) step.
    - Optionally you can set the **Compress Archive** input to `true`. This is useful if your cache folders are bigger.

## bitrise.yml

```yaml
- cache-pull@2: {}
- yarn@0:
    inputs:
    - cache_local_deps: 'yes'
    - command: install
- cache-push@2:
    inputs:
    - compress_archive: 'true'
```
