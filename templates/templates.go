// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package templates generated by go-bindata.// sources:
// index.html
// mech.html
// mechs.html
// mech.css
package templates

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _indexHTML = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x91\x31\x52\xeb\x30\x10\x86\xfb\x77\x8a\x7d\x2a\x52\xbd\x58\x49\xaa\xc7\x64\xed\x01\x02\x0c\x0d\x03\x13\x42\x41\xa9\x38\x9b\x48\x83\x1c\x0b\x69\x31\xf1\x91\xb8\x06\x27\x63\x64\x39\xc6\x95\xfc\xc9\xfb\x7f\xfe\x67\x8d\x7f\x6f\x1e\x57\x9b\xd7\xa7\x5b\xd0\x5c\xd9\xe2\x0f\xa6\x03\x00\x35\xa9\x5d\x7c\x00\xc0\x8a\x58\x81\x66\x76\x53\x7a\xff\x30\x4d\x2e\xca\xfa\xc8\x74\xe4\x29\xb7\x8e\x04\xf4\x94\x0b\xa6\x13\xcb\x28\x58\x42\xa9\x95\x0f\xc4\xf9\xcb\xe6\x6e\xfa\x5f\xf4\x1e\x36\x6c\xa9\x78\xa0\x52\x07\xf2\x0d\x79\x94\xe9\x26\xbd\xb5\xe6\xf8\x06\xda\xd3\x3e\x17\x52\x39\x23\x9b\xb9\x0c\xdc\x5a\xca\xca\x10\x04\x78\xb2\xb9\xe8\x38\x68\x22\xee\x94\x28\xcf\x25\x71\x5b\xef\xda\xde\x13\xef\xc8\x27\x88\xb8\x18\x7d\x11\x9a\x79\x36\xcb\x66\x28\xf5\xa2\x9f\x96\xe3\x71\x64\xb5\x3d\xf7\x89\xf4\x6b\xed\xd8\x15\xf7\xe4\x09\x3e\x09\x54\x3c\x0c\x6b\x60\x4d\x50\x0d\xfa\x80\xd2\x0d\x69\x39\x8a\xa3\x1c\x99\xd1\x41\x69\x55\x08\xc3\x1e\x45\x01\x43\x10\xf7\x75\xcd\xe4\xcf\x13\x89\x44\xb1\xaa\x5d\xeb\xcd\x41\x33\x7c\x7f\xc1\x62\x36\xbf\x80\x6b\x4f\xcc\xf0\x5c\xc5\x16\x13\xcb\xcb\xd3\xb6\x0c\x11\x2e\x0f\x95\x32\x36\x2b\xeb\x6a\x72\xe0\xe5\x3f\xc8\xe0\xca\x5a\x58\xc7\x6c\x80\x35\x75\x45\x77\x19\xca\x64\x4e\x6b\x4c\x45\x51\xa6\x9f\xff\x13\x00\x00\xff\xff\xac\x2c\x91\xe1\x14\x02\x00\x00")

func indexHTMLBytes() ([]byte, error) {
	return bindataRead(
		_indexHTML,
		"index.html",
	)
}

