package content

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"internal/config/config"
)

func GetClientes() {
	url := config.GetEnv("SUPABASE_URL") + "/rest/v1/clientes"
	apiKey := config.GetEnv("SUPABASE_ANON_KEY")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creando la solicitud:", err)
	}

	req.Header.Set("apikey", apiKey)
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Accept", "application/json")

	q := req.URL.Query()
	q.Add("select", "*")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error en la solicitud:", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Respuesta del servidor:\n", string(body))
}
