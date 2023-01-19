package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	SupabaseProjectRef string
	SupabaseAPIKey     string
	Datasource         string
	JWTSecret          string
}

var env Env

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("WARNING : No .env file")
	}

	env.SupabaseProjectRef = os.Getenv("SUPABASE_PROJECT_REF")
	env.SupabaseAPIKey = os.Getenv("SUPABASE_API_KEY")
	env.Datasource = os.Getenv("DATASOURCE")
	env.JWTSecret = os.Getenv("JWT_SECRET")

	if env.SupabaseProjectRef == "" || env.SupabaseAPIKey == "" || env.Datasource == "" || env.JWTSecret == "" {
		log.Fatal("Environment variables are incomplete")
	}
}

func Get() Env {
	return env
}
