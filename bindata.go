// Code generated by go-bindata.
// sources:
// examples/default/default.hs.js
// VERSION
// DO NOT EDIT!

package hotshell

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

var _examplesDefaultDefaultHsJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x58\xfd\x8e\xd4\x46\x12\xff\x9f\xa7\x28\x8d\x10\x9e\xd9\xc5\xf6\x2e\xda\xbb\xe3\xe0\x80\xbb\x13\x77\xb7\x17\x91\x80\xb4\x44\x89\x04\x04\xf5\xd8\x35\xe3\xce\xda\xdd\x4d\x77\x7b\x3e\x02\x9b\x67\xc9\xb3\xe4\xc9\x52\xdd\xf6\xf8\x6b\xc6\xc3\xae\x94\x05\x31\x3b\xfd\x51\x1f\xbf\xaa\xfa\x55\x35\xdc\x62\x31\xfd\x9c\xa2\x49\x9e\x40\x70\x29\xad\xc9\x30\xcf\x83\x9b\x87\xb0\x28\x45\x62\xb9\x14\x30\x9d\xc1\xe7\x7b\xf7\x80\x7e\x56\x4c\x43\xce\x45\xb9\x81\x67\x80\x1b\x4c\xa6\x41\x29\x58\x81\xc1\x2c\xe2\x22\xc5\xcd\xeb\xc5\x34\x78\xe5\xb6\x83\x19\x3c\x87\xf0\xbc\xb9\x33\xd7\x72\x6d\x50\xd3\xad\xea\xf6\x0b\x08\x0c\x0a\xc3\xe7\x39\x86\xf5\x5e\x00\xa4\x5e\x2a\x14\x41\xa5\xaa\x6b\xd6\xe4\x6d\xc6\x0d\xd0\xdf\xc6\x3c\x03\x29\x2e\x58\x99\x5b\x28\x50\x94\x0f\x21\xe5\x46\xe5\x6c\x8b\x29\xac\x33\x14\x20\xa4\xdb\xe7\x82\x7b\xfb\x17\x3c\x47\x77\x5b\x69\xb9\xe2\x29\xa6\xd1\x7b\x31\xb9\x99\x3d\x3d\xa0\x07\xe0\xa5\x14\x81\x85\x6b\x21\xd7\x24\x89\xd9\x46\xa3\xbb\xcf\xe6\xb2\xb4\xf0\xa2\xbe\xdd\x5e\xbe\xc6\x2d\xd9\xce\x02\x32\xa3\x12\x73\xc9\xad\xfb\x0e\x56\xc2\x12\xad\xfb\xf0\x02\xb9\x9d\x1c\x86\x75\xcf\x8e\x9e\x56\x48\x64\x51\x30\x91\x86\x84\x1d\x02\x53\x2a\xe7\x09\xf3\x12\x48\x30\x2e\x16\x3c\xe1\x28\x6c\xbe\x05\x8d\x09\xa3\x3b\x74\x14\x4c\xc6\x34\xee\x2e\x9a\x9e\xc3\xbb\x98\x70\x21\x50\xbf\xc4\x42\xc2\xb3\x66\xdd\xfd\x10\x08\x10\xc7\xbb\xbb\xde\x2e\x43\xf7\xe1\x74\xef\x54\xd7\x7b\xdd\x78\x1f\x68\x34\x96\x69\x4b\x96\xb2\x24\x43\x5a\x4f\x8a\x94\x96\x4d\x99\x4a\xa0\x40\xaf\x78\x82\xf5\xde\x23\xa8\xcf\x06\x37\xb3\xaf\xab\x30\xad\x0a\xb3\x15\x49\xa6\xa5\xe0\xbf\x20\x58\x5e\x0c\x94\xc4\xa5\xd1\xb1\x99\x73\x11\x0b\xab\x52\x66\x11\x94\x94\x79\x44\x5f\x22\xa9\x97\xa3\xba\x0e\xaf\x12\x14\xa6\x9c\xbb\x2c\xf3\xe1\xd4\xb2\x54\x90\xcb\x25\x59\x9e\x93\xe4\xb4\xc1\xf8\xeb\xf6\xe7\xad\xfd\x95\xfb\x4e\x8e\xd9\xab\xb4\xc3\x82\xc6\x72\x2d\x60\x49\x82\xc6\x44\x24\xaa\x01\x21\xa7\x05\x38\xfd\x2f\xc4\x14\xe6\x98\x36\xe2\x1a\xed\xb8\x73\x76\x0c\x84\xbe\x1e\x6c\xf5\xa0\xd6\x52\xdf\x4e\x4d\x7b\x74\x4c\x8b\x5f\xef\xa7\x63\xea\x33\x11\x26\x3d\x3a\x72\x8b\x87\x01\xea\xa4\xef\x29\x89\x1c\x08\xec\x15\xd3\xb7\x14\x3b\x2a\x23\xaa\x07\xcf\x09\x14\xb3\xd2\x70\xb1\xa4\xc2\xfa\x86\xad\xd8\x55\xa2\xb9\xb2\x14\x46\xb1\xe0\xcb\x52\x57\x85\xf5\xf2\xea\x55\x04\xff\xd9\xb0\x42\x11\x73\x3c\xe9\x14\xfc\x50\xb8\x33\xf0\x66\x76\x58\xaf\x27\xad\x44\x23\xa5\x89\x01\x4b\xe1\x5e\xc8\x3c\x97\x6b\xa7\xda\xa7\xd3\x50\x2e\xae\x58\x3e\x6d\xdc\x9a\x8d\x78\xf3\x5e\xc0\x0f\x4c\x78\x52\xc9\x91\x69\x01\x85\x24\xcf\x2a\x52\x1a\xdc\xa8\x62\x58\x0c\x78\xa9\xf0\xbc\xe4\x98\xb6\x4b\xa7\x8a\x2d\x71\x52\x47\x76\xc7\xd6\xa7\x10\x40\x66\xad\x32\x4f\xe2\x78\xc9\x6d\x56\xce\x23\xca\xf6\xf8\xe7\x32\x27\xc2\x29\x64\x59\x08\x8c\xb3\xa6\x61\x1c\x31\xd8\xe9\x75\x00\x18\x4a\x10\x9c\x53\xb0\x5d\x25\x49\xf7\xaf\x5b\x55\x1a\x57\x5c\x52\x88\x3c\x2a\x52\x43\x62\x75\x7e\x9a\xb8\xed\x4f\x25\x6f\x29\x78\xe7\xe0\x4e\xd3\x50\x0b\xc0\x25\x5b\x39\x74\x2d\x15\x29\x75\x16\x58\x93\xcd\xee\x76\xed\xe2\x8e\xd1\x3c\x8d\x4a\xe5\x02\x6d\xc6\xc8\x3c\x1b\x80\x96\x79\xd0\xea\x2e\xe3\xad\x26\xa1\x6a\x07\x58\x90\x19\x08\x43\xb7\x12\x8c\x1b\xb7\x8b\x9a\x41\xa4\xde\xe9\x53\xcb\x80\x5c\x34\x04\x32\x66\x0a\x0e\x4c\xc1\xe0\xab\x5d\xa4\xba\xb8\x6e\x8b\x77\x4d\x49\x98\x51\x44\x1f\x54\x6c\xb9\x77\xbf\x5b\x9f\x5d\x11\x49\x2b\x82\x2a\x3b\xb9\xf6\x9e\xd7\xc2\x1a\x1e\x48\x4a\x9d\xc3\xda\x5a\x4d\x43\x40\xd0\x49\xc2\xa1\xb0\x3f\x8f\xbf\x47\x55\xd8\x8e\xcb\xcc\x26\x99\xb7\xd7\x29\x70\xe9\x36\xdf\x36\x6a\xaa\xcd\x50\xc0\x39\x38\xf9\x5d\x91\xfb\x79\xbc\x67\xfd\xd5\xd5\x25\x6c\x65\xa9\xa9\xf4\x0c\xb5\x39\x4f\xaa\xc4\x2a\xae\xb7\xa1\xde\x27\xf4\x51\x6b\x3b\x0d\x21\x97\xd4\xba\x33\x12\xd7\x22\x61\x32\x68\x57\x8f\xf8\xbc\x6c\xa5\xfc\x4f\xe3\x06\xde\x50\xee\xf3\x04\xfe\xe5\xcd\x82\xef\xbf\xfb\xff\x8f\x3d\x99\x02\xd7\x25\x59\xfa\xcf\x25\x9d\x1d\xc2\x39\xe6\xfb\xa2\xd3\x06\x52\xca\x41\xef\x7c\x9f\x31\xdd\x8c\x75\x07\xd7\xe7\xad\xc4\x15\x2f\xe0\xd7\x38\x9a\x33\x93\x7d\xa4\x09\xcd\x09\x6a\x0c\x3e\xb4\x77\x04\x8a\x7f\x77\xec\x2c\x58\x62\x46\xe5\x1e\xde\x3d\x22\x39\xeb\xdb\x1b\xa3\x4d\x62\x17\x18\xd3\xcf\xdc\xc1\xde\x11\x81\x97\x43\x53\xc7\x44\xee\xed\xde\x22\x5e\xa2\x15\x2e\xd0\xae\xa5\xbe\xa6\xca\x37\x5b\x43\x47\xa0\xb4\x3c\xbf\x43\xa0\x3a\xa1\xa7\xde\x99\x82\xc5\x8d\xa5\xde\x5b\xc7\xbb\xc1\x33\xc9\xa4\x2b\xa7\xc9\x3b\x97\xb2\x4e\xea\x07\x78\xa7\x98\xb5\xa8\xe9\xb7\xc9\x53\x9a\x96\x18\x11\x2f\xa8\xa7\x34\x41\xa1\x82\x50\x8b\xb5\x81\xfb\x39\x84\x08\xf7\xd5\x31\x9c\xe4\x4e\xff\xef\xbf\xd5\x14\xe4\x8b\x82\x18\x41\xdb\x03\xea\xdd\x72\xab\x90\xd4\x91\x35\xf0\x0f\x88\x53\x5c\xc5\x36\x51\xf1\xf9\xa3\xbf\x45\x67\xf4\xe7\x3c\x3e\xae\xb6\xec\x72\x55\x85\x9b\xea\x11\x55\xfd\xf5\xa0\x88\xde\x92\xfb\x69\xc8\x74\x6f\xa7\xd6\xe1\x89\xa4\x34\x9d\x71\xf2\xc0\xd9\x8e\xaf\x13\x1a\x2d\xac\xd4\xdb\x09\x7c\x01\x97\xc2\x10\x72\x78\xf4\xdc\x7b\x29\x4a\x7a\x02\x7c\x01\x27\x6b\x62\xe2\x9f\xe0\xe4\xdd\x59\xf8\xf7\x0f\x27\x70\x12\xc7\xee\xb4\x21\x80\xe8\xa3\x14\xfc\x13\x84\xc9\x6e\x21\x14\x9a\x7e\xcd\x08\xb3\xa0\xa7\xf6\x08\x42\xaa\x45\x48\x55\x8c\xc3\x15\xb0\x34\xd5\x8e\x75\x1e\xd0\xbb\x47\xee\x52\xa1\xdf\x28\xdc\x40\x41\xf3\x04\x57\x5c\x2c\x64\xc4\xe5\x9d\x20\xe4\xe3\x10\x72\xe1\x92\x0d\xfd\x24\x27\xb0\xca\x6c\xa3\x10\xd3\x51\x24\xd7\xee\x6d\x16\xbe\x86\x16\xb6\xda\x36\x1a\xed\x91\x19\x34\x51\x39\x2f\x85\x2d\xfd\xd0\x73\x7e\x11\x9d\x5d\x44\x17\x71\xb5\x14\xd6\x5f\x43\x52\x7e\x6d\xa5\x0a\x59\x91\xfe\xf5\x22\xe2\x46\xde\x1a\xbf\x0e\x6d\x2b\x37\xb4\xf8\x87\xb1\x23\x64\x98\xb6\xf3\x0f\x05\x59\xcd\x1a\xfc\xfa\xe7\x6e\xdb\x6a\x5d\x57\x02\x82\x5e\xd3\x23\x91\x26\x18\x7a\x24\xba\xcc\x01\x82\xc7\x15\x0b\x3c\x3e\x7b\x7c\xde\x6a\xd8\xda\x8c\x36\xc2\x02\xae\xb8\x1b\x50\x2e\xdf\xbe\x7d\x73\xe5\xdb\x5a\x75\xee\x16\xdc\xd3\x6b\x6a\xcb\xb1\xbe\x70\xab\x9e\xc8\x5c\x51\xf8\x01\x20\xdd\xb1\x57\x2b\xb2\x79\x85\xd0\xe8\x95\xdb\xe6\x11\xb2\x4b\xe3\x23\xe0\x74\x1e\x4f\xfe\x09\xd3\x7b\xb6\xdc\xed\x75\xd3\xd3\xf1\x71\x5a\xfd\x97\x88\xb3\x28\x6d\x2d\x8a\x4f\xa2\x93\xda\x2a\xc7\x51\x7f\x09\x66\x11\x4d\x91\xdc\x4e\x83\xf7\x22\x98\xcd\x22\x24\xb1\xd3\x16\x1d\xcc\x1f\x02\xdf\x0c\x29\x79\xe0\x03\xdf\xf4\x5d\x08\x68\x56\x77\x17\xfb\xa6\x57\xab\x03\x2b\xfb\x21\xac\x3f\x6e\x66\x7f\x04\x00\x00\xff\xff\xef\xe2\x58\x98\x06\x12\x00\x00")

func examplesDefaultDefaultHsJsBytes() ([]byte, error) {
	return bindataRead(
		_examplesDefaultDefaultHsJs,
		"examples/default/default.hs.js",
	)
}

func examplesDefaultDefaultHsJs() (*asset, error) {
	bytes, err := examplesDefaultDefaultHsJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/default/default.hs.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _version = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x32\xd0\x33\xd6\x33\x00\x04\x00\x00\xff\xff\xdf\xf5\x1d\x2c\x05\x00\x00\x00")

func versionBytes() ([]byte, error) {
	return bindataRead(
		_version,
		"VERSION",
	)
}

func version() (*asset, error) {
	bytes, err := versionBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "VERSION", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"examples/default/default.hs.js": examplesDefaultDefaultHsJs,
	"VERSION": version,
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
	"VERSION": &bintree{version, map[string]*bintree{}},
	"examples": &bintree{nil, map[string]*bintree{
		"default": &bintree{nil, map[string]*bintree{
			"default.hs.js": &bintree{examplesDefaultDefaultHsJs, map[string]*bintree{}},
		}},
	}},
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

