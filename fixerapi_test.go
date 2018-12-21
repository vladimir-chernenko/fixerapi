package fixerapi

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestFixerClient(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey, ok := os.LookupEnv("FIXER_API_KEY")

	if !ok {
		t.Errorf("FIXER_API_KEY must be set for testing")
	}

	fc := NewFixerClient(apiKey)

	cr, err := fc.GetCurrencyRates()

	if cr.Base != "EUR" {
		t.Errorf("Expected base currency EUR, got %s", cr.Base)
	}

	if err != nil {
		t.Error(err)
	}
}
