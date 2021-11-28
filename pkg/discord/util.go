package discord

import "os"

func getEnv(keys ...string) string {
	for _, key := range keys {
		if len(key) == 0 {
			continue
		}

		if val := os.Getenv(key); len(val) > 0 {
			return val
		}
	}

	return ""
}
