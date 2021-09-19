// Code generated for package migrations by go-bindata DO NOT EDIT. (@generated)
// sources:
// pkg/database/migrations/20210918204516_create_users_table.down.sql
// pkg/database/migrations/20210918204516_create_users_table.up.sql
// pkg/database/migrations/20210918204526_create_files_table.down.sql
// pkg/database/migrations/20210918204526_create_files_table.up.sql
// pkg/database/migrations/20210918204532_create_metadata_table.down.sql
// pkg/database/migrations/20210918204532_create_metadata_table.up.sql
package migrations

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

var __20210918204516_create_users_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x2d\x4e\x2d\x2a\xb6\xe6\x02\x04\x00\x00\xff\xff\xcf\x0c\x8a\x87\x12\x00\x00\x00")

func _20210918204516_create_users_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20210918204516_create_users_tableDownSql,
		"20210918204516_create_users_table.down.sql",
	)
}

func _20210918204516_create_users_tableDownSql() (*asset, error) {
	bytes, err := _20210918204516_create_users_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20210918204516_create_users_table.down.sql", size: 18, mode: os.FileMode(420), modTime: time.Unix(1632010476, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __20210918204516_create_users_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcc\xb1\x0a\xc2\x30\x10\x87\xf1\x3d\x4f\xf1\x1f\x13\x70\x12\x3a\x39\x5d\xc3\x51\x83\xf1\x2a\xd7\x53\xc8\xe4\xd2\x0c\xa5\xa0\x62\x75\xf0\xed\x05\xed\xfa\xf1\xf1\x8b\xca\x64\x0c\xa3\x36\x33\xde\x4b\x7d\x2e\xde\x01\xc0\x34\xa2\x4d\xdd\xc0\x9a\x28\x6f\x7e\xe5\x75\x9f\xeb\x0d\x71\x4f\x4a\xd1\x58\x71\x21\x2d\x49\x3a\xbf\x6d\x9a\x00\xe9\x0d\x72\xce\xeb\x1a\x7b\x19\x4c\x29\x89\xfd\xc9\xeb\x63\xae\x1f\x9c\x34\x1d\x49\x0b\x0e\x5c\xe0\xa7\x31\xb8\xb0\x73\xee\x1b\x00\x00\xff\xff\xb5\x4e\xd3\xe8\x80\x00\x00\x00")

func _20210918204516_create_users_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20210918204516_create_users_tableUpSql,
		"20210918204516_create_users_table.up.sql",
	)
}

func _20210918204516_create_users_tableUpSql() (*asset, error) {
	bytes, err := _20210918204516_create_users_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20210918204516_create_users_table.up.sql", size: 128, mode: os.FileMode(420), modTime: time.Unix(1632011576, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __20210918204526_create_files_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\xcb\xcc\x49\x2d\xb6\xe6\x02\x04\x00\x00\xff\xff\x36\xa6\x2f\xb3\x12\x00\x00\x00")

func _20210918204526_create_files_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20210918204526_create_files_tableDownSql,
		"20210918204526_create_files_table.down.sql",
	)
}

func _20210918204526_create_files_tableDownSql() (*asset, error) {
	bytes, err := _20210918204526_create_files_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20210918204526_create_files_table.down.sql", size: 18, mode: os.FileMode(420), modTime: time.Unix(1632010680, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __20210918204526_create_files_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xd0\x41\x4b\xbd\x40\x14\x05\xf0\xbd\x9f\xe2\x2c\x15\xfe\x8b\x3f\xd1\x83\xa0\xd5\x6d\xbc\xda\x90\x8e\x72\x1d\x03\x57\x32\xe5\x3c\x92\x5e\x2f\x79\x1a\xd1\xb7\x0f\xf5\x51\x60\x6d\xda\xde\xf3\x9b\xe1\xde\xf3\x78\xf2\x6e\xf2\x98\xdc\xc3\xc1\x63\xdf\x1f\xfc\x18\x06\x00\xd0\x77\xb8\xd1\x69\xc5\xa2\x29\x83\x29\x2c\x4c\x9d\x65\xff\x96\xe8\x6d\xf4\xa7\x76\xcd\xb5\xb1\x9b\xf0\xe8\x5e\x3c\xd4\x2d\x09\x29\xcb\x82\x7b\x92\x46\x9b\x34\xbc\xd8\xed\xa2\x8d\x1c\xdc\xf4\xf4\x9b\xfc\x7f\x79\xb5\xa5\xeb\x92\x5d\xeb\x26\x58\x9d\x73\x65\x29\x2f\xbf\x08\x62\x4e\xa8\xce\x2c\x8e\xaf\xef\x61\x74\x5e\x71\xe8\xfe\xf6\x40\x15\xa6\xb2\x42\xf3\x3d\x4b\x09\xed\xf0\xec\x3f\x50\x8a\xce\x49\x1a\xdc\x71\x83\xb0\xef\x7e\xda\xb9\x8a\xb1\xdd\xcf\x36\x29\x84\x75\x6a\x56\x7b\xae\x28\x82\x70\xc2\xc2\x46\x71\xb5\xda\xe5\x1b\x14\x06\x31\x67\x6c\x19\x8a\x2a\x45\x31\xcf\x93\xba\x8c\xe9\x7b\x12\x44\xd7\xc1\x67\x00\x00\x00\xff\xff\x65\x5a\x9e\x5e\x9c\x01\x00\x00")

