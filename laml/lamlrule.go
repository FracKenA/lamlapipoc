package laml

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

func Create(
	config ConfigTree,
	reqBody ESRuleTree,
) (
	statusCode int,
	statusText string,
	id string,
	result string,
	err error,
) {
	log.Println("Create Rule...")

	server := BuildString(":", config.Server, config.Port)

	if config.Secure == true {
		server = BuildString("/", "https:/", server, ".intelligence_rules", "doc")
	} else {
		server = BuildString("/", "http:/", server, ".intelligence_rules", "doc")
	}

	// This doesn't make any sense right now since the API is 1-to-1 to ES, but...
	payload, err := json.Marshal(reqBody)
	if err != nil {
		return -1, "", "", "", err
	}

	req, err := http.NewRequest("POST", server, bytes.NewBuffer(payload))
	if err != nil {
		return -1, "", "", "", err
	}

	req.Header.Add("content-type", BuildString("/", "application", config.ContentType))
	req.Header.Add("accept", BuildString("/", "application", config.Accept))
	req.SetBasicAuth(config.User, config.Pass)

	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Println("Error dumping request.")
		return -1, "", "", "", err
	}
	log.Printf("Request Dump:\n%s\n", reqDump)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error sending request to ES, %s\n.", err)
		return -1, "", "", "", err
	}
	defer res.Body.Close()

	resDump, err := httputil.DumpResponse(res, true)
	if err != nil {
		return -1, "", "", "", err
	}
	log.Printf("Response Dump:\n%q\n", resDump)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return -1, "", "", "", err
	}

	bodyJSON := ESRuleCreatedTree{}
	err = json.Unmarshal(body, &bodyJSON)
	if err != nil {
		return -1, "", "", "", err
	}

	return res.StatusCode, res.Status, bodyJSON.ID, bodyJSON.Result, nil
}

func Delete(
	config ConfigTree,
	resource string,
) (
	statusCode int,
	statusText string,
	id string,
	result string,
	err error,
) {
	log.Println("Deleting Rule...")

	server := BuildString(":", config.Server, config.Port)

	if config.Secure == true {
		server = BuildString(
			"/",
			"https:/",
			server,
			".intelligence_rules",
			"doc",
			resource,
		)
	} else {
		server = BuildString(
			"/",
			"http:/",
			server,
			".intelligence_rules",
			"doc",
			resource,
		)
	}

	req, err := http.NewRequest("DELETE", server, nil)
	if err != nil {
		return -1, "", "", "", err
	}

	req.Header.Add("accept", BuildString("/", "application", config.Accept))
	req.SetBasicAuth(config.User, config.Pass)

	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Println("Error dumping request.")
		return -1, "", "", "", err
	}
	log.Printf("Request dump:\n%s\n", reqDump)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error sending request to ES, %s\n", err)
		return -1, "", "", "", err
	}
	defer res.Body.Close()

	resDump, err := httputil.DumpResponse(res, true)
	if err != nil {
		log.Println("Error dumping response.")
		return -1, "", "", "", err
	}
	log.Printf("Response dump:\n%s\n", resDump)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response body, %s\n", body)
	}

	bodyJSON := ESRuleCreatedTree{}
	err = json.Unmarshal(body, &bodyJSON)
	if err != nil {
		log.Printf("Error converting response body to JSON, %s\n", err)
		return -1, "", "", "", err
	}

	return res.StatusCode, res.Status, bodyJSON.ID, bodyJSON.Result, nil
}
