package aozorarodokuweb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const AppMetaJsonURL = "https://aozoraroudoku.jp/app/ar_app_meta.json"

type AppMeta struct {
	Version string `json:"version"`
	Url     string `json:"url"`
}

func GetAppMeta() (AppMeta, error) {
	res, err := http.DefaultClient.Get(AppMetaJsonURL)
	if err != nil {
		return AppMeta{}, err
	}
	if res.StatusCode >= 400 {
		return AppMeta{}, fmt.Errorf("bad response status code %d", res.StatusCode)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return AppMeta{}, err
	}
	var meta AppMeta
	if err := json.Unmarshal(body, &meta); err != nil {
		return AppMeta{}, err
	}
	return meta, nil
}

func GetContentsCSV(url string) ([]byte, error) {
	//return []Content{}
	res, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("bad response status code %d", res.StatusCode)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
