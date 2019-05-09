// +build production
// Code generated by go-bindata.
// sources:
// templates/error.tmpl
// templates/main.tmpl
// templates/partials/footer.tmpl
// templates/partials/header.tmpl
// redirects/redirects.csv
// DO NOT EDIT!

package assets

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

var _templatesErrorTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x52\xc1\x8e\xd3\x30\x10\xbd\xe7\x2b\x46\xbe\x3b\xd5\x5e\x51\x1a\x21\x01\x12\x48\x70\x5b\xc4\x71\x35\xb1\x27\xb5\x55\xc7\x63\x79\x26\xe9\x56\x4b\xfe\x1d\xa5\xb0\x2c\xad\x7a\xe2\xf4\xc6\xcf\x6f\x34\xe3\xf7\xdc\xf9\xb8\x80\x4b\x28\xb2\x37\x05\x0f\x64\x03\xa1\xa7\x0a\x92\x50\xc9\xf4\xcd\xbf\xf7\xa7\x8a\xa5\x50\xdd\xd8\xf0\x70\xa7\xe9\xe9\x49\xa3\x26\x82\x31\xcd\x12\xac\x0d\x98\x46\x6b\x07\x56\xe5\xc9\xf4\x2f\x2f\xd0\x7e\xaa\x95\x6b\xfb\x78\x11\xad\x6b\xb7\x0b\x0f\x7d\xd3\xed\x7c\x5c\xfa\xe6\x15\xef\xcc\x83\x82\x99\x92\xb5\x13\xd6\x9b\x85\x0e\x35\x7a\xbb\xa9\xee\xf1\x8e\x13\x78\x92\xa3\x72\xb1\x17\x62\x9c\x53\xb2\xa7\xe8\x35\x6c\x72\xac\x1a\x5d\xa2\xeb\xc6\x81\x9f\x61\xe0\x67\x6b\x0b\x7a\x4f\xfe\x77\x5d\xc9\x5b\x8f\xf5\x78\x7d\xb2\x56\xa8\x60\x45\x25\x6f\x13\x8d\x0a\x42\x4e\x23\x67\xf2\x66\x7b\xcd\x7f\x6f\xa9\x27\xb6\x1a\x62\xf5\x62\xfa\xe6\xcd\xb4\x8f\x24\xae\xc6\xb2\x4d\x80\x9f\x20\x38\xd2\xe7\xc7\x6f\x5f\x61\x5d\x9b\xae\xf4\x5f\x46\x38\xf3\x0c\xa2\x31\x25\xa0\xec\x78\xce\xba\xf9\x56\x79\x48\x34\x09\x94\x44\x28\x04\x1d\x42\xa8\x34\xee\xcd\x84\x31\x29\xbf\x3b\xd1\xd0\x3a\x9e\x26\xca\x2a\xef\x39\x4b\x7b\xe0\xa5\x9d\x8f\x06\x2e\x31\xee\xcd\x07\xce\x8a\x4e\xe1\xbb\x98\xde\xfd\xa9\x67\xe9\x76\xd8\xb7\xf0\x83\x00\x0b\x27\x3e\x44\x21\x18\xb9\x02\xe6\x33\xc4\xec\x38\x2f\x94\x23\x65\x47\xa0\x21\x0a\x4c\x78\x86\x80\x0b\x81\xc3\x59\xc8\xb7\xdd\xae\xfc\xcd\xfc\x16\xde\x22\xb9\xfe\x13\xaf\xf8\x2b\x00\x00\xff\xff\x38\x7a\x05\x98\xac\x02\x00\x00")

func templatesErrorTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesErrorTmpl,
		"templates/error.tmpl",
	)
}

