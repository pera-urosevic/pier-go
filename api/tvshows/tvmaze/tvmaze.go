package tvmaze

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const site = "http://api.tvmaze.com"

func Search(title string) (string, error) {
	var results string
	url := fmt.Sprintf("%s/search/shows?q=%s", site, url.QueryEscape(title))
	res, err := http.Get(url)
	if err != nil {
		return results, err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return results, err
	}
	results = string(bytes)
	return results, nil
}

func Get(id int64) (string, error) {
	var result string
	url := fmt.Sprintf("%s/shows/%d?embed=episodes", site, id)
	res, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}
	result = string(bytes)
	return result, nil
}
