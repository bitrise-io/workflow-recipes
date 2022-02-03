# GitHub pull request - send the build QR code

## Description
Send a comment to the GitHub pull request with a QR code to the build uploaded to Bitrise.io.

## Prerequisites

1. You have your iOS or Android app archived.
2. Generate a [GitHub personal access token](https://github.com/settings/tokens) and add it as a Secret (`$GITHUB_ACCESS_TOKEN`). Make sure to select the `repo` scope.

## Instructions

1. Add the [Deploy to Bitrise.io - Apps, Logs, Artifacts](https://www.bitrise.io/integrations/steps/deploy-to-bitrise-io) Step.
2. Add the [Create install page QR code](https://www.bitrise.io/integrations/steps/create-install-page-qr-code) Step.
3. Add the [Comment on GitHub Pull Request](https://www.bitrise.io/integrations/steps/comment-on-github-pull-request) Step. Set the following input variables:
    - **GitHub personal access token**: Set it to the previously created Secret, `$GITHUB_ACCESS_TOKEN`.
    - **Body**:

```
![QR code]($BITRISE_PUBLIC_INSTALL_PAGE_QR_CODE_IMAGE_URL)

$BITRISE_PUBLIC_INSTALL_PAGE_URL
```

## bitrise.yml

```yaml
- deploy-to-bitrise-io@2: {}
- create-install-page-qr-code@1: {}
- comment-on-github-pull-request@0:
    inputs:
    - body: |-
        ![QR code]($BITRISE_PUBLIC_INSTALL_PAGE_QR_CODE_IMAGE_URL)

        $BITRISE_PUBLIC_INSTALL_PAGE_URL
    - personal_access_token: "$GITHUB_ACCESS_TOKEN"
```