func templatesErrorTmpl() (*asset, error) {
	bytes, err := templatesErrorTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/error.tmpl", size: 684, mode: os.FileMode(420), modTime: time.Unix(1552397013, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x6f\x4f\xdb\x4c\x12\x7f\xfd\xf0\x29\xe6\x71\xa5\x73\x52\xbc\x76\x5c\x8e\x5c\x20\x31\x12\x6d\x51\x0f\xa9\x2d\xe8\x4a\x75\x77\x42\x48\xdd\xd8\x63\x7b\xdb\xf5\xae\xbb\x3b\x4e\x1a\x05\x7f\xf7\xd3\x3a\x09\x04\x4a\x39\x1e\xa1\xbe\xb1\xbd\xde\x99\xdf\x9f\xd9\xf1\x66\x33\xf9\xf3\xed\xd9\x9b\x8b\xff\x9e\x9f\x40\x49\x95\x3c\xda\x99\x6c\x6e\xc8\xb3\xa3\x1d\x00\x80\x09\x09\x92\x78\x74\xaa\x08\x8d\xe2\x12\x2c\x9a\x19\x1a\x40\x63\xb4\x01\x06\x67\x79\x2e\x52\x84\x5c\x1b\xf8\xc8\x49\x68\x17\xf2\x89\x38\x09\x4b\x22\xb5\x93\x68\x95\xbd\x42\xaa\x90\x38\x94\x44\x35\xc3\xef\x8d\x98\x25\xde\x7f\xd8\xe7\x63\xf6\x46\x57\x35\x27\x31\x95\xe8\x41\xaa\x15\xa1\xa2\xc4\x3b\x3d\x49\x30\x2b\x30\x48\x4b\xa3\x2b\x4c\x62\x0f\xa2\x6d\x90\xb4\xe4\xc6\x22\x25\x5e\x43\x39\x1b\x79\x77\xe7\x36\x18\x73\x91\x51\x99\x64\x38\x13\x29\xb2\x6e\x10\x08\x25\x48\x70\xc9\x6c\xca\x25\x26\x71\x38\x08\x1a\x8b\xa6\x1b\xf2\xa9\xec\x78\x14\xaf\x30\xf1\x66\x02\xe7\xb5\x36\xe4\x6d\x23\xaf\xa6\x72\x6d\x2a\x4e\x2c\x43\xc2\xd4\x19\xde\x52\x4d\x28\xb1\x2e\xb5\xc2\x44\xe9\x07\x32\xa9\xc4\x0a\x59\xaa\xa5\x36\x5b\x49\x2f\xf6\x47\xfb\x07\xfb\xaf\x1f\x88\xe7\x75\x2d\x91\x55\x7a\x2a\x24\xb2\x39\x4e\x19\xaf\x6b\x66\x89\x53\x63\xd9\x94\x1b\x66\x69\x71\xa7\x68\xbf\x46\xd2\xca\x1e\xba\x85\x13\xe9\x76\x7c\x56\xb3\xdc\x74\x83\x8c\x19\xdd\x10\x9a\x4d\xaa\x14\xea\x1b\x18\x94\x89\x67\x4b\x6d\x28\x6d\x08\x44\xea\xac\x96\x06\xf3\xc4\x73\x6b\x68\x0f\xa3\x28\xcd\x54\xa8\x95\x0d\x0b\x3d\x0b\x9b\x6f\x11\xb7\x16\xc9\x46\xa2\xe2\x05\xda\x28\xe7\x33\x97\x13\x8a\x54\x7b\x40\x8b\x1a\x13\xaf\x9b\x89\x7e\xb0\x0e\xeb\x68\x67\xe7\x8f\x3f\x96\xcb\x2f\x93\x3f\x19\xbb\x14\x39\x48\x42\x38\x3d\x81\xd1\xd5\xd1\x17\xb8\x06\xcb\x73\xfc\xe7\xc5\x87\xf7\x6d\xdb\x29\xba\xab\xc9\xf9\xb6\x25\x22\x6d\x04\x2d\x97\xe1\x39\x27\xd7\xa1\xef\xc5\xd4\x70\xb3\x38\xee\xa4\x9c\x73\x2a\xdb\x36\x4a\xad\x8d\xb4\xcc\x98\xc0\x30\xb5\xd6\x3b\xda\xf0\x5e\xa2\xca\x44\x7e\xc5\xd8\x43\x8c\x5b\xca\x4e\x4f\xe0\xe0\xf7\xa8\x12\xc8\x0e\xd6\x9a\x6e\x38\x7f\xa9\xea\x4e\xb5\x0a\x5a\xcb\x72\x2f\x7e\x8b\xb6\x8a\x0b\x75\x5f\x1b\x63\x8f\xe8\x9b\x44\xab\x5d\x63\x32\xd5\xd9\x02\x52\xc9\xad\x75\x3c\x10\x5e\x2c\x6a\x84\xb6\x75\x2b\xbe\x42\x02\xc2\xaa\x96\x9c\x10\xbc\x9a\x1b\xf7\x45\xda\x2e\x17\x8d\x07\x21\xb4\xed\xce\xba\x81\xb9\x50\x20\xb2\xc4\x73\x0f\x1e\x18\x2d\x71\xf3\x4c\x7c\x2a\x54\x86\x3f\x12\x8f\xc5\x6b\x81\x2b\xe8\x85\x40\x99\xad\x4b\x30\xe9\x3c\x3c\xca\x9a\x6b\x4d\xf7\x58\x6d\x6a\x44\x4d\x1b\xcc\x5e\xde\xa8\xee\x33\xef\x89\xc0\x06\x3a\x28\x02\x13\xf0\xa0\xea\x2f\xc5\xa5\xff\x4e\xeb\x42\xe2\xb1\xe2\x72\xe1\x36\xbb\xb3\xe9\x57\x4c\xc9\xbf\x4a\xcc\x58\x5c\x9a\xab\xc4\x5d\xae\xaf\x6f\xf2\xfb\xcb\x0d\xa4\x9b\x08\xbf\x27\xab\xdb\xf5\xf5\xe5\x55\x3f\xac\x1b\x5b\xf6\xb8\x29\x9a\x0a\x15\xd9\x7e\x1b\x74\x93\x32\x89\x5f\x2a\x9c\xc3\x5b\x4e\xd8\xeb\x8f\x79\x62\xc3\xd4\x20\x27\x3c\x91\xe8\x02\x7b\xba\x1f\xac\x41\xab\xc4\x86\x05\xd2\x7a\xc2\xbe\x5e\x5c\xf0\xe2\x23\xaf\xb0\xa7\xfb\x97\x83\xab\x31\x0f\xb9\x5d\xa8\x34\x89\xc7\x3c\xb4\x26\x4d\x8a\x71\x15\xd6\xdc\xa0\xa2\x8f\x3a\xc3\x50\x28\x8b\x86\x5e\x63\xae\x0d\xf6\x9c\xbd\x35\x6a\xdb\xef\xcd\x85\xca\xf4\x3c\xc8\x74\xda\x69\x0b\xfc\x55\x7d\xfc\xc0\x8f\xa2\xf9\x7c\x1e\x16\x5d\x11\x18\xdf\x54\x21\x4c\x75\x15\xdd\x8e\xbe\x5a\x3f\xf0\x0b\xee\xf7\xc7\x6b\xc8\x82\xf7\xfc\x95\x09\x3f\x00\xff\xf3\x31\xdb\x1f\x8e\x0e\x5e\x0d\xf6\xfe\xc1\xf6\xfc\x00\x96\x3e\x97\x52\xcf\x8f\x55\x5a\x6a\xe3\x1f\x02\x99\x06\xdb\x3b\xb9\x16\x55\xe6\x32\x6b\x5e\xa0\xdb\xa5\xbb\x24\x37\xf0\x0f\x41\xea\xb4\xfb\x0d\x0a\x6b\x4e\xa5\xdb\xf9\x60\x17\x0a\xa4\x4f\xc8\x4d\x5a\xf6\xfa\xb0\x7b\x1b\x51\x72\x5b\xde\x02\x6f\x16\x69\x3b\x7a\x79\xd3\x55\xae\x79\xa2\x97\x70\x71\xf6\xf6\x0c\x18\xfc\xbb\x44\x05\xb6\x0b\x02\x61\xa1\xd2\x33\xcc\x80\x34\x18\x54\x19\x1a\x34\x30\x47\x5f\x4a\x50\xb8\x7a\xcd\xb3\x6c\x13\x4d\x68\x2a\x10\x8a\x34\x38\xbd\xf0\xee\x18\x0c\xda\x5a\x2b\x8b\x5b\x54\x51\x04\x22\xef\xfd\xec\x24\x49\x12\xf0\xa3\x15\x92\x7f\x47\x5c\x14\x75\xb7\x19\x37\xa0\x9a\x6a\x8a\xe6\x2c\xff\x17\xda\x46\x92\x85\x04\x96\xcb\x17\x22\x77\x3c\x8d\xa4\xf0\xde\x74\xdb\xc2\x72\xf9\xc8\x14\x4a\x8b\x6d\x0b\x03\x67\x5f\xe4\x6d\x3b\xfe\x99\xd4\x20\x35\x46\xdd\x96\x75\xed\x74\x17\xfc\xbf\xdd\x43\x4c\x7c\xd8\xbd\xaf\xef\x1e\x60\x0b\x8e\x11\x1e\xf0\xf6\x30\xcd\x4f\xe9\x2f\xa3\x9b\xcd\xef\x36\x6b\xd3\xca\x3b\x3b\x37\x91\x1f\x90\xdb\xc6\x20\x90\xa8\x10\xb4\x5a\x2d\x07\x83\xd4\x60\x26\xc8\x2d\x9a\xfb\x81\x3b\x8c\xa2\x12\x65\x1d\xde\xf4\xb2\x3b\x90\x74\xdd\xbd\x6a\xf9\x9b\xf7\x51\xd5\xc1\x09\x55\x30\x07\xc8\xb4\x62\x53\xdd\xa8\x14\x99\xc3\x8d\xee\xb7\x98\x8b\x31\x71\xdc\xeb\x2f\xb7\xdb\x19\x67\xa8\xc8\x3d\x5c\x88\x0a\xcf\xd4\xb9\x6b\xe8\x00\xfc\xb8\xbb\xc4\x6c\x6f\x00\x16\x53\xad\x32\xeb\xfa\x1d\x7c\xa5\x55\x77\x22\xe3\x1d\xa6\x7f\x08\x31\xb4\xfd\x71\xfb\x20\xd7\xde\x13\xb9\x5e\xb9\xcb\x5e\xcc\x86\xcf\xe0\x1a\x3e\x91\xcb\x7d\xeb\xfe\x30\x66\xf1\xe8\x19\x64\xf1\xe8\x89\x6c\x7f\xef\xaa\x38\x72\xd6\x9e\xe3\x6d\xf0\x44\xba\xfd\xce\xdc\xa0\x73\xf7\x2c\x7b\x4f\x25\x1c\xae\xfc\x0d\xe2\xdd\xbf\xc2\xf6\x04\xe0\x41\x77\x61\xf1\x93\x5c\xac\x61\x2d\x92\x03\xd1\x0d\xf5\xd6\xad\x1e\xc4\xf1\x60\x30\xf8\x65\xc0\x5e\x1c\xec\x3d\x1a\x30\x8c\x83\xe1\xa3\x01\xf1\x28\x0e\xe2\xd1\xe3\x18\x83\x38\x18\x0e\xfe\x0f\xca\xc0\xc1\x6c\x07\x4d\xa2\xcd\x69\x60\x12\xb9\x53\x8d\xbb\xaf\xfe\x21\xfd\x2f\x00\x00\xff\xff\xee\x50\x5e\x8c\x39\x0d\x00\x00")

func templatesMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesMainTmpl,
		"templates/main.tmpl",
	)
}

