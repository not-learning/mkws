package main

import (
	"bytes"
	"log"
	"os"
)

// linx("index", "Матем", "maths")
// index.md: (Матем)[maths.html]

func linx(fPath, lName, lPath string) {
	mdPath := fPath + ".md"
	f, err := os.OpenFile(mdPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	doErr("linx: ", err)
	defer func() {
		err := f.Close()
		doErr("linx: ", err)
	}()

	// todo proper read
	bb, err := os.ReadFile(mdPath)
	doErr("linx: ", err)

	lName, err = ruName(lName)
	doErr("linx: ", err)

	htmlPath := "/" + lPath + ".html"
	if bytes.Contains(bb, []byte(htmlPath)) {
		return
	}
	htmlLink := "(" + lName + ")" + "[" + htmlPath + "]" + "\n"
	if _, err = f.Write([]byte(htmlLink)); err != nil {
		f.Close()
		log.Fatal("linx: ", err)
	}
}
