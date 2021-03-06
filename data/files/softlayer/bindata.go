// Code generated by go-bindata.
// sources:
// cloud.json
// DO NOT EDIT!

package softlayer

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

var _cloudJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5c\x5d\x53\xdc\x3a\x12\x7d\xe7\x57\xa8\xe6\x89\xad\x72\xd8\xf1\xe7\x9a\xfb\x36\x97\x90\x5c\x42\x48\xd8\x40\xb2\x9b\xda\x4a\x51\x8a\x2d\x66\x9c\xf1\x48\x5e\x49\x06\x26\x29\xfe\xfb\x2d\x1b\x98\x0f\xcb\xb4\xac\xd0\x54\xa5\xf2\x92\x80\x6c\xd4\xa7\xd5\xdd\xe7\x4c\x37\xd8\x3f\x76\x08\x19\x71\xba\x60\xa3\x3f\xc8\x48\x89\x4b\x5d\xd2\x25\x93\x23\xaf\x59\x66\xfc\x6a\xf4\x07\xf9\xdf\x0e\x21\x84\x8c\x72\x76\xd5\xae\x12\x32\xfa\x3f\x1d\xed\x10\xf2\xa5\xbd\x47\xb2\x69\x21\xb8\x5a\xdd\xf7\xa3\xfd\x97\x90\x51\x29\x32\xaa\x0b\xc1\x9b\x7d\x27\x0b\xa5\x99\xcc\xe9\xc2\x23\xef\x98\x9e\x31\x59\x52\x9e\xab\xfb\xed\x56\x9b\x34\x77\xd2\xc5\xc6\xf2\x77\xc1\xd9\x7a\xe7\x76\x89\x2e\xd4\xd8\x5f\xdd\x71\xbf\x10\x8e\xee\xbf\xff\xd2\xfe\x7f\xeb\x3d\x8e\xe4\x60\xc6\x38\xa7\x85\x47\x8e\x78\x5e\xd0\x3e\x04\xd9\x8c\x81\x08\xb2\x19\x1b\xfb\xc3\x0d\xbe\xa4\x65\x49\x95\x47\x3e\x9e\x4d\xfa\xac\xe5\xb4\x04\xad\xe5\xb4\xdc\xf6\xb7\x59\x88\xbb\x0b\x49\x77\x61\xbf\xb3\xe0\x8f\x87\x23\x7e\x25\x29\x9f\x5f\xd6\x52\x7b\xe4\x35\x93\x0b\xca\x97\x7d\xc0\x2f\x25\x05\x81\x5f\x4a\x3a\x0e\x86\x1b\xfd\x4b\xf0\x29\x39\x16\x7c\xea\x91\x83\x59\xc1\x7b\x23\x33\x9b\x4f\x41\x93\xb3\xf9\xd4\xcd\x64\xad\xb4\xe0\x8f\x86\x66\x26\x6a\xd8\x9c\xa8\x5d\xcc\xbd\x15\x3c\x6f\xad\x1d\xf7\x19\x2b\x05\x07\x8d\x95\x82\xbb\x18\x3b\x61\xe5\x57\x51\x4b\xce\x3c\x32\xa9\x95\x96\xb4\xec\x4f\xf6\x05\x83\xd3\x6f\xc1\x4a\x97\x64\xff\x77\xcd\x24\xd3\x54\x0a\x8f\x9c\xb0\x9b\x22\x13\xfd\x36\x6f\x2c\x36\x6f\x5c\x6c\x9e\x14\x25\xe5\x1e\x39\xd2\xb4\xec\x4d\xd4\x45\x61\x71\xb1\x70\x72\xf1\x44\x70\x2d\x19\x2d\x3d\x72\x40\x39\xcd\xfb\x4f\xd5\x12\xcc\x85\xe0\x2e\x26\xdf\xab\x52\x78\xe4\x9d\x90\xd7\xb4\xd7\x43\xa1\x60\x0f\x85\x72\xf2\xf0\x94\xca\x42\x79\xa4\xa1\x81\x8c\xf5\xd9\xab\xa8\x04\xed\x55\x54\xba\xd8\x3b\xa3\x9c\xbc\x11\x8a\x3d\x5a\x88\xea\x5b\x06\xda\x53\xdf\xb2\x6d\x8e\x6c\x16\x1c\x34\xe1\x8c\x0a\x72\x4a\xeb\xe6\x90\x1f\x43\x40\x05\x8c\x80\x0a\x27\x8f\x99\xa8\x4b\x8f\x1c\x0b\xc9\x7a\xf3\x47\x31\x8b\x39\xe6\x68\x8e\x6a\x5d\x02\xe7\xcb\x60\x2a\x57\x8c\x3a\x99\x2b\xf8\x94\x56\x42\x32\x8f\xac\xbe\xec\x35\xcb\x61\x3a\x57\x7c\xea\x64\x76\x99\x73\xb6\xb4\xf0\x9d\x5a\xe6\xb0\xcd\x65\xee\x62\xf3\x5c\xcc\x97\xc2\x23\x6f\x68\x45\x79\x9f\x39\x2d\xe6\xa0\x39\x2d\xe6\x2e\xac\x7e\x2e\xa4\xe0\x5a\x40\xd4\xa3\x05\x5c\x9b\x5a\x38\xd5\xe6\x7f\xa8\x9a\x15\x7c\xda\xca\xe4\xcb\x83\x47\x33\xe8\x3a\x87\x2b\xf4\x3a\xef\x54\x68\xb3\x10\x75\x60\xac\x3e\x52\x16\x5c\xe9\x86\x7b\xce\x97\x15\xeb\xf9\x60\xa9\xe6\x75\x63\xd3\xcf\xfc\xc5\xda\x68\xce\x54\x26\x8b\xea\x01\xb7\x4f\x32\x21\x19\xa1\x3c\x27\xfe\xeb\x3f\xc9\x87\xc9\xc9\xfa\xd6\x8c\x6a\x36\x15\x72\x79\xff\x19\x40\x96\x4b\xf2\xa9\x90\xba\xa6\x25\x39\x63\xf2\x8a\x6d\x1c\x60\x56\x35\xa6\xfc\xb5\xc7\x74\xd1\x7c\xdf\x7b\x6c\x2b\x58\xc1\x20\x58\x01\x3a\xac\x00\x86\x15\x0d\x82\x15\xa1\xc3\x8a\x60\x58\xc9\x20\x58\x09\x3a\xac\x04\x86\x95\x0e\x82\x95\xa2\xc3\x4a\x61\x58\xfe\xb0\xe4\xf2\xf1\xb3\xcb\xb7\xa4\x97\x3f\x2c\x90\x3e\x7e\x24\x7d\x30\x94\x01\x44\x13\x41\x0b\x4c\x21\xf1\x44\xe0\xc2\x13\x01\xc4\x13\x9b\xb8\x9e\x1e\xca\x2e\x2e\x30\x92\x01\x44\x14\x9b\xb8\x9e\xce\x14\x5d\x5c\x20\x53\x04\x10\x53\x6c\xe2\x7a\x7a\x82\x75\x71\x59\xf2\x0b\xa0\x8a\x4d\x5c\x4f\xe7\x8a\x2e\x2e\x90\x2b\x02\x90\x2b\xb6\x12\x1f\x3f\xc3\x60\xb2\x08\x40\xb2\xd8\x82\x86\x1f\x4c\x98\x2d\x22\x88\x2d\x22\x54\xb6\x88\x5c\xd8\x22\x82\xd8\x22\x42\x65\x8b\x2e\x2e\x30\x94\x11\xc4\x16\x11\x2a\x5b\x74\x71\x81\x6c\x11\x41\x6c\x11\xa1\xb2\x45\x17\x97\x25\xbf\x00\xb6\x88\x50\xd9\xa2\x8b\x0b\x64\x8b\x08\x64\x8b\x08\x97\x2d\x8c\xcc\xb7\xa4\x18\xc4\x16\x11\x2e\x5b\x18\xd0\x2c\xd1\x0c\x07\x9e\x5a\x88\x7f\x6a\xa1\xad\x30\x07\x66\x5a\x84\x9f\x6a\x91\x25\xd7\x92\x81\xa4\x91\xe0\xb3\x46\x02\xd2\x46\x0a\xd1\x7f\x8a\x4a\xff\xa9\x0b\xfd\xa7\x10\xfd\xa7\xa8\xf4\xdf\xc5\x05\x66\x59\x0a\xd1\x7f\x8a\x4a\xff\x5d\x5c\x96\x38\x02\x94\x91\xa2\xd2\x7f\x17\x17\x48\x18\x29\x44\xff\x29\x2a\xfd\x77\x71\x81\x25\x99\x82\xf4\x9f\xe2\xd2\xbf\x91\xf9\x96\x14\x83\xe8\x3f\xc5\xa5\x7f\x03\x9a\x25\x9a\x10\xfd\xa7\xb8\xf4\xdf\x85\x06\xd3\x7f\x0a\xd2\x7f\x8a\x4b\xff\x46\x69\x5a\x72\x0d\xa2\xff\x14\x97\xfe\x8d\xea\x84\xa7\x51\xe0\xb0\xc0\xc7\x9d\x16\xf8\x4e\xe3\x02\x1f\x9c\x17\xf8\xb8\x03\x03\x03\x1a\x3c\xfb\x01\x47\x06\x3e\xee\xcc\xc0\x80\x66\x0b\x28\x34\x96\xc2\x1d\x1b\x18\xd0\xe0\x11\x23\x38\x38\xf0\x71\x27\x07\x06\x34\x78\xcc\x08\xcf\x0e\x7c\xe4\xe1\x81\x59\x08\xb6\x74\x03\x67\x8d\xc8\xf3\x03\x13\x9d\x2d\xac\x90\x28\x6c\xa1\x43\x50\x05\x03\x1d\x2c\x0b\x4d\xa9\x0e\xcd\x3a\x04\x61\x30\x8b\xd5\x96\x77\x90\x34\x6c\x97\xeb\x33\x50\x89\x45\x1c\x12\x50\x1c\x12\x5c\x71\x48\x9c\xc4\x21\x01\xc5\x21\xc1\x15\x87\x2e\x34\x38\xe3\x12\x50\x1c\x12\x5c\x71\xe8\x42\xb3\x05\x14\xe2\x91\x04\x57\x1c\xba\xd0\x60\x16\x49\x40\x71\x48\x70\xc5\xa1\x0b\x0d\x2e\xd2\x04\x16\x87\x04\x59\x1c\x8c\x42\xb0\xa5\x1b\x28\x0e\x09\xb2\x38\x18\xe8\x6c\x61\x05\xc5\x21\x41\x16\x87\x2e\x3a\x8b\x38\x24\xb0\x38\x24\xc8\xe2\x60\x14\xab\x2d\xef\x40\x71\x48\x90\xc5\xc1\xa8\x57\xab\x38\x04\x43\x0f\xcf\x0f\x9e\xe1\xf4\xfc\x00\x3c\xbe\x10\x6c\x6d\x42\xdc\xd6\x26\x74\x6a\x6d\x42\xb0\xb5\x09\x71\x5b\x1b\x03\x1a\x58\x12\x21\xd8\xda\x84\xb8\xad\x8d\x01\x0d\xcc\xb8\x10\x6c\x6d\x42\xdc\xd6\xc6\x80\x06\xd2\x5c\x08\xb6\x36\x21\x6e\x6b\x63\x40\xb3\x96\xc1\xd0\x64\x43\x50\x2f\xb3\x10\x6c\xe9\x06\xa9\x57\x88\xdc\xda\x98\xe8\x6c\x61\x85\xd4\x2b\x44\x6e\x6d\x0c\x74\xb0\x7a\x85\x70\x6b\x13\x22\xb7\x36\x66\xb1\xda\xf2\x0e\x52\xaf\x10\xb9\xb5\x31\xeb\xd5\xc6\x25\xa0\x7a\x75\xca\xe2\x19\x4e\xcf\xae\x5e\x41\x34\x58\x24\xa2\xe7\x90\x89\x08\xcc\xbe\x18\x6c\x0e\x63\xdc\xe6\x30\x76\x6a\x0e\x63\xb0\x39\x8c\x71\x9b\x43\x03\x9a\xed\xd4\x80\xa2\x88\x71\x9b\x43\x03\x1a\x58\x12\x31\xd8\x1c\xc6\xb8\xcd\xa1\x01\x0d\xe4\xe1\x18\x6c\x0e\x63\xdc\xe6\xd0\x80\x06\xd6\x69\x0c\x37\x87\x31\x72\x73\x68\x16\x82\xb5\x48\x87\x06\x15\x41\x5e\x4d\x74\xb6\xb0\x42\xf2\x1a\x23\x37\x87\x06\x3a\x58\x5e\x63\xb8\x39\x8c\x91\x9b\x43\xb3\x58\x6d\x79\x07\xc9\x6b\x8c\xdc\x1c\x9a\xf5\x6a\xe3\x12\x50\x5e\x63\xec\xe6\xb0\xa7\x2e\x6c\xc7\x07\xca\xeb\xb6\x48\x20\xc8\xab\x29\x13\xb0\xbc\x1e\x71\xcd\x4a\xf2\x5f\x26\x38\x39\x0c\x5f\xf8\xc1\xbf\xc6\xe4\x2a\x7c\x14\xee\x59\xc1\xa7\x25\x23\xbd\x3f\x44\x76\x23\x72\xd0\xf8\xe2\x91\x70\x2f\x1e\x93\xd7\x7f\x7d\xff\x07\xe8\xc8\x9f\x54\x32\x72\xc2\x34\xd2\xdf\xe5\x6c\x80\x22\x87\xf1\x8b\x20\x09\x40\x57\x5e\x36\x67\xd8\xff\x33\x64\xd7\x0f\x1e\x5c\x09\xf6\x22\x0c\x57\x1c\x27\xe2\x9b\x07\xfc\x00\x2b\xfa\x09\x57\x22\xb2\xeb\x27\x6b\x57\x7c\x14\x57\xdc\x4a\xf4\xb7\x4c\xb0\x9f\x08\xca\xaf\x1a\x93\x06\x56\xec\x54\x29\xeb\x1f\x21\xbb\xc1\x78\xed\x49\x88\xe1\x49\x30\x76\x23\xd8\x2e\xae\x7d\x77\x57\xf6\xef\x5d\x89\xd6\xae\x24\x28\xae\x98\x7f\x8d\xe9\xec\x8a\x63\x7e\xed\xdf\xe5\x57\x90\x62\xbb\x62\xfe\x65\xd1\x6f\x13\x15\xb7\x52\xf9\x95\x3d\x09\x62\xf0\x63\xf0\x6f\x9c\x5f\xbf\x32\x15\xaf\x5c\x59\x3d\x85\x99\x49\x96\x33\xae\x0b\x5a\xf6\x3c\x83\x59\x49\x71\x55\xe4\x4c\xb6\x6a\xb8\xf5\xd2\x90\x3b\xd7\x0a\x55\x95\x74\xf9\x4a\xc8\x05\xd5\xcd\x3d\x97\x05\x2b\x37\x9e\xbc\xa5\x9c\x0b\xdd\x3e\x64\xda\xec\xfd\xb0\x6b\xb3\xef\x8c\xca\x05\x93\x7b\xb4\xaa\x54\x26\x72\xb6\x97\x89\xc5\x3f\xb3\xb2\x56\x9a\xc9\x17\x6b\x44\xcd\x96\x0f\x4f\x8f\xde\xae\x76\x6d\x8d\x6c\x3f\x79\xba\xde\xfa\xee\x75\x26\x99\xe0\x97\xc5\xb4\x45\xfd\xfe\xd5\xf9\xdb\xc9\xe7\xc3\x0f\x17\x93\xd3\xa3\x8b\x8f\x67\x87\x1f\xde\x4d\x4e\x0e\x57\x10\xef\x36\x14\xb2\x39\x9b\xf5\x6b\x51\x2e\x68\x55\x5c\xd4\x8a\xc9\xf6\x85\x29\x5b\xf7\x7e\x53\x77\x61\xec\xbf\x5a\xd2\xaf\xac\x45\x3d\x39\x3d\x22\x1f\x7b\x6f\x29\x78\x55\xb7\x67\xa5\xd9\x8d\x1e\xad\xae\xdc\x7a\xce\xce\x1c\x1f\x7e\x1e\xe2\xc7\x9c\x2d\xfb\x5d\xa0\x55\x71\xdc\xbd\xb6\xe5\x80\x71\x75\x85\xbd\xa2\x4a\x5d\x0b\x99\x6f\xe0\xbf\xff\xaa\xfb\x90\xef\xbc\xfe\xca\x24\x67\x9a\xa9\x4f\x4c\xaa\xfe\x57\xc8\x5c\xdd\x5d\x69\x7f\x81\xb6\x97\xee\x8d\x1f\xff\xf5\xda\xf6\xd5\xbb\xf7\xd6\x6c\xa4\x55\xce\x9a\x05\x2d\x6b\xb6\xb3\x89\xaa\x45\xb3\x73\xfb\x77\x00\x00\x00\xff\xff\xae\xc4\x4b\x54\x03\x47\x00\x00")

func cloudJsonBytes() ([]byte, error) {
	return bindataRead(
		_cloudJson,
		"cloud.json",
	)
}

func cloudJson() (*asset, error) {
	bytes, err := cloudJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cloud.json", size: 18179, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
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
	"cloud.json": cloudJson,
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
	"cloud.json": {cloudJson, map[string]*bintree{}},
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
