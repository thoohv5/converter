package ent

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _effect_ent_ent_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\x4f\x6b\x14\x31\x14\x3f\x4f\x3e\xc5\x33\x50\x99\x29\x4b\x72\xf2\xb2\xb0\x07\x2d\xad\x78\x68\x29\xf4\xe0\x71\x49\xa7\x6f\xba\xa1\x99\x64\x48\xde\x74\x5b\xca\x82\x07\x85\x4a\xbd\xea\x45\x3d\x88\x88\x9e\xbc\x89\x07\xed\xb7\x71\x6d\x3f\x86\x24\x33\xbb\x5d\xb0\x14\x65\x60\x32\x2f\xef\xbd\xdf\xef\xf7\xfe\x4c\xa3\xca\x23\x75\x88\x80\x96\x18\xd3\x75\xe3\x3c\x41\xce\x32\x5e\xd5\xc4\x59\xc6\xb5\x8b\x6f\xc2\x13\x92\x84\x75\x63\x14\x21\x67\x2c\xe3\x87\x9a\x26\xed\xbe\x28\x5d\x2d\x69\xe2\xdc\xe4\xf8\x81\x2c\x9d\x3d\x46\x4f\xe8\x25\x56\x15\x96\x29\xfd\xce\xb0\x96\xb4\x09\x9c\x15\x8c\x49\x09\xbf\xbf\x7c\xb8\x7a\x79\xce\xe8\xb4\x49\x5a\x20\x90\x6f\x4b\x82\x33\x36\x63\xac\x74\x36\x24\x55\xe3\x71\x07\xbd\xa3\x6a\x84\x11\xf0\x50\x4e\xb0\x56\x82\x1a\xc3\x6f\x9c\xbb\x8a\x26\xd1\x89\x96\xe4\x6a\x40\xc1\x58\xd5\xda\x12\x76\x70\x9a\x17\xd0\xc5\x8a\x27\x9b\xe9\x84\x33\x96\x79\xa4\xd6\x5b\xb8\x8f\x96\xce\x66\x91\x56\x4a\xb8\xfa\xf1\xfc\xfa\xe3\x8b\x2e\x2f\x27\x58\x47\x4b\x05\x3c\x0c\x01\xeb\x7d\x73\x9a\x4f\x3d\x68\x27\x9e\x7a\x4d\xe8\x07\x50\xbb\x03\x34\xb0\xde\x03\x6f\x47\xab\x00\xf4\xde\xf9\x88\x2e\x25\x48\x09\xf3\xef\x9f\x7e\x5d\xbe\xbb\x7a\xfb\x75\x7e\xf9\x3a\xdd\x4d\x9d\x3f\x8a\x82\x07\x31\x12\x86\x23\x70\x41\x3c\x46\x9a\x1e\xe4\x45\xf2\xeb\x0a\xac\x36\x70\x6f\x94\xfc\x1d\xce\x42\x69\x55\x93\xd8\x8c\xf8\x55\xce\x5d\x80\x94\x16\xc3\x12\xd6\x70\x6d\xca\xd3\x47\x87\x33\x63\xe9\xa0\xc6\xb0\x8c\x1a\xf3\xe8\x94\x70\x49\xd9\x29\x1e\xa3\xa5\x71\xd7\xaf\x31\x35\x26\xf2\xff\x45\x7e\x1b\x73\x9f\x9d\xaa\x8f\x89\xb7\x2a\xe8\xe9\xbb\x21\xcf\xcf\xdf\xcf\x3f\x5f\xcc\x5f\xbd\x49\x52\x96\x32\x16\xfb\x25\xe2\x80\x56\x27\x5d\x88\xad\xd6\x96\x21\x5f\x06\x44\x73\x5b\x35\x51\x10\x2f\x55\x8d\x66\x43\x05\xe4\x43\x48\x0b\x25\x36\x16\x37\x03\x96\xa5\xd2\x0b\xb1\xab\x7c\xc0\x2d\x6d\x30\xe4\x51\xfa\x5e\xe3\xb5\xa5\x2a\xe7\x6b\x21\x3e\x7c\xb0\x32\x86\x40\x5e\xdb\xc3\xdc\x05\x11\xed\x3d\x6c\x94\x57\xe4\x7c\x31\x80\xd5\xfd\x2a\x62\x4d\x3d\x6e\xde\xa7\xf4\x6d\x2d\xfe\xb5\x71\x8b\x72\xe0\x46\xdd\x5d\xbd\xbb\xfe\x76\x31\x7f\xf6\x93\x65\x11\x71\x14\x07\x29\x36\x4f\xb0\x6c\x09\xf3\xe9\x62\xf9\xfe\x9b\xb9\x47\x58\xa5\xed\xa1\x86\x6b\xc7\x9d\x80\x1b\xe8\xd9\xf2\x17\xb1\xda\xb0\x19\xfb\x13\x00\x00\xff\xff\x21\x5f\x55\x0f\x3a\x04\x00\x00")

func effect_ent_ent_go() ([]byte, error) {
	return bindata_read(
		_effect_ent_ent_go,
		"effect/ent/ent.go",
	)
}

var _effect_ent_schema_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func effect_ent_schema_go() ([]byte, error) {
	return bindata_read(
		_effect_ent_schema_go,
		"effect/ent/schema.go",
	)
}

