package bookmark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/adrg/xdg"
)

type Entries struct {
	Data []Entry `json:"entries"`
}

type Entry struct {
	Url       string   `json:"url"`
	Directory string   `json:"directory"`
}

func (e Entries) Add(urls... string) error {
	cwd, _ := os.Getwd()
	for _, url := range urls {
		e.Data = append(e.Data, Entry{
			Url: url,
			Directory: cwd,
		})
	}

	return e.save()
}

func (e Entries) save() error {
	buf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(buf)
	encoder.SetIndent("", "  ")
	
	if err := encoder.Encode(e); err != nil {
		return err
	}

	path, err := xdg.DataFile("tsesh/data.json")
	if err != nil {
		return err
	}

	fmt.Printf("writing: %v\n to file: %v\n", buf.String(), path)

	return os.WriteFile(path, buf.Bytes(), os.ModePerm)
}