func templatesMainTmpl() (*asset, error) {
	bytes, err := templatesMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/main.tmpl", size: 3385, mode: os.FileMode(420), modTime: time.Unix(1556015904, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPartialsFooterTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x57\x4d\x6f\xe3\x36\x10\xbd\xef\xaf\x20\x78\xae\xcc\x60\x03\xf4\x50\xc8\x02\x82\xb4\xd9\x2e\x1a\x24\x87\xa4\xe8\x31\x18\x93\x23\x71\x6a\x8a\x14\xc8\x91\x14\x77\xb1\xff\xbd\xa0\x15\xc7\x8e\xed\x60\x3f\xd0\x36\x46\x6f\xe4\xcc\xbc\xe1\x1b\xf3\x71\x34\x2e\xeb\x10\x18\xa3\xd0\x0e\x52\x9a\xcb\x2e\x92\xe7\xa2\xb0\x64\x50\x56\xef\x84\x10\xa5\x7d\xbf\xf1\x0d\x94\x7a\x70\x6e\x65\xc9\x18\xf4\xb2\xba\x9a\x90\x8e\xfc\x32\x95\xca\xbe\x9f\xe2\x0d\x0d\x1b\xc0\x94\x7a\xca\xb3\xe7\x1a\x23\x74\xdd\x8e\x2f\xbb\x3d\x0c\x3b\xdb\xa3\xb9\x0a\x0f\x83\xd0\xc1\x15\x19\x2f\x5f\x46\xef\x01\x74\x70\xeb\xc8\xc2\x35\x45\xf0\x58\xb0\xa5\x68\x26\x4b\x6b\xb6\x96\xc3\x24\xeb\xa2\xcf\x0f\xcf\x7d\x78\xb0\x08\x86\x7c\x23\xab\x5f\xd1\x75\xa5\xb2\xe7\x47\xb1\xbd\x3b\x86\x75\x94\xf8\xe8\x59\x19\xe2\xe8\x18\x84\x18\xdb\xd7\x20\x19\x05\xc2\x46\xac\xe7\xf2\xd3\xa7\xd9\x3d\x3c\x06\x1f\xda\xd5\xcf\xa1\x05\xf2\x9f\x3f\x2b\x8b\xae\x53\xa0\x35\xa6\x44\x0b\x72\xc4\x2b\x59\x5d\xec\x6e\x4b\x05\xaf\xb1\x51\x8e\xfe\x63\xa2\x3a\x84\x25\x61\x02\x6f\xba\x48\x03\xe8\x95\xac\x2e\x27\x93\x00\x6f\xc4\x93\xf1\xa4\x28\x33\xc6\x36\x13\xd6\xc1\x1b\x62\x0a\x3e\xc9\xea\x3e\xdb\xd6\x8c\xb7\xd6\x6f\x27\x5d\xaa\xde\x1d\x2a\x5b\x19\x1a\xde\x4c\xef\x17\x8b\xd0\xb3\xb8\xbd\xb9\x3b\x6d\xd1\x43\xa6\xd9\x27\x35\x5a\xe0\x11\x4d\x90\xd5\x1f\x16\x58\x8c\x28\x4c\x38\x15\xf1\x6c\x38\x6a\x88\x88\x31\xc9\xea\x72\x5a\x9c\x1c\xbf\xe0\x19\x34\xf7\x99\xe1\xb4\x14\xfd\xc9\x90\xf4\x38\x26\x59\xdd\xe0\x78\x32\x8c\x36\x3f\x1b\x47\xf0\xa9\x83\x88\x5e\xaf\xc0\x9b\x26\x0c\x18\x3d\x78\x8d\xaa\x8e\x88\x26\xb4\xa1\x26\x5f\x87\xd8\x42\xee\x0e\x75\x20\x59\x5d\x4d\x0e\x11\x6a\xf1\x71\xeb\xfa\x5f\xf4\x8d\xcb\xe0\x3d\x6a\x16\x23\xb1\x5d\xcb\xe7\x54\xba\x87\x65\xee\xd2\x4f\x4a\xf1\x48\xcc\x18\x67\x3a\xb4\xea\xf6\xe6\x4e\x6e\x72\x92\x0e\xfe\x69\x0a\x12\x0c\xb1\x41\x9e\xcb\x87\x85\x03\xbf\x94\xd5\xfd\x84\x79\x0b\xe5\x6d\x68\x8f\xe3\x38\xab\x41\xe3\x22\x84\xe5\x37\x71\xbf\x7a\x02\xbd\x35\xf9\x3c\x33\xa2\x21\xbf\x26\xaf\x43\xdb\x81\x5f\xa9\x50\xd7\xa4\xb1\xa8\x43\x4e\x9c\xdf\x00\xb8\x22\x31\x30\x25\x26\x9d\xbe\xaa\xbe\xeb\x75\xde\x8f\xdf\xf1\x7a\xfe\xb9\xfa\xba\x7e\xe1\x48\xcf\x9a\x30\x18\x74\x34\x60\x5c\xad\xab\x04\xad\x43\xef\x39\xa9\xdf\x7f\xbb\xbd\xb9\x53\xa9\x5f\x24\x1d\x69\x81\x31\xe5\x76\xf6\x55\xd5\xfd\xd2\x02\x39\x01\x0e\x23\xff\x8b\x73\xc5\xbe\xa9\x54\x3b\x83\xf9\x0b\xe7\x97\x06\xfa\xc3\xf1\xdd\x91\x46\x9f\x70\xef\xe7\x2c\xa9\x6d\x8e\x07\x3e\x3c\x50\xdb\x48\x01\x8e\xe7\xf2\xf6\xc3\xb5\x14\x23\x19\xb6\x73\xf9\xe3\x99\x14\x29\xea\xb9\x54\xd4\x36\x2a\x34\x6e\xd6\xe5\x7e\xf3\x32\x6b\xf7\x5a\x4e\xc6\x47\x16\x2d\xc4\x86\x7c\xe1\xb0\xe6\x22\xb5\x45\x71\x76\x78\xc7\x17\x2e\xb7\x46\xcf\xe8\x59\x50\x12\x30\x00\x39\x58\x38\x14\xbd\x37\x18\x05\x5b\xcc\xb7\x7f\xe4\xe2\xb6\x82\x78\xd2\xfb\x46\xcf\x10\xb5\xa5\x01\x53\x56\xc7\xac\x5f\x2a\x13\xb4\x0a\x1d\xfa\x62\xfa\x50\xb4\xe8\x79\xa2\xa9\x51\x0d\x18\x13\x05\xaf\xce\xd5\xa1\x12\x6e\x3b\xf4\xe2\xc3\x33\x46\x5c\x4f\x18\x31\x9c\xcf\xce\xb2\x32\x44\x99\x3a\xf0\xbb\xd4\xc4\x9a\x1f\x3e\x72\xfe\x1e\xe5\x6e\x4f\x8d\xcd\x85\x83\x73\xb2\x2a\x55\x0e\xaf\x7e\x10\xf8\xa8\xb1\x63\x31\x5a\x8c\x28\x02\x5b\x8c\x23\x25\x14\xf9\x15\xa2\xd9\x53\x49\xf7\x42\x23\xbb\xb2\x78\xde\xec\xac\xb2\x16\xc8\xe4\x3f\x91\x38\x76\x21\xe6\xa3\x9f\x35\xff\x67\x2a\xb6\x66\xfa\x0b\x33\xa1\xd7\x81\xad\xf9\x4e\xa0\x6b\xbe\x0c\x2c\xd5\x24\x96\xea\xdd\xdf\x01\x00\x00\xff\xff\x67\x15\x8a\xeb\x19\x0f\x00\x00")

func templatesPartialsFooterTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesPartialsFooterTmpl,
		"templates/partials/footer.tmpl",
	)
}

