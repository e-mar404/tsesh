package bookmark

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/adrg/xdg"
)

type Entries struct {
	Data []Entry `json:"entries"`
}

type Entry struct {
	Url       string `json:"url"`
	Directory string `json:"directory"`
}

func (e *Entries) Add(urls ...string) error {
	cwd, _ := os.Getwd()
	for _, url := range urls {
		e.Data = append(e.Data, Entry{
			Url:       url,
			Directory: cwd,
		})
	}
	return e.save()
}

func (e *Entries) Load() error {
	path, err := xdg.DataFile("tsesh/data.json")
	if err != nil {
		return err
	}

	f, err := os.Open(path)
	if !os.IsNotExist(err) && err != nil {
		return err
	}

	decoder := json.NewDecoder(f)
	err = decoder.Decode(e)
	if err != nil {
		return err
	}
	return nil
}

func (e *Entries) save() error {
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

	return os.WriteFile(path, buf.Bytes(), os.ModePerm)
}

func ValidateDataStorage() error {
	path, err := xdg.DataFile("tsesh/data.json")
	if err != nil {
		return err
	}

	_, err = os.Stat(path)
	if !os.IsNotExist(err) {
		return nil
	}

	e := &Entries{}
	return e.save()
}
