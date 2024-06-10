# Drinks
(*Drinks*)

## Overview

The drinks endpoints.

### Available Operations

* [GetDrink](#getdrink) - Get a drink.
* [ListDrinks](#listdrinks) - Get a list of drinks.

## GetDrink

Get a drink by name, if authenticated this will include stock levels and product codes otherwise it will only include public information.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/template-speakeasy-bar/pkg/models/shared"
	templatespeakeasybar "github.com/speakeasy-sdks/template-speakeasy-bar"
	"context"
	"log"
)

func main() {
    s := templatespeakeasybar.New(
        templatespeakeasybar.WithSecurity(shared.Security{
            APIKey: templatespeakeasybar.String("<YOUR_API_KEY>"),
        }),
    )
    var name string = "<value>"
    ctx := context.Background()
    res, err := s.Drinks.GetDrink(ctx, name)
    if err != nil {
        log.Fatal(err)
    }
    if res.Drink != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `name`                                                | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.GetDrinkResponse](../../pkg/models/operations/getdrinkresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 5XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListDrinks

Get a list of drinks, if authenticated this will include stock levels and product codes otherwise it will only include public information.

### Example Usage

```go
package main

import(
	templatespeakeasybar "github.com/speakeasy-sdks/template-speakeasy-bar"
	"github.com/speakeasy-sdks/template-speakeasy-bar/pkg/models/operations"
	"github.com/speakeasy-sdks/template-speakeasy-bar/pkg/models/shared"
	"context"
	"log"
)

func main() {
    s := templatespeakeasybar.New()
    security := operations.ListDrinksSecurity{
            BearerAuth: templatespeakeasybar.String("<YOUR_JWT>"),
        }

    var drinkType *shared.DrinkType = shared.DrinkTypeSpirit.ToPointer()
    ctx := context.Background()
    res, err := s.Drinks.ListDrinks(ctx, security, drinkType)
    if err != nil {
        log.Fatal(err)
    }
    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `security`                                                                         | [operations.ListDrinksSecurity](../../pkg/models/operations/listdrinkssecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |
| `drinkType`                                                                        | [*shared.DrinkType](../../pkg/models/shared/drinktype.md)                          | :heavy_minus_sign:                                                                 | The type of drink to filter by. If not provided all drinks will be returned.       |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.ListDrinksResponse](../../pkg/models/operations/listdrinksresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 5XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
