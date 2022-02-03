# Shallow clone the git repository

## Description

Speed up cloning the repository by specifying a clone depth. Keep in mind that in case of Pull Request you need to make sure that both the source and destination branch are available in the history.

## Instructions

1. Set the **Limit fetching to the specified number of commits** input variable:
    - In the case of builds running on a single branch you can set it to `1`.
    - In the case of Pull Request builds set it to a depth that will contain both the source and target branches. For example, `100`. 

## bitrise.yml

```yaml
- git-clone@6:
    inputs:
    - clone_depth: 1
```

## Relevant Links

* https://discuss.bitrise.io/t/git-clone-step-is-very-slow/3844/2
* https://stackoverflow.com/questions/6941889/is-it-safe-to-shallow-clone-with-depth-1-create-commits-and-pull-updates-aga
