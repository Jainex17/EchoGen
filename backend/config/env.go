package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ElevenAPIKey       string
	GeminiAPIKey       string
	GoogleClientID     string
	GoogleClientSecret string
	FrontendURL        string
	BackendURL         string
	JwtSecret          []byte
	CookieSecure       bool
	VercelBlobSroreId  string
	VercelStoreUrl     string

	DatabaseURL string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ElevenAPIKey = os.Getenv("ELEVEN_API_KEY")
	if ElevenAPIKey == "" {
		log.Fatal("ELEVEN_API_KEY is not set")
	}

	GeminiAPIKey = os.Getenv("GEMINI_API_KEY")
	if GeminiAPIKey == "" {
		log.Fatal("GEMINI_API_KEY is not set")
	}

	GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	if GoogleClientID == "" {
		log.Fatal("GOOGLE_CLIENT_ID is not set")
	}

	GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	if GoogleClientSecret == "" {
		log.Fatal("GOOGLE_CLIENT_SECRET is not set")
	}

	FrontendURL = os.Getenv("FRONTEND_URL")
	if FrontendURL == "" {
		log.Fatal("FRONTEND_URL is not set")
	}

	BackendURL = os.Getenv("BACKEND_URL")
	if BackendURL == "" {
		log.Fatal("BACKEND_URL is not set")
	}

	JwtSecret = []byte(os.Getenv("JWT_SECRET"))
	if len(JwtSecret) == 0 {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	DatabaseURL = os.Getenv("DatabaseURL")
	if DatabaseURL == "" {
		log.Fatal("DatabaseURL is not set")
	}

	CookieSecure = os.Getenv("COOKIE_SECURE") == "true"

	VercelBlobSroreId = os.Getenv("BLOB_READ_WRITE_TOKEN")
	if VercelBlobSroreId == "" {
		log.Fatal("BLOB_READ_WRITE_TOKEN is not set")
	}

	fmt.Println("Environment variables loaded successfully")
}
