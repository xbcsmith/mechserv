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

// Mode return file modify time
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

var _indexHTML = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x91\x31\x52\xeb\x30\x10\x86\xfb\x77\x8a\x7d\x2a\x52\xbd\x58\x49\xaa\xc7\x64\xed\x01\x02\x0c\x0d\x03\x13\x42\x41\xa9\x38\x9b\x48\x83\x1c\x0b\x69\x31\xf1\x91\xb8\x06\x27\x63\x64\x39\xc6\x95\xfc\xc9\xfb\x7f\xfe\x67\x8d\x7f\x6f\x1e\x57\x9b\xd7\xa7\x5b\xd0\x5c\xd9\xe2\x0f\xa6\x03\x00\x35\xa9\x5d\x7c\x00\xc0\x8a\x58\x81\x66\x76\x53\x7a\xff\x30\x4d\x2e\xca\xfa\xc8\x74\xe4\x29\xb7\x8e\x04\xf4\x94\x0b\xa6\x13\xcb\x28\x58\x42\xa9\x95\x0f\xc4\xf9\xcb\xe6\x6e\xfa\x5f\xf4\x1e\x36\x6c\xa9\x78\xa0\x52\x07\xf2\x0d\x79\x94\xe9\x26\xbd\xb5\xe6\xf8\x06\xda\xd3\x3e\x17\xca\x19\xd9\xcc\x65\xe0\xd6\x52\x56\x86\x20\xc0\x93\xcd\x45\xc7\x41\x13\x71\x67\x44\x79\xee\x88\xdb\x7a\xd7\xf6\x9a\x78\x47\x3e\x41\xc4\xc5\xe8\x83\xd0\xcc\xb3\x59\x36\x43\xa9\x17\xfd\xb4\x1c\x8f\x23\xab\xed\xb9\x4e\xa4\x5f\x6b\xc7\xae\xb8\x27\x4f\xf0\x49\xa0\xe2\x61\x58\x03\x6b\x82\x6a\xd0\x07\x94\x6e\x48\xcb\x51\x1c\xe5\xc8\x8c\x0e\x4a\xab\x42\x18\xd6\x28\x0a\x18\x82\xb8\xaf\x6b\x26\x7f\x9e\x48\x24\x8a\x55\xed\x5a\x6f\x0e\x9a\xe1\xfb\x0b\x16\xb3\xf9\x05\x5c\x7b\x62\x86\xe7\x2a\xb6\x98\x58\x5e\x9e\xb6\x65\x88\x70\x79\xa8\x94\xb1\x59\x59\x57\x93\x03\x2f\xff\x41\x06\x57\xd6\xc2\x3a\x66\x03\xac\xa9\x2b\xba\xcb\x50\x26\x73\x5a\x63\x2a\x8a\x32\xfd\xfb\x9f\x00\x00\x00\xff\xff\x5c\xb6\x45\xad\x13\x02\x00\x00")

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

	info := bindataFileInfo{name: "index.html", size: 531, mode: os.FileMode(436), modTime: time.Unix(1601250142, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mechHTML = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x94\xcf\x8e\xd3\x30\x10\xc6\xef\x3c\xc5\xe0\xc3\x9e\x68\xdc\xf6\x04\xea\x24\x02\xb6\x20\xed\x81\x3f\x2a\x0b\x12\x27\xe4\x3a\xb3\xb5\x85\x13\x07\x7b\xa8\xb6\x5a\xed\x03\xf1\x1a\x3c\x19\x72\x9c\xa8\xe1\xb2\xca\xa1\xa7\xfa\xe7\x7e\xf9\x7e\x8e\xe4\x09\x3e\xdf\x7e\xba\xbe\xfd\xfe\xf9\x1d\x18\x6e\x5c\xf5\x0c\xf3\x0f\x00\x1a\x52\x75\x5a\x00\x60\x43\xac\xc0\x30\x77\x0b\xfa\xf5\xdb\x1e\x4b\xa1\x7d\xcb\xd4\xf2\x82\x4f\x1d\x09\x18\xa8\x14\x4c\xf7\x2c\x53\xc1\x06\xb4\x51\x21\x12\x97\x5f\x6f\xdf\x2f\x5e\x8a\xa1\x87\x2d\x3b\xaa\x3e\x90\x36\x28\xf3\x3a\xef\x3b\xdb\xfe\x04\x13\xe8\xae\x14\xaa\xb3\xf2\xb8\x92\x91\x4f\x8e\x0a\x1d\xa3\x80\x40\xae\x14\x3d\x47\x43\xc4\x7d\x17\xca\xf1\x74\xb8\xf7\xf5\x69\xa8\x49\x7b\x14\x32\x24\x5c\xf7\xaa\x48\xe1\x48\x01\x8e\xab\x62\x59\x2c\x51\x9a\xf5\x90\x96\xd3\x38\xb2\xda\x8f\xc7\x49\x74\x6e\xcd\x1c\xce\x90\xd0\x80\x76\x2a\xc6\x52\x34\xa4\xcd\x8f\xfe\x55\x44\x75\xb3\x45\xc9\xe6\xff\x60\x3d\x0d\x8a\xea\xe1\x01\x8a\x9b\x2d\x3c\x3e\xa2\xe4\x7a\xd2\x2f\xa7\x82\x79\xb6\x8f\xaa\xa1\x59\xbe\x14\xbc\x88\xf1\x1b\x85\x68\x7d\x0b\xb3\xac\x63\xf8\x12\xe2\x1d\x39\x52\x91\xe6\x89\xc7\xf0\x25\xc4\x5b\x8a\x3a\xd8\x8e\xad\x6f\x67\xb9\x27\xf9\x27\xfd\x28\x27\x17\x0c\xe5\xe4\xee\x61\x37\xb6\x0e\x43\x25\x2a\x40\xd9\x0d\x7f\xde\x79\xcf\x14\xc6\x44\x26\x51\x5d\xfb\xee\x14\xec\xc1\x30\xfc\xfd\x03\xeb\xe5\xea\x15\xbc\x0d\xc4\x0c\x5f\x1a\xcb\x06\xae\x1c\x6f\xee\xf7\x3a\x26\x78\x7d\x68\x94\x75\x85\xf6\xcd\xd5\x81\x37\x2f\xa0\x80\x37\xce\xc1\x2e\x3d\x1b\x61\x47\xfd\xa4\xd4\x05\xca\xdc\x9c\x07\x2d\x1f\x14\x65\xfe\x2e\xfc\x0b\x00\x00\xff\xff\x88\x47\x0f\xa4\x2f\x04\x00\x00")

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

	info := bindataFileInfo{name: "mech.html", size: 1071, mode: os.FileMode(436), modTime: time.Unix(1601250233, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mechsHTML = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x53\x4b\x8e\x13\x31\x10\xdd\x73\x8a\xc2\x8b\x59\x11\x3b\xc9\x0a\x14\x77\x0b\x98\x80\x34\x0b\x3e\x0a\x03\x12\x2b\xe4\xb8\x2b\xb1\x85\xfb\x83\x5d\x13\x4d\x14\xe5\x40\x5c\x83\x93\x21\xb7\x9d\xa4\x95\x84\xd9\xb0\x6a\xbf\xaa\xe7\xd7\x55\xaf\x5c\xf2\xf9\xfc\xd3\xed\xfd\xf7\xcf\xef\xc0\x50\xed\xca\x67\x32\x7d\x00\xa4\x41\x55\xc5\x03\x80\xac\x91\x14\x18\xa2\x6e\x84\xbf\x1e\xec\xa6\x60\xba\x6d\x08\x1b\x1a\xd1\xb6\x43\x06\x19\x15\x8c\xf0\x91\x44\x14\x98\x81\x36\xca\x07\xa4\xe2\xeb\xfd\xfb\xd1\x4b\x96\x75\xc8\x92\xc3\xf2\x03\x6a\x13\xd0\x6f\xd0\x4b\x91\x22\x29\xeb\x6c\xf3\x13\x8c\xc7\x55\xc1\x54\x67\xc5\x66\x22\x02\x6d\x1d\x72\x1d\x02\x03\x8f\xae\x60\x3d\x0e\x06\x91\x7a\x45\x29\x0e\x35\xca\x65\x5b\x6d\xb3\x4c\x8c\xa1\x4f\x20\xc2\xe9\xe0\x87\xb0\x99\xf0\x31\x1f\x4b\x61\xa6\x99\x2d\x86\x74\x49\x6a\x79\x28\x27\xa2\x93\x6a\xc2\x1e\xb4\x53\x21\x14\xac\x46\x6d\x7e\xf4\xb5\xb3\x53\x3e\x32\x4c\x79\x37\x97\x82\xcc\x79\xf4\xa3\xaa\xf1\x5a\xfc\x1b\xfa\x60\xdb\xe6\x5a\x6a\x81\x0e\x55\xb8\x7a\x6b\x8e\x41\x7b\xdb\xd1\xc5\x4d\x29\xc8\x9f\xd0\x6e\x07\x5e\x35\x6b\x04\xde\x3b\x00\xfb\xfd\xbf\x7a\x39\xeb\xa2\x1a\xc2\xff\x9b\xff\x03\xad\x8e\xf3\x3f\xea\x29\xb0\x55\xc1\xe2\xa8\x59\x1e\x38\x17\xb1\x8c\x20\x76\x3b\xe0\x77\x73\xd8\xef\x59\x79\x3c\x4a\xa1\xca\xd8\x59\x75\x5e\x64\x64\x44\x63\x7b\xce\xf5\x74\xf6\xf7\x09\x46\xb6\xf9\x09\xc6\xc0\xed\x0b\xd6\x85\xe1\xd8\x54\x27\x9f\xa5\x18\x3c\x21\x29\x06\xaf\x4b\x76\x07\xfb\xb3\x79\x2c\xb6\xd8\xe5\xe4\xaa\x6d\x09\x8f\x03\x4a\x88\x95\xb7\x6d\xb7\xf5\x76\x6d\x08\xfe\xfc\x86\xe9\x78\xf2\x0a\xde\x7a\x24\x82\x2f\xb5\x25\x03\x37\x8e\x66\x8f\x4b\x1d\x22\x78\xbd\xae\x95\x75\x5c\xb7\xf5\xcd\x9a\x66\x2f\x80\xc3\x1b\xe7\x60\x11\xef\x06\x58\x60\xbf\x0b\x15\x97\x22\x29\xa7\x55\x4a\x85\x4a\x91\xf6\xff\x6f\x00\x00\x00\xff\xff\xd9\xde\x70\xcb\x17\x04\x00\x00")

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

	info := bindataFileInfo{name: "mechs.html", size: 1047, mode: os.FileMode(436), modTime: time.Unix(1601250158, 0)}
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
	"index.html": &bintree{indexHTML, map[string]*bintree{}},
	"mech.css":   &bintree{mechCSS, map[string]*bintree{}},
	"mech.html":  &bintree{mechHTML, map[string]*bintree{}},
	"mechs.html": &bintree{mechsHTML, map[string]*bintree{}},
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
