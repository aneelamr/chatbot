# chatbot

This is a chatbot that uses the messenger api to gather personal reviews from customers.

## Continuous Integration

Continuous integration is enabled through github actions.
The following workflows are enabled-

* Build and Test: This workflow uses the go cli to validate that source code can be built and all unit tests pass.
* golangci-lint: This package uses golangci-lint to run quick linters on go. Please check
the workflow under `.github/workflows/golangci-lint.yml` for configuration of linter.

## References

* Github Actions <https://github.com/features/actions>
* Messaging API <https://developers.facebook.com/docs/messenger-platform/>
* golangci-lint <https://github.com/golangci/golangci-lint>
