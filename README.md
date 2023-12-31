# chatbot

This is a chatbot that uses the messenger api to gather personal reviews from customers.

## Creating Lambda zip file locally

The below step is automated in the continuous integration pipeline.

To create the lambda deployment zip file locally

```shell
GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
zip lambda-handler.zip bootstrap
```

## Lex API

Amazon Lex is a fully managed artificial intelligence (AI) service with advanced natural language models to design, build, test, and deploy conversational interfaces in applications.

Lambda responds to events from a Lex Bot.

## Package Structure

A quick tour of this package's contents

* intents: Intents are the goals that your users want to accomplish, such as ordering flowers or booking a hotel. Your bot must have at least one intent. This package has handlers for handling various user intents. See more: <https://docs.aws.amazon.com/lexv2/latest/dg/add-intents.html>

* lex: Lex contains event and response structs that interact with the lex API.

## Continuous Integration/ Continuous Deployment

CI/CD is enabled through github actions.
The following workflows are enabled-

* CI: Build and Test: This workflow uses the go cli to validate that source code can be built and all unit tests pass.
* CI: golangci-lint: This package uses golangci-lint to run quick linters on go. Please check
the workflow under `.github/workflows/golangci-lint.yml` for configuration of linter.
* CD: deploy-lambda: deploys the lex-bot lambda

## References

* Lex <https://aws.amazon.com/lex/>
* Github Actions <https://github.com/features/actions>
* Messaging API <https://developers.facebook.com/docs/messenger-platform/>
* golangci-lint <https://github.com/golangci/golangci-lint>
