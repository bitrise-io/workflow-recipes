# Advanced key-based cache recipes

These workflow recipes are based on the **Save cache** and **Restore cache** Steps.

For recipes about the most popular platforms and dependency managers, check out the **Key-based caching** section in the [README](../README.md).

## Key templates

The **Save cache** and **Restore cache** Steps use a string key when uploading and downloading a cache archive. To always download the most relevant cache archive for each build, the **Cache key** input can contain template elements. The Steps evaluate the key template at runtime and the final key value can change based on the build environment or files in the repo.

The following variables are supported in the **Cache key** input:

- `cache-key-{{ .Branch }}`: Current git branch the build runs on
- `cache-key-{{ .CommitHash }}`: SHA-256 hash of the git commit the build runs on
- `cache-key-{{ .Workflow }}`: Current Bitrise workflow name (eg. `primary`)
- `{{ .Arch }}-cache-key`: Current CPU architecture (`amd64` or `arm64`)
- `{{ .OS }}-cache-key`: Current operating system (`linux` or `darwin`)

Functions available in a template:

`checksum`: This function takes one or more file paths and computes the SHA256 [checksum](https://en.wikipedia.org/wiki/Checksum) of the file contents. This is useful for creating unique cache keys based on files that describe content to cache.

Examples of using `checksum`:
- `cache-key-{{ checksum "package-lock.json" }}`
- `cache-key-{{ checksum "**/Package.resolved" }}`
- `cache-key-{{ checksum "**/*.gradle*" "gradle.properties" }}`

`getenv`: This function returns the value of an environment variable or an empty string if the variable is not defined.

Examples of `getenv`:
- `cache-key-{{ getenv "PR" }}`
- `cache-key-{{ getenv "BITRISEIO_PIPELINE_ID" }}`

## Skip saving the cache in PR builds (restore only)

If you want builds triggered by pull requests to only restore the cache and skip saving it, you can run the **Save cache** Step conditionally:

```yaml
steps:
- restore-cache@1:
    inputs:
    - key: node-modules-{{ checksum "package-lock.json" }}

# Build steps

- save-cache@1:
    run_if: ".IsCI | and (not .IsPR)" # Condition that is false in PR builds
    inputs:
    - key: node-modules-{{ checksum "package-lock.json" }}
    - paths: node_modules
```

## Separate caches for each OS and architecture

Cache is not guaranteed to work across different Bitrise Stacks (different OS or same OS but different CPU architecture). If a Workflow runs on different stacks, it's a good idea to include the OS and architecture in the **Cache key** input:

```yaml
steps:
- save-cache@1:
    inputs:
    - key: '{{ .OS }}-{{ .Arch }}-npm-cache-{{ checksum "package-lock.json" }}'
    - path: node_modules
```

## Multiple independent caches

You can add multiple instances of the cache steps to a Workflow:

```yaml
steps:
- save-cache@1:
    title: Save NPM cache
    inputs:
    - paths: node_modules
    - key: npm-cache-{{ checksum "package-lock.json" }}
- save-cache@1:
    title: Save Python cache
    inputs:
    - paths: venv/
    - key: pip-cache-{{ checksum "requirements.txt" }}
```

## Cache warm-up for pull requests

Caching works best when the cached content is up to date and contains useful data for dependency managers and build systems. It's a good idea to run a Workflow periodically that builds the project from the latest code on the main branch and saves the result in the cache. This way, other builds triggered by pull requests can restore an up-to-date cache.

By including a checksum in the **Cache key** input, the **Save cache** Step will save multiple unique cache archives when the project files change (instead of overriding the previous cache). This way PRs not targeting the latest state of the main branch can still download a relevant cache archive.

```yaml
workflows:
    pr-validation:
        steps:
        - restore-cache@1:
            inputs:
            - key: |-
                node-modules-{{ checksum "package-lock.json" }}
                node-modules-
        # Rest of the PR validation workflow
    
    cache-warm-up:
        description: This Workflow should either be run on a scheduled basis or triggered by a push event on the main branch.
        steps:
        # Build steps
        - save-cache@1:
            inputs:
            - key: node-modules-{{ checksum "package-lock.json" }}
            - paths: node_modules/
```

## Restore cache from the PR target branch

If you have a setup where the cache key is based on the current workflow (such as `cache-{{ .Workflow }}`), then you can configure the Restore Step with the following keys:

```
cache-{{ .Workflow }}
cache-{{ getenv "BITRISEIO_GIT_BRANCH_DEST" }}
cache-
```

The keys listed in the step input are processed in priority order. If there is a cache entry for the exact same branch, the first rule will match that. You can also compute the cache key of the pull requests's target branch (such as `main` or `trunk`) via the `BITRISEIO_GIT_BRANCH_DEST` env var, which is automatically set for PR builds. Restoring the cache from the target branch can be useful when there are multiple long-lived branches and PRs are targeting different branches.