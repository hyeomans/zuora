[![Build Status](https://travis-ci.com/hyeomans/zuora.svg?branch=master)](https://travis-ci.com/hyeomans/zuora)

A Go client library to consume [Zuora API](https://www.zuora.com/developer/api-reference/).

This is a __WIP__ and has minimal endpoints covered but it is really easy to add new ones.


- [Requirements](#requirements)
- [Basic usage](#basic-usage)
  * [Declaring a custom http client.](#declaring-a-custom-http-client)
  * [Declaring a custom Token Store](#declaring-a-custom-token-store)
  * [Error Handling](#error-handling)
    + [Error ordering](#error-ordering)
  * [ZOQL](#zoql)
    + [All Products example](#all-products-example)
    + [Filtering example](#filtering-example)
- [Available endpoints](#available-endpoints)
  * [Accounts](#accounts)
  * [Actions](#actions)
  * [Billing Documents](#billing-documents)
  * [Describe](#describe)
  * [Payments](#payments)
  * [Products](#products)
  * [Subscriptions](#subscriptions)


# Requirements

* Go >1.7
* Zuora client ID (Use Environment variables as best practice)
* Zuora client secret (Use Environment variables as best practice)
* Zuora api url (Use Environment variables as best practice)

# Basic usage

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hyeomans/zuora"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	zuoraClientID := os.Getenv("ZUORA_CLIENT_ID")
	zuoraClientSecret := os.Getenv("ZUORA_CLIENT_SECRET")
	zuoraURL := os.Getenv("ZUORA_URL")

	zuoraAPI := zuora.NewAPI(&http.Client{}, zuoraURL, zuoraClientID, zuoraClientSecret)
	ctx := context.Background()

	object, err := zuoraAPI.AccountsService.Summary(ctx, "customerAccount")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%+v\n", object)
	}
}
```

## Declaring a custom http client.
```go
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/hyeomans/zuora"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	zuoraClientID := os.Getenv("ZUORA_CLIENT_ID")
	zuoraClientSecret := os.Getenv("ZUORA_CLIENT_SECRET")
	zuoraURL := os.Getenv("ZUORA_URL")
	httpClient := newHTTPClient()

	zuoraAPI := zuora.NewAPI(httpClient, zuoraURL, zuoraClientID, zuoraClientSecret)
	ctx := context.Background()

	object, err := zuoraAPI.AccountsService.Summary(ctx, "customerAccount")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%+v\n", object)
	}
}

func newHTTPClient() *http.Client {
	keepAliveTimeout := 600 * time.Second
	timeout := 2 * time.Second
	defaultTransport := &http.Transport{
		Dial: (&net.Dialer{
			KeepAlive: keepAliveTimeout,
		}).Dial,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
	}

	return &http.Client{
		Transport: defaultTransport,
		Timeout:   timeout,
	}
}
```

## Declaring a custom Token Store

By default this package uses an in-memory backing store, but you can bring your own backing store, you only need to fullfill the interface:

```
//TokenStorer handles token renewal with two simple methods.
//Token() returns a boolean to indicate a token is valid and if valid, it will return the active token.
//Update() causes a side-effect to update a token in whichever backing store you choose.
type TokenStorer interface {
	Token() (bool, *Token)
	Update(*Token)
}
```


## Error Handling

Zuora API is not consistent with their error responses, this package tries to unify all error responses in a single one. One of the most important 
error responses from Zuora is __Request exceeded limit__ and this package follows ["Errors as behaviour"](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully) to identify when this happens.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hyeomans/zuora"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	zuoraClientID := os.Getenv("ZUORA_CLIENT_ID")
	zuoraClientSecret := os.Getenv("ZUORA_CLIENT_SECRET")
	zuoraURL := os.Getenv("ZUORA_URL")

	zuoraAPI := zuora.NewAPI(&http.Client{}, zuoraURL, zuoraClientID, zuoraClientSecret)
	ctx := context.Background()

	object, err := zuoraAPI.AccountsService.Summary(ctx, "customerAccount")

	if err != nil {
    temporary, ok := result.(err)
    if ok && temporary.Temporary() {
      fmt.Println("You could continue making requests after cool off time")
    } else if ok && temporary.Temporary() {
      log.Fatal("This is not a temporary error, modify your request")
    } else {
      log.Fatalf("an error ocurred %v", err)
    }
	} else {
		fmt.Printf("%+v\n", object)
	}
}
```

Errors that are temporary according to this package are:

```
http.StatusTooManyRequests
http.StatusLocked
http.StatusInternalServerError
```

More about error as behaviour:
[https://dave.cheney.net/2014/12/24/inspecting-errors](https://dave.cheney.net/2014/12/24/inspecting-errors)
[https://www.ardanlabs.com/blog/2014/10/error-handling-in-go-part-i.html](https://www.ardanlabs.com/blog/2014/10/error-handling-in-go-part-i.html)


### Error ordering
There could be a possibilty to that a response has multiple error messages encoded as described by Zuora documentation:

> If the JSON success field is false, a JSON "reasons" array is included in the response body with at least one set of code and message attributes that can be used to code a response.

Example:

```json
{

  "success": false,
  "processId": "3F7EA3FD706C7E7C",
  "reasons":  [
    {
      "code": 53100020,
      "message": " {com.zuora.constraints.either_or_both}"
    },
    {
      "code": 53100320,
      "message": "'termType' value should be one of: TERMED, EVERGREEN"
    }
  ]
}
```

The problem this presents is that, if you have a __Request exceeded limit__ inside here, you might take different approaches to handle it. This package resolves this issue be setting a priority on errors, here is the list from highest (top) to lowest priority (bottom):

```
http.StatusTooManyRequests
http.StatusUnauthorized
http.StatusForbidden
http.StatusNotFound
http.StatusLocked
http.StatusInternalServerError
http.StatusBadRequest //<-- If not in list, is considered a BadRequest
```


## ZOQL

Zuora allows to query tables by using a query language they call ZOQL, this package contains a helper struct/function to make easier to query whatever table you want.

### All Products example
Here is an example where:

* We wrap the API client
* Create a custom struct to Unmarshal the raw response from our Query endpoint:

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hyeomans/zuora"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	zuoraClientID := os.Getenv("ZUORA_CLIENT_ID")
	zuoraClientSecret := os.Getenv("ZUORA_CLIENT_SECRET")
	zuoraURL := os.Getenv("ZUORA_URL")

	zuoraAPI := zuora.NewAPI(&http.Client{}, zuoraURL, zuoraClientID, zuoraClientSecret)
	ctx := context.Background()

	myWrapper := myZuoraClient{zuoraAPI: zuoraAPI}

	products, err := myWrapper.GetProducts(ctx)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%+v\n", products)
	}
}

type myZuoraClient struct {
	zuoraAPI *zuora.API
}

type products struct {
	Records []zuora.Product `json:"records"`
}

//GetProducts Returns all products
func (m *myZuoraClient) GetProducts(ctx context.Context) ([]zuora.Product, error) {
	fields := []string{"ID", "Name"}
	zoqlComposer := zuora.NewZoqlComposer("Product", fields)
	rawProducts, err := m.zuoraAPI.ActionsService.Query(ctx, zoqlComposer)

	if err != nil {
		return nil, err
	}

	jsonResponse := products{}

	if err := json.Unmarshal(rawProducts, &jsonResponse); err != nil {
		return nil, err
	}

	return jsonResponse.Records, nil
}
```

### Filtering example

`ZoqlComposer` uses [functional options](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis) that allow you to compose a query that require:

* Single filter
* OR Filter
* AND filter
* Combination of those 3

Here is a minimal example:

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hyeomans/zuora"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	zuoraClientID := os.Getenv("ZUORA_CLIENT_ID")
	zuoraClientSecret := os.Getenv("ZUORA_CLIENT_SECRET")
	zuoraURL := os.Getenv("ZUORA_URL")

	zuoraAPI := zuora.NewAPI(&http.Client{}, zuoraURL, zuoraClientID, zuoraClientSecret)
	ctx := context.Background()

	myWrapper := myZuoraClient{zuoraAPI: zuoraAPI}

	products, err := myWrapper.GetProductById(ctx, "an-id")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%+v\n", products)
	}
}

type myZuoraClient struct {
	zuoraAPI *zuora.API
}

type products struct {
	Records []zuora.Product `json:"records"`
}

//GetProductById Will return a product by ID with many filters.
func (m *myZuoraClient) GetProductById(ctx context.Context, id string) ([]zuora.Product, error) {
	fields := []string{"ID", "Name"}
	filter := zuora.QueryFilter{Key: "ID", Value: id}
	singleFilter := zuora.QueryWithFilter(filter)
	andFilter := zuora.QueryWithAndFilter([]zuora.QueryFilter{filter, filter})
	orFilter := zuora.QueryWithOrFilter([]zuora.QueryFilter{filter, filter})
	zoqlComposer := zuora.NewZoqlComposer("Product", fields, singleFilter, andFilter, orFilter)
	fmt.Println(zoqlComposer) //You can print to see returning query.
	//{ "queryString" : "select ID, Name from Product where ID = 'an-id'  and  ID = 'an-id' and ID = 'an-id'  or  ID = 'an-id' or ID = 'an-id'" }
	rawProducts, err := m.zuoraAPI.ActionsService.Query(ctx, zoqlComposer)

	if err != nil {
		return nil, err
	}

	jsonResponse := products{}

	if err := json.Unmarshal(rawProducts, &jsonResponse); err != nil {
		return nil, err
	}

	return jsonResponse.Records, nil
}

```


# Available endpoints

## Accounts

| Zuora reference     | How to call it               | Link                                                                        |
|---------------------|------------------------------|-----------------------------------------------------------------------------|
| Get account summary | AccountsService.Summary(...) | https://www.zuora.com/developer/api-reference/#operation/GET_AccountSummary |

## Actions 

| Zuora reference | How to call it            | Link                                                                      |
|-----------------|---------------------------|---------------------------------------------------------------------------|
| Query           | ActionsService.Query(...) | https://www.zuora.com/developer/api-reference/#operation/Action_POSTquery |

## Billing Documents

| Zuora reference       | How to call it                   | Link                                                                          |
|-----------------------|----------------------------------|-------------------------------------------------------------------------------|
| Get billing documents | BillingDocumentsService.Get(...) | https://www.zuora.com/developer/api-reference/#operation/GET_BillingDocuments |

## Describe

| Zuora reference | How to call it             | Link                                                        |
|-----------------|----------------------------|-------------------------------------------------------------|
| Describe object | DescribeService.Model(...) | https://www.zuora.com/developer/api-reference/#tag/Describe |

## Payments

| Zuora reference   | How to call it                    | Link                                                                       |
|-------------------|-----------------------------------|----------------------------------------------------------------------------|
| CRUD: Get payment | PaymentsService.ByIdThroughObject | https://www.zuora.com/developer/api-reference/#operation/Object_GETPayment |

## Products

| Zuora reference        | How to call it           | Link                                                                       |
|------------------------|--------------------------|----------------------------------------------------------------------------|
| CRUD: Retrieve Product | ProductsService.Get(...) | https://www.zuora.com/developer/api-reference/#operation/Object_GETProduct |

## Subscriptions

| Zuora reference              | How to call it                       | Link                                                                            |
|------------------------------|--------------------------------------|---------------------------------------------------------------------------------|
| Get subscriptions by account | SubscriptionsService.ByKey(...)      | https://www.zuora.com/developer/api-reference/#operation/GET_SubscriptionsByKey |
| CRUD: Retrieve Subscription  | SubscriptionsService.ByAccount(...)  | https://www.zuora.com/developer/api-reference/#operation/Object_GETSubscription |
| CRUD: Update Subscription    | SubscriptionsService.Update(...)     | https://www.zuora.com/developer/api-reference/#operation/Object_PUTSubscription |
| CRUD: Update Subscription    | SubscriptionsService.UpdateFull(...) | https://www.zuora.com/developer/api-reference/#operation/Object_PUTSubscription |
| Cancel subscription          | SubscriptionsService.Cancel(...)     | https://www.zuora.com/developer/api-reference/#operation/PUT_CancelSubscription |