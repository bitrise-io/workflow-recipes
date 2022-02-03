# (React Native) Run tests

## Description

Run tests (for example, Jest).

## Instructions

1. Add either the [Run yarn command](https://www.bitrise.io/integrations/steps/yarn) or the [Run npm command](https://github.com/bitrise-steplib/steps-npm) Step based on your project setup.
2. Set the **The yarn command to run** or **The npm command with arguments to run** input to `test`.

## bitrise.yml

Using `yarn`:

```yaml
- yarn@0:
    inputs:
    - command: test
```

Using `npm`:

```yaml
- npm@1:
    inputs:
    - command: test
```