var _effect_ent_schema_tpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x4d\x6f\xd4\x30\x10\x3d\xc7\xbf\x62\xe4\x53\x22\xb5\x8e\x04\x88\x43\xaf\x0b\x5c\xa0\x54\x68\x7b\xab\x2a\xd5\x9b\x4c\xb2\x56\xfd\x11\x12\x47\x62\x65\xfc\xdf\xd1\x38\xfb\x91\xa5\x5d\xd4\x05\xba\x97\x64\xc6\x9e\xf7\xde\xbc\x9d\x49\x08\x35\x36\xca\x22\xf0\xa1\x5a\xa3\x91\xc2\x77\x9a\xc7\xd8\xc9\xea\x51\xb6\x08\x53\x92\x31\x65\x3a\xd7\x7b\xc8\x19\x00\x00\x6f\x95\x5f\x8f\x2b\x51\x39\x53\x36\xb2\xc2\x95\x73\x8f\x25\x5a\x5f\xd6\x4a\x6a\xac\x3c\x67\x2c\x3b\x75\x87\x9f\x3e\x2a\x27\xb2\xb2\x51\xa8\x6b\xce\x0a\xc6\xca\x12\x42\x10\xb7\x72\xa5\x51\xdc\xca\x36\x46\x58\x3b\x5d\x0f\xe0\xd7\x3b\x65\x90\xd4\x2b\xaf\x9c\x85\xc6\xf5\xe9\xe4\xb7\x12\xb4\x5e\xf9\x8d\x60\x7e\xd3\x3d\x39\x1b\x7c\x3f\x56\x1e\x02\xcb\xd0\x7a\xb1\x9c\x9a\x8d\x89\xf8\x5a\xfd\x50\x16\x5c\xf3\x0c\xa4\x60\xcd\x68\x2b\xc8\x8f\xb3\xc5\x54\x92\x17\x70\x77\x4f\x68\x13\x40\x60\x59\x8f\x7e\xec\xed\x3c\x1b\x58\x16\x89\xe6\x79\x98\x85\xb3\x8d\x6a\xf3\x82\x94\x8b\x29\x98\xc1\x1c\x92\x81\x65\x59\x2a\xbc\x02\xbe\xc7\xf8\x2a\x0d\xc6\xc8\x2f\xb6\x0c\x65\x09\x9f\xc8\xce\xe1\xbc\x4e\xa6\x9a\x7d\x2b\x29\x7c\xd2\x4a\xca\x06\x08\xa1\x97\xb6\x45\x10\x0b\xa7\x47\x63\x87\x18\x59\x46\x53\x12\xc2\x25\xa8\x06\xf0\x3b\x24\x51\xc0\xb5\x5c\x21\xcd\x56\x9a\x21\xfa\xa5\x3f\x5a\x10\xf7\xa6\xc3\x9f\x95\x34\xa8\x17\x72\xc0\x18\x73\xba\xcb\x8b\x10\x54\x03\xe2\x03\x36\x72\xd4\x3e\xc6\xdd\x1b\xa9\xdd\x27\x8b\x10\x50\x53\x8d\xb8\xe9\x68\x0a\xa4\xce\x29\x65\xeb\x18\xc5\xc2\x19\x83\xd6\xe7\x0f\x21\xec\xde\x63\x7c\x28\xc4\xd2\xbb\x5e\xb6\xf8\x19\x37\x39\x19\xb7\xb5\xac\xb8\x38\xe8\x26\xc8\x63\xf1\xaa\x7e\xa1\xf2\x19\xe0\x29\x01\x27\x88\x08\x0b\xb8\xb2\xfe\xed\x1b\x9e\x9e\xef\xdf\x9d\xcf\xf9\x4a\x9e\xfd\x51\xb2\x57\x06\xff\xc2\x9d\x03\xf9\x76\xf3\xe8\x7a\x6e\x64\x77\x37\xf8\x5e\xd9\xf6\x7e\x7a\x04\xd8\x7e\x56\xc4\xf5\x66\xf9\xed\xcb\x15\x61\xdc\xf4\xaa\x55\x56\x6a\xaa\xa0\x69\x8f\x2f\xf5\xfa\xbf\xf9\xc9\xe7\x86\xf2\x7f\x74\xd4\xd6\x70\x49\x6b\x73\x14\xec\xf6\xf7\x63\xdd\xe2\x99\xeb\x9b\x4a\xf6\xdb\x4b\xd1\x6c\x79\xad\xd2\x2c\xb2\xad\xba\x5f\x01\x00\x00\xff\xff\x4c\xbf\xb2\xc0\xff\x05\x00\x00")

func effect_ent_schema_tpl() ([]byte, error) {
	return bindata_read(
		_effect_ent_schema_tpl,
		"effect/ent/schema.tpl",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"effect/ent/ent.go":     effect_ent_ent_go,
	"effect/ent/schema.go":  effect_ent_schema_go,
	"effect/ent/schema.tpl": effect_ent_schema_tpl,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"effect": &_bintree_t{nil, map[string]*_bintree_t{
		"ent": &_bintree_t{nil, map[string]*_bintree_t{
			"ent.go":     &_bintree_t{effect_ent_ent_go, map[string]*_bintree_t{}},
			"schema.go":  &_bintree_t{effect_ent_schema_go, map[string]*_bintree_t{}},
			"schema.tpl": &_bintree_t{effect_ent_schema_tpl, map[string]*_bintree_t{}},
		}},
	}},
}}
