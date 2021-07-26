package md

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

var _effect_md_md_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\xc1\x6b\xd4\x4e\x14\xc7\xcf\x99\xbf\xe2\xfd\x06\xfa\x23\x29\x4b\xe6\xe4\x65\x61\x0f\x5a\x5a\xf1\xd0\x52\xe8\xc1\xe3\x92\x26\x2f\xdd\xa1\x33\x99\x30\xf3\xd2\x6d\x29\x0b\x1e\x14\x2a\xf5\xaa\x17\xf5\x20\x22\x7a\xf2\x26\x1e\xb4\xff\x8d\xb1\xfd\x33\x64\x26\xc9\xee\x8a\xa5\x28\x81\x4c\x66\xde\x7b\xdf\xf7\x79\xef\x4d\xea\x2c\x3f\xce\x8e\x10\x74\xc1\x98\xd4\xb5\xb1\x04\x31\x8b\x78\xa9\x89\xb3\x88\x4b\xe3\xdf\x84\xa7\x24\x08\x75\xad\x32\x42\xce\x58\xc4\x8f\x24\xcd\x9a\xc3\x34\x37\x5a\xd0\xcc\x98\xd9\xc9\x3d\x91\x9b\xea\x04\x2d\xa1\x15\x58\x96\x98\x87\xf0\x3b\xdd\x1a\x92\xca\x71\x96\x30\x26\x04\xfc\xfc\xf4\xee\xfa\xf9\x05\xa3\xb3\xda\xa3\x80\x23\xdb\xe4\x04\xe7\x6c\xc1\x58\x6e\x2a\x17\xa0\xa6\xd3\x4e\x79\x2f\xd3\x08\x13\xe0\x03\x51\x4a\xb5\xe2\x2b\xf3\x7e\x46\x33\x6f\xd6\x85\xf8\xdd\x23\x61\xac\x6c\xaa\x1c\xf6\x70\x1e\x27\xd0\x39\xa7\x8f\xb6\xc3\x0a\xe7\x2c\xb2\x48\x8d\xad\xe0\x7f\x5d\x9c\x2f\x7c\x62\x21\xe0\xfa\xdb\xd3\x9b\xf7\xcf\xba\xb0\x98\x60\x53\x17\x09\xdc\x77\x0e\xf5\xa1\x3a\x8b\xe7\x16\xa4\x49\x1f\x5b\x49\x68\x47\xa0\x4d\x81\x0a\x36\x7b\xd9\x5d\xbf\x4b\x00\xad\x35\xd6\x6b\x0b\x01\x42\x40\xfb\xf5\xc3\x8f\xab\x37\xd7\xaf\x3f\xb7\x57\x2f\xc3\xd9\xdc\xd8\x63\xcf\x3b\xf2\x9e\x30\x9e\x80\x71\xe9\x43\xa4\x79\x11\x27\xc1\x2e\x4b\xa8\xa4\x82\xff\x26\xc1\xde\xe9\x0c\x9c\xa5\xa6\x74\xdb\xeb\x97\x31\x37\x0e\x42\x98\x77\x0b\x5a\xe3\x8d\x39\x0f\x1f\x9d\xce\x82\x85\x85\x6a\xc5\x22\xaa\xd5\x83\x33\xc2\x65\xca\x8e\x78\xaa\x8b\xe9\xd0\xae\x29\xd5\xca\x03\xfc\x91\xfd\xb6\xd4\x43\xb8\x2f\xd8\x07\xde\x8a\xd0\xe7\xef\xa6\xdc\x5e\xbc\x6d\x3f\x5e\xb6\x2f\x5e\x05\x96\x25\xc7\x72\x58\x7e\x3e\xeb\xb3\x4e\xd2\x9d\xa6\xca\x5d\xbc\x74\xf0\xdb\xdd\xac\xf6\x40\x3c\xcf\x34\xaa\xad\xcc\x21\x1f\x43\xb8\x51\xe9\xd6\x70\x32\x62\x51\xa8\x3d\x49\xf7\x33\xeb\x70\x47\x2a\x74\xb1\x47\x3f\xa8\xad\xac\xa8\x8c\xf9\x86\xf3\x0f\x1f\xad\xcd\xc1\x91\x95\xd5\x51\x6c\x5c\xea\xf7\x07\x58\x67\x36\x23\x63\x93\x11\xac\xdf\xaf\xc4\xd7\xd4\xeb\xc6\x7d\x48\xdf\xd7\xe4\x6f\x1b\x37\x94\x03\x2b\xba\xbb\x7a\x77\xf3\xe5\xb2\x7d\xf2\x9d\x45\x5e\x71\xe2\x27\x99\x6e\x9f\x62\xde\x10\xc6\xf3\xe1\xf6\xfd\x73\xe6\x5e\x61\x3d\x6d\x2f\x35\xde\x38\xe9\x00\x56\xd2\x8b\xe5\x1f\x52\x49\xc5\x16\xec\x57\x00\x00\x00\xff\xff\x97\x79\x55\xa9\x3a\x04\x00\x00")

func effect_md_md_go() ([]byte, error) {
	return bindata_read(
		_effect_md_md_go,
		"effect/md/md.go",
	)
}

var _effect_md_template_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func effect_md_template_go() ([]byte, error) {
	return bindata_read(
		_effect_md_template_go,
		"effect/md/template.go",
	)
}

var _effect_md_template_tpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\xd6\x0b\x49\x4c\xca\x49\xd5\x73\xce\xcf\xcd\x4d\xcd\x2b\xa9\xad\x7d\xbf\x67\x16\x5c\xd0\x2f\x31\x37\xb5\xb6\x96\x8b\xab\x46\xe1\xe9\xda\xe9\xcf\xd6\x6d\x55\x80\x80\x1a\x85\xa7\x4b\xda\x9f\x6d\x5e\x01\x66\x3e\xdf\xb8\xfb\xe9\xbc\x6e\x88\xf0\xf3\x2d\x8b\x9e\xee\x99\xaa\x50\xc3\x55\xa3\x60\xa5\x0b\x03\x0a\x30\x0e\x9c\x01\x63\x29\xd4\x70\x55\x57\xeb\x2a\x14\x25\xe6\xa5\xa7\x2a\xe8\x39\xe7\xe7\x94\xe6\xe6\x15\x2b\xd4\xd6\x72\xd5\x28\x54\x57\x43\x2d\x07\x19\x5b\x5d\x8d\x70\x1e\x84\xeb\x5f\x94\x99\x9e\x99\x97\x98\x13\x52\x59\x90\x0a\x16\x83\xb8\x0b\x6c\x5c\x6a\x5e\x8a\x82\x6e\x6d\x2d\x20\x00\x00\xff\xff\x33\xfd\x86\x74\xda\x00\x00\x00")

func effect_md_template_tpl() ([]byte, error) {
	return bindata_read(
		_effect_md_template_tpl,
		"effect/md/template.tpl",
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
	"effect/md/md.go":        effect_md_md_go,
	"effect/md/template.go":  effect_md_template_go,
	"effect/md/template.tpl": effect_md_template_tpl,
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
		"md": &_bintree_t{nil, map[string]*_bintree_t{
			"md.go":        &_bintree_t{effect_md_md_go, map[string]*_bintree_t{}},
			"template.go":  &_bintree_t{effect_md_template_go, map[string]*_bintree_t{}},
			"template.tpl": &_bintree_t{effect_md_template_tpl, map[string]*_bintree_t{}},
		}},
	}},
}}
