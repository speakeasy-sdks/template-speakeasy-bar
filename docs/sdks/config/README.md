# Config
(*Config*)

### Available Operations

* [SubscribeToWebhooks](#subscribetowebhooks) - Subscribe to webhooks.

## SubscribeToWebhooks

Subscribe to webhooks.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/template-speakeasy-bar/pkg/models/shared"
	templatespeakeasybar "github.com/speakeasy-sdks/template-speakeasy-bar"
	"github.com/speakeasy-sdks/template-speakeasy-bar/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := templatespeakeasybar.New(
        templatespeakeasybar.WithSecurity(shared.Security{
            APIKey: templatespeakeasybar.String("<YOUR_API_KEY>"),
        }),
    )
    request := []operations.RequestBody{
        operations.RequestBody{},
    }
    ctx := context.Background()
    res, err := s.Config.SubscribeToWebhooks(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [[]operations.RequestBody](../../.md)                        | :heavy_check_mark:                                           | The request object to use for the request.                   |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |


### Response

**[*operations.SubscribeToWebhooksResponse](../../pkg/models/operations/subscribetowebhooksresponse.md), error**
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.APIError   | 5XX                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |
