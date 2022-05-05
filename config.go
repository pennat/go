package config

import (
	"flag"
)

type Config struct {
	Port         int
	Db_url       string
	Jaeger_url   string
	Sentry_url   string
	Kafka_broker string
	some_app_id  string
	some_app_key string
}

func ReadbyFlag() Config {
	port := flag.Int("port", 8080, "")
	dbURL := flag.String("db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "")
	jaegerURL := flag.String("jaeger_url", "http://jaeger:16686", "")
	sentryURL := flag.String("sentry_url", "http://sentry:9000", "")
	kafka := flag.String("kafka_broker", "kafka:9092", "")
	id := flag.String("some_app_id", "testid", "")
	key := flag.String("some_app_key", "testkey", "")

	flag.Parse()

	/*
		fmt.Println("port", *port)
		fmt.Println("db_url", *dbURL)
		fmt.Println("jaeger_url", *jaegerURL)
		fmt.Println("sentry_url", *sentryURL)
		fmt.Println("kafka_broker", *kafka)
		fmt.Println("some_app_id", *id)
		fmt.Println("some_app_key", *key)

		return Config{}
	*/

	return Config{
		*port,
		*dbURL,
		*jaegerURL,
		*sentryURL,
		*kafka,
		*id,
		*key,
	}
}
