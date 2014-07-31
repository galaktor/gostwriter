#!/bin/bash

curl https://raw.githubusercontent.com/torvalds/linux/master/include/uapi/linux/input.h > /tmp/input.h
./get_keycodes.sh /tmp/input.h
