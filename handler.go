package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Parse entire multipart request
func retrieveMulipartPayload(r *http.Request) ([]byte, error) {
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

func handleEvent(p *PlexPayload) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	cmd := exec.CommandContext(ctx, CommandPath)
	env := os.Environ()
	env = append(env, fmt.Sprintf("PLEX_EVENT=%s", p.Event))
	env = append(env, fmt.Sprintf("PLEX_USER=%s", p.Account.Title))
	env = append(env, fmt.Sprintf("PLEX_SERVER=%s", p.Server.Title))
	env = append(env, fmt.Sprintf("PLEX_PLAYER=%s", p.Player.Title))
	cmd.Env = env
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalf("failed to open stdin: %v", err)
	}
	go func() {
		defer stdin.Close()
		raw, err := json.Marshal(p)
		if err != nil {
			log.Fatal(err)
		}
		io.WriteString(stdin, string(raw))
	}()
	go func() {
		defer cancel()
		if err := cmd.Start(); err != nil {
			log.Fatalf("failed to exec cmd: %v", err)
		}
		if err := cmd.Wait(); err != nil {
			log.Printf("command failed: %v", err)
			return
		}
		log.Println("command exited with 0")
	}()
	<-ctx.Done()
}

// Hook is the handler for the /plex endpoint
func Hook(w http.ResponseWriter, r *http.Request) {
	payload, err := retrieveMulipartPayload(r)
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
	go handleEvent(&data)
}
