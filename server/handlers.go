package server

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type AuthorizationResponse struct {
	AccessToken  string `json:"access_token"`
	IDToken      string `json:"id_token"`
	ExpiresAt    int64  `json:"expires_at"`
	RefreshToken string `json:"refresh_token"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type UserClaimsJWT struct {
	AtHash        string `json:"at_hash"`
	Aud           string `json:"aud"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Iat           int    `json:"iat"`
	Iss           string `json:"iss"`
	Name          string `json:"name"`
	Sub           string `json:"sub"`
}

type ApplicationHanlder struct {
	state           string
	oauth2Config    oauth2.Config
	idTokenVerifier *oidc.IDTokenVerifier
}

func (a ApplicationHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		a.handleLogin(w, r)
	case "/callback":
		a.handleCallback(w, r)
	case "/favicon.ico":
		w.Write([]byte{})
	default:
		log.Println("called unknown URL:", r.URL.String())
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte{})
		log.Fatal(http.StatusNotFound)
	}
}

func (a ApplicationHanlder) authorizationMiddleware(r *http.Request) (string, error) {
	var userClaims UserClaimsJWT
	var userId string

	rawIDToken := r.URL.Query().Get("token")
	if rawIDToken == "" {
		return userId, errors.New("missing token")
	}

	idToken, err := a.idTokenVerifier.Verify(r.Context(), rawIDToken)
	if err != nil {
		return userId, err
	}

	if err := idToken.Claims(&userClaims); err != nil {
		return userId, err
	}

	if !userClaims.EmailVerified {
		return userId, errors.New("email not verified")
	}

	userId, err = decodeBase64(userClaims.Sub)
	if err != nil {
		return userId, err
	}

	log.Printf("%+v\n", userClaims)

	return userId, nil
}

func (a ApplicationHanlder) handleExample(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("you're the user with id"))
}

func (a ApplicationHanlder) handleLogin(w http.ResponseWriter, r *http.Request) {
	providerURL := a.oauth2Config.AuthCodeURL(a.state)
	http.Redirect(w, r, providerURL, 307)
}

func (a ApplicationHanlder) handleCallback(w http.ResponseWriter, r *http.Request) {
	var err error
	var token *oauth2.Token

	ctx := context.Background()

	switch r.Method {
	case http.MethodGet: //Oauth2 flow
		if errMsg := r.FormValue("error"); errMsg != "" {
			http.Error(w, errMsg+": "+r.FormValue("error_description"), http.StatusBadRequest)
			return
		}

		code := r.FormValue("code")
		if code == "" {
			http.Error(w, fmt.Sprintf("no code in request: %q", r.Form), http.StatusBadRequest)
			return
		}

		if state := r.FormValue("state"); state != a.state {
			http.Error(w, "state mismatch", http.StatusBadRequest)
			return
		}

		/// CREATE USER RECORD IN MAIN DATABASE HERE??

		token, err = a.oauth2Config.Exchange(ctx, code)
	case http.MethodPost: // Form request from frontend to refresh a token.
		refresh := r.FormValue("refresh_token")
		if refresh == "" {
			http.Error(w, fmt.Sprintf("no refresh_token in request: %q", r.Form), http.StatusBadRequest)
			return
		}

		t := &oauth2.Token{
			RefreshToken: refresh,
			Expiry:       time.Now().Add(-time.Hour),
		}

		token, err = a.oauth2Config.TokenSource(ctx, t).Token()
	default:
		http.Error(w, fmt.Sprintf("method not implemented: %s", r.Method), http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get token: %v", err), http.StatusInternalServerError)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "no id_token in token response", http.StatusInternalServerError)
		return
	}

	if _, err := a.idTokenVerifier.Verify(r.Context(), rawIDToken); err != nil {
		http.Error(w, fmt.Sprintf("failed to verify ID token: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AuthorizationResponse{
		ExpiresAt:    token.Expiry.Unix(),
		IDToken:      rawIDToken,
		RefreshToken: token.RefreshToken,
	})

	log.Println("authentication succeeded:")
	fmt.Println("access_token:", token.AccessToken)
	fmt.Println("refresh_token:", token.RefreshToken)
	fmt.Println("id_token:", rawIDToken)
	fmt.Println("expire_at:", token.Expiry.Unix())

	go func() {
		time.Sleep(300 * time.Millisecond)
		os.Exit(0)
	}()
}

func decodeBase64(encoded string) (string, error) {
	rawDecodedText, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	return string(rawDecodedText), nil
}
