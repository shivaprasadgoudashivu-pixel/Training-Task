package keycloak

import (
	"context"
	"io"
	"keycloak-demo/config"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func tokenEndpoint(cfg config.Config) string {
	return strings.TrimRight(cfg.KeycloakURL, "/") + "/realms/" + cfg.Realm + "/protocol/openid-connect/token"
}

func PostForm(ctx context.Context, endpoint string, form url.Values) (*http.Response, []byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return resp, body, nil
}

func Login(ctx context.Context, cfg config.Config, username, password string) (*http.Response, []byte, error) {
	form := url.Values{}
	form.Set("grant_type", "password")
	form.Set("client_id", cfg.ClientID)
	if cfg.ClientSecret != "" {
		form.Set("client_secret", cfg.ClientSecret)
	}
	form.Set("username", username)
	form.Set("password", password)
	return PostForm(ctx, tokenEndpoint(cfg), form)

}
