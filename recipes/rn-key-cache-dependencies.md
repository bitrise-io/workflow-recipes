# (React Native) Cache NPM dependencies (Beta)

## Description

Cache the contents of the `node_modules` folder with the new key-based caching Steps, **Save Cache** and **Restore Cache**.

## Instructions

1. Add the [Restore cache](https://github.com/bitrise-steplib/bitrise-step-restore-cache) Step to the Workflow.
1. Add the following keys to the **Cache keys** input:
    ```
    npm-cache-{{ checksum "package-lock.json" }}
    npm-cache-
    ```
    The first key will result in a unique string based on the exact dependencies defined in `package-lock.json` (make sure to commit the file!). If there is no cache to restore with that key, the Step will move on to the second key and will restore a cache with a key that starts with `npm-cache-`. This will result in a cache that was saved for slightly different dependencies, but it's still better than not restoring any cache.
1. Add either the [Run yarn command](https://www.bitrise.io/integrations/steps/yarn) Step or the [Run npm command](https://github.com/bitrise-steplib/steps-npm) Step based on your project setup. Set the input variables:
    - Set the **The yarn command to run** or **The npm command with arguments to run** input to `install`.
1. Add the [Save cache](https://github.com/bitrise-steplib/bitrise-step-restore-cache) Step.
    - Add `npm-cache-{{ checksum "package-lock.json" }}` to the **Cache key** input. The checksum at the end guarantees a new cache archive when dependencies change.
    - Set the **Paths to cache** input to `node_modules` (or adjust it if your project has a different folder structure)

## bitrise.yml

```yaml
- restore-cache@1:
    inputs:
    - key: |
        npm-cache-{{ checksum "package-lock.json" }}
        npm-cache-
- npm@1:
    inputs:
    - command: install
- save-cache@1:
    inputs:
    - key: npm-cache-{{ checksum "package-lock.json" }}
    - paths: node_modules/
```
