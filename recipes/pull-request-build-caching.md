# Make caching efficient for pull request builds

## Description

Bitrise caching is branch based, which means it's the most efficient when the cache is kept up-to-date on the given branch.

When a branch doesn't have a cache entry yet, the repo's default branch is used as a fallback to pull a cache entry. When a pull request targets the main branch, this fallback mechanism is used to pull the default branch's cache (since the PR branch doesn't have any cache entry yet). This means that caching can still be efficient if the PR destination is the default branch and the cache is up-to-date on the default branch.

## Instructions

1. Make sure that any Workflow that runs frequently on the default branch contains the [Bitrise.io Cache:Push](https://www.bitrise.io/integrations/steps/cache-push) Step. This will keep the cache up to date with content from successful builds. To run this Workflow frequently we recommend two approaches:
    - Run the Workflow after every commit (for example, a merge) to the default branch, triggered by the push event on the default branch.
    - Run the Workflow as a scheduled nighly build every day. This warms up the cache by pushing content based on the latest state of the default branch.

2. Optionally, you can set the **Compress Archive** input to `true` on the Step. This is useful if your cached folders are bigger.

3. Make sure that Workflow which runs for pull requests contains the [Bitrise.io Cache:Pull](https://www.bitrise.io/integrations/steps/cache-pull) Step. This will pull cache from successful builds on the default branch. Note that it's not recommended for PR builds to push content into the cache for security and efficiency reasons. Even if the Workflow contains the Cache Push step, it's skipped by default for PR builds.

## bitrise.yml

Workflow running on the default branch:

```yaml
# Add steps that produce the cached content (e.g. dependecies, builds)

- cache-push@2: {}
```

Pull Request workflow:
```yaml
- cache-pull@2: {}

# Add steps that can utilise the restored cache content
```
