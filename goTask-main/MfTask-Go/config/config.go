package config

import (
	"os"
)

type Config struct {
	KeycloakURL  string
	Realm        string
	ClientID     string
	ClientSecret string
	ListenAddr   string
}

func mustGetEnv(name, def string) string {
	v := os.Getenv(name)
	if v == "" {
		return def
	}
	return v
}

func Load() Config {
	return Config{
		KeycloakURL:  mustGetEnv("KEYCLOAK_URL", "keycloak:8084"),
		Realm:        mustGetEnv("KEYCLOAK_REALM", "myrealm"),
		ClientID:     mustGetEnv("KEYCLOAK_CLIENT_ID", "myclient"),
		ClientSecret: mustGetEnv("KEYCLOAK_CLIENT_SECRET", ""),
		ListenAddr:   mustGetEnv("LISTEN_ADDR", ":8081"),
	}
}
