#!/usr/bin/env bash

json=$(cat)


echo "Event: $PLEX_EVENT"
echo "User: $PLEX_USER"
echo "Server: $PLEX_SERVER"
echo "Player: $PLEX_PLAYER"

echo

echo "Raw input:"
echo "$json"
