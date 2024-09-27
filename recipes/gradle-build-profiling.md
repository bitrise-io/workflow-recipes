# Turn on Gradle build profiling

## Description
Generate and store a performance report of every Gradle build to spot build speed issues or compare different builds.

## Instructions

No matter what Android or Gradle Step you use in your Bitrise Workflow, there is an option to define additional command line arguments for Gradle. Add `--profile` to this input to generate a performance report of the Gradle tasks. In the example below, we are adding the argument to the Android Unit Test Step.

To sum up the procedure:
1. Add the Android Unit Test Step to your Workflow.
2. Add a Script Step to compress the reports and copy the ZIP file to the deploy directory. 
3. Trigger a manual build and download and open the HTML report.
4. Check the various aspects of the build in the report. 

### Adding the Android Unit Test Step

Add an Android Unit Test Step to your Workflow. Set the necessary input values:

- **Project location**: "$PROJECT_LOCATION"
- **Module**: "$MODULE"
- **Variant**: "$VARIANT"
- **Arguments**: "--profile"

### Compressing the report files and copying the ZIP file

Add a Script Step to the end of the Workflow in order to compress the report files and copy the ZIP file to the deploy directory:

```
#!/usr/bin/env bash
# fail if any commands fails
set -e
# debug log
set -x

zip -r $BITRISE_DEPLOY_DIR/gradle-profile.zip $PROJECT_LOCATION/build/reports/profile
```
Gradle creates the HTML report in build/reports/profile/, so we need to take all files in that folder (HTML, CSS and JS files), compress them, and move the ZIP archive to $BITRISE_DEPLOY_DIR. Files in this folder can be accessed on the build page’s Apps & Artifacts tab.

### Downloading the report file

Trigger a manual build of the Workflow you edited previously. Download and unarchive `gradle-profile.zip`, then open the HTML report in your browser.

![image](https://user-images.githubusercontent.com/5689177/149338897-3ab06e6a-9b58-465f-95fc-45235afff259.png)

![image](https://user-images.githubusercontent.com/5689177/149339159-0ad35634-518e-464c-a460-2c7c636cae4a.png)



### Checking the build report

You can check various aspects of a build in the report:

- The **Summary** tab shows time spent on things other than task execution

- The **Task execution** tab lists all tasks sorted by execution time

- Cached tasks are marked as **UP-TO-DATE**. This helps to fine-tune the [Bitrise Cache Steps](https://devcenter.bitrise.io/builds/caching/about-caching-index/) by comparing the reports of multiple builds.

For Gradle optimization ideas, check out [this article by Google](https://developer.android.com/studio/build/profile-your-build#using-the-gradle---profile-option).

If you only want to display task execution times only in the build log, you can use the [build-time-tracker](https://github.com/asarkar/build-time-tracker) project.

## bitrise.yml

```yaml
- android-unit-test@1:
    inputs:
    - project_location: $PROJECT_LOCATION
    - module: $MODULE
    - arguments: "--profile"
    - variant: $VARIANT
- script@1:
    title: Collect Gradle profile report
    inputs:
    - content: |-
        #!/usr/bin/env bash
        # fail if any commands fails
        set -e
        # debug log
        set -x

        zip -r $BITRISE_DEPLOY_DIR/gradle-profile.zip $PROJECT_LOCATION/build/reports/profile
- deploy-to-bitrise-io@2: {}
```
