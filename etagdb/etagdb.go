package etagdb

import "net/http"

func CheckEtag(url, etag string) (bool, error) {
	resp, err := http.Head(url)
	if err != nil {
		return false, nil
	}
	if etag == resp.Header.Get("Etag") {
		return true, nil
	}
	return false, nil
}
