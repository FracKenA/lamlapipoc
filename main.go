package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"lamlapipoc/laml"
)

// Responses to requests for monitoring reasons. We're a monitoring company, and
// we should build the tools for easy monitoring into our products. :)
func APIPing(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ret_val := map[string]interface{}{
		"result":      "OK",
		"version":     vars["version"],
		"time":        time.Now().String(),
		"statuscode":  0,
		"description": "It's alive!",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ret_val)
}

func APIRoot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("APIRoot")
	ret_val := map[string]interface{}{
		"result":      "OK",
		"version":     vars["version"],
		"time":        time.Now().String(),
		"statuscode":  0,
		"description": "API Root",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ret_val)
}

func APIRule(w http.ResponseWriter, r *http.Request) {
	var err error

	vars := mux.Vars(r)
	version := r.URL.Query().Get("version")
	log.Println("APIRule")
	log.Println(vars)

	// TODO: Create a config file to manage this.
	config := laml.ConfigTree{
		Server:      "server.domain.tld",
		Port:        "8888",
		Secure:      false,
		User:        "accountname",
		Pass:        "password",
		ContentType: "json",
		Accept:      "json",
	}

	ret_val := map[string]interface{}{
		"result":      "OK",
		"version":     version,
		"time":        time.Now().String(),
		"statuscode":  0,
		"description": "API Rule",
		"es_id":       "",
		"es_result":   "",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	switch r.Method {
	case "GET":
		log.Printf("Method: %s\n", r.Method)
	case "DELETE":
		log.Printf("method: %s\n", r.Method)
		resource := vars["resource"]
		ret_val["statuscode"], ret_val["result"], ret_val["es_id"], ret_val["es_result"], err = laml.Delete(config, resource)
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Unable to parse body, %s\n", err)
			http.Error(w, "Unable to parse body.", http.StatusBadRequest)
			return
		}

		if vars["resource"] == "create" {
			bodyJSON := laml.ESRuleTree{}
			err = json.Unmarshal(body, &bodyJSON)
			if err != nil {
				log.Printf("Unable to parse body, %s\n", err)
				http.Error(w, "Unable to parse body.", http.StatusBadRequest)
				return
			}
			ret_val["statuscode"], ret_val["result"], ret_val["es_id"], ret_val["es_result"], err = laml.Create(config, bodyJSON)
		}
	default:
		log.Printf("Method: %s\n not implemented.", r.Method)
	}

	jsonRetPayload, err := json.Marshal(ret_val)
	if err != nil {
		panic(err)
	}

	w.Write(jsonRetPayload)
}

// Deals with the UI stuff, if any.
func UIRoot(w http.ResponseWriter, r *http.Request) {
	ret_val := fmt.Sprintf("Hello. LA ML API PoC UI")
	fmt.Fprintf(w, ret_val)
}

func main() {
	router := mux.NewRouter()
	router.Schemes("https")
	router.HandleFunc("/", APIRoot).Methods("GET")
	router.HandleFunc("/ping", APIPing).Methods("GET")
	router.HandleFunc("/rule/{resource}", APIRule).Methods(
		"GET",
		"POST",
		"PUT",
		"PATCH",
		"DELETE",
		"OPTIONS",
		"HEAD",
	)
	routerAPI := router.PathPrefix("/ui").Subrouter()
	routerAPI.HandleFunc("", UIRoot).Methods("GET")
	log.Fatal(http.ListenAndServeTLS(":1820", "cert.pem", "key.pem", router))
}
