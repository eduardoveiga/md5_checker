/*
Copyright (C) 2019
Eduardo Kluwe Veiga: eduardoveiga@protonmail.ch
SPDX-License-Identifier: WTFPL
*/

package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func readfile(filename string) [][]string {
	file, err := os.Open(filename)
	var lines [][]string
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		lines = append(lines, strings.Split(string(line), " "))

	}
	return lines
}

func check(hash string, file string) {
	Readfile, _ := ioutil.ReadFile(file)
	hasher := md5.New()
	hasher.Write(Readfile)
	file_hash_string := hex.EncodeToString(hasher.Sum(nil))
	if hash != file_hash_string {

		fmt.Printf("%s hash is invalid:\t is %s and should  be %s\n", file, hash, file_hash_string)
	} else {
		fmt.Printf(".")
	}

}

func main() {
	var hash_table [][]string
	filename := os.Args[1]
	hash_table = readfile(filename)
	for _, file_to_check := range hash_table {

		check(file_to_check[0], file_to_check[len(file_to_check)-1])
	}

}