func templatesPartialsFooterTmpl() (*asset, error) {
	bytes, err := templatesPartialsFooterTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/partials/footer.tmpl", size: 3865, mode: os.FileMode(420), modTime: time.Unix(1556015904, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPartialsHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x57\x5f\x6f\xdb\x36\x10\x7f\xef\xa7\xb8\x71\x2f\xdd\x03\xad\x61\x2f\x5b\x0b\x59\xc0\xd6\x05\x68\x80\xb6\x1b\x92\xf6\x61\x28\x82\xf4\x22\x9e\x25\xce\x14\xa9\x91\xb4\x13\x43\xd3\x77\x1f\x68\xfd\x8d\x2d\x39\x09\x56\x6c\xad\x9f\xa4\xd3\xfd\xbf\x1f\x7f\x3c\xc7\x39\xa1\x20\x9b\x3c\x03\x00\x88\x11\x52\x85\xce\x2d\x99\x5b\xcb\x52\x49\xbd\x66\x90\x5b\x5a\x2d\xd9\xb7\x05\x4a\xcd\xc0\xe3\x8d\xd4\x82\xee\x96\xec\x7b\x96\x5c\xae\x65\x09\xde\x40\xf8\x04\xa9\xd1\x9e\xb4\x8f\x23\x6c\x5d\x09\xb9\x05\x29\x96\xac\xc4\x8c\x7e\x47\x9f\xb3\xce\x75\x2e\x05\xb1\xa4\xaa\x16\x1f\x2e\xce\xeb\x3a\x8e\x84\xdc\x26\xcf\x06\x9b\x56\xeb\xd6\x62\x59\x92\x65\x8d\xb7\xc3\xaf\x4d\xd2\x90\x1a\xc5\x83\xe2\x48\xeb\x50\x33\x35\x6a\xaf\xc6\x55\xc6\x8d\x26\xee\x73\x69\x45\x23\x29\xc4\x20\x39\xf0\xd0\x36\x23\xe4\xaf\x4c\x66\xf8\xb8\x15\xd1\x84\x6e\xf8\x55\xd5\xa7\xf8\x1b\xce\x3f\xca\x15\x28\x4f\x70\x7e\x06\x3f\x5d\x25\x9f\xe0\x6f\x70\xb8\xa2\xd7\xef\xdf\xbe\xa9\xeb\x49\xbb\x7d\x2c\x59\x64\x5d\xc6\x21\x20\x03\x67\xd3\x25\xcb\xbd\x2f\xdd\xcb\x28\x4a\x85\x5e\x18\xed\x16\x99\xd9\x2e\x36\xeb\x08\x9d\x23\xef\x22\x59\x60\x46\x2e\x32\xda\xf1\x60\x13\x6d\x7f\xe8\x9f\x17\xa5\xce\x18\xa0\xf2\x4b\xf6\xdb\x6a\x25\x53\x82\x95\xb1\xf0\x0e\xbd\x34\x1a\x15\x5c\x7a\xf4\xd2\x79\x99\xba\x53\xb5\x7c\x24\x2d\xe4\xea\x8a\xf3\xc7\x54\x31\xaa\x3e\x6b\xaa\x7f\x71\x95\x04\xc9\xff\xd6\x02\xb7\xfd\x0f\x5b\xd0\xe3\x7e\x10\xec\x71\x7d\x4f\x34\x03\x4b\x7f\x6b\x1a\x10\xba\x1e\x97\x23\x51\x38\x2e\x9c\xbb\x02\x4a\x2b\xb5\xe7\xbc\x39\x3e\xc7\x09\x08\xd5\xc0\x15\x75\xb6\xc1\x8c\xfa\xe3\xd6\x0b\xa6\xcb\x8c\x85\x3f\xd4\xbc\xbe\xf6\xd2\x2b\x62\xc9\xab\x1c\x75\x46\xd0\xc9\x5f\xc6\x91\xf0\x73\x5e\xc4\xb1\x17\xe9\xa9\x98\x89\xda\x34\x58\xae\x80\xfe\x82\xc5\x9b\xd6\x02\x18\x69\x56\xd7\x03\x09\x0d\xae\xc2\xf9\x03\x99\x1a\xdd\xd6\xdf\x1d\xc5\x28\xdd\x2d\xaa\x6a\x71\x29\x3d\xfd\x6a\x02\x0f\xd5\x75\x47\x2c\x2c\x79\xb5\x2b\x2c\x52\x06\xcf\x5f\xfd\xf1\x5d\x98\x4f\x55\x91\x16\x27\x10\x38\x91\x50\xba\x7b\x6a\x42\xb3\xd9\x9c\xe9\x4c\x49\x97\xc3\xf3\xb3\x77\x0f\x66\x13\x47\x42\x4c\x8c\x38\x12\xea\x49\x20\x73\x94\x1a\x2d\xd0\xee\xb8\xc6\x2d\x3c\x1e\x72\x0f\x20\x6d\xd3\x20\x4d\xe3\x96\xf7\x11\xd8\x64\xcc\xd0\x27\xe7\xe1\x4f\x17\x5e\x78\xaa\x8c\xa6\x46\x34\x07\x46\x25\x67\xfc\x3c\x80\xa5\xd1\xc5\x75\x18\x5f\xaf\xe1\x28\x81\x81\xcc\xab\x0a\x16\xef\xf1\xce\x68\x53\xec\x9a\x99\x41\x5d\x47\x96\x14\xa1\xa3\x14\x15\x05\x5f\x2c\xb9\x68\x04\xd0\x49\x8e\x8e\xfb\x30\x11\x25\xbf\xec\xda\x0a\xf2\xb9\x11\x46\x99\x6c\xc7\x92\xb7\xc3\xcb\x57\x5c\x92\xa6\x5b\x17\x6a\x11\x12\xbf\xe2\x2a\xf0\xc6\x6c\xfc\xc6\xb1\xe4\xe7\xf0\xf0\xe5\x14\x32\xb0\xdc\x64\x49\x1e\x6d\x46\x7e\xc9\xae\x6f\x14\x86\x77\x4b\x6a\xc9\xb4\x31\x25\x69\xb2\xa0\x8d\xa5\x15\x59\x4b\xb6\x2b\xbe\xbb\xd0\x6f\x94\xc9\xc6\x37\x3a\x4b\x7e\x51\x26\x7b\x62\xd9\x71\xb4\x39\x4d\x89\xa3\xd7\xf1\xe3\x88\x24\x4b\x2b\x8b\x8e\x22\xa7\x89\x2f\xd6\x78\x48\xb2\x1b\xd5\x99\x87\x7e\xf0\xb0\xfc\x5a\xa3\xa6\x56\x89\xd1\x50\xee\xa9\x9e\x1c\x4a\x8c\xdd\xda\x1d\x6c\xda\x0c\xd9\x9e\x75\x0b\xd2\x1b\xee\x4d\x96\x29\x62\x80\x56\x62\xef\xb1\x09\xd0\x2b\x4f\x07\x0d\xe6\xa7\x90\xe0\x4a\xd4\x33\xa6\x9e\xee\x7c\x38\x65\x7a\x13\x47\x41\x6d\x6e\x4e\x13\x03\x9c\x19\xde\xc9\xd6\xc0\xa3\x7a\xe3\x08\x6d\x9a\x37\xad\x69\x9e\x4f\x35\xa7\xd3\x9e\x8e\xda\x7e\xfd\x37\xdd\xb9\xdc\xbb\xf8\x0c\xfd\x99\x00\xf6\xf1\x9f\x23\x08\x39\x04\xfa\xd8\x43\x56\x90\x66\xfd\xc5\xdc\xc3\x60\xdf\x05\xba\x2b\x51\x0b\x12\x4b\xb6\x42\xe5\xe6\x2e\xf5\xe3\x03\xf1\xd8\xcb\xfa\x9e\xc9\x7e\x78\xfd\x2e\xf5\xe1\xe2\x1c\x58\xc4\xea\xfa\x48\x87\x73\x4c\xbd\xdc\x52\xbb\x08\xb5\xe4\x32\xec\xbb\x46\x09\x2e\x89\x73\x21\x5d\xa9\x70\xc7\x6f\x94\x49\xd7\xb3\xd3\x19\x38\xec\x20\x7f\xbd\x1e\x36\x9f\x42\xf0\x1f\xfb\x1d\xe8\xc5\xe8\x0f\xdd\x6b\x53\xd0\xe7\xa7\x9e\x9e\x37\xda\x2f\x71\xd4\xfd\xd9\xfe\x27\x00\x00\xff\xff\x4e\xec\x16\x9e\x76\x0f\x00\x00")

func templatesPartialsHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesPartialsHeaderTmpl,
		"templates/partials/header.tmpl",
	)
}

func templatesPartialsHeaderTmpl() (*asset, error) {
	bytes, err := templatesPartialsHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/partials/header.tmpl", size: 3958, mode: os.FileMode(420), modTime: time.Unix(1556015904, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _redirectsRedirectsCsv = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5c\xcd\x92\xdb\xba\x8e\xde\xfb\x59\xda\x51\x92\x7b\xcf\x4c\xd5\xec\x66\x52\xb3\xb8\x9b\x73\x16\x79\x80\x14\x4c\x42\x12\x62\x8a\x50\xf3\xc7\x1d\xcf\xd3\x4f\x91\xd4\xaf\x2d\xc9\xb2\xad\xce\xe9\xbb\x49\x2c\x91\xc4\xf7\x81\x3f\x20\x08\x42\x9d\xa1\x60\xcd\xd5\x39\x2b\x0c\x5b\x2b\xb9\x42\xeb\x48\xd4\x86\xa5\x17\xae\x90\x75\xf6\x5a\x91\x9d\x2b\x7c\xad\xe8\xe5\xa6\x80\x0a\x5d\xc9\x92\x15\x17\x84\x4b\x92\x76\x59\x8d\x5c\x2b\xac\xb9\xf6\x0a\x1c\xb1\x06\x2d\x05\x57\x95\xd7\xe4\xce\x59\xc9\xde\x92\x2e\x12\x9f\xf0\x80\xb5\x21\x81\xd6\x81\xa3\x20\xcb\xe6\x6c\x6c\x05\x4a\x81\x41\xb0\x91\xd9\x2a\x79\x63\x7a\x6b\x04\xef\x3a\x95\x75\x12\xab\x40\x08\xf6\xda\xd9\xcc\x1f\x2d\x0a\xc7\xa6\x7b\x11\xd9\x92\x96\xf8\x8b\x73\x8b\xe6\x14\x04\x8f\x7a\xed\xb6\x88\x31\xc1\x6b\x59\xbb\x0c\xab\x5a\xf1\xb9\x42\xed\x40\x4b\x05\x07\xf6\xa6\x02\x73\x44\xd7\xf4\x00\xe9\x37\x36\xc7\x2c\x95\x34\x7d\x4e\xa7\xd0\x09\x91\xde\xf5\xfb\xc4\xf0\x61\xa9\x63\xc6\x93\xe2\xfb\x3e\x2c\xf8\x84\x46\x07\x98\xda\x1f\x14\x89\x46\x7b\x2d\x1d\xfc\x42\x9b\x0d\x5f\xe6\xa4\x41\x0b\x4c\xac\x27\x0a\x2c\x68\xd9\x8b\x93\x98\x93\xa0\x40\x5e\xe2\xc1\x79\x2d\xd1\xb8\x12\x2b\x00\xeb\x0c\x89\xd2\x39\x83\xd0\xaa\xfa\x38\x97\xb1\xae\xdb\x91\xea\x3b\x88\x74\xde\x4d\xdf\x38\x33\x49\xcb\x30\xf4\x4d\x37\xc4\x7e\x45\x33\x2c\x19\xe9\x34\xdb\xfc\x82\xf9\xb4\x9c\x5d\x76\x08\xcb\x04\xad\x25\x2d\xbd\x75\xe6\x1c\x7a\xc3\x80\xc4\xcc\xa0\x03\x52\xed\xeb\xc4\x26\xbd\xb3\xa0\xd0\xc6\x89\x1a\x99\xac\x95\x30\x26\x74\x29\xca\x58\xba\xa3\x5b\xfa\x75\x1c\x1b\x97\x35\x3d\xd2\x29\xfe\x38\x29\xe7\x51\x03\xd0\xae\x01\xd6\x4f\x99\x00\x7f\x9c\x92\xb7\x40\xea\x00\x2a\xcc\x44\xce\x6b\x88\xeb\xb9\x61\xe5\x8f\x71\x18\x96\xb9\x5c\xb7\xbd\x24\xd3\x0a\xb9\x61\xbf\x47\xaf\x2b\x2a\x4c\xfc\x3d\x78\x1d\x76\x84\x0a\x5c\x3b\x82\x5f\x3f\x7f\xf9\x22\x50\x5b\x6f\x47\x96\x18\x75\xa1\x40\x4b\xd0\xf2\x2d\x4c\x8d\x0a\x8c\x28\x43\xd5\xdb\x06\xff\x0e\x02\x63\x15\xef\x67\x72\x6b\x2b\x43\x50\xae\x04\x2d\x2d\x0b\x02\x25\xc0\x60\x26\xc0\x5b\xb4\x9c\x4b\x04\x57\xa6\x1e\x00\x25\xb8\x64\x65\x50\x81\x43\x19\x0b\x2c\x69\x57\xa2\x3f\xae\xd8\xde\x6e\x62\x8c\x95\x5c\x06\x5b\x30\x03\x6d\x41\xf7\xa3\xdd\x9a\x1a\x25\xb4\xf6\xa0\xba\x32\x6f\x4e\x78\x5e\xb6\x0a\xf3\x02\x2f\x18\x4f\x4b\xde\x78\x1a\x26\x94\x8a\xe4\x19\xc1\x4c\xd4\x7a\xc7\x79\xb7\x02\x7a\xf5\xde\x3f\xaa\x96\x1e\x10\xdd\xb9\x1e\xab\xd9\x83\xa4\xee\x84\xda\xde\xe3\x0a\x2c\x80\x4c\x69\x36\x8d\xb6\xd1\xda\x39\x31\x49\x38\x28\xac\xd8\x38\x50\xe4\xce\xa4\x2f\x16\xec\x3b\xac\xa1\x55\xa0\x37\x14\x3c\x90\x71\xa5\x4d\x0b\x30\x4c\x18\x30\x86\xa0\x40\x9b\xa5\x57\x49\xbd\x58\x29\x2c\x58\xe3\x5c\x00\x0c\xf6\x88\x74\x0e\xc1\xc1\x08\xb5\x6e\xeb\x76\x0b\x66\xac\xd9\x0a\xbc\x35\x26\x02\x1a\x07\xd0\xd2\xff\x61\x98\x49\x2c\xd2\x5a\x48\x3a\x35\xb5\x24\x56\x5c\x18\xa8\xcb\xb5\x76\x62\x4e\xea\x85\x0a\x53\xe2\x87\xee\xc4\x09\xad\x8b\x3b\x5c\x8d\xda\x86\x49\x19\x91\xbc\x6d\xb7\xcb\x9e\x48\x5b\xf3\xc2\xa5\x58\x10\x30\x4d\x65\x24\x69\xb3\xa5\xdc\x4a\x37\x58\x90\x75\x68\xfa\xfa\x69\x99\x1d\x0c\xbe\xcf\xaa\x5e\x09\xfc\xc8\x02\x67\x2d\xc9\x35\x5d\x2a\xc9\x22\xd8\x56\x5b\x11\xbc\x13\x93\x20\x93\x65\xed\x37\xe9\x07\x57\xf8\x24\xd8\x58\xd5\x1b\xa8\xef\xa1\x62\xe8\x43\x3a\x81\xea\x80\xc2\xde\xa7\x14\x3a\xd2\xef\xae\xe8\x0d\xec\xbb\xfd\x61\x01\x35\xb9\x20\x8e\xc5\x31\x80\x37\xcf\x22\xec\x04\x55\xfd\xb4\x83\xbc\x4a\xfc\x16\x66\x58\x94\xa4\x64\x67\xec\xef\x99\x78\xf7\x99\xdf\x05\x9c\xd5\x66\x63\x78\x34\xad\xd1\x58\xd6\x1a\x55\xa3\x06\x9d\x48\xb5\x5e\xd6\x58\x89\xa7\x84\x5f\x28\x31\x87\xb2\x99\xe5\x13\x0a\xa8\x02\xed\xe2\xa4\x78\x17\x13\x77\x89\xf0\xe8\x24\x12\xac\x05\xd6\x4d\xfd\x1c\x8d\xa3\x30\xb0\xa6\x77\x3b\xfb\x0a\x9b\xcc\xaa\x25\xbc\x0b\x0d\xa7\x81\x17\x96\xb8\x05\x87\x4a\x91\xc3\x8b\x35\x1e\x57\x1b\x1a\x67\x50\xcb\x1b\x31\xaf\x6b\x11\x57\xa4\xc6\xb2\xb6\x58\xbc\xed\x19\xe7\xc4\xea\x44\xba\x10\x8a\xad\x33\x24\xc9\x57\x92\xf2\x9c\x04\x29\x5c\xe9\x27\xde\xb7\x98\xef\xc0\x7d\x07\x35\x2b\x63\xe1\x77\xa8\x35\xc4\xd9\x4e\x8d\xe6\x78\xea\x58\x1a\x5f\xd4\x4c\x96\x35\xe9\xe2\x01\xd7\xfe\x11\xdd\xd6\x83\x3f\xac\x30\x9d\xd8\xb4\x01\xce\xe6\xc1\x6e\xaa\x5d\x03\x70\xa1\xde\x02\xd2\x82\x5b\x1f\xc3\x98\xa4\x1d\x1a\x1d\x3c\x81\x61\x1c\x10\x03\x17\x34\x02\x97\x5d\xf8\x69\x09\x63\x72\x43\x51\xbd\x15\x42\x7d\x22\xc3\x31\xa4\x3a\xb0\x23\x09\x7b\xaa\x88\x35\x90\xc1\x8a\x6c\x74\xcc\x47\xd6\x68\x5a\xd4\x05\x89\x55\x32\x9f\xa1\x87\x1a\x4d\x71\xde\x8a\x58\x27\xed\x39\x4a\x83\xd7\x31\x02\xbe\x1d\xbd\x09\xc9\x5b\x58\x09\xfc\x25\xd0\xda\xb7\x38\xa5\x9e\x39\xf9\xdf\x67\x1e\x56\xa3\xde\x13\x46\x23\xad\xf9\x34\x38\x1e\xe7\x6c\x90\x0a\x2d\xc9\xa0\x70\xfd\xc9\x31\x97\x74\x67\x38\x6d\x20\x78\xac\xc6\x22\xc2\x36\x61\x99\x22\x79\x7e\xac\x6d\x8d\x82\xc2\x66\x17\x83\x08\x5d\xa7\x6d\x1f\x93\xb9\x85\x38\x3d\x4c\x6b\xfd\x46\x30\x61\x0b\x08\x53\x24\x3c\x93\x2e\x4a\xf6\x66\x1c\xf4\x4f\xad\x05\x5b\x67\x6b\x34\xa1\xfc\x2e\xcf\x74\x06\x61\xf2\x52\x70\x12\x6a\xed\x8d\x49\x77\xa7\x40\x27\x70\xd8\xdc\x8e\x9a\xb8\x42\xd3\x7d\xe8\x23\x17\x28\x6b\x84\xde\x73\xef\x1b\x96\x58\x11\xdc\x56\x19\x2f\x66\x4a\x56\x32\xc5\x16\x68\x4d\xcc\x6b\xfa\xca\x77\x56\xe6\xc3\xc1\xe4\xee\xd7\x1b\xb9\xb2\x09\xa2\xf7\xfc\x8d\x06\xd5\xd5\x78\x3e\x8e\x3c\x09\x76\xad\xe1\x2c\xea\x60\x87\x88\xff\x93\x60\xef\x6a\xef\xe2\xc0\x0e\x2e\x75\x87\x0f\x15\x82\xf5\xa6\x9f\x39\x51\x7c\xf2\xe9\x05\x57\x35\x98\xe0\x1b\xd9\xfe\x82\xaa\xbf\x53\x7e\x1c\x6a\x52\xa5\xdb\x98\x37\x46\x50\x21\x05\xf1\xc1\x5a\xb2\x37\x64\xab\x09\x95\x6a\xb0\x16\x75\x91\xc2\x20\x78\xa6\x7a\xc5\x60\x5d\xcb\x5d\xe0\x7f\x01\x70\xc3\xeb\x1a\xb6\x4c\xaf\xae\x29\xc7\xf7\xa4\x47\x19\x07\xf7\x48\x5c\x20\x7b\x2d\x7a\xc5\x02\x8e\xab\x4a\x94\x60\x40\x38\x34\xe9\x54\x99\x95\x5c\x61\xeb\xfc\x75\x06\xbd\x42\x49\xe0\x2d\x14\x23\xb5\xd0\x81\x08\x7b\x6c\xbf\x42\xa3\xe3\x28\xe9\x44\xd2\x83\x5a\xb9\xf6\x1f\x62\x31\xd5\x15\x2b\xe8\xac\xde\x43\xc2\x3f\xb5\x02\x81\x92\x6c\xed\x1d\x0e\x4c\x7d\x1f\xa0\x1b\x26\x6a\xb4\xf5\x06\x77\x69\x9b\x42\x4d\x65\x6f\x5c\x61\x6e\x16\xb3\x49\x2d\xf3\x78\xfe\x88\xc2\x55\xfe\x3e\xd1\xe9\x19\xa0\x8d\x15\x49\x2d\xdf\x51\x81\x0e\xe0\x01\x87\x4c\x92\x85\x43\x8c\xfc\x34\x94\x29\x47\xae\x6b\x36\x2e\xb4\xa1\xd1\x94\x7a\x46\xf8\x05\xef\x59\x94\x5b\xbb\x6b\x0c\x24\x42\x70\xce\xba\x85\xd6\x26\xd5\x64\xa4\x05\x57\xc1\xbe\xbe\x45\x2e\xad\x42\x31\xa2\x12\x1c\x20\xd0\x32\x67\x96\x6b\x35\xba\x0b\xeb\x52\xbf\x59\xd0\x67\xef\xa2\x6b\xc3\x3f\x51\x0c\x2c\x40\x05\x26\x45\xb6\xc1\x79\x3b\x59\xef\xd9\xcb\xe8\x21\xe4\x58\xcf\xdb\xd8\xcb\x1b\x97\x28\x41\x07\x6f\x9c\xbb\x13\x49\x85\x26\x6c\x7a\xc1\xf9\x16\xaf\x9e\xec\xd0\xd6\x4d\x97\x55\xb0\xbc\x95\xad\xc6\xb8\x50\x6d\x1e\xec\xf1\xcc\x87\x8a\xb5\x2b\xd5\xf9\x3d\x52\x1f\xe6\x44\x6f\x66\xcc\x42\xa5\x68\x2a\x7f\xf2\xe1\x7d\xcc\xf1\x25\xc2\x63\x4b\x65\xe4\x99\xf4\xaf\xa3\x0e\x36\xde\x56\xa3\xa9\xa6\x2b\x75\x3e\x70\xce\x46\xb1\x00\x05\xde\x95\x6c\xa2\x85\x7a\xd8\x19\x9f\xe1\x33\x56\xfd\x0e\x62\x77\x87\x16\xdf\xb7\x9f\xee\x8e\xa3\xfc\x8d\xdd\xf4\xfc\xb5\xd8\xb0\xa8\x97\xb6\xf5\xb5\xd8\x2c\xca\xbf\x99\xe5\xb9\xdb\x61\xe8\xde\x05\x07\x01\x7f\xd5\x28\x1c\x68\x41\xad\x12\x5c\x93\x6e\xee\xa6\x43\x05\xeb\xce\x0a\x9f\xf1\x4f\xe6\xe1\xc6\x1a\x2e\xe2\x3e\xac\x66\xf8\x69\xcf\xd6\x61\x73\xc0\x6c\xa7\xf3\x01\x0d\xc2\x09\xfb\x24\x89\x13\x77\x07\xb6\xe7\xa1\xc6\x9a\xdd\xc6\xdc\x6c\xff\x38\x81\x00\x2d\xce\xf7\x1f\x51\x56\xef\x1f\x97\x08\xab\xc3\x18\xcd\x72\x4b\xb3\xfb\x3a\x19\x7e\x54\x3e\x2c\xee\xec\x0d\x4a\x9f\x12\x8c\xee\x89\x68\xcc\xa3\x4e\x5b\x83\x75\xf0\x6b\x63\x7b\x63\xe0\xe7\x93\xc5\x67\xe4\x35\x6d\x7f\xeb\x68\x38\x76\xa0\xc6\x62\xfe\x9e\x71\x99\x26\xb2\xf9\x11\xc7\x95\x88\x79\x8e\xc2\x59\xce\xe3\xcd\x09\x68\x79\x40\x8d\x39\x39\xcb\xba\x13\x92\xda\xbe\xeb\xb9\x67\x99\x49\x93\x3a\xdf\x73\xb9\xd1\x13\xa8\xda\x53\x47\xfc\xc5\x06\xd4\x30\xbd\xaa\xcd\x58\xef\x0a\xef\x49\x8b\xb8\x25\xfb\x32\xa3\x7d\x1a\xe4\x9e\x5c\xc1\xea\xf5\x8f\xbe\xd2\xe1\x4c\xda\x7a\x13\xfa\x34\x46\x2a\x35\x61\xd7\x2e\xf7\x5a\xf6\x8d\x1f\x4d\x28\x7c\x1c\xae\x57\x2a\xf4\xc8\x28\x29\x23\x7d\xa9\x45\xb6\x66\x0b\x07\x85\x17\xc3\xd9\x7e\xef\x91\x1a\x2d\xd7\x1d\xa9\x75\x2f\xce\xe5\x57\x21\x2b\x01\xef\x19\x2d\x7e\xd3\x68\x6c\x49\x35\xe7\xfe\xf8\xea\xd9\xa1\xb4\x25\x98\x2b\xdb\xb8\x7a\x3c\x96\x04\x3e\x9d\x27\xeb\x8f\x6d\xbd\x4d\xf3\x63\x7b\xb1\x33\xf5\x9f\xbc\xbf\xf4\xc7\xd1\xa7\x7f\xc1\x03\xd3\x92\x9c\x37\xc8\xda\xa0\x45\x30\xe1\xf0\x2e\x25\x9e\x50\x71\xdd\x79\xda\x1b\xdc\x6b\x3e\x84\xbc\x1b\x7f\xd7\x78\x02\xe5\x11\xa4\x44\x59\x9c\x60\x62\xee\xf7\xe5\xfd\x47\x39\x50\xd7\x86\x41\x94\xd7\x5f\x49\x8e\xa5\x2d\xcc\xf0\x1b\x62\x77\xeb\xbe\xde\x7c\xf5\x60\x1c\x1a\x75\x1e\x46\x37\x52\xce\x67\x58\x3f\xca\x72\xde\xa4\x48\x82\xb5\xe8\x46\x91\xe5\xbb\x3e\xee\x7c\x18\x68\x97\xa5\x9f\x36\x2b\xd0\xed\xad\x0b\x62\xe4\x4b\xe9\x5c\x6d\xff\x2b\x6b\xcb\x3e\xb1\xb6\x9f\x0a\x3e\x7d\xf2\xc7\x5d\x76\x22\xeb\x41\x91\x4d\x77\xbb\x59\x98\x00\x82\xb5\x43\xed\x32\x89\xa7\x7f\x64\x7f\xfd\xf9\xfd\xc7\xf7\xbf\xbe\xfd\x60\x21\x7c\x1d\x2b\xfd\x10\x2c\x49\x17\x3f\x1c\xb3\xfa\x54\xba\x4a\x75\xf2\x59\x5b\x49\x45\xa0\xf5\xa9\x20\x57\xfa\xc3\x27\xe2\x4c\xd6\x7b\x11\x68\x52\x4e\x69\xfe\xef\x43\x43\x9b\xd9\xe0\x44\x80\x91\xfb\x5e\x32\xa8\x8b\xaa\x6b\xd0\xd7\xa9\xf0\xe7\xf7\xef\xff\xfb\xed\x87\x24\x2b\xf8\x84\xe6\xfc\xdb\xc8\xcf\xe2\x3e\x99\xa5\x71\xb1\x32\x2d\x9a\xc2\x93\x44\xc7\x13\xd9\xad\x9b\xe6\x68\x2c\x43\xfd\xf4\xea\xfc\xf5\xf3\x97\xff\x5c\x1c\x94\xef\xff\xfa\x96\x06\xf6\x5f\xdf\x7e\x94\x84\x26\xd8\x8d\xf3\x8f\x13\xe1\xdb\x16\x23\xd2\x18\x38\x9a\x99\x4c\xd3\x98\xbb\xcc\xa0\x6a\xf2\xb6\xd3\xc9\x3c\xa7\xc2\x1b\xb4\xac\xdb\x9c\xb5\x26\x17\x5f\x1e\xce\x60\x10\x38\xf7\x41\x3f\x83\x96\x24\x6a\x81\x17\x91\x93\xaf\x9f\xbf\xfc\x91\x3e\xce\x79\xd9\x58\xf2\x4f\xaf\xf1\xeb\xe7\x2f\xff\x51\x1b\x3e\x91\x8d\xf3\x6e\x97\xa5\xaf\xea\xb2\x12\x55\x9d\x39\x34\x55\x4c\xd6\xee\x6e\xa2\x5e\xda\xf2\xf4\x9f\x33\xa0\x6d\xce\xa6\x6a\xc3\xce\x85\x81\xaa\xc2\x89\x86\x63\xc1\x82\xf9\x48\xd1\x3b\x8d\x69\x08\xe2\xbc\x52\x6e\x53\x3b\x4c\x92\x78\x56\x1e\x49\x5d\x29\x23\xcc\xa9\xa6\x06\x5a\xb7\xeb\x4e\x23\x82\x75\xfe\x92\x85\x53\xb0\xf3\x36\x7b\x2b\xc1\xbd\xa1\xe4\x0c\x4f\xf1\x0b\xcb\x61\x2d\x34\xa1\x33\x77\x99\xb6\xe5\xdb\x4b\x67\x25\x49\x77\x60\x39\x9b\xb9\x6b\xc8\xfe\xd2\x73\xfc\xbe\x11\xd2\xc6\x03\xac\xf3\xf2\xcc\x79\x17\x51\x78\x43\xa5\x0e\x48\xba\x88\x69\xe7\x06\xc3\xb2\x3b\xb3\xd7\x45\x5a\x8e\x8f\xb6\x4b\x3a\x94\xe6\x23\x6a\x11\xe0\xaf\x1a\x8e\xe9\x0d\xda\xa6\x1b\x1f\xd2\x25\xa8\xb0\xa5\x75\x53\x7e\x17\x96\x7e\x58\x37\x7d\x52\xc6\x97\xfe\xe7\xbe\x37\x37\xfb\x6e\xa3\xdc\x1b\xac\xd9\xb8\x2c\xc7\x83\xf1\x60\xce\xfb\xb0\x46\x32\xeb\x0e\xfb\xca\xbe\x9a\xfd\xe8\x75\xb2\x32\x5b\x06\x5c\xbb\x8f\x38\x7a\x96\x83\x23\x4f\xcb\xf1\x82\x62\xa0\xb2\x99\xa6\x31\x41\xe8\x43\x2a\xa6\x2a\x9b\xc5\x31\xdf\x27\x5e\x7b\xcd\x6e\x4f\x7a\xdf\xc5\x5f\xf6\xfb\x3e\x58\xb5\x67\xb3\x77\x06\x48\x93\x2e\xf6\x7b\x8d\xe8\xec\xfe\x7a\x4c\x1b\x3d\x6f\x85\xc2\x34\xbb\x26\x1a\xe6\x75\x5f\x77\xa0\xd3\x60\x2a\xc6\xba\x1d\xa5\xbe\x36\x9b\x96\x4e\x20\x33\xa3\xa2\x06\xf8\x9a\xd9\x60\x6d\xe4\xbe\x8d\x63\xec\x39\xdf\x47\xc7\xf1\x9f\x81\xf6\x1f\x71\x2a\x7e\x4d\xe5\xe1\xfd\xbe\x79\xdf\xaa\x72\xc3\x3d\xec\x39\x27\x98\x16\x85\xf3\x88\x91\x06\xe2\x9f\x2c\x9c\x63\x89\x22\x08\xee\xd9\x1d\xbc\xdd\xa7\x33\x57\xe7\xe2\xef\xfb\x33\x58\x4b\x70\x3f\xd8\x50\xf6\x06\xad\x57\xce\x8e\x66\xd5\x7a\x8a\xd7\x9f\xe5\x4d\x32\x1c\x00\x36\x78\xe3\x49\x93\x46\x74\x9f\x86\x74\xb0\x14\x16\x66\xfe\xf3\xc1\xd1\x5e\x0b\x7f\x1c\x09\xb8\x1e\x79\xa8\xa9\x09\x50\x65\x6f\x78\x18\x3e\x96\x5c\xe1\x4b\xda\x34\xe3\x8d\x58\xcf\x3d\xb5\x1c\x52\x5e\x1b\x7a\x6c\x3f\xdd\xb8\x0a\x11\xde\x1d\x74\x9c\x96\x44\x5a\x28\x2f\xd1\x82\x52\xff\x68\xda\x89\x9a\x4a\x51\x13\x68\x69\x46\x7f\xae\x81\x44\xbb\x26\x26\xb6\xde\x6e\xcf\x4e\x8e\x42\xbc\xd3\xee\x77\xe2\xbe\x27\xc6\x3b\x7d\x86\xba\x80\x82\x74\xf1\x46\xae\xbc\xae\x7c\xb9\xa1\x8f\xc4\x44\x26\x3d\xb7\x28\x29\x3a\x1a\xbf\x8f\xdd\xd5\xab\xc4\x77\x40\x2a\x3e\x7f\x00\x42\xdd\x55\xcc\x6a\x2e\xac\x15\x69\x9c\xbc\xcd\x09\x3b\xfe\x6e\x5b\x2f\x64\x2a\xbb\x68\x63\x47\xe7\x0a\x62\x63\x15\x6e\xe4\x3c\x6e\xac\xcd\x12\xda\xd6\x63\x33\x93\xa0\xa3\x44\xbe\xf5\x10\xcd\x20\x6d\xac\x50\x5f\x3a\x88\xa2\x94\xb0\xf5\x10\x4d\xc2\x6c\xac\x4a\x0e\x15\x05\x47\xcc\xb2\x0f\xf3\x3a\xbd\xce\xcd\xd6\xaa\x4c\xc2\x6c\xac\xca\xfc\xc5\x31\xd7\x7a\x63\x7d\xe6\xb1\x36\x56\x4a\x18\xaa\x1a\xc9\x7d\x9e\x45\x3c\xd0\x4f\x17\xb5\xa5\x39\xbc\x6e\x3d\x86\x8b\x78\xbb\xac\x3a\x47\xbb\xbe\x10\x34\x6c\xab\xec\x32\x07\x47\xac\xc1\xac\xdf\x69\xd3\x66\xd2\x13\x4f\x62\xf0\xd5\x83\x8a\x09\x43\xab\x05\x09\xd4\xce\x60\xe0\xdf\xb5\x8d\xca\x0a\xe5\x83\x4b\xb9\xcb\xc0\x96\x38\xd7\x71\xad\x83\x3a\xf8\x5b\x0b\x6d\xc5\x14\xb4\x69\xa6\x5b\x1e\xbf\x11\x09\x3e\x62\xfb\xf5\x88\x2d\x71\xd7\x51\x8c\x7b\x72\x0d\x06\xb5\x38\x77\x7f\x9c\x2c\xfe\x31\x33\x09\x0e\x6a\xc3\x2e\x5d\x97\xb5\x11\x90\x8b\xd1\x1b\xe5\x51\xc4\xed\x39\x8c\x54\x60\xff\xb2\x11\x44\x60\x1f\x86\x28\x74\x35\x1a\xbb\x86\x39\x6b\x1b\x25\xb3\x22\x31\x1a\x8f\xf9\x26\x8a\xf9\x48\xba\x80\xdc\xa1\x01\x2d\xe3\xa7\x21\x41\x46\xce\x26\x5d\xe3\x36\x97\x99\x59\x2b\xf4\x01\x1e\x59\xc9\x6f\x6f\x28\x58\x29\x14\x2e\x14\x7c\x2c\x62\x47\xc4\x3a\xbc\xb5\x28\xbc\xc1\x8f\xc5\xcd\x5b\xfc\x78\x1d\x16\xc5\xd9\x72\x18\xa3\xfe\x50\xf4\xd0\x54\xd4\x7c\x9b\xe0\xb8\xc1\x68\xaf\xb8\x3e\x0c\x55\x85\x05\xd9\x74\x7a\x63\xfd\x8e\x63\xfc\x3f\xa4\x8b\xef\xe4\xf0\xbf\xbd\x2b\x3f\xfd\x1a\x04\xe8\x85\xd4\xc3\x6d\x21\xb9\x37\xd9\x09\x14\xc9\x74\x4c\xb9\x6c\xb9\xcb\x72\x3b\xef\x01\x2f\x98\xe5\x94\x4a\x41\xbd\x65\x6e\x6f\x8e\x45\x89\x15\xda\xff\x0f\x00\x00\xff\xff\x60\x3f\xce\xfb\x7d\x55\x00\x00")

