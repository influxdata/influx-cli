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
* Add checks for `isByteArray` to generate `[]byte` request fields instead of `*string`
* Update creation of `GenericOpenAPIError` to track sub-error models by reference
* Add checks for `isResponseBinary` to directly return the response-body-reader, instead of
  pulling the entire body into memory and transforming it into an `*os.File`
  

`client.mustache`
* Removed use of `golang.org/x/oauth2` to avoid its heavy dependencies
* Fixed error strings to be idiomatic according to staticcheck (lowercase, no punctuation)
* Use `strings.EqualFold` instead of comparing two `strings.ToLower` calls
* GZip request bodies when `Content-Encoding: gzip` is set
* Update the `GenericOpenAPIError` type to enforce that error response models implement the `error` interface

`configuration.mustache`
* Deleted `ContextOAuth2` key to match modification in client
* Fixed error strings to be idiomatic according to staticcheck (lowercase, no punctuation)
