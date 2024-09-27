# (React Native) Cache NPM dependencies (Beta)

## Description

Cache the contents of the `node_modules` folder with the new key-based caching Steps, **Save Cache** and **Restore Cache**.

## Instructions

1. Add the [Restore NPM Cache](https://github.com/bitrise-steplib/bitrise-step-restore-npm-cache) Step to the Workflow.
1. Add either the [Run yarn command](https://www.bitrise.io/integrations/steps/yarn) Step or the [Run npm command](https://github.com/bitrise-steplib/steps-npm) Step based on your project setup. Set the input variables:
    - Set the **The yarn command to run** or **The npm command with arguments to run** input to `install`.
1. Add the [Save NPM Cache](https://github.com/bitrise-steplib/bitrise-step-save-npm-cache) Step.

### Fine tune cache behaviour

The NPM specific cache steps use optimal cache key and path configurations maintained by Bitrise. If you want full control over what should be cached then please check out the generic [Restore Cache](https://github.com/bitrise-steplib/bitrise-step-restore-cache) and [Save Cache](https://github.com/bitrise-steplib/bitrise-step-save-cache) Steps.

You can always check out what key and path settings the NPM cache step uses:
[Github code snippet](https://github.com/bitrise-steplib/bitrise-step-save-npm-cache/blob/main/step/step.go#L13-L25)

## bitrise.yml

```yaml
- restore-npm-cache@2: {}
- npm@1:
    inputs:
    - command: install
- save-npm-cache@1: {}
```
