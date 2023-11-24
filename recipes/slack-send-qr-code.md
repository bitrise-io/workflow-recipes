# (iOS/Android) Send QR code to Slack

## Description
Sending a QR code of the iOS or Android build uploaded to bitrise.io to Slack.

## Prerequisites

1. You have your iOS or Android app archived.
2. You have Slack webhook set up and added to Env Vars (for example, `$SLACK_WEBHOOK`). See [Configuring Slack integration](https://devcenter.bitrise.io/en/builds/configuring-build-settings/configuring-slack-integration.html)

## Instructions

1. Add the [Deploy to Bitrise.io - Apps, Logs, Artifacts](https://www.bitrise.io/integrations/steps/deploy-to-bitrise-io) step.
2. Add the [Create install page QR code](https://www.bitrise.io/integrations/steps/create-install-page-qr-code) step.
3. Add the [Send a Slack message](https://www.bitrise.io/integrations/steps/slack) step. Set the input variables:
    - **Slack Webhook URL**: for example, `$SLACK_WEBHOOK`.
    - **Target Slack channel, group or username**: for example, `#build-notifications`.
    - **A URL to an image file that will be displayed as a thumbnail**: `$BITRISE_PUBLIC_INSTALL_PAGE_QR_CODE_IMAGE_URL`.

## bitrise.yml

```yaml
- deploy-to-bitrise-io@2: {}
- create-install-page-qr-code@1: {}
- slack@4:
    inputs:
    - channel: "#build-notifications"
    - thumb_url: $BITRISE_PUBLIC_INSTALL_PAGE_QR_CODE_IMAGE_URL
    - webhook_url: $SLACK_WEBHOOK
```

## Relevant Links

* https://devcenter.bitrise.io/en/deploying/ios-deployment/deploying-an-ios-app-to-bitrise-io.html
* https://devcenter.bitrise.io/en/deploying/android-deployment/deploying-android-apps-to-bitrise-and-google-play.html#deploying-an-android-app-to-bitrise-io-43303
* https://devcenter.bitrise.io/en/builds/configuring-build-settings/configuring-slack-integration.html
