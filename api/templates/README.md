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
* Add checks for `isBinary` to generate `io.ReadCloser` request fields instead of `**os.File`
* Update creation of `GenericOpenAPIError` to track sub-error models by reference
* Add checks for `isResponseBinary` to directly return the raw `*http.Response`, instead of
  pulling the entire body into memory and transforming it into an `*os.File`
* GUnzip non-binary response bodies before unmarshalling when `Content-Encoding: gzip` is set
* Remove `*http.Response`s from the return values of generated `Execute` operations & provide a separate method
  `ExecuteWithHttpInfo` for returning the `*http.Response` if needed. The standard `Execute` operation should normally be
  used unless specific header values or response codes need to be examined with the `ExecuteWithHttpInfo`. The response
  body cannot be read from the `*http.Reponse` returned from ExecuteWithHttpInfo` as it will have already have been read
  and closed; access to the response body content should be achieved through the returned response model if applicable

`client.mustache`
* Removed use of `golang.org/x/oauth2` to avoid its heavy dependencies
* Fixed error strings to be idiomatic according to staticcheck (lowercase, no punctuation)
* Use `strings.EqualFold` instead of comparing two `strings.ToLower` calls
* Update the `GenericOpenAPIError` type to enforce that error response models implement the `error` interface
* Update `setBody` to avoid buffering data in memory when the request body is already an `io.ReadCloser`

`configuration.mustache`
* Deleted `ContextOAuth2` key to match modification in client
* Fixed error strings to be idiomatic according to staticcheck (lowercase, no punctuation)

`model_enum.mustache`
* Add a function to build and return a slice of the valid values for the enum, named <enum_name>Values

`model_oneof.mustache`
* Fixed error strings to be idiomatic according to staticcheck (lowercase, no punctuation)

`model_simple.mustache`
* Added `yaml:` tags to all model fields to support unmarshalling camelCase
* Added support for `x-go-field-type` vendor extension, to explicitly override the type generated for model fields
