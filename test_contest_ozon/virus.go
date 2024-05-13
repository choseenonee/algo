package test_contest_ozon

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func hasHack(files []string) bool {
	for _, fileName := range files {
		if strings.HasSuffix(fileName, ".hack") {
			return true
		}
	}

	return false
}

func checkDirHack(folders []Folder) bool {
	for _, folder := range folders {
		if hasHack(folder.Files) {
			//TODO: implement tree elements calculation
			return true
		}
	}
	for _, folder := range folders {
		if checkDirHack(folder.Folders) {
			//TODO: implement tree elements calculation
			return true
		}
	}

	return false
}

type Folder struct {
	Dir     string   `json:"dir"`
	Files   []string `json:"files"`
	Folders []Folder `json:"folders"`
}

func Virus() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var lineCount int
	var line []byte
	fmt.Fscan(in, &lineCount)

	var jsonBytes []byte

	for lineCount >= 0 {
		lineCount--
		line, _, _ = in.ReadLine()
		opa := string(line)
		_ = opa
		jsonBytes = append(jsonBytes, line...)
	}

	var data Folder

	err := json.Unmarshal(jsonBytes, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
