package zuora

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// AuthHeaderProvider an interface that defines
// a common get method for different types of Authentication
// for Zuora.
type AuthHeaderProvider interface {
	AuthHeaders(ctx context.Context) (string, error)
}

//TokenStorer handles token renewal with two simple methods.
//Token() returns a boolean to indicate a token is valid and if valid, it will return the active token.
//Update() causes a side-effect to update a token in whichever backing store you choose.
type TokenStorer interface {
	Token() (bool, *Token)
	Update(*Token)
}

//Token represents the OAuth token returned by Zuora.
type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	Jti         string `json:"jti"`
}

// BasicAuthHeader represents a basic HTTP auth header
// holder.
type BasicAuthHeader struct {
	clientID     string
	clientSecret string
}

// NewBasicAuthHeader initialize with clientID & clientSecret from Zuora.
func NewBasicAuthHeader(clientID, clientSecret string) *BasicAuthHeader {
	return &BasicAuthHeader{clientID: clientID, clientSecret: clientID}
}

// AuthHeaders returns a string that will be added to each request going out
// to Zuora
func (t *BasicAuthHeader) AuthHeaders(ctx context.Context) (string, error) {
	toEncode := fmt.Sprintf("%v:%v", t.clientID, t.clientSecret)
	return fmt.Sprintf("Basic %v", b64.StdEncoding.EncodeToString([]byte(toEncode))), nil
}

// OAuthHeader is the holder of information to retrieve an OAuth token from Zuora
type OAuthHeader struct {
	clientID     string
	clientSecret string
	http         Doer
	baseURL      string
	tokenStorer  TokenStorer
}

// NewOAuthHeader initialize OAuthHeader struct with information coming from Zuora.
func NewOAuthHeader(doer Doer, tokenStorer TokenStorer, clientID, clientSecret, baseURL string) *OAuthHeader {
	return &OAuthHeader{
		http:         doer,
		clientID:     clientID,
		clientSecret: clientSecret,
		baseURL:      baseURL,
		tokenStorer:  tokenStorer,
	}
}

// AuthHeaders returns a string that will be added to each request going out
// to Zuora.
func (t *OAuthHeader) AuthHeaders(ctx context.Context) (string, error) {
	isValid, token := t.tokenStorer.Token()

	if isValid {
		return fmt.Sprintf("Bearer %v", token.AccessToken), nil
	}

	tokenURL := fmt.Sprint(t.baseURL, "/oauth/token")
	data := url.Values{}

	data.Add("grant_type", "client_credentials")
	data.Add("client_id", t.clientID)
	data.Add("client_secret", t.clientSecret)

	values := strings.NewReader(data.Encode())

	req, err := http.NewRequest(http.MethodPost, tokenURL, values)

	if err != nil {
		return "", responseError{isTemporary: false, message: fmt.Sprintf("error while trying to create an HTTP request: %v", err)}
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := t.http.Do(req.WithContext(ctx))

	if err != nil {
		return "", responseError{isTemporary: false, message: fmt.Sprintf("error while trying to make request: %v", err)}
	}

	defer res.Body.Close()

	if err != nil {
		return "", responseError{isTemporary: false, message: fmt.Sprintf("error while trying to make request: %v", err)}
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", responseError{isTemporary: false, message: fmt.Sprintf("error while trying to read body response into memory: %v", err)}
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		var isTemporary bool
		if http.StatusRequestTimeout == res.StatusCode ||
			http.StatusTooManyRequests == res.StatusCode ||
			http.StatusInternalServerError == res.StatusCode ||
			http.StatusServiceUnavailable == res.StatusCode {
			isTemporary = true
		}

		return "", responseError{isTemporary: isTemporary, message: fmt.Sprintf("error while trying to read body response into memory: %v - data: %v", err, string(body))}
	}

	jsonResponse := Token{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return "", responseError{isTemporary: false, message: fmt.Sprintf("error while Unmarshal json response. Error: %v. JSON: %v", err, string(body))}
	}

	t.tokenStorer.Update(&jsonResponse)

	return fmt.Sprintf("Bearer %v", jsonResponse.AccessToken), nil
}
