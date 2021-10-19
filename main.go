package main

import (
	_ "embed"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type ConfigEntry struct {
	From string `yaml:"from"`
	To   string `yaml:"to"`
}

type Config []ConfigEntry

var (
	//go:embed config.yaml
	data   []byte
	config Config
)

func main() {
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	for _, entry := range config {
		tmp := entry
		http.HandleFunc(tmp.From, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, tmp.To, http.StatusTemporaryRedirect)
		})
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
