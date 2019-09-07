package zuora

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthHeaders(t *testing.T) {
	ctx := context.Background()
	authHeaders := NewBasicAuthHeader("testClientID", "testClientSecret")
	want := "Basic dGVzdENsaWVudElEOnRlc3RDbGllbnRTZWNyZXQ="
	got, err := authHeaders.AuthHeaders(ctx)
	if err != nil {
		t.Errorf("AuthHeaders() return an error: %v", err)
	}

	if got != want {
		t.Errorf("AuthHeaders() = %q, want %q", got, want)
	}
}

func TestOAuthHeadersWhenTokenValid(t *testing.T) {
	ctx := context.Background()
	mockServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		t.Logf("Sending fake request to %v", req.URL)
		rw.WriteHeader(200)
		rw.Write([]byte(`{"access_token": "reallylongaccesstoken", "token_type": "fake", "expires_in": 100, "scope": "all", "jti": ""}`))
	}))
	defer mockServer.Close()

	oauthHeader := NewOAuthHeader(mockServer.Client(), &MemoryTokenStore{}, "testClientID", "testClientSecret", mockServer.URL)
	want := "Bearer reallylongaccesstoken"
	got, err := oauthHeader.AuthHeaders(ctx)

	if err != nil {
		t.Errorf("OAuthHeader.AuthHeaders() returned an error: %v", err)
	}

	if got != want {
		t.Errorf("OAuthHeader.AuthHeaders() = %q, want %q", got, want)
	}
}

func TestOAuthHeadersWhenTokenAlreadyCached(t *testing.T) {
	ctx := context.Background()
	mockServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		t.Logf("Sending fake request to %v", req.URL)
		rw.WriteHeader(200)
		rw.Write([]byte(`{"access_token": "reallylongaccesstoken", "token_type": "fake", "expires_in": 200, "scope": "all", "jti": ""}`))
	}))

	oauthHeader := NewOAuthHeader(mockServer.Client(), &MemoryTokenStore{}, "testClientID", "testClientSecret", mockServer.URL)
	want := "Bearer reallylongaccesstoken"
	got, err := oauthHeader.AuthHeaders(ctx)

	if err != nil {
		t.Errorf("OAuthHeader.AuthHeaders() returned an error: %v", err)
	}

	if got != want {
		t.Errorf("OAuthHeader.AuthHeaders() = %q, want %q", got, want)
	}

	mockServer.Close() //Closing to make sure that it does not make another HTTP request

	again, err := oauthHeader.AuthHeaders(ctx)

	if err != nil {
		t.Errorf("OAuthHeader.AuthHeaders() returned an error: %v", err)
	}

	if again != want {
		t.Errorf("OAuthHeader.AuthHeaders() = %q, want %q", again, want)
	}
}

func TestOAuthHeadersWhenTokenFails(t *testing.T) {
	ctx := context.Background()
	mockServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		t.Logf("Sending fake request to %v", req.URL)
		rw.WriteHeader(401)
		rw.Write([]byte(`{"something": "happened"}`))
	}))
	defer mockServer.Close()

	oauthHeader := NewOAuthHeader(mockServer.Client(), &MemoryTokenStore{}, "testClientID", "testClientSecret", mockServer.URL)
	want := `error while trying to read body response into memory: <nil> - data: {"something": "happened"}`
	g, err := oauthHeader.AuthHeaders(ctx)

	t.Log(g, err)
	if err.Error() != want {
		t.Errorf("OAuthHeader.AuthHeaders() wanted an error: %v but got: %v", want, err)
	}
}
