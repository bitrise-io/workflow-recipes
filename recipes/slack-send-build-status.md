# (iOS/Android) Slack - send build status

## Description
Sending a slack massage to Slack with the build status after a build has finished.

## Prerequisites

2. You have a Slack webhook set up and added to Env Vars (for example, `$SLACK_WEBHOOK`). For details, see [Configuring Slack integration](https://devcenter.bitrise.io/en/builds/configuring-build-settings/configuring-slack-integration.html)

## Instructions

1. Add the [Send a Slack message](https://www.bitrise.io/integrations/steps/slack) step. Set the input variables:
    - **Slack Webhook URL**: for example, `$SLACK_WEBHOOK`.
    - **Target Slack channel, group or username**: for example, `#build-notifications`.
    - Check out the other optional input variables in the Workflow Editor or in the Step description.

## bitrise.yml

```yaml
- slack@3:
    inputs:
    - channel: "#build-notifications"
    - webhook_url: $SLACK_WEBHOOK
```

## Relevant Links

* https://devcenter.bitrise.io/en/builds/configuring-build-settings/configuring-slack-integration.html
