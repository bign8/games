#! /bin/sh

trap 'kill $(jobs -p)' EXIT

gin -a 8000 &
gopherjs serve -m &

wait
