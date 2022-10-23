package config

import "os"

func Config(k string) string {
	return os.Getenv(k)
}
