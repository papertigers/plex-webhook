package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"
)

// Parse entire multipart request
func _retrieveMulipartPayload(r *http.Request) ([]byte, error) {
	mediaType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	if !strings.HasPrefix(mediaType, "multipart/") {
		return nil, errors.New("request is not multipart")
	}

	mr := multipart.NewReader(r.Body, params["boundary"])
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if p.FormName() != "payload" {
			continue
		}
		slurp, err := ioutil.ReadAll(p)
		if err != nil {
			return nil, err
		}
		return slurp, nil
	}
	return nil, nil
}

// Hook is the handler for the /plex endpoint
func Hook(w http.ResponseWriter, r *http.Request) {
	payload, err := _retrieveMulipartPayload(r)
	if err != nil {
		log.Printf("Failed to retrieve data from plex payload: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var data PlexPayload
	if err := json.Unmarshal(payload, &data); err != nil {
		log.Println("bad data from plex found")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO: Do something with the event
	if b, err := json.Marshal(data); err == nil {
		log.Printf("%v\n", string(b))
	}
}
