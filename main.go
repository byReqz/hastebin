package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mattn/go-isatty"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	url := os.Getenv("HASTEBIN_SERVER_URL")
	if url == "" {
		url = "https://www.toptal.com/developers/hastebin"
	}
	if isatty.IsTerminal(os.Stdin.Fd()) {
		fmt.Println("Running interactively. Press Ctrl-d to submit.")
		fmt.Println("")
	}
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(stdin)
	if reader.Len() == 0 {
		log.Fatal("hastebin: Empty stdin, exiting")
	}
	req, err := http.NewRequest("POST", url+"/documents", reader)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("User-Agent", "hastebin")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response map[string]string
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url + "/" + response["key"])
}
