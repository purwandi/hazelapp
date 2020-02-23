// Code generated by go-bindata.
// sources:
// schema/schema.graphql
// schema/types/issue.graphql
// schema/types/project.graphql
// DO NOT EDIT!

package schema

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

var _schemaSchemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x31\x8e\xc3\x20\x10\x45\xfb\x39\xc5\xa7\xf3\x4a\x7b\x02\xda\xdd\xc6\x45\xa4\x44\x51\x0e\x80\xf0\x28\x21\x92\x31\x81\xa1\xb0\x22\xdf\x3d\x02\xec\x38\x45\x2a\xe6\x0f\x6f\x9e\x7e\xb2\x37\x1e\x0d\x9e\x04\x3c\x32\xc7\x59\xe3\x54\x1e\x02\xc6\x2c\x46\xdc\xe4\x35\x0e\xeb\x44\x0b\x91\xcc\x81\x1b\x52\x6f\x42\x9c\xee\x6c\x25\x75\x3f\x1a\xc7\x36\xff\x4d\xde\xb3\x2d\xbc\xda\x81\xce\x0d\x1a\xfd\xbf\xda\x31\xf5\xb6\x6d\xfa\x2a\xdc\x24\x91\x8d\x70\xe7\xcd\xc8\x1a\x67\x89\xce\x5f\xd5\x2f\x06\x4e\x36\xba\xd0\x5a\xb5\xed\x2e\x24\xa0\x4f\x29\xf3\x7a\xea\x7c\xc8\xa2\xd1\x52\xfd\xe8\xcb\xa6\x34\xa8\x69\xc3\x2f\x61\xf8\xc0\x5b\xfa\x86\x2f\xf4\x0a\x00\x00\xff\xff\x43\x13\x5c\xd2\x2b\x01\x00\x00")

func schemaSchemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaSchemaGraphql,
		"schema/schema.graphql",
	)
}

func schemaSchemaGraphql() (*asset, error) {
	bytes, err := schemaSchemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/schema.graphql", size: 299, mode: os.FileMode(420), modTime: time.Unix(1582450747, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaTypesIssueGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xc1\x4e\xc3\x30\x0c\x86\xef\x79\x8a\x1f\x71\xef\x03\xf4\x36\xad\x3b\x44\x42\x0c\xa9\xf0\x00\x5d\xe3\x51\xa3\x92\x44\xb1\x7b\xa8\xd0\xde\x1d\x25\x11\x6c\x8c\xed\xe6\x38\xbf\xfd\xff\xfe\x74\x8d\x04\x2b\xb2\x10\xbe\x0c\xc0\xae\x85\xed\x1e\x72\x55\x4a\xaf\xb9\x56\xd6\x99\x5a\xf4\x9a\xd8\xbf\xe7\xc6\x21\xb8\xf5\xf2\x3d\x26\x1a\x94\xdc\x46\xcf\xcd\x93\x31\xe4\x97\xcf\xba\xbb\xd7\x41\xab\xc1\x23\x36\x1e\x5c\xfc\x74\x1a\x14\x2c\x10\xe5\x79\x46\x88\xe4\x0d\xb0\x7f\xd9\x3d\xdf\x90\x4d\x83\xe0\x40\xe4\x31\xce\x41\xc8\x19\x60\xfb\xb4\xef\x77\x5d\xb6\x61\x1f\x17\xc5\xb6\x44\x28\x6e\xb6\x34\xaa\xdb\xeb\x44\x88\x29\x7c\xd0\xa8\xb0\x1d\xc2\x11\x3a\x11\x12\xc5\x20\xac\x21\xad\x8d\xc1\xcf\xbf\xfd\xbd\xbd\x8e\x95\xab\x71\x0c\xa9\x8c\x94\x30\xcd\x2d\x18\x55\x9d\x91\xfc\x15\xc3\x91\x8c\x89\xa3\x72\xf0\xcd\x15\xb4\x73\xee\xb7\xe8\xfe\xe7\xbe\x0b\xff\x7a\x0d\x20\x19\x6d\x7b\x81\xd9\x9c\xcc\x77\x00\x00\x00\xff\xff\xff\x58\xe3\x53\xd5\x01\x00\x00")

func schemaTypesIssueGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaTypesIssueGraphql,
		"schema/types/issue.graphql",
	)
}

func schemaTypesIssueGraphql() (*asset, error) {
	bytes, err := schemaTypesIssueGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/types/issue.graphql", size: 469, mode: os.FileMode(420), modTime: time.Unix(1582451037, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaTypesProjectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8e\xb1\x0a\x02\x41\x0c\x44\xfb\x7c\xc5\xf8\x1b\xdb\xc9\x69\x71\x9d\x60\x29\x16\xc7\x26\x1c\x2b\x9a\x1c\xd9\x5c\x21\x72\xff\x2e\x2b\x9e\x8a\x65\x86\x79\x2f\x13\xf7\x49\x70\x70\xbb\x48\x0e\x3c\x08\x28\x9c\xd0\xef\x36\x04\xe8\x70\x93\x84\x63\x78\xd1\xb1\xdd\x2c\x35\x7b\x99\xa2\x98\xae\x31\x01\xd9\x65\x08\xe1\x6d\x7c\xab\x0b\xd1\xaf\xb6\x33\x55\xc9\x0d\x7b\x3d\x08\x8b\xe1\xda\xd9\xac\x91\xd0\x6b\x34\xb3\xf0\x28\x35\xe1\xf4\x06\xf6\x3c\xca\xf9\xdf\xd2\xc2\x86\xe7\xd9\xab\xf9\x67\xa3\xb1\xa4\xb5\x42\x0b\x3d\x03\x00\x00\xff\xff\x08\xdf\xa7\x3c\xcf\x00\x00\x00")

func schemaTypesProjectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaTypesProjectGraphql,
		"schema/types/project.graphql",
	)
}

func schemaTypesProjectGraphql() (*asset, error) {
	bytes, err := schemaTypesProjectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/types/project.graphql", size: 207, mode: os.FileMode(420), modTime: time.Unix(1582347194, 0)}
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
	"schema/schema.graphql": schemaSchemaGraphql,
	"schema/types/issue.graphql": schemaTypesIssueGraphql,
	"schema/types/project.graphql": schemaTypesProjectGraphql,
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
	"schema": &bintree{nil, map[string]*bintree{
		"schema.graphql": &bintree{schemaSchemaGraphql, map[string]*bintree{}},
		"types": &bintree{nil, map[string]*bintree{
			"issue.graphql": &bintree{schemaTypesIssueGraphql, map[string]*bintree{}},
			"project.graphql": &bintree{schemaTypesProjectGraphql, map[string]*bintree{}},
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

