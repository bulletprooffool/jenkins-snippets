# Postman Collections Repository

This repository contains a series of Postman collections that can be triggered using Newman in a Jenkins pipeline. The collections are designed to be used for testing different environments and have separate branches aligned to each environment.

## Test Environments

The repository consists of the following branches, each aligned to a different test environment:

- `foo` (Production Environment)
- `foo-dev` (Development Environment)
- `foo-uat` (User Acceptance Testing Environment)
- `foo-sap` (SAP Integration Environment)

## Updating Postman Collections

To update the Postman collections, follow these steps:

1. Clone the repository locally: `git clone <repository_url>`.
2. Switch to the branch representing the desired environment: `git checkout <branch_name>`.
3. Make the necessary changes to the Postman collection using the Postman desktop application or the Postman API.
4. Export the updated collection as a JSON file and replace the existing collection in the repository.
5. Commit and push the changes to the repository.

**Note:** It's important to ensure that no sensitive information, such as passwords, is included in the Postman collections or any other public documents.

## Environment Variables (env-vars.json)

Each branch in the repository contains an `env-vars.json` file specific to the corresponding environment. This file should include environment-specific variables required for running the Postman collections.

To update the `env-vars.json` file, follow these guidelines:

- Use variable names that are self-descriptive and meaningful.
- Avoid hard-coding any environment-specific values in the `env-vars.json` file.
- Reference the variables stored in the global var files stored in Jenkins for credentials or sensitive information.
- For local execution and testing purposes, you can utilize Postman's local/global variables.

**Note:** Ensure that all credential mappings in the `env-vars.json` file reference the variables stored in the global var files stored in Jenkins. This helps maintain security and prevents sensitive information from being exposed in the repository.

## Running Tests Locally

To verify the tests locally, follow these steps:

1. Ensure that you have Postman installed on your local machine.
2. Clone the repository locally: `git clone <repository_url>`.
3. Switch to the branch representing the desired environment: `git checkout <branch_name>`.
4. Open Postman and import the Postman collection JSON file from the cloned repository.
5. Set the required environment variables in Postman using the `env-vars.json` file.
6. Execute the individual requests within the collection to verify their functionality.

**Note:** Avoid hard-referencing any URIs or specific endpoints in the tests. This ensures that the tests remain flexible and can be easily adapted to different endpoints or environments.

## Managing Longer Running Tasks

For longer running tasks or tests, it is recommended to refer to existing tests within the repository. These tests may provide guidance on how to manage tasks that take a considerable amount of time to complete.

To refer to existing tests, follow these steps:

1. Open the Postman collection in the repository using the Postman desktop application.
2. Navigate to the specific test or request that demonstrates the management of longer running tasks.
3. Review the test documentation and associated scripts to understand the approach taken.
4. Apply the relevant strategies and techniques to manage longer running tasks in your own tests.

Remember to document any additional information or considerations specific to your tests and update the `README.md` file accordingly.

