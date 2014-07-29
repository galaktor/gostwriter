#/bin/bash

cat ../COPYRIGHT > keycodes.go
echo "" >> keycodes.go
echo "package uinput" >> keycodes.go
echo "" >> keycodes.go
echo "/*" >> keycodes.go
echo "  #include <linux/input.h>" >> keycodes.go
echo "*/" >> keycodes.go
echo "import \"C\"" >> keycodes.go
echo "" >> keycodes.go
echo "const(" >> keycodes.go
curl -s https://raw.githubusercontent.com/torvalds/linux/master/include/uapi/linux/input.h | grep -e KEY_ | awk '{printf("    %-21s = C.%-23s  /* %-5s */\n", $2, $2, $3)}' >> keycodes.go
echo ")" >> keycodes.go
gofmt -w keycodes.go
