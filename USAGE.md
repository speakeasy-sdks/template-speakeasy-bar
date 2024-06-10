<!-- Start SDK Example Usage [usage] -->
### Sign in

First you need to send an authentication request to the API by providing your username and password.
In the request body, you should specify the type of token you would like to receive: API key or JSON Web Token.
If your credentials are valid, you will receive a token in the response object: `res.object.token: str`.

```go
package main

import (
	"context"
	templatespeakeasybar "github.com/speakeasy-sdks/template-speakeasy-bar"
	"github.com/speakeasy-sdks/template-speakeasy-bar/pkg/models/operations"
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

### Browse available drinks

Once you are authenticated, you can use the token you received for all other authenticated endpoints.
For example, you can filter the list of available drinks by type.

```go
package main

import (
	"context"
	templatespeakeasybar "github.com/speakeasy-sdks/template-speakeasy-bar"
	"github.com/speakeasy-sdks/template-speakeasy-bar/pkg/models/operations"
	"github.com/speakeasy-sdks/template-speakeasy-bar/pkg/models/shared"
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

### Create an order

When you submit an order, you can include a callback URL along your request.
This URL will get called whenever the supplier updates the status of your order.

```go
package main

import (
	"context"
	templatespeakeasybar "github.com/speakeasy-sdks/template-speakeasy-bar"
	"github.com/speakeasy-sdks/template-speakeasy-bar/pkg/models/shared"
	"log"
)

func main() {
	s := templatespeakeasybar.New(
		templatespeakeasybar.WithSecurity(shared.Security{
			APIKey: templatespeakeasybar.String("<YOUR_API_KEY>"),
		}),
	)
	requestBody := []shared.OrderInput{
		shared.OrderInput{
			ProductCode: "APM-1F2D3",
			Quantity:    26535,
			Type:        shared.OrderTypeDrink,
		},
	}

	var callbackURL *string = templatespeakeasybar.String("<value>")
	ctx := context.Background()
	res, err := s.Orders.CreateOrder(ctx, requestBody, callbackURL)
	if err != nil {
		log.Fatal(err)
	}
	if res.Order != nil {
		// handle response
	}
}

```

### Subscribe to webhooks to receive stock updates

```go
package main

import (
	"context"
	templatespeakeasybar "github.com/speakeasy-sdks/template-speakeasy-bar"
	"github.com/speakeasy-sdks/template-speakeasy-bar/pkg/models/operations"
	"github.com/speakeasy-sdks/template-speakeasy-bar/pkg/models/shared"
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
<!-- End SDK Example Usage [usage] -->