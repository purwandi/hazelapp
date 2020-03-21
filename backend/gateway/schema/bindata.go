// Code generated by go-bindata.
// sources:
// schema/schema.graphql
// schema/types/issue.graphql
// schema/types/milestone.graphql
// schema/types/pagination.graphql
// schema/types/project.graphql
// schema/types/user.graphql
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

var _schemaSchemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8d\x41\xce\x82\x30\x14\x84\xf7\xef\x14\x43\xfe\x0d\xff\x15\xba\x75\xc5\x82\x44\x63\x38\x00\xc1\x17\xac\x81\x16\xdb\x57\x0d\x31\xdc\xdd\x50\xa0\xba\xc0\x55\x3b\xf3\xf5\x9b\xfa\xe6\xca\x7d\x8d\x17\x01\xf7\xc0\x6e\x54\x38\xcd\x07\x01\x7d\x90\x5a\xb4\x35\x0a\xe5\x7a\xa3\x89\x48\xc6\x81\x97\x27\xd1\x79\x68\x7e\xb2\xcb\xff\x15\x2a\xcf\x2e\x23\x20\x78\x76\x79\x67\x5b\x6d\x14\xce\xe2\xb4\x69\xb3\x44\x37\x7f\x1b\x8c\x13\x47\x67\x6f\xdc\xc8\xc1\x71\x2d\x9c\x6b\x33\x04\x51\x58\xd2\x8a\x8a\xb9\x9b\x57\xd6\x4c\xc0\x1f\x0a\xef\x03\xef\x49\x11\x24\x25\xa6\x8f\x50\x0d\x97\x2f\x61\x49\xbf\x84\x52\x77\xec\xc5\x9a\xdd\x5f\x12\x4c\x62\x6a\x68\xa2\x77\x00\x00\x00\xff\xff\x0e\x85\x32\xf0\x56\x01\x00\x00")

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

	info := bindataFileInfo{name: "schema/schema.graphql", size: 342, mode: os.FileMode(420), modTime: time.Unix(1584765701, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaTypesIssueGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x91\x41\x4e\xf3\x30\x10\x85\xf7\x3e\xc5\xfb\xf5\xef\x73\x80\xec\xaa\xa6\x0b\x4b\x88\x22\x05\x0e\x90\xc6\x53\x32\x28\xd8\x96\x67\xb2\x88\x50\xef\x8e\x6c\x0b\x1a\xa0\x0b\xc4\x6e\x32\x79\xe3\xf7\xcd\x1b\x5d\x23\xc1\x8a\x2c\x84\x37\x03\xb0\x6b\x61\xbb\x7f\xb9\x2a\xa5\xd7\x5c\x2b\xeb\x4c\x2d\x7a\x4d\xec\x9f\x73\xe3\x14\xdc\xba\xfd\x1e\x13\x0d\x4a\x6e\xa7\xd7\xe6\xc5\x18\xf2\xcb\x6b\x7d\xbb\xd7\x41\xab\xc1\x7f\xec\x3c\xb8\xf8\xe9\x34\x28\x58\x20\xca\xf3\x8c\x10\xc9\x1b\xe0\xf8\x70\xb8\xbf\x21\x9b\x06\xc1\x89\xc8\x63\x9c\x83\x90\x33\xc0\xfe\xee\xd8\x1f\xba\x6c\xc3\x3e\x2e\x8a\x7d\x41\x28\x6e\xb6\x34\xaa\xdb\xe3\x44\x88\x29\xbc\xd0\xa8\xb0\x1d\xc2\x19\x3a\x11\x12\xc5\x20\xac\x21\xad\x8d\xc1\xc7\x7f\xfb\xb9\x7b\x1d\x2b\x5b\xe3\x1c\x52\x19\x29\x30\xcd\xad\x30\xaa\x3a\x47\xf2\x55\x0c\x47\x32\x26\x8e\xca\xc1\x37\xdf\x42\xbb\x72\x3f\x45\xf7\x93\x7b\x1b\xfe\xef\x61\xfe\xca\x02\x48\xbe\x4f\xbb\xb9\x95\xb9\x98\xf7\x00\x00\x00\xff\xff\x0a\x10\x07\x40\x1a\x02\x00\x00")

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

	info := bindataFileInfo{name: "schema/types/issue.graphql", size: 538, mode: os.FileMode(420), modTime: time.Unix(1582451082, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaTypesMilestoneGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\xcc\xb1\x0a\xc2\x50\x0c\x85\xe1\x3d\x4f\x71\x7c\x8d\xae\x0a\x72\x07\x07\xf1\x09\x4a\xef\x41\x22\x9a\x1b\xd2\x74\x28\xc5\x77\x17\x05\xdb\xc5\xd1\x31\x3f\x39\x5f\xce\x4e\x9c\xf4\xce\x31\x9b\x11\x8b\x00\x5a\x3b\x94\xc3\x4e\x00\xeb\x1f\xec\x70\xc9\x50\xbb\xbe\xef\xca\x71\x08\xf5\xd4\x66\xdf\x2c\x80\x47\xbb\x71\xc8\x52\xb7\xd7\xa7\x88\x9a\x4f\x89\x7d\xb0\x4f\xae\x7e\xf9\xc4\xe5\x2f\xf4\x91\xb9\xba\xe7\x89\x31\x6f\xf8\xcf\xd9\x2b\x00\x00\xff\xff\x96\xce\x27\x33\xea\x00\x00\x00")

func schemaTypesMilestoneGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaTypesMilestoneGraphql,
		"schema/types/milestone.graphql",
	)
}

func schemaTypesMilestoneGraphql() (*asset, error) {
	bytes, err := schemaTypesMilestoneGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/types/milestone.graphql", size: 234, mode: os.FileMode(420), modTime: time.Unix(1582938477, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaTypesPaginationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x48\x4c\x4f\xf5\xcc\x4b\xcb\x57\xa8\xe6\x52\x50\x28\x2e\x49\x2c\x2a\x71\x2e\x2d\x2a\xce\x2f\xb2\x52\x08\x2e\x29\xca\xcc\x4b\xe7\x52\x50\x48\xcd\x4b\xc1\x10\xcb\x48\x2c\xf6\x4b\xad\x28\x01\xe9\xb6\x52\x70\xca\xcf\xcf\x49\x4d\xcc\x53\x84\x88\x07\x14\xa5\x96\x65\xe6\x97\x16\xa3\xc9\xd5\x72\x01\x02\x00\x00\xff\xff\x82\xe4\x59\xe2\x70\x00\x00\x00")

func schemaTypesPaginationGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaTypesPaginationGraphql,
		"schema/types/pagination.graphql",
	)
}

