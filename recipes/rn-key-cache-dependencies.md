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

For example, to setup the cache key value based on your need you simply need to enter this to the **Restore Cache** and **Save Cache** step **Cache keys** input:
```
npm-cache-{{ checksum "package-lock.json" }}
npm-cache-
```
The first key will result in a unique string based on the exact dependencies defined in `package-lock.json` (make sure to commit the file!). If there is no cache to restore with that key, the Step will move on to the second key and will restore a cache with a key that starts with `npm-cache-`. This will result in a cache that was saved for slightly different dependencies, but it's still better than not restoring any cache.

And if you need to fine tune what gets saved then you need to set the **Paths to cache** input to the `node_modules` folder.

## bitrise.yml

```yaml
- restore-npm-cache@1: {}
- npm@1:
    inputs:
    - command: install
- save-npm-cache@1: {}
```
