package bookmark

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/url"
	"os"
	"strings"

	"github.com/adrg/xdg"
)

var InvalidUrlScheme = errors.New("Invalid URL scheme provided")
var EmptyData = errors.New("Data has no saved info, add something before removing")

type Data map[string][]Bookmark

type Bookmark struct {
	Url string `json:"url"`
}

func (d Data) Add(urls ...string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	for _, rawUrl := range urls {
		if err := validate(rawUrl); err != nil {
			return err
		}

		if has(d[cwd], rawUrl) {
			continue
		}

		d[cwd] = append(d[cwd], Bookmark{
			Url: rawUrl,
		})
	}
	return nil
}

func (d Data) Remove(partial string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	var newData []Bookmark
	for _, bookmark := range d[cwd] {
		if !strings.Contains(bookmark.Url, partial) {
			newData = append(newData, bookmark)
		}
	}
	d[cwd] = newData

	return nil
}

func (d *Data) Load() error {
	path, err := xdg.DataFile("tsesh/data.json")
	if err != nil {
		return err
	}

	f, err := os.Open(path)
	if !os.IsNotExist(err) && err != nil {
		return err
	}

	decoder := json.NewDecoder(f)
	err = decoder.Decode(d)
	if err != nil {
		return err
	}
	return nil
}

func (d *Data) Save() error {
	buf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(buf)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(d); err != nil {
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

	e := &Data{}
	return e.Save()
}

func validate(rawUrl string) error {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return err
	}

	if u.Scheme == "" {
		return InvalidUrlScheme
	}

	return nil
}

func has(bookmarks []Bookmark, rawUrl string) bool {
	for _, bookmark := range bookmarks {
		if bookmark.Url == rawUrl {
			return true
		}
	}
	return false
}
