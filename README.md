[![Build Status](https://travis-ci.com/hyeomans/zuora.svg?branch=master)](https://travis-ci.com/hyeomans/zuora)

A Go client library to consume [Zuora API](https://www.zuora.com/developer/api-reference/).

This is a __WIP__ and has minimal endpoints covered but it is really easy to add new ones.

# Requirements

* Go >1.7
* Zuora client ID (Use Environment variables as best practice)
* Zuora client secret (Use Environment variables as best practice)
* Zuora api url (Use Environment variables as best practice)

## Missing types

Why return an array of bytes? 
Zuora responses vary from company to company. The variation comes from Custom Fields defined into your company definition of Zuora.

The package could send typed information, let's take the example of `Product`. Zuora defines "Product" entity like this:

```
type Product struct {
  AllowFeatureChanges *bool   `json:"AllowFeatureChanges,omitempty"`
  Category            *string `json:"Category,omitempty"`
  CreatedByID         *string `json:"CreatedById,omitempty"`
  CreatedDate         *string `json:"CreatedDate,omitempty"`
  Description         *string `json:"Description,omitempty"`
  EffectiveEndDate    string  `json:"EffectiveEndDate"`
  EffectiveStartDate  string  `json:"EffectiveStartDate"`
  ID                  *string `json:"Id,omitempty"`
  Name                string  `json:"Name"`
  SKU                 *string `json:"SKU,omitempty"`
  UpdatedByID         *string `json:"UpdatedById,omitempty"`
  UpdatedDate         *string `json:"UpdatedDate,omitempty"`
}
```

But, how would you define custom fields if the signature of the method is:

```
func (t *catalogService) GetProduct(ctx context.Context, pageSize int) (*Product, error) {....}
```

Well, you can't. That's why we return the raw bytes, and then you can marshal into your own struct. We include common types into the package, so you don't have to guess. 
Imagine you have a custom field named "DisplayName__c", you can define your own struct using the power of struct embedding. For example:

```
type myProduct struct {
  zuora.Product
  DisplayName     *string `json:"DisplayName__c,omitempty"`
}
```

Now marshal the JSON into your custom struct. Let's see a practical example with the Account Summary endpoint.

## Account Summary Example

Account summary response retrieves a great overview of an account state. The problem is that if you defined custom properties, the nested payload would include those properties.

In the following code example, you will find custom structs that define those custom properties, that later will be bound to a custom struct.

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/hyeomans/zuora"
	"github.com/joho/godotenv"
)

type myRatePlan struct {
	zuora.RatePlan
	MyCustomProperty *string            `json:"MyCustomProperty__c,omitempty"`
	RatePlanCharges  []myRatePlanCharge `json:"ratePlanCharges"`
}

type myRatePlanCharge struct {
	zuora.RatePlanCharge
	MyCustomProperty *string `json:"MyCustomProperty,omitempty"`
}

type mySubscription struct {
	zuora.Subscription
	MyCustomProperty *string      `json:"MyCustomProperty__c,omitempty"`
	RatePlans        []myRatePlan `json:"ratePlans"`
}

type myInvoice struct {
	zuora.Invoice
	MyCustomProperty *string `json:"MyCustomProperty__c,omitempty"`
}

type myAccount struct {
	zuora.Account
	DefaultPaymentMethod zuora.PaymentMethod `json:"defaultPaymentMethod"`
	MyCustomProperty     *string             `json:"MyCustomProperty__c,omitempty"`
}

type summary struct {
	BasicInfo     myAccount        `json:"basicInfo"`
	BillToContact zuora.Contact    `json:"billToContact"`
	SoldToContact zuora.Contact    `json:"soldToContact"`
	TaxInfo       zuora.Account    `json:"taxInfo"`
	Subscriptions []mySubscription `json:"subscriptions"`
	Invoices      []myInvoice      `json:"invoices"`
	Usage         []zuora.Usage    `json:"usage"`
	Payments      []zuora.Payment  `json:"payments"`
	Success       bool             `json:"success"`
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	zuoraClientID := os.Getenv("ZUORA_CLIENT_ID")
	zuoraClientSecret := os.Getenv("ZUORA_CLIENT_SECRET")
	zuoraURL := os.Getenv("ZUORA_URL")
	httpClient := newHTTPClient()

	if err != nil {
		log.Fatal(err)
	}

	zuoraOAuthHeaderProvider := zuora.NewOAuthHeader(httpClient, &zuora.MemoryTokenStore{}, zuoraClientID, zuoraClientSecret, zuoraURL)
	zuoraAPI := zuora.NewAPI(httpClient, zuoraOAuthHeaderProvider, zuoraURL)

	r, err := zuoraAPI.V1.AccountsService.Summary(ctx, "accountIdFromYouZuoraInstance")

	if err != nil {
		log.Fatal(err)
	}
	s := summary{}

	if err = json.Unmarshal(r, &s); err != nil {
		log.Fatal(err)
	}

	fmt.Println(*s.TaxInfo.VATId)
}

func newHTTPClient() *http.Client {
	keepAliveTimeout := 600 * time.Second
	timeout := 3 * time.Second
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

You can apply this pattern to other endpoints, for example, "Subscriptions," "Products," or ZOQL calls to the Query endpoint.


## Updating an Account
Updating an account through Zuora requires to send a custom JSON payload. According to documentation, it is only necessary to submit those properties that need to be changed.

This package also allows modifying custom properties. Here is an example:

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

type myAccountUpdate struct {
	zuora.AccountUpdate
	CustomField *string `json:"CustomField__c,omitempty"`
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	zuoraClientID := os.Getenv("ZUORA_CLIENT_ID")
	zuoraClientSecret := os.Getenv("ZUORA_CLIENT_SECRET")
	zuoraURL := os.Getenv("ZUORA_URL")
	httpClient := newHTTPClient()

	if err != nil {
		log.Fatal(err)
	}

	zuoraOAuthHeaderProvider := zuora.NewOAuthHeader(httpClient, &zuora.MemoryTokenStore{}, zuoraClientID, zuoraClientSecret, zuoraURL)
	zuoraAPI := zuora.NewAPI(httpClient, zuoraOAuthHeaderProvider, zuoraURL)

	if err != nil {
		log.Fatal(err)
	}

	myAcc := myAccountUpdate{}

	tr := "My Custom Field"
	myAcc.CustomField = &tr
	myAcc.Name = "New Name"

	response, err := zuoraAPI.V1.AccountsService.Update(ctx, "AccountNumber", myAcc)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Done: %+v", response)
}
```

## Updating a Subscription

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

	ctx := context.Background()
	zuoraClientID := os.Getenv("ZUORA_CLIENT_ID")
	zuoraClientSecret := os.Getenv("ZUORA_CLIENT_SECRET")
	zuoraURL := os.Getenv("ZUORA_URL")
	httpClient := newHTTPClient()

	if err != nil {
		log.Fatal(err)
	}

	zuoraOAuthHeaderProvider := zuora.NewOAuthHeader(httpClient, &zuora.MemoryTokenStore{}, zuoraClientID, zuoraClientSecret, zuoraURL)
	zuoraAPI := zuora.NewAPI(httpClient, zuoraOAuthHeaderProvider, zuoraURL)

	mySub := zuora.SubscriptionUpdate{}
	note := "some notheeees"
	mySub.Notes = &note
	r, err := zuoraAPI.V1.SubscriptionsService.Update(ctx, "A-S000XXXXX", mySub)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", r)
}

func newHTTPClient() *http.Client {
	keepAliveTimeout := 600 * time.Second
	timeout := 3 * time.Second
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

## Cancelling a subscription

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

	ctx := context.Background()
	zuoraClientID := os.Getenv("ZUORA_CLIENT_ID")
	zuoraClientSecret := os.Getenv("ZUORA_CLIENT_SECRET")
	zuoraURL := os.Getenv("ZUORA_URL")
	httpClient := newHTTPClient()

	if err != nil {
		log.Fatal(err)
	}

	zuoraOAuthHeaderProvider := zuora.NewOAuthHeader(httpClient, &zuora.MemoryTokenStore{}, zuoraClientID, zuoraClientSecret, zuoraURL)
	zuoraAPI := zuora.NewAPI(httpClient, zuoraOAuthHeaderProvider, zuoraURL)

	mySub := zuora.SubscriptionCancellation{InvoiceCollect: false, CancellationPolicy: "EndOfLastInvoicePeriod"}
	r, err := zuoraAPI.V1.SubscriptionsService.Cancel(ctx, "A-S000XXX00", mySub)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", r)
}

func newHTTPClient() *http.Client {
	keepAliveTimeout := 600 * time.Second
	timeout := 3 * time.Second
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

## Getting Expired Subscriptions with Zoql

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

	ctx := context.Background()
	zuoraClientID := os.Getenv("ZUORA_CLIENT_ID")
	zuoraClientSecret := os.Getenv("ZUORA_CLIENT_SECRET")
	zuoraURL := os.Getenv("ZUORA_URL")
	httpClient := newHTTPClient()

	if err != nil {
		log.Fatal(err)
	}

	zuoraOAuthHeaderProvider := zuora.NewOAuthHeader(httpClient, &zuora.MemoryTokenStore{}, zuoraClientID, zuoraClientSecret, zuoraURL)
	zuoraAPI := zuora.NewAPI(httpClient, zuoraOAuthHeaderProvider, zuoraURL)

	zoql := zuora.NewZoqlComposer().
		Fields("Name").From("Subscription").
		Where("status", "cancelled").
		And("termEndDate", time.Now().UTC().Format("2006-01-02"))

	t, err := zuoraAPI.V1.ActionsService.Query(ctx, zoql)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(t))
}

func newHTTPClient() *http.Client {
	keepAliveTimeout := 600 * time.Second
	timeout := 3 * time.Second
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