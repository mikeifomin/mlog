package main

import (
	"github.com/mikeifomin/mlog/server"
	"os"
	"strconv"
)

func main() {
	s := server.Server{
		Bind:     ":8092",
		Access:   map[string]string{"mike": "q"},
		Tokens:   []string{"tokenA"},
		PgHost:   mustEnv("PG_HOST"),
		PgUser:   mustEnv("PG_USER"),
		PgPass:   mustEnv("PG_PASS"),
		PgName:   mustEnv("PG_NAME"),
		PgPort:   mustEnvUInt16("PG_PORT"),
		AdminDir: mustEnv("ADMIN_DIR"),
	}

	s.Run()
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("env " + key + " must be set")
	}
	return value
}

func mustEnvUInt16(key string) uint16 {
	valueStr := os.Getenv(key)

	value, _ := strconv.ParseInt(valueStr, 10, 16)
	if value == 0 {
		panic("env " + key + " must be set")
	}
	return uint16(value)
}
