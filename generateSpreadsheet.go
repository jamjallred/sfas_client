package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func (cfg *Config) sendSpreadsheetRequest(filepath, saveAs string) error {
	if _, err := os.Stat(filepath); err != nil {
		return err
	}

	if len(filepath) <= 5 || filepath[len(filepath)-5:] != ".xlsx" {
		return errors.New("file must be an .xlsx")
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	url := "http://100.113.55.39:52431/api/create_sheet"
	resp, err := http.Post(url, "application/octet-stream", f)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	cfg.Log("File successfully sent.")

	return cfg.decodeSpreadsheetRequest(resp, filepath, saveAs)

}

func (cfg *Config) decodeSpreadsheetRequest(resp *http.Response, filepath, saveAs string) error {

	filepathWords := strings.Split(filepath, "/")
	filepathWords[len(filepathWords)-1] = saveAs + ".xlsx"
	filepath = strings.Join(filepathWords, "/")

	f, err := os.Create(filepath)
	if err != nil {
		cfg.Log(fmt.Sprintf("error creating save file location: %s", err))
		return err
	}
	defer f.Close()

	if _, err := io.Copy(f, resp.Body); err != nil {
		cfg.Log(fmt.Sprintf("error copying file contents: %s", err))
		return err
	}

	cfg.Log("Successfully saved contents")

	return nil
}