func indexHTML() (*asset, error) {
	bytes, err := indexHTMLBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 532, mode: os.FileMode(436), modTime: time.Unix(1603026852, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mechHTML = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x94\xcf\x8e\xd3\x30\x10\xc6\xef\x3c\xc5\xe0\xc3\x9e\x68\xdc\xf6\x04\xea\x24\x02\xb6\x20\xed\x81\x3f\x2a\x0b\x12\x27\xe4\x3a\xb3\xb5\x85\x13\x07\x7b\xa8\xb6\x5a\xed\x03\xf1\x1a\x3c\x19\x72\x9c\xa8\xe1\xb2\xca\xa1\xa7\xfa\xe7\x7e\xf9\x7e\x8e\xe4\x09\x3e\xdf\x7e\xba\xbe\xfd\xfe\xf9\x1d\x18\x6e\x5c\xf5\x0c\xf3\x0f\x00\x1a\x52\x75\x5a\x00\x60\x43\xac\xc0\x30\x77\x0b\xfa\xf5\xdb\x1e\x4b\xa1\x7d\xcb\xd4\xf2\x82\x4f\x1d\x09\x18\xa8\x14\x4c\xf7\x2c\x53\xc1\x06\xb4\x51\x21\x12\x97\x5f\x6f\xdf\x2f\x5e\x8a\xa1\x87\x2d\x3b\xaa\x3e\x90\x36\x28\xf3\x3a\xef\x3b\xdb\xfe\x04\x13\xe8\xae\x14\x52\x75\x56\x1e\x57\x32\xf2\xc9\x51\xa1\x63\x14\x10\xc8\x95\xa2\xe7\x68\x88\xb8\x2f\x43\x39\x1e\x0f\xf7\xbe\x3e\x0d\x3d\x69\x8f\x42\x86\x84\xeb\xde\x15\x29\x1c\x29\xc0\x71\x55\x2c\x8b\x25\x4a\xb3\x1e\xd2\x72\x1a\x47\x56\xfb\xf1\x3c\x89\xce\xad\x99\xc3\x19\x12\x1a\xd0\x4e\xc5\x58\x8a\x86\xb4\xf9\xd1\xbf\x8b\xa8\x6e\xb6\x28\xd9\xfc\x1f\xac\xa7\x41\x51\x3d\x3c\x40\x71\xb3\x85\xc7\x47\x94\x5c\x4f\xfa\xe5\x54\x30\xcf\xf6\x51\x35\x34\xcb\x97\x82\x17\x31\x7e\xa3\x10\xad\x6f\x61\x96\x75\x0c\x5f\x42\xbc\x23\x47\x2a\xd2\x3c\xf1\x18\xbe\x84\x78\x4b\x51\x07\xdb\xb1\xf5\xed\x2c\xf7\x24\xff\xa4\x1f\xe5\xe4\x82\xa1\x9c\xdc\x3d\xec\xc6\xd6\x61\xaa\x44\x05\x28\xbb\xe1\xcf\x3b\xef\x99\xc2\x98\xc8\x24\xaa\x6b\xdf\x9d\x82\x3d\x18\x86\xbf\x7f\x60\xbd\x5c\xbd\x82\xb7\x81\x98\xe1\x4b\x63\xd9\xc0\x95\xe3\xcd\xfd\x5e\xc7\x04\xaf\x0f\x8d\xb2\xae\xd0\xbe\xb9\x3a\xf0\xe6\x05\x14\xf0\xc6\x39\xd8\xa5\x67\x23\xec\xa8\x9f\x94\xba\x40\x99\x9b\xf3\xa0\xe5\x83\xa2\xcc\x1f\x86\x7f\x01\x00\x00\xff\xff\x43\xd2\xba\xc3\x30\x04\x00\x00")

func mechHTMLBytes() ([]byte, error) {
	return bindataRead(
		_mechHTML,
		"mech.html",
	)
}

func mechHTML() (*asset, error) {
	bytes, err := mechHTMLBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mech.html", size: 1072, mode: os.FileMode(436), modTime: time.Unix(1603026852, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mechsHTML = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x54\x4b\x8e\x13\x31\x10\xdd\x73\x8a\xc2\x8b\x59\x11\x3b\xc9\x0a\x14\x77\x0b\x98\x80\x34\x0b\x3e\x0a\x03\x12\x2b\xe4\xb8\x2b\xb1\x85\xfb\x83\x5d\x13\x4d\x14\xe5\x40\x5c\x83\x93\x21\xb7\x9d\xa4\x95\x84\xd9\xb0\x8a\x5f\xd5\xf3\x4b\xd5\xab\x72\xcb\xe7\xf3\x4f\xb7\xf7\xdf\x3f\xbf\x03\x43\xb5\x2b\x9f\xc9\xf4\x03\x20\x0d\xaa\x2a\x1e\x00\x64\x8d\xa4\xc0\x10\x75\x23\xfc\xf5\x60\x37\x05\xd3\x6d\x43\xd8\xd0\x88\xb6\x1d\x32\xc8\xa8\x60\x84\x8f\x24\xa2\xc0\x0c\xb4\x51\x3e\x20\x15\x5f\xef\xdf\x8f\x5e\xb2\xac\x43\x96\x1c\x96\x1f\x50\x9b\x80\x7e\x83\x5e\x8a\x14\x49\x59\x67\x9b\x9f\x60\x3c\xae\x0a\x26\x54\x67\xc5\x66\x22\x02\x6d\x1d\x72\x1d\x02\x03\x8f\xae\x60\x3d\x0e\x06\x91\x7a\x49\x29\x0e\x45\xca\x65\x5b\x6d\xb3\x4e\x8c\xa1\x4f\x20\xc2\xe9\xe0\x1f\x61\x33\xe1\x63\x3e\x96\xc2\x4c\x33\x5b\x0c\xe9\x92\xd4\xf2\x50\x4f\x44\x27\xd5\x84\x3d\x68\xa7\x42\x28\x58\x8d\xda\xfc\xe8\x8b\x67\xa7\x7c\x64\x98\xf2\x6e\x2e\x05\x99\xf3\xe8\x47\x55\xe3\xb5\xf8\x37\xf4\xc1\xb6\xcd\xb5\xd4\x02\x1d\xaa\x70\xf5\xd6\x1c\x83\xf6\xb6\xa3\x8b\x9b\x52\x90\x3f\xa1\xdd\x0e\xbc\x6a\xd6\x08\xbc\x77\x00\xf6\xfb\x7f\xf5\x72\xd6\x45\x35\x84\xff\xb7\x00\x0f\xb4\x3a\x2e\xc0\x51\x4f\x81\xad\x0a\x16\x67\xcd\xf2\xc4\xb9\x88\x65\x04\xb1\xdb\x01\xbf\x9b\xc3\x7e\xcf\xca\xe3\x51\x0a\x55\xc6\xce\xaa\xf3\x22\x23\x23\x1a\xdb\x73\xae\xa7\xb3\xbf\x4f\x30\xb2\xcd\x4f\x30\x06\x6e\x5f\xb0\x2e\x0c\xc7\xa6\x3a\xf9\x2c\xc5\x60\x85\xa4\x18\x6c\x97\xec\x0e\xf6\x67\xf3\x58\x6c\xb1\xcb\xc9\x55\xdb\x12\x1e\x07\x94\x10\x2b\x6f\xdb\x6e\xeb\xed\xda\x10\xfc\xf9\x0d\xd3\xf1\xe4\x15\xbc\xf5\x48\x04\x5f\x6a\x4b\x06\x6e\x1c\xcd\x1e\x97\x3a\x44\xf0\x7a\x5d\x2b\xeb\xb8\x6e\xeb\x9b\x35\xcd\x5e\x00\x87\x37\xce\xc1\x22\xde\x0d\xb0\xc0\xfe\x2d\x54\x5c\x8a\xa4\x9c\x9e\x52\x2a\x54\x8a\xf4\x01\xf8\x1b\x00\x00\xff\xff\x2a\x0f\x21\x5e\x18\x04\x00\x00")

func mechsHTMLBytes() ([]byte, error) {
	return bindataRead(
		_mechsHTML,
		"mechs.html",
	)
}

func mechsHTML() (*asset, error) {
	bytes, err := mechsHTMLBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mechs.html", size: 1048, mode: os.FileMode(436), modTime: time.Unix(1603026852, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mechCSS = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\xdb\x4e\x1b\x31\x10\x7d\x5e\xbe\xc2\x52\x85\x04\x2a\xae\xf6\x42\x12\x30\x4f\x24\x94\xdf\xa8\x1c\x7b\x36\xb6\xf0\x7a\x56\xde\xc9\x0d\xc4\xbf\x57\xbb\xf6\x86\xa4\x40\xab\xaa\x7d\x8b\x66\xce\x8e\xcf\x65\x26\x59\x66\x40\x6a\x08\xec\xe5\x2c\xcb\x18\x5b\x4a\xf5\xb4\x0a\xb8\xf6\x9a\x2b\x74\x18\x04\xfb\x52\x16\xe5\xa4\x9c\xdf\x0d\xed\x56\x6a\x6d\xfd\x4a\xb0\x2a\x6f\x77\xb1\x44\xb0\x23\x2e\x9d\x5d\x79\xc1\x14\x78\x82\x10\xeb\x35\x7a\xe2\x9d\x7d\x06\xc1\xaa\xc9\x08\x1e\x67\x4e\xef\x6f\x1f\xe7\x93\xbe\xf6\x7a\x96\x65\x35\x22\xfd\x35\x83\xe2\x8f\x0c\xc6\xcf\xef\xa7\x93\x9b\x6a\x16\x6b\x06\xec\xca\x90\x60\x93\xe1\xeb\xf4\xbe\xa1\xc6\x5d\xb1\x25\xea\x7d\xe2\x30\xa2\x8a\x3c\x3f\x1f\x49\x1e\xda\x59\x23\xc3\xca\x7a\xc1\xf2\xbe\x95\x0d\x3a\x6b\xd9\x58\xb7\x17\xac\x93\xbe\xe3\x1d\x04\x5b\x0f\x3d\xbe\x85\xe5\x93\x25\x1e\xbd\x68\x10\xc9\x0c\xdc\xa5\x27\x2b\x9d\x95\x1d\xe8\x88\x6b\xf0\x99\x63\xb7\x7b\x07\x5c\x05\xb9\xef\x94\x74\x30\xc0\x3e\xb2\xe6\xa6\x5c\x54\xd7\x51\x9b\xb6\x5d\xeb\xe4\x5e\xb0\xda\x41\xb2\xa6\xff\xc5\xb5\x0d\xa0\xc8\x62\x6f\x0f\xba\x75\xe3\x47\x49\xa6\x48\x7a\xc7\x69\xb7\xf3\x2a\x4d\x1b\xda\xe5\x15\x33\xd5\x15\x33\xd7\x51\xf7\x88\xaa\xeb\xfa\x4d\x79\x4c\x58\x49\xa7\x2e\x8a\x69\xbb\x63\x5f\x59\xb9\x69\xac\xbf\x8c\xef\x47\xab\x38\x61\x2b\x58\x51\x8e\x81\xa5\x6a\x48\x26\xff\x5a\x77\x50\x1f\x95\x93\xdd\x7c\x89\x44\xd8\x08\x96\xff\x7e\x97\x14\x6a\x88\x74\x4f\x82\x69\xd0\x63\xd7\x4a\x05\x23\x8e\xe4\xd2\xc1\x71\x9e\xff\x95\xe4\xa1\xbe\xb5\x9a\x8c\x60\xb7\x37\xe7\x09\xb6\xe3\xa9\xe4\xd1\xa7\x54\x31\x68\x08\xbc\x67\x37\x64\x1e\xd7\x0a\x9a\x96\xf6\x5c\x81\x73\x9d\x60\xc6\xea\x88\xc5\x35\x39\xeb\x21\x81\x18\xc3\x0d\x84\xda\xe1\x96\xef\x84\x5c\x13\x9e\x1a\x93\x72\x7a\x53\x4b\x26\x0a\x3e\x3e\x98\x5e\xc7\xdd\xa7\x87\xb7\xa8\xca\x6a\x71\x3a\xb5\x7a\x98\x3d\x7c\x9f\x1d\x06\x07\xe1\xc9\x70\x65\xac\xd3\x17\xa8\xf5\xe5\xcb\xa7\x63\xde\xc1\x61\x03\xfe\x43\x7c\x75\x5f\x3d\x5e\xcf\x7b\xbc\x4c\x0b\x3a\x30\xd6\xa0\x30\xc8\xb8\xc8\xa3\x7b\x6f\xbc\x66\xd3\xc5\x62\x7a\xd8\xdd\x6f\x0d\x28\xf3\x83\x2c\x1d\x42\xb6\x9e\x1f\x9d\xf5\xc6\x7c\xbc\xc2\xf9\xbb\x15\x1e\x20\xdb\xf4\xe5\x12\x9d\x3e\x79\xe2\x9f\x86\x0f\x53\x14\x7a\x02\x4f\x49\x68\x7f\xb0\x82\x15\x2c\x67\x63\x9e\x03\xe8\xe4\x1f\x72\x38\xea\xce\x04\xeb\x9f\xd2\x22\xbc\x9e\xfd\x0c\x00\x00\xff\xff\xd1\x1b\x1d\x7f\xc7\x05\x00\x00")

func mechCSSBytes() ([]byte, error) {
	return bindataRead(
		_mechCSS,
		"mech.css",
	)
}

func mechCSS() (*asset, error) {
	bytes, err := mechCSSBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mech.css", size: 1479, mode: os.FileMode(436), modTime: time.Unix(1602556043, 0)}
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
	"index.html": indexHTML,
	"mech.html":  mechHTML,
	"mechs.html": mechsHTML,
	"mech.css":   mechCSS,
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
	"index.html": {indexHTML, map[string]*bintree{}},
	"mech.css":   {mechCSS, map[string]*bintree{}},
	"mech.html":  {mechHTML, map[string]*bintree{}},
	"mechs.html": {mechsHTML, map[string]*bintree{}},
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
