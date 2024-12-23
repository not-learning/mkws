package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const mapFile = "map.mp"

func main() {
	data, e := os.ReadFile(mapFile)
	if e != nil {
		fmt.Println("Can't read "+mapFile, e)
	}

	tree := parse(string(data))
	for _, k := range tree {
		myMkdirs(k.name)
		linx("index", k.name, filepath.Join(k.name, k.name)) // homepage
		linx(filepath.Join(k.name, k.name),                  // links home in subject pages
			"home",
			"index")

		for i, v := range k.knot {
			myMkdirs(filepath.Join(k.name, v.name))
			linx(filepath.Join(k.name, v.name, v.name),
				"home",
				"index")
			linx(filepath.Join(k.name, k.name), // links to topics in subject pages
				v.name,
				filepath.Join(k.name, v.name, v.name))

			if i+1 < len(k.knot) { // "next" links in topic pages
				linx(filepath.Join(k.name, v.name, v.name),
					"next",
					k.knot[i+1].name)
			}
			if i-1 > -1 { // "prev" links in topic pages
				linx(filepath.Join(k.name, v.name, v.name),
					"prev",
					k.knot[i-1].name)
			}
		}
	}
}

// linx("index", "Матем", "maths")
// index.md: (Матем)[maths.html]

func doErr(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
