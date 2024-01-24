package env

import "os"

type Env struct {
	DatabaseURL        string
	Port               string
	ApplicationBaseUrl string
}

func GetEnv() Env {

	databaseUrl := os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		panic("DATABASE_URL environment variable is not set")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	applicationBaseUrl := os.Getenv("APPLICATION_BASE_URL")

	if applicationBaseUrl == "" {
		applicationBaseUrl = "http://localhost:" + port
	}

	return Env{
		DatabaseURL:        databaseUrl,
		Port:               port,
		ApplicationBaseUrl: applicationBaseUrl,
	}
}
