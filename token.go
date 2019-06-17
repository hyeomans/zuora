package zuora

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/url"
// 	"strconv"
// 	"strings"

// 	"github.com/hyeomans/zuora/errors"
// )

// type TokenService struct {
// 	config       *Config
// 	errorHandler errors.RequestHandler
// }

// func newTokenService(config *Config, errorHandler errors.RequestHandler) *TokenService {
// 	return &TokenService{
// 		config:       config,
// 		errorHandler: errorHandler,
// 	}
// }

// //Token Returns a token from TokenStore or from calling Zuora http endpoint.
// //Possible return error BadRequest or InvalidResponse
// func (t *TokenService) Token(ctx context.Context) (*Token, error) {
// 	if valid, cachedToken := t.config.tokenStore.Token(); valid {
// 		return cachedToken, nil
// 	}

// 	token, err := t.create(ctx)

// 	if err != nil {
// 		return nil, err
// 	}

// 	t.config.tokenStore.Update(token)
// 	return token, nil
// }

// func (t *TokenService) create(ctx context.Context) (*Token, error) {
// 	tokenURL := fmt.Sprint(t.config.BaseURL, "/oauth/token")
// 	data := url.Values{}
// 	data.Add("grant_type", "client_credentials")
// 	data.Add("client_id", t.config.ClientID)
// 	data.Add("client_secret", t.config.ClientSecret)

// 	values := strings.NewReader(data.Encode())

// 	req, err := http.NewRequest(http.MethodPost, tokenURL, values)
// 	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
// 	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
// 	req.WithContext(ctx)

// 	if err != nil {
// 		return nil, t.errorHandler.BadRequest(err)
// 	}

// 	res, err := t.config.HTTPClient.Do(req)

// 	if res != nil {
// 		defer res.Body.Close()
// 	}

// 	if err != nil {
// 		return nil, t.errorHandler.BadRequest(err)
// 	}

// 	body, err := ioutil.ReadAll(res.Body)

// 	if res.StatusCode != http.StatusOK {
// 		return nil, t.errorHandler.InvalidResponse(body, res.StatusCode)
// 	}

// 	TokenJSON := &Token{}

// 	if err := json.Unmarshal(body, TokenJSON); err != nil {
// 		return nil, t.errorHandler.BadRequest(err)
// 	}

// 	return TokenJSON, nil
// }
