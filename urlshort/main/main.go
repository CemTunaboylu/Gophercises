package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"urlshort"

	"github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {

	var yaml_config, json_config, redis_key string
	flag.StringVar(&yaml_config, "yaml", "conf.yaml", "The YAML file path that consists of path, url key value pair list")
	flag.StringVar(&json_config, "json", "conf.json", "The JSON file path that consists of path, url key value pair")
	flag.StringVar(&redis_key, "redis", "", "The redis key to the hash table path:url")
	flag.Parse()

	read_from_DB := false
	mux := defaultMux()
	var file_to_open string
	var marshaller string

	decide_marshaller := func(f *flag.Flag) {
		switch f.Name {
		case "yaml":
			marshaller = "yaml"
			file_to_open = yaml_config
		case "json":
			marshaller = "json"
			file_to_open = json_config
		case "redis":
			read_from_DB = true
		}
	}

	flag.Visit(decide_marshaller)

	var handler http.HandlerFunc

	switch read_from_DB {
	case true:
		urls_map, err := retrieve_map_from_redis(ctx)
		if err != nil {
			panic(err)
		}
		handler = urlshort.MapHandler(urls_map, mux)

	case false:
		bytes, err := file_opener(file_to_open)
		if err != nil {
			panic(err)
		}
		handler, err = urlshort.Handler(bytes, marshaller, mux)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func file_opener(file_name string) ([]byte, error) {
	f, err := ioutil.ReadFile(file_name)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func populate_redis(ctx context.Context) {
	k_val_pair := map[string]interface{}{
		"/urlshort":       "https://github.com/gophercises/urlshort",
		"/urlshort-final": "https://github.com/gophercises/urlshort/tree/solution",
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	intcmd := rdb.HMSet("urls", k_val_pair).Err()
	fmt.Printf("intcmd : %v\n", intcmd)

	val, err := rdb.HGetAll("urls").Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("val : %v\n", val)

}

func retrieve_map_from_redis(ctx context.Context) (map[string]string, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	val, err := rdb.HGetAll("urls").Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}
