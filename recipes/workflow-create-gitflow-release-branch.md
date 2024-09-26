# Example Workflow: Create Gitflow release branch

## Description
An example Workflow that creates a Gitflow release branch for a specific version. The version can be passed as an Environment Variable for the Workflow.

## Prerequisites

Make sure that Bitrise has write access to your repository. You need to [manually add an SSH key](https://devcenter.bitrise.io/en/apps/configuring-ssh-keys.html#configuring-ssh-keys) with **write** permission on GitHub. 

## bitrise.yml

```yaml
# Run the workflow with $VERSION env set up to, for examepl, '2.4.3'
create-release-branch:
  steps:
  - activate-ssh-key@4:
      run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
  - git-clone@8: {}
  - script@1:
      inputs:
      - content: |-
          #!/usr/bin/env bash
          # fail if any commands fails
          set -e
          # debug log
          set -x

          git checkout -b release-$VERSION
          git push origin release-$VERSION
```
