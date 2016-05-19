#/bin/bash

#  Copyright 2014, Raphael Estrada
#  Author email:  <galaktor@gmx.de>
#  Project home:  <https://github.com/galaktor/gostwriter>
#  Licensed under The GPL v3 License (see README and LICENSE files)

inputh="/usr/include/linux/input-event-codes.h"

if [[ $# -ne 0 ]]
then
  inputh=$1
elif test -f /usr/include/linux/input-event-codes.h
then
  # Linux >= 4.4 factored input-event-codes.h out of input.h
  inputh=/usr/include/linux/input-event-codes.h
elif test -f /usr/include/linux/input.h
then
  inputh=/usr/include/linux/input.h
else
  echo "no input header found; use ./get_keycodes_kerneltip.sh"
  echo "or install /usr/include/linux headers (e.g. debian: linux-libc-dev;"
  echo "arch linux: linux-api-headers)"
  exit 1
fi

echo "using KEYS from file at: $inputh"
rm ./codes.go
cp ./codes.template ./codes.go
codes=$(cat $inputh | grep -e KEY_ | sed "s/KEY_//" | awk '{printf("    CODE_%-21s \\\= Code\\\(C\\\.KEY_%-23s\\\)  \\\/\\\* %-5s \\\*\\\/\\n", $2, $2, $3)}')  # have to escape it all for sed...madness
#echo "$codes"
# WHY WON'T SED ACCEPT WHAT IS IN $CODES??
sed -i "s/\/\*KEYCODES\*\//$codes/" codes.go
gofmt -w codes.go
