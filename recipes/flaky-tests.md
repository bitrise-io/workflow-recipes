# Detect and quarantine flaky tests

## Description

Rerun failed tests automatically, so that a test that fails only now and then doesn't fail the build. A test that fails and then passes on a rerun is marked as flaky on the **Tests** tab of the build page, and the build can still go green. Quarantined tests are skipped until you take them out of quarantine.

## Prerequisites

1. The tests run with one of the Steps that support reruns: Xcode Test for iOS, Xcode Test without building, iOS Device Testing, Virtual Device Testing for Android, Android Unit Test or Maestro Test.
1. A Deploy to Bitrise.io Step of version 2.17.0 or newer in the Workflow, to show flaky tests.
1. Test Quarantine needs the Enterprise plan.

## Instructions

1. Turn on reruns in the Step that runs the tests:
    - Xcode Test for iOS: reruns failed tests by default (**Test Repetition Mode**: `retry_on_failure`, **Maximum Test Repetitions**: `3`).
    - Xcode Test without building: set **Test Repetition Mode** to `retry_on_failure`. Its default is `none`.
    - iOS Device Testing and Virtual Device Testing for Android: set **Number of times a test execution is reattempted**, and set **Download files** to `true`. These Steps rerun the whole test suite, not only the failed tests.
    - Maestro Test: set **Number of times a failed flow is reattempted**.
    - Android Unit Test: apply the [Gradle Test Retry plugin](https://github.com/gradle/test-retry-gradle-plugin) in your project.
2. Add a [Deploy to Bitrise.io](https://bitrise.io/integrations/steps/deploy-to-bitrise-io) Step, so that the flaky tests show up on the **Tests** tab. If the pull request comment is turned on with the **Flaky tests** block, they are also posted on the pull request.
3. Optionally, act on the flaky tests found in the build. Xcode Test for iOS, Android Unit Test and the two device testing Steps list them in the `$BITRISE_FLAKY_TEST_CASES` Environment Variable, one per line. For example, send a Slack message.
4. To quarantine a flaky test, click **Add to quarantine** next to it on the **Flaky** tab of the **Tests** tab, and scope it to branches or Workflows if needed. The testing Steps skip quarantined tests without any further setup: they read them from `$BITRISE_QUARANTINED_TESTS_JSON`.

## bitrise.yml

```yaml
- xcode-test@6:
    inputs:
    - test_repetition_mode: retry_on_failure
    - maximum_test_repetitions: 3
- deploy-to-bitrise-io@2: {}
- slack@4:
    run_if: '{{getenv "BITRISE_FLAKY_TEST_CASES" | ne ""}}'
    inputs:
    - webhook_url: $SLACK_WEBHOOK
    - text: |-
        Flaky tests in build $BITRISE_BUILD_NUMBER:
        $BITRISE_FLAKY_TEST_CASES
```

## Relevant Links

* https://docs.bitrise.io/en/bitrise-ci/testing/detecting-and-quarantining-flaky-tests
