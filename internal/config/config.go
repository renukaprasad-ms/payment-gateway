package config

import (
	"bufio"
	"net"
	"net/url"
	"os"
	"strings"
)

type Config struct {
	Port  string
	DBUrl string
}

func LoadConfig() Config {
	loadDotEnv(".env")

	return Config{
		Port:  getEnv("PORT", "8080"),
		DBUrl: getDBURL(),
	}
}

func getEnv(key, fallback string) string {

	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}

func loadDotEnv(path string) {
	file, err := os.Open(path)

	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, ok := strings.Cut(line, "=")

		if !ok {
			continue
		}

		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		if key == "" {
			continue
		}

		if _, exists := os.LookupEnv(key); exists {
			continue
		}

		_ = os.Setenv(key, value)
	}
}

func getDBURL() string {
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		return dbURL
	}

	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "password")
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	name := getEnv("DB_NAME", "payments")
	sslMode := getEnv("DB_SSLMODE", "disable")

	dbURL := &url.URL{
		Scheme: "postgres",
		Host:   net.JoinHostPort(host, port),
		Path:   name,
	}

	if password == "" {
		dbURL.User = url.User(user)
	} else {
		dbURL.User = url.UserPassword(user, password)
	}

	query := dbURL.Query()
	query.Set("sslmode", sslMode)
	dbURL.RawQuery = query.Encode()

	return dbURL.String()
}
