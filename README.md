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

Now marshal the JSON into your custom struct.

# Basic usage

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

type mySubscription struct {
	zuora.Subscription
	AppVersion                 *string `json:"AppVersion__c,omitempty"`
	BusinessLine               *string `json:"BusinessLine__c,omitempty"`
	EmployeesFromAccountObject *string `json:"Employees_From_Account_Object__c,omitempty"`
	PONumber                   *string `json:"PO_Number__c,omitempty"`
	QuoteOpportunityID         *string `json:"Quote_Opportunity_ID__c,omitempty"`
	ResellerCompanyName        *string `json:"ResellerCompanyName__c,omitempty"`
	Segment                    *string `json:"Segment__c,omitempty"`
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

	zuoraOAuthHeaderProvider := zuora.NewOAuthHeader(httpClient, &zuora.MemoryTokenStore{}, zuoraClientID, zuoraClientSecret, zuoraURL)

	zuoraAPI := zuora.NewAPI(httpClient, zuoraOAuthHeaderProvider, zuoraURL)

	t, err := zuoraAPI.V1.SubscriptionsService.ByKey(ctx, "A-S00019063")
	if err != nil {
		log.Fatalf("an error: %v", err)
	}

	my := mySubscription{}
	if err := json.Unmarshal(t, &my); err != nil {
		log.Fatalf("could not unmarshal json: %v", err)
	}

	fmt.Println((my.(&AppVersion))
}

func newHTTPClient() *http.Client {
	keepAliveTimeout := 600 * time.Second
	timeout := 10 * time.Second
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