func _20210918204526_create_files_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20210918204526_create_files_tableUpSql,
		"20210918204526_create_files_table.up.sql",
	)
}

func _20210918204526_create_files_tableUpSql() (*asset, error) {
	bytes, err := _20210918204526_create_files_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20210918204526_create_files_table.up.sql", size: 412, mode: os.FileMode(420), modTime: time.Unix(1632011645, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __20210918204532_create_metadata_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\xcb\xcc\x49\x8d\xcf\x4d\x2d\x49\x4c\x49\x2c\x49\xb4\xe6\x02\x04\x00\x00\xff\xff\xfa\x58\xf3\x5b\x1a\x00\x00\x00")

func _20210918204532_create_metadata_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20210918204532_create_metadata_tableDownSql,
		"20210918204532_create_metadata_table.down.sql",
	)
}

func _20210918204532_create_metadata_tableDownSql() (*asset, error) {
	bytes, err := _20210918204532_create_metadata_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20210918204532_create_metadata_table.down.sql", size: 26, mode: os.FileMode(420), modTime: time.Unix(1632010991, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __20210918204532_create_metadata_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcf\x41\x4b\xc4\x30\x10\x05\xe0\x7b\x7f\xc5\x3b\xb6\xe0\x41\xc4\x05\xc1\xd3\x98\xce\xd6\x60\x9d\x2e\xd3\xac\xd0\x53\x89\x36\x0b\xc5\x0a\xe2\x46\xc1\x7f\x2f\x69\x15\xa1\xf4\xfa\xde\x97\xbc\xe4\xe5\x23\xf8\x18\x10\xfd\xf3\x14\x70\x1a\xa7\x70\xee\xdf\x42\xf4\x83\x8f\x3e\xcf\x00\x60\x1c\x70\x67\xab\x96\xd5\x52\x0d\x69\x1c\xe4\x58\xd7\x17\x73\x95\x78\xbf\xf4\x56\xdc\xaa\x7c\x0d\xdf\x30\xf7\xa4\x64\x1c\x2b\x9e\x48\x3b\x2b\x55\x7e\xb5\xdb\x15\x2b\xf8\xe5\xa7\xcf\xb0\x45\x2f\xaf\x6f\xd6\xd6\x34\xd2\x3a\xa5\xb4\x36\x8f\xff\x3d\xb5\x7f\x4f\x73\x07\xb5\x8f\xa4\x1d\x1e\xb8\x43\x3e\x0e\xc5\xe6\x99\x73\x7f\x4a\x76\xdf\x28\xdb\x4a\x16\xfb\xfb\x91\x02\xca\x7b\x56\x16\xc3\xed\x62\xe7\x6b\xd0\x08\x4a\xae\xd9\x31\x0c\xb5\x86\x4a\x4e\xc9\xf1\x50\xd2\x7f\x92\x15\xb7\xd9\x4f\x00\x00\x00\xff\xff\x99\x44\xaa\x5b\x4b\x01\x00\x00")

func _20210918204532_create_metadata_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20210918204532_create_metadata_tableUpSql,
		"20210918204532_create_metadata_table.up.sql",
	)
}

func _20210918204532_create_metadata_tableUpSql() (*asset, error) {
	bytes, err := _20210918204532_create_metadata_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20210918204532_create_metadata_table.up.sql", size: 331, mode: os.FileMode(420), modTime: time.Unix(1632011685, 0)}
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
	"20210918204516_create_users_table.down.sql":    _20210918204516_create_users_tableDownSql,
	"20210918204516_create_users_table.up.sql":      _20210918204516_create_users_tableUpSql,
	"20210918204526_create_files_table.down.sql":    _20210918204526_create_files_tableDownSql,
	"20210918204526_create_files_table.up.sql":      _20210918204526_create_files_tableUpSql,
	"20210918204532_create_metadata_table.down.sql": _20210918204532_create_metadata_tableDownSql,
	"20210918204532_create_metadata_table.up.sql":   _20210918204532_create_metadata_tableUpSql,
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
	"20210918204516_create_users_table.down.sql":    &bintree{_20210918204516_create_users_tableDownSql, map[string]*bintree{}},
	"20210918204516_create_users_table.up.sql":      &bintree{_20210918204516_create_users_tableUpSql, map[string]*bintree{}},
	"20210918204526_create_files_table.down.sql":    &bintree{_20210918204526_create_files_tableDownSql, map[string]*bintree{}},
	"20210918204526_create_files_table.up.sql":      &bintree{_20210918204526_create_files_tableUpSql, map[string]*bintree{}},
	"20210918204532_create_metadata_table.down.sql": &bintree{_20210918204532_create_metadata_tableDownSql, map[string]*bintree{}},
	"20210918204532_create_metadata_table.up.sql":   &bintree{_20210918204532_create_metadata_tableUpSql, map[string]*bintree{}},
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
