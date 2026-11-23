# (Android, iOS) Run Maestro UI tests

## Description

Run [Maestro](https://maestro.dev) flows on an Android emulator or an iOS simulator on the build machine. Each flow shows up as a test case on the **Tests** tab of the build page, with its screenshots, recordings and logs attached.

## Prerequisites

1. Maestro flows in your repository. By default, the Step runs the flows in the `.maestro` folder.
1. For Android, a Linux stack: Android emulators don't run on macOS stacks.
1. For iOS, a macOS stack.

## Instructions

1. Build the app:
    - Android: add an [Android Build](https://bitrise.io/integrations/steps/android-build) Step and set the **Module** and **Variant** inputs. An emulator can't install an `.aab`, so build an `.apk`, or convert the bundle with the [Export Universal APK](https://bitrise.io/integrations/steps/bitrise-step-export-universal-apk) Step.
    - iOS: add an [Xcode Build for Simulator](https://bitrise.io/integrations/steps/xcode-build-for-simulator) Step. It builds the project and scheme set in `$BITRISE_PROJECT_PATH` and `$BITRISE_SCHEME`.
1. Start exactly one device. The Maestro Test Step fails if there is none or more than one.
    - Android: add an [AVD Manager](https://bitrise.io/integrations/steps/avd-manager) Step and set **Android API Level**, **OS Tag** and **ABI**. The ABI must match the CPU architecture of the stack. The Step waits for the emulator to boot.
    - iOS: add an [Xcode Start Simulator](https://bitrise.io/integrations/steps/xcode-start-simulator) Step and set **Simulator boot timeout** so that it waits for the simulator to boot.
1. Add a [Maestro Test](https://bitrise.io/integrations/steps/maestro-test) Step and set **App path** to `$BITRISE_APK_PATH` (Android) or `$BITRISE_APP_DIR_PATH` (iOS). Optional inputs:
    - **Include tags** and **Exclude tags**: run only the flows that have any of the listed tags, or skip them.
    - **Number of times a failed flow is reattempted**: retry failed flows. A flow that passes on a retry is reported as flaky and doesn't fail the Step. See [Detect and quarantine flaky tests](./flaky-tests.md).
    - Quarantined flows are skipped without any setup (Enterprise plan).
    - Environment Variables and Secrets starting with `MAESTRO_` are passed to the flows, for example `${MAESTRO_PASSWORD}`.
1. Add a [Deploy to Bitrise.io](https://bitrise.io/integrations/steps/deploy-to-bitrise-io) Step. It makes the results available on the **Tests** tab.

To spread the flows across several machines, run parallel copies of the Workflow in a Pipeline, tag the flows `shard0`, `shard1` and so on, and set **Include tags** to `shard$BITRISE_IO_PARALLEL_INDEX`. See [Running flows on several devices in parallel](https://docs.bitrise.io/en/bitrise-ci/testing/running-maestro-tests-on-bitrise#running-flows-on-several-devices-in-parallel), and [UI tests in parallel shards](./android-parallel-ui-test-shards.md) for a Pipeline with parallel copies of a Workflow.

## bitrise.yml

Android:

```yaml
- android-build@1:
    inputs:
    - module: app
    - variant: debug
- avd-manager@2:
    inputs:
    - api_level: 34
    - tag: google_apis
    - abi: x86_64
- maestro-test@1:
    inputs:
    - app_path: $BITRISE_APK_PATH
    - num_flaky_test_attempts: 2
- deploy-to-bitrise-io@2: {}
```

iOS:

```yaml
- xcode-build-for-simulator@3: {}
- xcode-start-simulator@0:
    inputs:
    - destination: platform=iOS Simulator,name=Bitrise iOS default,OS=latest
    - wait_for_boot_timeout: 120
- maestro-test@1:
    inputs:
    - app_path: $BITRISE_APP_DIR_PATH
    - num_flaky_test_attempts: 2
- deploy-to-bitrise-io@2: {}
```

## Relevant Links

* https://docs.bitrise.io/en/bitrise-ci/testing/running-maestro-tests-on-bitrise
