#/bin/bash

curl -s https://raw.githubusercontent.com/torvalds/linux/master/include/uapi/linux/input.h

inputh="/tmp/input.h"

if [[ $# -eq 0 ]]
then
  sudo curl -s https://raw.githubusercontent.com/torvalds/linux/master/include/uapi/linux/input.h > $inputh
else
  sudo cp $1 $inputh
fi

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
