package main

import (
	"bytes"
	"errors"
	"os"
)

const dictFile = "dict.txt"

/*func main() {
	data, err := os.ReadFile(dictFile)
	// doErr("runame: ", err)
	if err != nil {
		fmt.Println("runame:", err)
	}

ph, err := ruName("uniaccel", mkMap(data))
if err != nil { fmt.Println(err) }
fmt.Println(ph)

econ, err := ruName("econ", mkMap(data))
if err != nil { fmt.Println(err) }
fmt.Println(econ)
}*/

func mkMap(data []byte) map[string]string {
	dict := make(map[string]string)
	bb := bytes.Split(data, []byte("\n"))
	for _, v := range bb {
		ss := bytes.Split(v, []byte("	"))
		if len(ss) > 1 {
			dict[string(ss[0])] = string(ss[len(ss)-1])
		}
	}
	return dict
}

func ruName(eng string) (string, error) {
	data, err := os.ReadFile(dictFile)
	doErr("runame: ", err)
	/*if err != nil {
		fmt.Println("runame:", err)
	}*/

	dict := mkMap(data)
	for k, v := range dict {
		if eng == k {
			return v, nil
		}
	}
	return "", errors.New("dict: no entry for " + eng)
}
