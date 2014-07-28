#/bin/bash

curl -s https://raw.githubusercontent.com/torvalds/linux/master/include/uapi/linux/input.h | grep -e KEY_ | awk '{print $2}'
