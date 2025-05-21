#!/bin/sh
./server &
nginx -g 'daemon off;'
