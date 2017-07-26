#! /bin/sh

trap 'kill $(jobs -p)' EXIT

gopherjs build -v -w &
gin -i -a 8000 &

wait
