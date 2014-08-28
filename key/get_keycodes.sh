#/bin/bash

#  Copyright 2014, Raphael Estrada
#  Author email:  <galaktor@gmx.de>
#  Project home:  <https://github.com/galaktor/gostwriter>
#  Licensed under The GPL v3 License (see README and LICENSE files)

inputh="/usr/include/linux/input.h"

if [[ $# -ne 0 ]]
then
  inputh=$1
fi

echo "using KEYS from input.h at: $inputh"
rm ./codes.go
cp ./codes.template ./codes.go
codes=$(cat $inputh | grep -e KEY_ | sed "s/KEY_//" | awk '{printf("    CODE_%-21s \\\= Code\\\(C\\\.KEY_%-23s\\\)  \\\/\\\* %-5s \\\*\\\/\\n", $2, $2, $3)}')  # have to escape it all for sed...madness
#echo "$codes"
# WHY WON'T SED ACCEPT WHAT IS IN $CODES??
sed -i "s/\/\*KEYCODES\*\//$codes/" codes.go
gofmt -w codes.go
