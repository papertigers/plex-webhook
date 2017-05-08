# plex-webhook

plex-webhook listens for events on the `/plex` endpoint of the server. 
Upon receiving an event, plex-webhook will call the user provided command with the
following environmental variables:

```
PLEX_EVENT
PLEX_USER
PLEX_SERVER
PLEX_PLAYER
```

For more advanced usage plex-webhook will also send the raw json payload to the command over stdin.
An example script is provided in this repository called `event.sh`



**Usage**
```
$ ./plex-webhook -h
Usage of ./plex-webhook:
  -command string
    	path to the command that is execd upon each event (default "./event.sh")
  -listen string
    	address to listen on (default "127.0.0.1")
  -port string
    	port to listen on (default "8080")
```

[Plex documentation on using webhooks](https://support.plex.tv/hc/en-us/articles/115002267687-Webhooks)
