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
# ./plex-webhook -h
Usage of ./plex-webhook:
  -command string
    	path to the command that is execd upon each event (default "./event.sh")
  -listen string
    	address to listen on (default "127.0.0.1")
  -port string
    	port to listen on (default "8080")
  -timeout uint
    	amount of time in seconds to allow the command to run (default 5)
```

## Example script that I use to dim lights
```
#!/usr/bin/env bash

starttime=203000 #8:30pm
endtime=080000 #8:00am

function changelights {
        local timestamp=$(TZ=America/New_York date +"%H%M%S")
        if [[ $timestamp > $starttime || $timestamp < $endtime ]]
        then
                return 0
        fi
        return 1
}

[[ $PLEX_SERVER == "papertigers" ]] || exit 0
[[ $PLEX_USER == "papertigers" ]] || exit 0
[[ $PLEX_PLAYER == "Living Room" ]] || exit 0


if [[ $PLEX_EVENT == "media.play" || $PLEX_EVENT == "media.resume" ]]; then
        changelights || exit 0
        # turn lights on and make them dim
        /opt/local/bin/ghue -c /opt/local/etc/ghue/config.json lights state 7 3 6 18  --on true --bri 50
        exit $?
fi

if [[ $PLEX_EVENT == "media.pause"  || $PLEX_EVENT == "media.stop" ]]
then
        changelights || exit 0
        # turn lights on and make them bright
        /opt/local/bin/ghue -c /opt/local/etc/ghue/config.json lights state 7 3 6 18  --on true --bri 254
        exit $?
fi
```

[Plex documentation on using webhooks](https://support.plex.tv/hc/en-us/articles/115002267687-Webhooks)
