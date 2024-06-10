# Authentication
(*Authentication*)

## Overview

The authentication endpoints.

### Available Operations

* [Login](#login) - Authenticate with the API by providing a username and password.

## Login

Authenticate with the API by providing a username and password.

### Example Usage

```go
package main

import(
	templatespeakeasybar "github.com/speakeasy-sdks/template-speakeasy-bar"
	"github.com/speakeasy-sdks/template-speakeasy-bar/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := templatespeakeasybar.New()
    request := operations.LoginRequestBody{
        Type: operations.TypeAPIKey,
    }

    security := operations.LoginSecurity{
            Password: "<PASSWORD>",
            Username: "<USERNAME>",
        }
    ctx := context.Background()
    res, err := s.Authentication.Login(ctx, request, security)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.LoginRequestBody](../../pkg/models/operations/loginrequestbody.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.LoginSecurity](../../pkg/models/operations/loginsecurity.md)       | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.LoginResponse](../../pkg/models/operations/loginresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 5XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
