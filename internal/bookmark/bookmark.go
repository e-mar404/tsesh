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
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (d Data) Add(name, rawUrl string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	log.Debug("getting directory", "cwd", cwd)

	u, err := convertToUrl(rawUrl)
	if err != nil {
		return err
	}

	if has(d[cwd], rawUrl) {
		log.Info("duplicate url in list, ignoring", "rawUrl", rawUrl)
		return nil
	}

	if name == "" {
		name = u.Hostname()
	}

	d[cwd] = append(d[cwd], Bookmark{
		Name: name,
		Url:  rawUrl,
	})
	log.Info("added to bookmark list", "name", name, "rawUrl", rawUrl)

	return nil
}

func (d Data) Remove(partial string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	log.Debug("getting cwd", "cwd", cwd)

	var newData []Bookmark
	for _, bookmark := range d[cwd] {
		if !strings.Contains(bookmark.Url, partial) {
			newData = append(newData, bookmark)
		}
	}
	log.Debug("new data after removing matching entries", "newData", newData)
	d[cwd] = newData

	return nil
}

func (d Data) List() ([]Bookmark, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	log.Debug("getting directory", "cwd", cwd)

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

func (d Data) OpenIndex(idx int) error {
	list, err := d.List()
	if err != nil {
		return err
	}
	log.Debug("list from cwd", "bookmarks", list)

	if idx > len(list)-1 {
		return OutofBounds
	}

	return openUrl(list[idx].Url)
}

func (d Data) OpenName(name string) error {
	list, err := d.List()
	if err != nil {
		return err
	}

	for _, item := range list {
		if item.Name == name {
			return openUrl(item.Url)
		}
	}

	log.Info("no bookmark matched", "name", name)
	return nil
}

func (d Data) OpenAll() error {
	list, err := d.List()
	if err != nil {
		return err
	}
	log.Debug("list from cwd", "bookmarks", list)

	for _, item := range list {
		if err := openUrl(item.Url); err != nil {
			return err
		}
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
	log.Debug("got data file", "path", path)

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

func convertToUrl(rawUrl string) (*url.URL, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}

	if u.Scheme == "" {
		return nil, InvalidUrlScheme
	}

	return u, nil
}

func has(bookmarks []Bookmark, rawUrl string) bool {
	for _, bookmark := range bookmarks {
		if bookmark.Url == rawUrl {
			return true
		}
	}
	return false
}

func openUrl(rawUrl string) error {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", rawUrl}
	case "darwin":
		cmd = "open"
		args = []string{rawUrl}
	default:
		cmd = "xdg-open"
		args = []string{rawUrl}
	}
	log.Info("opening bookmark", "rawUrl", rawUrl)
	log.Debug("executing command based on GOOS", "cmd", cmd, "args", args)

	return exec.Command(cmd, args...).Run()
}
