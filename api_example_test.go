package zuora_test

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hyeomans/zuora"
)

//ExampleOAuthHeader how to initialize the Zuora client with OAuth header provider.
func ExampleOAuthHeader() {
	ctx := context.Background()

	httpClient := http.DefaultClient

	oauthHeaderProvider := zuora.NewOAuthHeader(httpClient, &zuora.MemoryTokenStore{}, "clientID", "clientSecret", "https://apisandbox.zuora.com/")

	zuoraAPI := zuora.NewAPI(httpClient, oauthHeaderProvider, "https://apisandbox.zuora.com/")
	fmt.Println(zuoraAPI.V1.AccountsService.Summary(ctx, "ObjectID"))
}

//ExampleBasicAuthHeader how to initialize the Zuora client with Basic header provider.
func ExampleBasicAuthHeader() {
	ctx := context.Background()
	httpClient := http.DefaultClient

	basicAuthHeader := zuora.NewBasicAuthHeader("yourClientID", "yourClientSecret")

	zuoraAPI := zuora.NewAPI(httpClient, basicAuthHeader, "https://apisandbox.zuora.com/")

	fmt.Println(zuoraAPI.V1.AccountsService.Summary(ctx, "ObjectID"))
}
