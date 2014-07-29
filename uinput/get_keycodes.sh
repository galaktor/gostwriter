#/bin/bash

inputh="/usr/include/linux/input.h"

if [[ $# -ne 0 ]]
then
  inputh=$1
fi

echo "using KEYS from input.h at: $inputh"

cat ../COPYRIGHT > keycodes.go
echo "" >> keycodes.go
echo "package uinput" >> keycodes.go
echo "" >> keycodes.go
echo "/*" >> keycodes.go
echo "  #include <linux/input.h>" >> keycodes.go
echo "*/" >> keycodes.go
echo "import \"C\"" >> keycodes.go
echo "" >> keycodes.go
echo "" >> keycodes.go
echo "type KeyCode C.__u16" >> keycodes.go
echo "" >> keycodes.go
echo "const(" >> keycodes.go
cat $inputh | grep -e KEY_ | awk '{printf("    %-21s = KeyCode(C.%-23s)  /* %-5s */\n", $2, $2, $3)}' >> keycodes.go
echo ")" >> keycodes.go
gofmt -w keycodes.go
