package passage

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"net/http"
	"strings"
)

type httpResponseError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Extract authentication token from an HTTP request. Note that the following precedence is applied:
//   1. Authorization Header: "Authorization: Bearer XXXX"
//   2. Cookie: "psg_auth_token"
func getAuthTokenFromRequest(r *http.Request) (string, error) {
	// First, try to extract a token from an HTTP header.
	authHeaderFields := strings.Fields(r.Header.Get("Authorization"))
	if len(authHeaderFields) == 2 && authHeaderFields[0] == "Bearer" {
		authToken := authHeaderFields[1]
		return authToken, nil
	}

	// Second, try to extract a token from a cookie.
	authTokenCookie, err := r.Cookie("psg_auth_token")
	if err == nil {
		authToken := authTokenCookie.Value
		return authToken, nil
	}

	// Could not fine an authentication token in either header or cookie. Return an error.
	return "", errors.New("missing authentication token")
}

// decodeRSAPublicKey takes a base-64 encoded RSA public key and parses it into an rsa.PublicKey type
func decodeRSAPublicKey(pubKey string) (*rsa.PublicKey, error) {
	// Decode base-64 public key
	keyBytes, err := base64.RawURLEncoding.DecodeString(pubKey)
	if err != nil {
		return nil, errors.New("public key must be valid base-64")
	}

	// Parse RSA public key from the raw public key
	pemBlock, _ := pem.Decode(keyBytes)
	if pemBlock == nil {
		return nil, errors.New("public key malformed")
	}
	pk, err := x509.ParsePKCS1PublicKey(pemBlock.Bytes)
	if err != nil {
		return nil, errors.New("public key malformed")
	}

	return pk, nil
}
