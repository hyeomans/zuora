[![Build Status](https://travis-ci.com/hyeomans/zuora.svg?branch=master)](https://travis-ci.com/hyeomans/zuora)

A Go client library to consume [Zuora API](https://www.zuora.com/developer/api-reference/).

This is a __WIP__ and has minimal endpoints covered but it is really easy to add new ones.

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
