package platform

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

func platform_js() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xc4, 0x56,
		0x4d, 0x6f, 0xe3, 0x36, 0x13, 0xbe, 0xeb, 0x57, 0xcc, 0xeb, 0xc3, 0x4a,
		0x4e, 0x04, 0x65, 0xdf, 0xab, 0x5d, 0x77, 0x8b, 0x4d, 0xb7, 0x40, 0x0a,
		0x6c, 0x13, 0x20, 0x45, 0x7b, 0x30, 0x82, 0x82, 0x96, 0xa8, 0x98, 0xbb,
		0xb2, 0xa8, 0x8a, 0x94, 0x1d, 0xc3, 0xd6, 0x7f, 0xef, 0x0c, 0xa9, 0x0f,
		0x4a, 0x76, 0x3e, 0xb6, 0x2d, 0x50, 0x5f, 0x2c, 0x0d, 0xe7, 0xe3, 0x99,
		0x99, 0x47, 0x33, 0x0c, 0xd2, 0x2a, 0x8f, 0xb5, 0x90, 0x79, 0x50, 0x4a,
		0xa9, 0xa7, 0x70, 0xf0, 0x3c, 0x00, 0x91, 0x42, 0xf0, 0xbf, 0x9d, 0xc8,
		0x13, 0xb9, 0x5b, 0x92, 0xf8, 0x81, 0xe4, 0x80, 0x3f, 0x57, 0x06, 0x0b,
		0xab, 0x4c, 0xbf, 0xab, 0x8b, 0x8b, 0xe6, 0x09, 0x2e, 0xe0, 0x87, 0x82,
		0x95, 0x6c, 0x03, 0x9a, 0x95, 0x8f, 0x5c, 0xc3, 0xe1, 0x76, 0xf5, 0x85,
		0xc7, 0xba, 0x06, 0xbd, 0xe6, 0xad, 0x4c, 0x1a, 0x11, 0x68, 0x09, 0xb1,
		0x2c, 0xf6, 0x20, 0x72, 0x2d, 0x4f, 0xcc, 0x95, 0xac, 0xca, 0x98, 0x0f,
		0xcd, 0x1b, 0xd9, 0xc8, 0x3c, 0x2d, 0xe5, 0xe6, 0xc4, 0x5c, 0x6e, 0x79,
		0xb9, 0x2b, 0x85, 0x46, 0x0f, 0x1f, 0xa5, 0xcc, 0x38, 0xcb, 0x6b, 0x8c,
		0x93, 0x88, 0x98, 0xa1, 0x08, 0xf3, 0xdb, 0xa1, 0xb7, 0xb5, 0xac, 0xb2,
		0xc4, 0xd1, 0x5c, 0xde, 0x16, 0x54, 0x09, 0x96, 0x3d, 0x9c, 0x26, 0x53,
		0xb2, 0x5c, 0xa5, 0xb2, 0xdc, 0xc0, 0xe1, 0xa7, 0xa6, 0x60, 0x75, 0x2f,
		0x64, 0xf4, 0x0e, 0x6d, 0x25, 0x01, 0x45, 0xc0, 0x59, 0xbc, 0x06, 0x74,
		0xba, 0x39, 0xe7, 0xd5, 0x71, 0x5f, 0x72, 0x5d, 0x95, 0xf9, 0x30, 0xcd,
		0x8d, 0x4c, 0x44, 0x2a, 0x78, 0x32, 0x2c, 0x57, 0x67, 0x74, 0xd5, 0x3c,
		0xfd, 0x41, 0xd9, 0xcf, 0xba, 0xb0, 0x81, 0xd5, 0x0e, 0x9b, 0x2a, 0x85,
		0x7d, 0x62, 0x61, 0x8f, 0xb4, 0x6d, 0x24, 0xfd, 0x08, 0x66, 0xb0, 0x65,
		0x25, 0x7c, 0xe5, 0xd4, 0x83, 0xc6, 0xce, 0xd5, 0xb0, 0x4c, 0xe8, 0x0b,
		0x74, 0x3c, 0x82, 0xde, 0x17, 0x5c, 0xa6, 0x0d, 0xb2, 0x25, 0x5a, 0x22,
		0x0f, 0x16, 0x0b, 0xf0, 0xab, 0x3c, 0xe1, 0xa9, 0xc8, 0x79, 0xe2, 0x0f,
		0x1d, 0xc0, 0x50, 0xd5, 0x29, 0xe4, 0x87, 0xfe, 0x39, 0xb0, 0xa1, 0x8d,
		0xce, 0x14, 0x66, 0xe0, 0xbc, 0xce, 0x1d, 0x5f, 0xb5, 0xd7, 0x3f, 0x75,
		0x8f, 0x4d, 0x01, 0x6d, 0x94, 0x56, 0xbb, 0x0e, 0x5f, 0x62, 0x66, 0xce,
		0x36, 0x0e, 0xb1, 0xd2, 0x2a, 0xcb, 0xf6, 0xf0, 0x67, 0xc5, 0x32, 0x5b,
		0x74, 0x3a, 0x3d, 0x31, 0xd9, 0xb2, 0xac, 0x72, 0x6c, 0xec, 0x2b, 0x12,
		0x50, 0x61, 0x7b, 0xbe, 0xbd, 0xc1, 0x71, 0xc9, 0x91, 0x86, 0xc9, 0xf3,
		0x8d, 0x35, 0xe7, 0x4e, 0x6b, 0x09, 0x53, 0x68, 0xa3, 0xba, 0xe5, 0xa5,
		0xe6, 0xe5, 0x32, 0xe1, 0x80, 0x85, 0x75, 0x3f, 0xcd, 0xd0, 0x24, 0xa1,
		0x50, 0x6a, 0x52, 0xfd, 0x60, 0xfe, 0x22, 0x55, 0x64, 0x42, 0x07, 0x7e,
		0xe4, 0x53, 0x89, 0x97, 0x58, 0xd9, 0x53, 0x26, 0x08, 0x34, 0x79, 0x1f,
		0x42, 0xc6, 0xf3, 0xc6, 0x56, 0x45, 0xf8, 0xfc, 0xa8, 0xd7, 0x73, 0x3c,
		0xfa, 0x8e, 0xe4, 0xf8, 0x70, 0x79, 0x39, 0x6c, 0x31, 0x19, 0x62, 0x95,
		0x74, 0x6b, 0xb2, 0x14, 0x04, 0x40, 0x49, 0x7a, 0x47, 0x70, 0x4b, 0x3a,
		0x73, 0xa3, 0x35, 0xf3, 0x05, 0x35, 0xc6, 0x54, 0xb1, 0x46, 0x81, 0xad,
		0xee, 0xbb, 0x77, 0x18, 0xea, 0xff, 0xc8, 0xae, 0x01, 0x90, 0x29, 0x66,
		0x63, 0xcf, 0x67, 0x70, 0xa8, 0xe7, 0x43, 0xf3, 0x2e, 0x1a, 0x85, 0x56,
		0xf2, 0x3c, 0x77, 0xac, 0xde, 0x58, 0xa3, 0x1e, 0x13, 0x8a, 0x94, 0x5e,
		0xa1, 0xd3, 0xa7, 0x27, 0xcd, 0xf3, 0x44, 0x99, 0x96, 0x62, 0x58, 0x9e,
		0x6b, 0x0b, 0xb5, 0x60, 0x38, 0xa4, 0x76, 0x42, 0xaf, 0x81, 0x0d, 0x47,
		0x56, 0xd4, 0xdb, 0xde, 0xa4, 0xc6, 0xac, 0xd7, 0x6f, 0xbe, 0xf5, 0x44,
		0x72, 0x95, 0xfb, 0x1a, 0xf8, 0x93, 0x50, 0xf8, 0x2d, 0x0b, 0x8d, 0x8e,
		0xb2, 0x0c, 0x56, 0x3d, 0x69, 0x58, 0xa5, 0x25, 0x4d, 0x9c, 0x98, 0x21,
		0x6f, 0xa3, 0x33, 0x94, 0xbb, 0xa3, 0xf1, 0xf3, 0x85, 0x6f, 0x8a, 0x4c,
		0xce, 0xce, 0x1c, 0xff, 0x2a, 0xd1, 0x39, 0x01, 0x37, 0x00, 0x88, 0x2f,
		0x3d, 0x8a, 0xb3, 0xfa, 0xeb, 0xb2, 0x52, 0x3a, 0xb2, 0x36, 0xc1, 0xa1,
		0x9e, 0xce, 0x9f, 0x3b, 0xeb, 0xe8, 0x8a, 0x7d, 0x6d, 0xcb, 0x88, 0x3d,
		0x82, 0x17, 0x4c, 0x26, 0x93, 0x10, 0x5e, 0x72, 0x49, 0xe7, 0x6f, 0x71,
		0x7b, 0x2e, 0xbd, 0x78, 0x2d, 0xb2, 0x04, 0xbb, 0xd2, 0xa7, 0xa7, 0x5e,
		0xcf, 0x6f, 0x72, 0x67, 0x1a, 0xf9, 0x1a, 0xac, 0x4e, 0xeb, 0xef, 0xe4,
		0x6c, 0x8d, 0xa3, 0x6b, 0xc2, 0xf7, 0xb6, 0x40, 0x9d, 0xee, 0xb7, 0xd5,
		0xe2, 0xec, 0x26, 0x3e, 0xb6, 0x0b, 0xec, 0x78, 0xaf, 0x4b, 0x91, 0x3f,
		0xfe, 0xc3, 0xcd, 0x7c, 0x74, 0xf6, 0xe1, 0x7f, 0xb2, 0xa2, 0xff, 0x9d,
		0x65, 0x6a, 0xeb, 0xfd, 0x96, 0x6d, 0xea, 0x0e, 0x2d, 0x1a, 0x64, 0x36,
		0x58, 0x54, 0x94, 0x52, 0x4b, 0xda, 0x8e, 0x91, 0x96, 0xb6, 0xb0, 0x11,
		0x7d, 0xa1, 0x8d, 0x9f, 0xa9, 0xdd, 0x92, 0xcb, 0xa6, 0x2e, 0xd6, 0xe6,
		0xc1, 0x3f, 0xd9, 0xa8, 0x56, 0xad, 0x05, 0x31, 0xda, 0xa5, 0x4d, 0x6d,
		0x17, 0xed, 0xb2, 0xeb, 0x6c, 0x30, 0x51, 0xa1, 0xe6, 0x2f, 0x6d, 0x47,
		0x3c, 0x8f, 0xcc, 0x7d, 0x21, 0x70, 0xfc, 0x9d, 0x89, 0xac, 0x0c, 0x72,
		0x9f, 0x96, 0xb3, 0x35, 0x31, 0x43, 0xa7, 0x4b, 0x62, 0xd6, 0x28, 0x87,
		0x27, 0xa8, 0x5c, 0x49, 0x57, 0xaa, 0x4e, 0xd6, 0xb1, 0xd3, 0x8e, 0x59,
		0x33, 0xb7, 0x11, 0x65, 0x3d, 0x0d, 0x26, 0x96, 0xef, 0x13, 0xd4, 0xf0,
		0x86, 0xd4, 0xf7, 0x6f, 0xee, 0xae, 0xfd, 0xd0, 0xde, 0x31, 0x95, 0xae,
		0x56, 0xb3, 0x01, 0xf9, 0x3d, 0x27, 0xbb, 0xad, 0x14, 0x49, 0xf0, 0x7e,
		0xea, 0xd9, 0x41, 0x4d, 0x9f, 0x82, 0x77, 0x75, 0xe1, 0xdd, 0x0b, 0x9c,
		0x80, 0x1c, 0x7e, 0xbc, 0xfd, 0x0c, 0x9f, 0x59, 0x2e, 0x8a, 0x2a, 0xb3,
		0xd7, 0xb4, 0xa2, 0x14, 0x1b, 0xa1, 0xc5, 0x16, 0x57, 0x8a, 0x87, 0xed,
		0xff, 0xf9, 0xe6, 0xe3, 0xa7, 0x2e, 0x24, 0x2a, 0x9b, 0x90, 0xb4, 0x83,
		0xb1, 0x79, 0x2b, 0x16, 0x7f, 0x55, 0xb4, 0x67, 0xcc, 0xfc, 0x5f, 0x33,
		0x75, 0x9d, 0x31, 0xa5, 0x1c, 0x20, 0xc8, 0x8b, 0x98, 0x44, 0xbf, 0xe0,
		0x84, 0x19, 0x81, 0x0a, 0x7c, 0xf0, 0x2f, 0x81, 0x47, 0xdd, 0x39, 0x5c,
		0xa2, 0x64, 0x1a, 0x21, 0xc3, 0xf9, 0xd3, 0x6d, 0x6a, 0x8f, 0x47, 0x87,
		0xf0, 0x3d, 0xae, 0xdf, 0xb9, 0xd7, 0x2e, 0x1c, 0x96, 0x24, 0x6f, 0x09,
		0x68, 0xb6, 0xa9, 0x69, 0x57, 0x8b, 0x70, 0xa8, 0xd6, 0x93, 0xc8, 0x45,
		0xb3, 0x18, 0x62, 0x03, 0x8c, 0x0f, 0x0e, 0x20, 0xdb, 0xb0, 0xba, 0xc3,
		0x82, 0x17, 0x9d, 0x7b, 0xbd, 0xcf, 0xf8, 0x10, 0x0b, 0x92, 0xbe, 0xe0,
		0xa5, 0xde, 0x8f, 0xee, 0x26, 0x78, 0xd7, 0x20, 0xdd, 0x65, 0x7b, 0x4c,
		0xfb, 0xd8, 0x28, 0xcc, 0x4f, 0xfc, 0x8d, 0x92, 0x33, 0x76, 0xaa, 0xf5,
		0xd3, 0xdd, 0x4b, 0x8c, 0xd8, 0xdc, 0x51, 0x07, 0xe7, 0x60, 0x59, 0xda,
		0xfa, 0xea, 0x1c, 0xb4, 0x7e, 0x96, 0xe6, 0xef, 0x61, 0x3a, 0x4e, 0xc6,
		0xb2, 0xfa, 0x77, 0xbe, 0xfa, 0x4d, 0xf0, 0x9d, 0x03, 0x40, 0x16, 0x5a,
		0x3d, 0xcf, 0x2d, 0xcb, 0xac, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7d,
		0x74, 0x9e, 0x8a, 0x31, 0x0d, 0x00, 0x00,
	},
		"platform.js",
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
	"platform.js": platform_js,
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
	"platform.js": &_bintree_t{platform_js, map[string]*_bintree_t{}},
}}
