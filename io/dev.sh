#! /bin/sh

trap 'kill $(jobs -p)' EXIT

gopherjs build -w &
gin -i -a 8000 &

wait