func schemaTypesPaginationGraphql() (*asset, error) {
	bytes, err := schemaTypesPaginationGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/types/pagination.graphql", size: 112, mode: os.FileMode(420), modTime: time.Unix(1584779905, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaTypesProjectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x51\xcb\x4e\x03\x31\x0c\xbc\xe7\x2b\x06\xf5\xce\x07\xe4\x56\xb5\x08\xed\xad\x12\x70\x42\x1c\xc2\xc6\xd9\x06\x95\x78\xe5\xb8\xaa\x10\xea\xbf\xa3\xa4\x9b\xee\xaa\xc0\x29\xf1\x63\x3c\x9e\xb1\x7e\x8d\x84\x9d\xf0\x07\xf5\x8a\x6f\x03\x44\x6f\xd1\x6d\xef\x0c\x90\xdc\x27\x59\x3c\xa9\xc4\x34\x94\x98\x4f\x89\xc4\xe2\x25\x93\x94\xd0\x53\xee\x25\x8e\x1a\x39\xb5\x2e\x03\xf4\x42\x4e\xc9\xaf\x75\x46\x9e\x8d\x59\xb2\x6c\x38\x25\xea\x0b\xac\xf2\x91\x1f\x28\x5b\xbc\x4e\xd5\x07\x3f\xd0\x5b\x21\x67\xbf\x4c\x97\xd4\xe8\x06\xea\x52\x60\x8b\xdd\xf4\x2b\x6b\x28\xab\x3b\x6c\xf8\x98\xd4\xa2\x4b\xfa\x8b\xae\x0c\xac\x44\xfd\x51\x32\xcb\x55\x1c\x7b\xb2\xad\xa7\x60\x56\xe8\xd2\x78\x54\x63\x62\x79\xb0\xa9\x3a\xa6\x7a\xad\xd4\x21\x2b\x3c\xef\x09\xdd\x16\x1c\xa0\x7b\xba\x78\x82\xc0\x52\xa3\x44\x27\x8c\x17\xc8\x7d\x33\xac\xf3\x75\xaf\x2b\xb6\xb8\xda\xd0\x37\xfd\xb7\x86\xaf\xb0\x46\xde\xb3\xe8\xd2\xeb\x7f\xb0\x7f\x5d\xe3\xdc\xc4\x3c\x92\x4e\x4a\xf2\x2c\x25\x44\xc9\xda\x76\x3b\xb8\xf9\xef\x82\x96\x3b\x5f\x4f\xfa\x4e\x81\x85\x16\x53\x7f\x02\x00\x00\xff\xff\xf1\xc7\xa3\x46\x34\x02\x00\x00")

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

	info := bindataFileInfo{name: "schema/types/project.graphql", size: 564, mode: os.FileMode(420), modTime: time.Unix(1584779669, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaTypesUserGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xc8\x31\x0a\xc2\x40\x10\x05\xd0\x7e\x4f\xf1\xd3\xe9\x15\xa6\x55\x90\x74\x82\x78\x80\x10\xbf\x32\xb2\x99\x5d\x66\x67\x0b\x11\xef\x2e\x81\x14\x62\xfb\xe2\x55\x89\x6b\xa3\xe3\x9d\x00\xbd\x09\xc6\xe3\x90\x80\x7b\xcf\xd9\xa6\x85\x82\x4b\xb8\xda\x63\xb5\xde\xe8\xff\xc6\x65\xd2\xfc\x0b\xd5\xcb\x93\x73\xb4\x9d\x5a\xed\x21\x38\x31\xce\x1b\x8d\xab\x0c\x7b\xc1\x06\x87\x62\xc6\x39\xb4\x58\xfa\xa4\x6f\x00\x00\x00\xff\xff\x5e\x0b\xab\xeb\x89\x00\x00\x00")

func schemaTypesUserGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaTypesUserGraphql,
		"schema/types/user.graphql",
	)
}

func schemaTypesUserGraphql() (*asset, error) {
	bytes, err := schemaTypesUserGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/types/user.graphql", size: 137, mode: os.FileMode(420), modTime: time.Unix(1584779174, 0)}
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
	"schema/types/milestone.graphql": schemaTypesMilestoneGraphql,
	"schema/types/pagination.graphql": schemaTypesPaginationGraphql,
	"schema/types/project.graphql": schemaTypesProjectGraphql,
	"schema/types/user.graphql": schemaTypesUserGraphql,
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
			"milestone.graphql": &bintree{schemaTypesMilestoneGraphql, map[string]*bintree{}},
			"pagination.graphql": &bintree{schemaTypesPaginationGraphql, map[string]*bintree{}},
			"project.graphql": &bintree{schemaTypesProjectGraphql, map[string]*bintree{}},
			"user.graphql": &bintree{schemaTypesUserGraphql, map[string]*bintree{}},
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

