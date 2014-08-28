#!/bin/bash

#  Copyright 2014, Raphael Estrada
#  Author email:  <galaktor@gmx.de>
#  Project home:  <https://github.com/galaktor/gostwriter>
#  Licensed under The GPL v3 License (see README and LICENSE files)

curl https://raw.githubusercontent.com/torvalds/linux/master/include/uapi/linux/input.h > /tmp/input.h
./get_keycodes.sh /tmp/input.h
