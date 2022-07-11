# Workflow Recipes

Workflow Recipes provide ready-made solutions for common Workflow tasks. Here you will find a range of different Recipes along with examples of entire Workflows.

## Using Recipes

You can use Workflow Recipes in two ways:

* By adding the Steps into your Workflow via the Workflow Editor
* By copy-pasting the `bitrise.yml` snippet into your [app's `bitrise.yml` file](https://devcenter.bitrise.io/en/builds/configuring-build-settings/managing-an-app-s-bitrise-yml-file.html)

### Adding Steps via the Workflow Editor

All you need to do here is follow the step-by-step instructions in the Recipe.

![recipes_ui_4](https://user-images.githubusercontent.com/5689177/150490578-788aac88-3ac1-404d-900b-c94e131af3d4.gif)


### Copy-pasting the bitrise.yml snippet

You can also simply copy-paste the snippet to your `bitrise.yml` file directly. Don't forget to:

1. Check the formatting of the copy-pasted YAML
2. Read the instructions as they may contain some important information on configuration
3. Check and customize the input variables

![recipe_yml](https://user-images.githubusercontent.com/5689177/150491534-fdde2ba1-aa3a-4f6f-9895-6cd8200bd6f1.gif)

## Please Share your Feedback

**Workflow Recipes** is currently an MVP feature, and we’d love to hear your feedback on it. Share your thoughts on Workflow Recipes by filling out our feedback survey below (~3-5min).

### [Fill out the Feedback Survey](https://docs.google.com/forms/d/e/1FAIpQLSfBQZCB02uOMsjp1kTLJ9Bv2tm0o39w4ez638m3y3kN5KQH_w/viewform?usp=sf_link)

  

## The Recipes

### Cloning & Setup

* [Cloning the Repository](recipes/ssh-and-clone.md)
* [(Flutter) Install Flutter SDK](recipes/flutter-install-flutter-sdk.md)

### Dependencies

* [(iOS) Install CocoaPods Dependencies](recipes/ios-install-cocoapods-dependencies.md)
* [(iOS) Install Carthage Dependencies](recipes/ios-install-carthage-dependencies.md)
* [(React Native) Install Dependencies](recipes/rn-install-dependencies.md)

### Testing

* [(iOS) Run Tests on Simulator](recipes/ios-simulator-test.md)
* [(iOS) Run Tests on a Physical Device](recipes/ios-device-testing.md)
* [(Android) Run Unit Tests](recipes/android-unit-test.md)
* [(Android) Run UI / Instrumentation Tests on Virtual Device](recipes/android-virtual-device-tests.md)
* [(Android) Run UI / Instrumentation Tests on Local Emulator](recipes/android-emulator-test.md)
* [(React Native) Run Tests](recipes/rn-tests.md)
* [(Flutter) Run Tests](recipes/flutter-test.md)

### Building

* [(React Native) Expo: Build using Turtle CLI](recipes/rn-expo-turtle-build.md)

### Linting

* [(Android) Run Lint](recipes/android-lint.md)
* [(Flutter) Run Dart Analyzer](recipes/flutter-dart-analyzer.md)

### Deploying

* [(iOS) Deploy to App Store Connect / TestFlight](recipes/ios-deploy-to-appstore.md)
* [(Android) Deploy to Google Play (Internal, Alpha, Beta, Production)](recipes/android-deploy-to-google-play.md)
* [(iOS) Deploy to Bitrise.io](recipes/ios-deploy-to-bitrise.md)
* [(Android) Deploy to Bitrise.io](recipes/android-deploy-to-bitrise.md)
* [(iOS) Deploy to Firebase App Distribution](recipes/ios-deploy-to-firebase.md)
* [(Android) Deploy to Firebase App Distribution](recipes/android-deploy-to-firebase.md)
* [(iOS) Deploy to Visual Studio App Center](recipes/ios-deploy-to-appcenter.md)
* [(Android) Deploy to Visual Studio App Center](recipes/android-deploy-to-appcenter.md)

### Notifications

* [Slack - Send Build Status](recipes/slack-send-build-status.md)
* [Slack - Send the Build QR Code](recipes/slack-send-qr-code.md)
* [GitHub Pull Request - Send the Build QR Code](recipes/github-pull-request-build-qr-code.md)

### Optimisation & Caching

* [Make Caching Efficient for Pull Request Builds](recipes/pull-request-build-caching.md)
* [Shallow Clone the git Repo](recipes/shallow-clone-repo.md)
* [(iOS) Cache CocoaPods Dependencies](recipes/ios-cache-cocoapods.md)
* [(Android) Turn on Gradle build profiling](recipes/gradle-build-profiling.md)
* [(React Native) Cache Dependencies (node_modules)](recipes/rn-cache-dependencies.md)

### Running Steps & Workflows

* [Run a Step Conditionally](recipes/run-step-conditionally.md)
* [Start (Parallel) Builds from the Workflow](recipes/start-builds.md)


## Example Workflows

### iOS

* [(iOS) Pull Request](recipes/ios-pull-request-workflow.md)
* [(iOS) CI](recipes/ios-ci-workflow.md)
* [(iOS) Nightly](recipes/ios-nightly-workflow.md)
* [(iOS) Release](recipes/ios-release-workflow.md)

### Android

* [(Android) Pull Request](recipes/android-pull-request-workflow.md)
* [(Android) CI](recipes/android-ci-workflow.md)
* [(Android) Nightly](recipes/android-nightly-workflow.md)
* [(Android) Release](recipes/android-release-workflow.md)

### Other
* [Create Gitflow release branch](recipes/workflow-create-gitflow-release-branch.md)

## Example Pipelines

### iOS

* [(iOS) Run tests in parallel on multiple simulators](./recipes/ios-run-tests-in-parallel-on-multiple-simulators.md)
* [(iOS) Run test groups in parallel](./recipes/ios-run-test-groups-in-parallel.md)
* [(iOS) Merging test results and deploying to the Test Reports add-on](./recipes/ios-merging-test-results-and-deploying-to-the-test-reports-add-on.md)

### Android

* [(Android) Parallel testing of unit test shards by module](recipes/android-parallel-testing-unit-test-shards.md)
* [(Android) Parallel UI tests on multiple devices](recipes/android-parallel-ui-tests-on-multiple-devices.md)
* [(Android) Parallel unit and UI tests](recipes/android-parallel-unit-and-ui-tests.md)
