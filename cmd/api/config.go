package main

type Config struct {
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string
	dbSSLMode  string
	secret     string
	addr       string
	cacheAddr  string
	dbUrl      string
}

func NewConfig() *Config {
	return &Config{
		dbHost:     getEnv("DB_HOST", "localhost"),
		dbPort:     getEnv("DB_PORT", "5432"),
		dbUser:     getEnv("DB_USER", "user"),
		dbPassword: getEnv("DB_PASSWORD", "password"),
		dbName:     getEnv("DB_NAME", "myecommerce"),
		dbSSLMode:  getEnv("DB_SSLMODE", "disable"),
		secret:     getEnv("JWT_SECRET", "secret"),
		addr:       getEnv("API_ADDR", ":8080"),
		cacheAddr:  getEnv("CACHE_ADDR", ""),
		dbUrl:      getEnv("DB_URL", ""),
	}
}
