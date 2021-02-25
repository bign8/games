#! /bin/bash -ex

go build -o games -v -ldflags="-w -s" ./cmd/server
gzip -f games
scp games.gz me.bign8.info:/opt/bign8
ssh me.bign8.info -- sudo systemctl restart games
rm -f games.gz
