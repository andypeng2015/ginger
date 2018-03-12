// Code generated by go-bindata.
// sources:
// misc/ale
// misc/main.go.template
// DO NOT EDIT!

package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _ale = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x93\xc1\x92\xa3\x36\x10\x86\xef\x79\x1a\x90\xed\xad\xf5\x61\x0f\x96\xc6\x08\x31\x86\x04\x18\x24\xa1\x1b\x92\xb6\x1a\x8c\xb0\x9d\xc2\x18\x8b\xa7\x4f\xe1\x49\x66\x93\xdc\x54\x5d\xdd\xad\xbf\xff\xfe\xba\xa4\xfb\xbb\x2a\x71\x56\x8b\xc2\xb1\xa3\xbd\xe9\xc1\x3a\xd3\x61\xac\x69\xf5\x4e\xba\xf4\x7c\x2a\x71\xf2\xd1\xdb\xb2\xa8\x8a\xa4\xa8\x76\x55\x35\x5f\xef\xec\xed\x08\x69\xf9\xdd\x33\x9a\x85\x26\x4e\xa1\x46\xed\xc3\xc4\xd8\x29\x82\x67\x25\xf8\xa8\x44\x0e\x0a\xb9\x49\x21\xee\x19\x81\x0d\xa3\xfb\x50\x0f\x99\x33\xbe\x7f\x3f\x95\x87\xaf\x3a\x8b\xa2\x40\x49\xb6\xc6\x96\xd3\x92\x43\xbd\xe1\x33\x8b\xb3\x50\xa1\xc8\x93\x2e\x80\x5c\x24\x0f\x2b\x73\x48\x4b\xfc\xb4\x22\xf2\x96\xe0\x73\x43\xdd\xa8\x29\xef\x59\x9c\x39\x1d\x17\xdf\x94\x64\xa0\x37\x0c\x6a\xf4\x0c\xeb\x0e\x2f\x1a\x15\xed\x7f\x6a\x3d\x0e\x6a\x91\x8c\x4a\x66\xb3\x46\xfb\xc9\x78\x3c\x98\x81\x2f\x0d\xc1\x63\x23\x42\xc7\xe8\x2d\x6c\x44\xe6\x08\x5c\xcf\xcc\xe3\xbc\xea\x79\x9e\xf3\x04\xf3\xa3\xfb\xe3\x03\xae\xef\xcc\xa7\x67\x76\x0c\xdb\x06\x55\x60\x36\xce\x5b\x79\xf8\x46\xfa\xec\xa1\xd1\xf8\x6b\x46\xea\x26\x16\xdb\xd6\xd2\xbf\xdf\xf4\x08\x06\x85\xad\xa6\x33\x18\x14\x85\x35\xe2\x73\x2d\xb6\x50\xcb\x1c\xea\x12\x8f\x7a\x63\xc0\x20\x77\xd7\x82\xfb\x13\xc1\x73\x2d\x93\xa0\x11\xd1\xa8\xe3\x1e\x6a\xb4\x47\x4a\x26\x4e\x91\x19\xb2\xb7\x0a\xb4\x70\x93\x95\x85\x33\x7e\x0b\xd5\xc0\xef\x7a\xa3\x1c\xa3\xca\x6b\x14\x40\x43\x79\x6b\x09\x6e\xf5\x90\x83\xa6\x3c\x60\x71\x16\x28\xc1\xe7\x13\xc1\xe7\x35\xcf\x0c\xbc\x3f\x91\x83\x4f\x09\xbe\x37\x62\x17\x5a\xca\x97\x13\xdc\x3e\x6c\x9c\xb4\x8d\xd8\xae\xbd\xb7\x36\xe6\x5e\x95\x38\x68\xe2\xe4\x61\x85\xbd\x7e\x6a\x77\xce\x0e\x15\x34\x62\x17\x68\x8f\x5b\x46\x93\x87\x45\xf3\xc8\x62\xec\x95\xcc\x96\x46\xec\x2e\x8c\xee\x27\x46\xbf\x78\x79\x69\xb0\xb4\x75\xba\xc3\x7d\x23\xb3\x73\x2d\x93\x7e\xcd\x3b\x75\xb8\x54\xb2\x08\xcd\xb0\x05\x4d\xdd\xd3\x0a\xd7\xb3\xb8\x78\xb0\x38\x6b\xad\xc8\x9c\xa1\xd1\xc4\x68\x34\x29\x82\x5b\x45\x73\x30\x1b\x7e\xa9\x25\x83\x5a\xec\xfa\x97\x16\xea\xa6\x1a\x01\x68\x54\xaf\x5e\x8e\x96\xcc\xff\xfc\x73\x55\x22\x0a\x18\x5d\xfd\x62\xa0\x05\xef\x1b\xc9\xef\x8c\xb6\xae\x96\xf9\xb8\x7a\xd1\xc8\xc4\xbf\x34\x10\x1c\xea\x4b\x71\xd3\x04\x2f\x56\xd8\xd6\x74\xf8\x5a\xcb\x14\x14\x75\x8b\x41\xfb\xd1\xae\x3e\x75\xf8\xed\xc5\x16\xc1\xcb\x4f\x99\x84\x86\xe0\x9b\xee\xd6\xd9\xb3\x87\x1d\xb8\x5f\x99\x6e\x86\xc8\xb3\x98\x4f\x96\xba\x71\x8d\x6b\x9a\x4f\x9f\xdc\xa4\x90\xcb\x6c\x51\x22\xec\x34\xad\x40\xc5\xc9\x4d\x0f\xf7\xe5\x77\xb8\x65\x8d\x84\x5f\x9c\xc4\xd9\xce\x5c\x56\xc6\xed\xcd\x52\x00\x83\xf8\x68\xe3\xdb\x97\x7f\x9a\xba\xf5\x16\xff\xb4\xc2\x9d\x55\x79\xb8\x9a\x4d\xd1\x9a\x4b\x0e\x16\xb9\xa0\x21\x87\xe7\xbf\x6e\xe4\xb3\xcf\xcb\xb7\xc3\x93\xbd\x1d\x1f\x69\x87\x03\x25\xa2\xc5\xd0\xfd\x43\x5f\xd2\xff\xf5\xc2\xf3\xea\x51\xea\xbf\x07\xaf\xdb\x23\x78\x51\xe2\x19\xfc\x1c\x5e\xfc\x05\x0d\xe5\xeb\x0e\xfa\xe6\xc2\x17\x4b\xf0\xba\xf3\xa0\x96\x59\xa0\xca\x7e\x22\xf0\xe3\xc7\x6f\x7f\x05\x00\x00\xff\xff\xbf\xd5\x55\x92\x25\x04\x00\x00")

