package bookmark

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/adrg/xdg"
	"github.com/charmbracelet/log"
)

var InvalidUrlScheme = errors.New("Invalid URL scheme provided")
var EmptyData = errors.New("Data has no saved info, add something before removing")
var OutofBounds = errors.New("Index given is out of bounds for current bookmark list")

type Data map[string][]Bookmark

type Bookmark struct {
	Url string `json:"url"`
}

func (d Data) Add(urls ...string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	log.Debugf("using directory: %v", cwd)

	for _, rawUrl := range urls {
		if err := validate(rawUrl); err != nil {
			return err
		}

		if has(d[cwd], rawUrl) {
			log.Infof("ignoring %s, duplicate url in list", rawUrl)
			continue
		}

		d[cwd] = append(d[cwd], Bookmark{
			Url: rawUrl,
		})
		log.Infof("added %s to bookmark list", rawUrl)
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

func (d Data) List() ([]Bookmark, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	log.Debugf("using cwd: %s", cwd)

	return d[cwd], nil
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

func (d Data) Open(idx int) error {
	list, err := d.List()
	if err != nil {
		return err
	}

	if idx > len(list)-1 {
		return OutofBounds
	}

	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", list[idx].Url}
	case "darwin":
		cmd = "open"
		args = []string{list[idx].Url}
	default:
		cmd = "xdg-open"
		args = []string{list[idx].Url}
	}

	return exec.Command(cmd, args...).Run()
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
	log.Debugf("using file at %s as a data file", path)

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