func redirectsRedirectsCsvBytes() ([]byte, error) {
	return bindataRead(
		_redirectsRedirectsCsv,
		"redirects/redirects.csv",
	)
}

func redirectsRedirectsCsv() (*asset, error) {
	bytes, err := redirectsRedirectsCsvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "redirects/redirects.csv", size: 21885, mode: os.FileMode(420), modTime: time.Unix(1556099962, 0)}
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
	"templates/error.tmpl": templatesErrorTmpl,
	"templates/main.tmpl": templatesMainTmpl,
	"templates/partials/footer.tmpl": templatesPartialsFooterTmpl,
	"templates/partials/header.tmpl": templatesPartialsHeaderTmpl,
	"redirects/redirects.csv": redirectsRedirectsCsv,
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
	"redirects": &bintree{nil, map[string]*bintree{
		"redirects.csv": &bintree{redirectsRedirectsCsv, map[string]*bintree{}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"error.tmpl": &bintree{templatesErrorTmpl, map[string]*bintree{}},
		"main.tmpl": &bintree{templatesMainTmpl, map[string]*bintree{}},
		"partials": &bintree{nil, map[string]*bintree{
			"footer.tmpl": &bintree{templatesPartialsFooterTmpl, map[string]*bintree{}},
			"header.tmpl": &bintree{templatesPartialsHeaderTmpl, map[string]*bintree{}},
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

