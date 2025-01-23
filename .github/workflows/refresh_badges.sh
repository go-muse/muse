#!/bin/bash

BADGES=$(curl -sL https://github.com/go-muse/muse | \
  grep -oE '<img src="https?://camo.githubusercontent.com/[^"]+' | \
  sed -e 's/<img src="//')

echo ${BADGES} | xargs -I % curl -sX PURGE %