#!/bin/sh

export GIN_MODE=release

/app/ocserv_api migrate

/app/ocserv_api serve

exec "$@"