func aleBytes() ([]byte, error) {
	return bindataRead(
		_ale,
		"ale",
	)
}

func ale() (*asset, error) {
	bytes, err := aleBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ale", size: 1061, mode: os.FileMode(420), modTime: time.Unix(1518375939, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mainGoTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8e\x41\xca\x83\x30\x14\x84\xd7\xce\x29\x1e\x81\x40\x02\xbf\x7a\x88\x7f\xd3\x7d\x4f\xf0\x8c\xa9\x95\x9a\x44\x92\x27\x15\x8a\x77\x2f\x12\xe9\x62\x98\x81\x61\x3e\x66\x65\xf7\xe2\xc9\x53\xe0\x39\x02\x73\x58\x53\x16\x32\x68\x94\x4b\x51\xfc\x2e\x0a\x68\xd4\x34\xcb\x73\x1b\x3a\x97\x42\xcf\xef\x72\xaa\x5d\x38\x0c\x23\xb7\x53\xea\x6b\x52\xba\xc0\x02\x8f\x2d\x3a\xd2\xe5\xc6\x71\x5c\x7c\x36\x4e\x76\xba\x40\xdd\x7f\xf5\x3f\xd2\xc5\x92\x2e\xf4\x41\x93\xbd\x6c\x39\x92\x2e\x38\xae\xe9\x79\xc3\xd8\xb3\xab\xd8\xee\x2e\x9c\xc5\xfc\x88\x16\x07\xbe\x01\x00\x00\xff\xff\x88\xe7\x78\xaa\xb3\x00\x00\x00")

func mainGoTemplateBytes() ([]byte, error) {
	return bindataRead(
		_mainGoTemplate,
		"main.go.template",
	)
}

func mainGoTemplate() (*asset, error) {
	bytes, err := mainGoTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "main.go.template", size: 179, mode: os.FileMode(420), modTime: time.Unix(1519484462, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"ale": ale,
	"main.go.template": mainGoTemplate,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"ale": &bintree{ale, map[string]*bintree{}},
	"main.go.template": &bintree{mainGoTemplate, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

