# Custom OpenAPI Templates

This directory contains custom mustache templates used by the OpenAPI code generator.
The original templates were extracted by running:
```shell
openapi-generator author template -g go
```
NOTE: This command extracts a copy of every template used by the generator, but we only
track templates that we've modified here. The generator can handle sourcing templates from
multiple locations.

## What have we changed?

`api.mustache`
* Add `GetX()` methods for each request parameter `X`, for use in unit tests

`client.mustache`
* Removed use of `golang.org/x/oauth2` to avoid its heavy dependencies
* Fixed error strings to be idiomatic according to staticcheck (lowercase, no punctuation)
* Use `strings.EqualFold` instead of comparing two `strings.ToLower` calls

`configuration.mustache`
* Deleted `ContextOAuth2` key to match modification in client
* Fixed error strings to be idiomatic according to staticcheck (lowercase, no punctuation)
