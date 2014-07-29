#/bin/bash

curl -s https://raw.githubusercontent.com/torvalds/linux/master/include/uapi/linux/input.h | grep -e KEY_ | awk '{printf("    %-21s = C.%-23s  /* %-5s */\n", $2, $2, $3)}' 
