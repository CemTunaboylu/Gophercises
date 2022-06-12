package urlshort

import (
	json "encoding/json"
	"errors"
	"net/http"
	"fmt"
	yaml "gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		relative_path := r.URL.Path
		fmt.Printf("Request received : %v", relative_path)
		if redirect, ok := pathsToUrls[relative_path]; ok {
			// give 302
			http.Redirect(w, r, redirect, 302)
		} else {
			fallback.ServeHTTP(w, r)
		}

	}

}

func Handler(bytes []byte, t string, fallback http.Handler) (http.HandlerFunc, error) {
	map_list, err := Unmarshal(bytes, t)
	if err != nil {
		return nil, err
	}

	reduced_map := reduce_to_single_map(map_list)

	return MapHandler(reduced_map, fallback), nil

}

func Unmarshal(b []byte, t string) ([]map[string]string, error) {
	unmarshallers := map[string]func([]byte, any) error{
		"yaml": yaml.Unmarshal,
		"json": json.Unmarshal,
	}
	f, ok := unmarshallers[t]
	if !ok {
		return nil, errors.New("Type does not have an unmarshaller assigned to it")
	}

	var map_list_to_put_into []map[string]string
	err := f(b, &map_list_to_put_into)
	if err != nil {
		return nil, err
	}

	return map_list_to_put_into, nil
}

func reduce_to_single_map(map_list []map[string]string) map[string]string {
	m := map[string]string{}
	for _, pair_map := range map_list {
		m[pair_map["path"]] = pair_map["url"]
	}
	return m
}
