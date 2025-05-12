// internal/auth/auth.go
package auth

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

const (
	SupabaseURL    = "https://your-project-id.supabase.co" // Reemplaza con tu URL de Supabase
	SupabaseAPIKey = "your-supabase-api-key"               // Reemplaza con tu API key
)

func LoginUser(email, password string) (string, error) {
	client := resty.New()

	// Realizamos la solicitud de login
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("apikey", SupabaseAPIKey).
		SetBody(fmt.Sprintf(`{"email":"%s", "password":"%s"}`, email, password)).
		Post(SupabaseURL + "/auth/v1/token")

	if err != nil {
		return "", err
	}

	// Si todo va bien, retornamos el JWT
	return resp.String(), nil
}
