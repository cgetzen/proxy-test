package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

var serveRoot = "assets"
var templateEnvironment = "TEMPLATE_DATA"

type TemplateData struct {
	Data string
}

var Environment = TemplateData{}

func main() {
	fmt.Println("Starting...")
	loadEnv()

	r := mux.NewRouter()
	// Enhancement #2 - Handle redirect configurations
	redirects, ok := loadConfig()
	if ok {
		for source, dest := range redirects {
			fmt.Printf("Adding redirect from /%s to /%s\n", source, dest)
			r.Handle(fmt.Sprintf("/%s", source), http.RedirectHandler(fmt.Sprintf("/%s", dest), http.StatusFound))
		}
	}

	// Enhancement #2 - Add assets without code change
	files, err := os.ReadDir("assets/")
	if err != nil {
		fmt.Printf("Unable to list assets: %s. Ignoring.\n", err.Error())
	}
	for _, file := range files {
		name := file.Name()
		if redirects[name] != nil {
			fmt.Printf("Redirect already set up for path %s. Will not serve file.\n", name)
			continue
		}
		if strings.Contains(name, ".tpl") {
			r.HandleFunc(fmt.Sprintf("/%s", name[:len(name)-4]), serveTemplate(name))
		} else {
			r.HandleFunc(fmt.Sprintf("/%s", name), serveFile(name))
		}
	}

	log.Fatal(http.ListenAndServe(":8000", r))
}

func loadEnv() {
	Environment = TemplateData{
		Data: os.Getenv(templateEnvironment),
	}
}

// Enhancement 3 - Environment Variable Templates
func serveTemplate(name string) func(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Caching %s\n", name[:len(name)-4])

	tmplFile, err := template.ParseFiles(fmt.Sprintf("%s/%s", serveRoot, name))
	if err != nil {
		fmt.Printf("Templating error: %s. Ignoring\n", err.Error())
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tmplFile.ExecuteTemplate(w, name, Environment)
	}

}

func serveFile(name string) func(w http.ResponseWriter, r *http.Request) {
	// serveFile loads the file on boot and caches it.
	fmt.Printf("Caching %s\n", name)
	f, err := os.ReadFile(fmt.Sprintf("%s/%s", serveRoot, name))
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeContent(w, r, name, time.Now(), bytes.NewReader(f))
	}
}

// Enhancement #2 - Handle redirect configurations
func loadConfig() (ret map[string]interface{}, ok bool) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Unable to load config: %s. Ignoring.\n", err.Error())
		return nil, false
	}
	// Check the format of the config file
	ret, ok = viper.Get("redirects").(map[string]interface{})
	if !ok {
		fmt.Printf("Unable to parse config. Ignoring.\n")
	}
	return
}
