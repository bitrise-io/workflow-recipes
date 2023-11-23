# (React Native) Install dependencies

## Description
Install dependencies using either yarn or npm.

## Instructions

1. Add either the [Run yarn command](https://www.bitrise.io/integrations/steps/yarn) or the [Run npm command](https://bitrise.io/integrations/steps/npm) Step based on your project setup.
2. Set the **The yarn command to run** or **The npm command with arguments to run** input to `install`.

## bitrise.yml

Using `yarn`:
```yaml
- yarn@0:
    inputs:
    - command: install
```

Using `npm`:
```yaml
- npm@1:
    inputs:
    - command: install
```
