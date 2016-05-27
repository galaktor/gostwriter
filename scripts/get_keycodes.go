// To run: in key/, ./get_keycodes.sh

//  Copyright 2014 Raphael Estrada, 2016 Isaac Dupree
//  Author email:  <galaktor@gmx.de>
//  Project home:  <https://github.com/galaktor/gostwriter>
//  Licensed under The GPL v3 License (see README and LICENSE files)

package main

import (
	"fmt"
	"log"
	"os/exec"
	"io/ioutil"
	"regexp"
	"flag"
	"strings"
)


func main() {
	flag.Parse()
	var inputHeaderName string
	var inputHeader []byte
	var err error
	if flag.NArg() > 0 {
		inputHeaderName = flag.Arg(0)
		inputHeader, err = ioutil.ReadFile(inputHeaderName)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// Linux >= 4.4 factored input-event-codes.h out of input.h
		inputHeaderName = "/usr/include/linux/input-event-codes.h"
		inputHeader, err = ioutil.ReadFile(inputHeaderName)
		if err != nil {
			inputHeaderName = "/usr/include/linux/input.h"
			inputHeader, err = ioutil.ReadFile(inputHeaderName)
			if err != nil {
				log.Fatal("no input header found; use ./get_keycodes_kerneltip.sh\n" +
				"or install /usr/include/linux headers (e.g. debian: linux-libc-dev;\n" +
				"arch linux: linux-api-headers)")
			}
		}
	}
	log.Print("using KEYS from file at: " + inputHeaderName)
	template, err := ioutil.ReadFile("codes.template")
	if err != nil {
		log.Fatalf("%v\n%v", err, "Can't find codes.template (are you in the right directory?)")
	}
	var codes = []string{}
	keyDefine := regexp.MustCompile(`^\s*#\s*define\s+((KEY_|BTN_)(\S*))\s+(\S+)`)
	for _,line := range strings.Split(string(inputHeader), "\n") {
		keyMatch := keyDefine.FindStringSubmatch(line)
		if keyMatch != nil {
			defineName := keyMatch[1]
			prefix := keyMatch[2]
			name := keyMatch[3]
			if prefix == "BTN_" {
				name = defineName
			}
			// value can be the name of another #define,
			// not just a numeric literal
			value := keyMatch[4]
			code := fmt.Sprintf("    CODE_%-21s = Code(C.%-27s)  /* %-5s */", name, defineName, value)
			codes = append(codes, code)
		}
	}
	if len(codes) == 0 {
		log.Fatal("no codes found in file?! " + inputHeaderName)
	}
	codesStr := strings.Join(codes, "\n")
	codesFile := strings.Replace(string(template), "/*KEYCODES*/", codesStr, 1)
	ioutil.WriteFile("codes.go", []byte(codesFile), 0644)
	err = exec.Command("gofmt", "-w", "codes.go").Run()
	if err != nil {
		log.Fatal(err)
	}
}

