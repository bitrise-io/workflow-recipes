# Testing on Multiple Destinations

## Description

Sometimes you will want to use a single test command to run on multiple destinations. For example, you may want to run your tests on a simulator with iOS 15 and a simulator with iOS 14.

You'll primarily use our Xcode Test for iOS step for this. The step comes with out of the box support for a single destination. We'll add a bit of scripting to add another destination (or destinations) to the test run.

## Prerequisites

1. The source code is cloned and the dependencies (for example, Cocoapods, Carthage) are installed.
1. Your project is set up with a test target
1. You use `xcodebuild` for log formatting because `xcpretty` does not support parallel test execution.

## Instructions

1. Add a script step to your workflow. This will be used to get the UDID of the additional simulator on which your tests will run. An example for iOS 15.2 and the iPhone 11 Pro is below. Use this script as a starting point for modification to use the target you desire. Customize by setting a different OS runtime and name. You can run `xcrun simctl list` to get the full list of options.
1. Add an [Xcode Test for iOS](https://github.com/bitrise-steplib/steps-xcode-test) Step. Override any of the following inputs if needed:
    - **Project path**: The default value is `$BITRISE_PROJECT_PATH` and in most cases you don't have to change it.
    - **Scheme**: The default value is `$BITRISE_SCHEME`, this variable stores the scheme that you set when adding the app on Bitrise. You can specify a different scheme if you want but it must be a shared scheme.
    - **Device destination specifier** (default: `platform=iOS Simulator,name=iPhone 8 Plus,OS=latest`). This is our first destination.
    - **Additional options for `xcodebuild build test` call** will set the additional destination(s)
1. Add a [Deploy to Bitrise.io - Apps, Logs, Artifacts](https://www.bitrise.io/integrations/steps/deploy-to-bitrise-io) Step that makes the test results available in the [Test Reports add-on](https://devcenter.bitrise.io/en/testing/test-reports.html).


## bitrise.yml

```yaml
    - script@1:
        title: Get iPhone 11 and iOS 15.2 udid
        inputs:
        - content: |-
            #!/usr/bin/env bash
            set -x
            envman add --key IPHONE_11_UDID --value "$(xcrun simctl list --json | jq -r '.devices["com.apple.CoreSimulator.SimRuntime.iOS-15-2"] | .[] |  select(.name == "iPhone 11 Pro") | .udid')"
    - xcode-test@4.2:
        inputs:
        - project_path: "$BITRISE_PROJECT_PATH"
        - xcodebuild_test_options: -destination 'id="$IPHONE_11_UDID"'
        - log_formatter: xcodebuild
        - scheme: "$BITRISE_SCHEME"
    - deploy-to-bitrise-io@1: {}
```
