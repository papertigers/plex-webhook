package main

// PlexAccount is the Plex Account associated with the webhook
type PlexAccount struct {
	ID    int    `json:"id"`
	Thumb string `json:"thumb,omitempty"`
	Title string `json:"title"`
}

// PlexServer is the Plex Server associated with the webhook
type PlexServer struct {
	Title string `json:"title"`
	UUID  string `json:"uuid"`
}

// PlexPlayer is the Plex Server associated with the webhook
type PlexPlayer struct {
	Local         bool   `json:"local"`
	PublicAddress string `json:"publicAddress"`
	Title         string `json:"title"`
	UUID          string `json:"uuid"`
}

// PlexPayload maps the json from Plex's POST body
type PlexPayload struct {
	Event   string `json:"event"`
	User    bool   `json:"user"`
	Owner   bool   `json:"owner"`
	Account PlexAccount
	Server  PlexServer
	Player  PlexPlayer
}
