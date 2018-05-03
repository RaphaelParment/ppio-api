package routes

import (
	"net/http"
	_ "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"github.com/coreos/go-oidc"
	"encoding/json"
	"context"
	"log"
)



func login() http.HandlerFunc {

	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "http://10.1.253.208:54117/auth/realms/macq")
	if err != nil {
		log.Fatal(err)
	}

	config := oauth2.Config{
		ClientID:     "macq",
		ClientSecret: "7793fdd3-b852-4f27-a598-149c0fa732c8",
		RedirectURL:  "http://localhost:9001/ppio/login/callback",

		// Discovery returns the OAuth2 endpoints.
		Endpoint: provider.Endpoint(),

		// "openid" is a required scope for OpenID Connect flows.
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}


	state := "123456"



	fn := func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, config.AuthCodeURL(state), http.StatusFound)

	}

	return http.HandlerFunc(fn)
}

func loginCallback() http.HandlerFunc {

	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "http://10.1.253.208:54117/auth/realms/macq")
	if err != nil {
		log.Fatal(err)
	}

	config := oauth2.Config{
		ClientID:     "macq",
		ClientSecret: "7793fdd3-b852-4f27-a598-149c0fa732c8",
		RedirectURL:  "http://localhost:9001/ppio/login/callback",

		// Discovery returns the OAuth2 endpoints.
		Endpoint: provider.Endpoint(),

		// "openid" is a required scope for OpenID Connect flows.
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}

	oidcConfig := &oidc.Config{
		ClientID: "macq",
	}

	state := "123456"

	verifier := provider.Verifier(oidcConfig)

	fn := func(w http.ResponseWriter, r *http.Request) {



		if r.URL.Query().Get("state") != state {
			http.Error(w, "state did not match", http.StatusBadRequest)
			return
		}

		oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
			return
		}
		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
			return
		}


		oauth2Token.AccessToken = "*REDACTED*"


		resp := struct {
			OAuth2Token   *oauth2.Token
			IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
		}{oauth2Token, new(json.RawMessage)}

		if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}


		data, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)

	}

	return http.HandlerFunc(fn)
}


func logout() http.HandlerFunc {

	fn := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "http://10.1.253.208:54117/auth/realms/macq/protocol/openid-connect/logout?redirect_uri=http://localhost:9001/ppio/login", http.StatusFound)
	}

	return http.HandlerFunc(fn)
}