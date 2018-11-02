package zuora

import "net/http"

func newConfig(httpClient *http.Client, baseURL, clientID, clientSecret string, tokenStorer TokenStorer, options []ConfigOption) *Config {
	config := &Config{
		HTTPClient:   httpClient,
		BaseURL:      baseURL,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		tokenStore:   tokenStorer,
	}

	for _, option := range options {
		option(config)
	}

	return config
}

func WithTokenStorer(tokenStorer TokenStorer) ConfigOption {
	return func(config *Config) {
		config.tokenStore = tokenStorer
	}
}
