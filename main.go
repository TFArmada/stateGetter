package main

import (
	"context"
	"errors"
	"flag"
	tfe "github.com/hashicorp/go-tfe"
	"io"
	"log"
	"net/http"
	"os"
)

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	org := flag.String("organization", "", "TFE Organization (required)")
	workspace := flag.String("workspace", "", "TFE Workspace (required)")
	filename := flag.String("filename", "stateGetter.tfstate", "Output file name")
	flag.Parse()

	required := []string{"organization", "workspace"}
	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			log.Fatal("missing required argument/flag; use -h for help\n")
		}
	}

	token := os.Getenv("TFE_TOKEN")
	if token == "" {
		log.Fatal("missing TFE_TOKEN env variable")
	}
	config := &tfe.Config{
		Token: token,
	}

	client, err := tfe.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	state, err := client.StateVersions.List(context.Background(),tfe.StateVersionListOptions{
		ListOptions:  tfe.ListOptions{},
		Organization: org,
		Workspace:    workspace,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = downloadFile(state.Items[0].DownloadURL, *filename)
	if err != nil {
		log.Fatal(err)
	}


}
