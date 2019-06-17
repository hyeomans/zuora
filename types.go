package zuora

import (
	"net/http"
)

// Doer is a common interface for Http clients
// https://golang.org/pkg/net/http/#Client.Do
type Doer interface {
	Do(request *http.Request) (*http.Response, error)
}

//ConfigOption helper function to modify current Config struct
type ConfigOption func(*Config)

//Config is the base configuration to return ZuoraApi
type Config struct {
	HTTPClient   *http.Client
	BaseURL      string
	ClientID     string
	ClientSecret string
	tokenStore   TokenStorer
}

//Querier One who, or that which, queries actions
type Querier interface {
	Build() string
}

//ContextKey is a helper to define Zuora context values
type ContextKey string

func (c ContextKey) String() string {
	return "zuora context key " + string(c)
}

//ContextKeyZuoraEntityIds will be added as a header on requests
const ContextKeyZuoraEntityIds = ContextKey("Zuora-Entity-Ids")

//ContextKeyZuoraTrackID will be addeed as a header on requests
const ContextKeyZuoraTrackID = ContextKey("Zuora-Track-Id")

//ContextKeyZuoraVersion will be added as header when an endpoint supports it.
const ContextKeyZuoraVersion = ContextKey("zuora-version")

//Response is a generic catch all struct to check if a response was successfull
type Response struct {
	Success bool `json:"success"`
}
