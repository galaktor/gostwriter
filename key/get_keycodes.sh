#/bin/bash

inputh="/usr/include/linux/input.h"

if [[ $# -ne 0 ]]
then
  inputh=$1
fi

echo "using KEYS from input.h at: $inputh"
rm ./keycodes.go
cp ./keycodes.go.template ./keycodes.go
codes=$(cat $inputh | grep -e KEY_ | sed "s/KEY_//" | awk '{printf("    CODE_%-21s \\\= Code\\\(C\\\.KEY_%-23s\\\)  \\\/\\\* %-5s \\\*\\\/\\n", $2, $2, $3)}')  # have to escape it all for sed...madness
#echo "$codes"
# WHY WON'T SED ACCEPT WHAT IS IN $CODES??
sed -i "s/\/\*KEYCODES\*\//$codes/" keycodes.go
gofmt -w keycodes.go
