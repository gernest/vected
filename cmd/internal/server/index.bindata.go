// Code generated by go-bindata.
// sources:
// cmd/server/index.html
// DO NOT EDIT!

package server

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

var _cmdServerIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3c\x6b\x73\xdb\x38\x92\xdf\xfd\x2b\x7a\x54\xbb\x5b\xd4\x58\xa1\x64\xc7\xe7\xf5\x78\xac\x54\x65\x62\x27\xeb\xab\x8c\xed\x8a\x33\x33\xb7\xe5\x75\x65\x21\xb2\x49\x21\x26\x01\x1e\x00\x4a\x56\xe5\xfc\xdf\xaf\x00\x90\xd4\x8b\x4f\xd9\x35\x33\x55\x77\xfc\x60\x8b\x24\xba\xd1\xef\x6e\xbc\x78\x36\x55\x71\xf4\x66\x6f\xef\x6c\x8a\xc4\x7f\xb3\x07\x00\x70\x26\x3d\x41\x13\x65\x6f\xf4\x35\x1c\xc2\x3b\x9e\x2c\x04\x0d\xa7\x0a\x0e\x47\x07\x27\xf0\x79\x8a\xf0\x81\xc3\xdb\x54\x4d\xb9\x90\x2e\xbc\x8d\x22\x30\xaf\x25\x08\x94\x28\x66\xe8\xbb\xab\xe0\xbf\x48\x04\x1e\x80\x9a\x52\x09\x92\xa7\xc2\x43\xf0\xb8\x8f\x40\x25\x84\x7c\x86\x82\xa1\x0f\x93\x05\x10\xf8\xe9\xf6\xfc\x95\x54\x8b\x08\x57\xa1\x23\xea\x21\x93\x08\x6a\x4a\x14\x78\x84\xc1\x04\x21\xe0\x29\xf3\x81\x32\x50\x53\x84\x8f\x97\xef\x2e\xae\x6e\x2f\x20\xa0\x11\xba\x7b\x05\xa8\xe3\xf4\x61\xfc\x06\xbe\x15\x0f\x32\x7c\x3f\x93\x04\xe6\x38\x81\x89\xe0\x73\x89\x02\xde\xde\x5c\x02\x61\x3e\x5c\x71\x1f\xdd\xaf\xd2\xdc\x2b\x0e\x04\x24\x65\x61\xa4\x49\x8d\x63\xce\xcc\x63\x27\x11\x18\xa0\x10\x94\x85\x06\x85\x54\x84\xf9\x44\xf8\x12\x34\x1b\xab\x18\xfa\xee\x5a\xb7\x1e\x67\x52\x01\x95\xba\xc5\x7f\xde\xc2\x18\xd4\x22\x41\x1e\x40\x22\xb8\x87\x52\xc2\x77\xe3\x31\xf4\x52\xe6\x63\x40\x19\xfa\xbd\x1f\xd7\x80\x69\x00\x4e\x0e\xda\xdf\xe0\x47\x5f\x61\xc4\x27\x24\x72\x05\xfe\x77\x4a\x05\xc2\x18\xb2\x5f\x3f\x56\xb5\x0c\xe4\xb2\x91\xd3\x0b\x64\xaf\xff\xe3\xde\x56\x5b\x4b\x32\xe3\x3e\xbe\x13\x8b\x44\xf1\x55\x10\xcf\x3c\xd1\x60\x15\x3d\x78\x39\xc8\x36\xb5\xa6\x15\xaa\x4f\x84\xf9\x3c\xfe\x95\x44\x29\x4a\x67\x52\xc6\x56\x7e\x2d\x49\x70\x85\x01\x7a\x4f\xa3\xe8\x76\xc1\x3c\x67\x52\x42\x80\xbe\x9e\x06\x5b\x8f\x9f\x4a\x38\xcc\x68\x4d\x50\x04\x5c\xc4\x84\x79\x58\x49\x30\xe3\x73\xa7\x8e\x46\x2b\xac\x3b\x89\xde\x00\x98\x44\xef\x1e\xc6\xb9\x6e\xdd\xa9\x50\x34\x46\xa7\x82\x56\x7d\x09\x54\xa9\x60\x20\xd1\x83\xef\xe1\x60\x34\x1a\xc1\xbe\xc1\x02\x43\x73\x37\x1a\x8d\x9e\xc5\xa7\x25\x2e\x55\x34\x5a\xd5\xa1\xbe\xaf\xd1\xe0\x67\x7c\x54\x17\x4c\xbb\xa9\x80\xb1\x01\x5e\x7d\x54\x0b\x76\x8e\x5b\x60\xd9\xa3\x75\xb0\x27\xc0\x48\x62\x89\x54\xb5\xc1\x67\x0e\x32\xa7\xcc\xe7\xf3\x4d\xff\xa8\x52\x85\x6d\xed\x5a\x5a\x60\x9c\xdd\x6f\x13\x9b\xf5\xbc\xd2\x8f\xc4\x28\x68\xdb\x8b\x6e\xbb\xec\x43\xdf\x55\xf6\x50\x8e\x40\x4d\x05\x9f\x03\xc3\x39\x5c\x08\xc1\x85\xd3\xf3\x08\x63\x5c\x01\x3e\x26\x5c\x28\x1d\x5a\x1d\x86\x54\x4d\x51\xe4\xfc\x33\x2e\x2c\x89\x54\x42\x46\x5e\xbf\x4c\x79\x4f\xdb\xda\x8f\x50\x01\x4f\x55\x92\xaa\x9f\xd2\x00\xc6\xd0\xeb\xd5\xc7\x85\x72\x92\x8d\x0d\x11\xa6\xe4\x29\x7c\x83\xeb\x2f\xbf\x7d\xba\xbe\xfa\xf8\xcf\x53\x78\x75\x30\x80\xeb\x2f\x9f\xce\x7f\xfb\x94\xff\x7e\xf7\xe9\xe2\xed\xe7\xfc\xe6\xf3\xa7\x5f\xae\xde\xe5\x37\x6f\x6f\x6e\x2e\xae\xce\xf3\xbb\x8b\xff\x7a\xf7\x51\xff\x86\xa7\x81\x0e\xca\x29\x4b\x25\xfa\xe5\x4a\x15\x54\xa1\xf1\xf7\xc0\x1f\xc0\x24\x0d\xea\x1c\x71\xc9\xe9\xfe\x18\x7c\x6b\x75\xae\xfd\xef\x68\xd0\x6a\x27\xcc\xe2\x9d\x56\x69\x81\xc4\x8d\x88\x54\x97\xcc\xc7\xc7\xeb\xc0\xe9\xfd\x8b\x95\x89\x3c\xbf\xb4\x31\xb1\x08\xbe\x1b\xc3\xab\x83\x3a\x0a\xf3\xbe\x78\x84\x6e\xc4\x43\x67\xd9\x99\x4c\x27\x52\x09\x67\x34\x00\x16\xf5\x6b\x7a\x5a\xe7\x73\x95\xdc\x0c\x03\x8b\x60\x1f\x0e\x6a\x30\x3c\x35\x85\xa2\x89\xe6\x1d\x59\xa8\xa6\xad\x23\x8f\xa1\x2a\x41\x66\x34\x95\x10\x35\x1d\x40\x10\x91\x50\x0e\x20\xe6\x3e\x36\x07\x4f\x14\x3a\x62\xac\x38\x85\xf6\x08\x1a\x27\x11\xc6\xc8\x94\x76\xc7\x6a\x7e\x50\x08\xd7\x54\x14\x63\xe8\x5d\x5c\x5d\xdf\xfe\xf3\xb6\xc4\xc8\xf3\xcb\x7a\x1f\x8a\x92\x20\x56\xc1\xda\xd3\x46\xdc\xda\x2b\xc9\xed\x58\xc4\x4a\xcd\xc2\x4a\xa8\xd4\x91\x36\x78\x75\xb2\x49\xbe\x85\xf2\x71\x13\x2a\x8b\x94\x2b\x50\x6b\x60\x99\xa7\x7e\xd0\xb9\xd5\x8b\x88\x94\x25\x72\x35\xa8\x45\xea\x29\x2e\x2a\x93\x96\xae\xc5\x5c\x22\xc2\x19\x8c\xe1\xae\xf7\x55\xf6\xee\xcb\xa5\x61\xda\x21\xd3\xcd\xbe\x3d\xd5\x35\x79\xa4\x0a\xc6\xe0\x78\x46\xd7\x5b\x65\xd7\xea\xa5\x1d\xc5\x28\x4b\x87\xdb\x51\x5b\x57\x99\x13\xc1\x9c\x9e\xe9\x46\x03\x9f\xf6\x06\xe6\x7f\x67\x23\xaf\x63\xe2\x8b\x47\xa2\x68\x42\xbc\x87\xcf\x34\x46\x9e\x2a\x99\xe9\xe5\x67\x92\x54\xa5\x6f\x0b\xc7\xf0\x51\xbd\x5b\x87\xbd\x3c\x87\x31\x1c\x94\x64\x63\x28\x94\x1f\x63\xac\x45\xd6\x20\xae\xe1\xd0\xd4\xdb\x93\x34\x08\x50\x40\x4c\x16\xe0\x4d\x09\x0b\x11\xe6\x53\x64\x26\x9d\xa3\x54\xba\x20\x8d\xb9\x40\x8d\x93\x8b\x85\xdb\xe4\xdf\x9a\xab\x73\xa2\xc8\xaf\x14\xe7\x8e\xe5\x81\x32\xa9\x5c\x9b\x81\xa4\x1b\x63\xec\xda\x1e\xab\x4a\xac\x3a\xc6\x24\xaa\x4b\xa6\x8e\x8f\x34\x77\xc4\xf7\xc5\x00\x66\x0d\x4c\xc6\x18\x3b\x7d\x57\xa2\xfa\x85\x32\xf5\xfa\xd0\x40\xc1\x3e\x8c\x06\x30\x1b\x80\x12\x69\x9d\x9e\x2b\x60\x8f\x06\xf0\x33\x51\x53\x37\x88\x38\x17\xce\x0c\x86\x70\x74\xf8\xc3\xd1\x0f\xc7\x7f\x3f\xfc\xe1\xb8\x5f\x8b\xb4\x96\xb7\x70\x83\xb7\x06\xc6\x2c\x50\xc4\xe7\x30\xce\x08\x0d\x4b\x98\x6c\xe0\xd0\x22\x99\xd2\x70\xba\x8a\xe5\x72\x83\xdb\x06\x24\x99\xe6\x35\x29\xfb\x16\xd7\xf7\x2b\x12\xd9\x41\x12\x11\x27\xbe\xa9\xdf\x3b\x89\x22\x58\x65\xe1\x7d\xc4\x89\x3a\x3e\xca\x8c\xa4\x81\x01\x1d\x37\xbe\xa3\xf2\x8a\x5c\x39\x41\xbf\x29\x6c\x64\xdc\x96\x54\x66\x0d\xbc\x2d\x29\xa5\x7e\x85\xce\xda\x8a\xda\xfa\xd5\xcc\x0c\x71\xee\xa8\x5f\x11\x61\xeb\x3d\x49\x71\x81\x6b\x42\x6e\xf6\xa5\xac\x90\x21\xec\x1f\x48\x34\x0b\xa3\xc7\xbf\xbf\x7f\x7f\x62\x47\x11\xb5\xc2\xcd\x4a\xe1\x19\x8c\x75\x1d\xcc\xd2\x78\x82\xa2\xb2\x08\x5e\x05\xb4\x4a\x99\x35\x2a\x05\x6a\xbd\x35\xa3\xb8\x51\xba\x75\x98\x06\x2d\xbc\x29\xbf\xac\x96\xea\xdb\x55\x17\x4a\x6b\x14\xac\xdb\x71\x73\xc8\x6a\xd3\x7d\x8d\x7d\xca\x39\x55\xde\x14\x9c\x59\x63\xf2\x24\x12\xa1\x18\xd0\x9c\xfe\xe1\xca\x39\x78\x51\xe5\x18\xee\x58\x1a\x45\x7f\x3c\x63\x87\x2f\xcf\x98\xc6\xf7\xc7\x33\xf6\xfa\xe5\x19\x0b\x48\x24\xff\x04\x9c\x35\x67\xcc\xb6\x9c\xd5\x78\xaa\x1e\x7c\x0b\xd4\x39\xcf\x26\x03\x81\x81\xd4\xd9\xc4\x99\x35\x24\x3a\x03\x34\x1e\x2f\x9d\xb7\x39\xe1\x2d\x7b\xb1\x29\xa7\x76\x00\x97\x5f\x6b\x00\x49\x2a\xa7\xb5\x94\x2d\x21\x0c\x23\x52\x33\x32\xd0\x5d\xef\x34\xd6\xd4\xc2\xd1\x39\xe7\x7d\x44\x42\x9d\xa7\xaa\x71\xe4\x01\x2f\xcf\x50\xad\xe2\x5e\x4f\x2a\x41\x59\xd8\x6b\x36\xb5\x15\x22\x0e\x9a\xcd\x61\x22\x90\x3c\xb4\xb0\xf3\x9e\x5c\xc4\x13\x1e\x75\xeb\xff\xf0\x05\xfb\x0f\x52\xe6\x29\xca\x59\x37\x0a\x5e\x3f\x9b\x82\x6a\x8d\x37\xfa\x34\xfc\x4f\x41\xca\x4e\x75\xbf\x31\xc6\xdd\x8b\x7b\x5d\xd2\xde\x46\xd4\xeb\x56\xd2\x12\x21\xc8\x02\xc6\xc5\xd0\xa0\x28\xec\x1b\x4b\xfa\x08\x59\x09\xdc\x49\x73\x69\xa9\xc7\x6f\x9a\xed\x93\xb7\xba\xef\xfa\x11\xdc\xc0\x12\x38\xd0\xbd\x3d\x4b\x28\xd7\x81\x9d\xae\xff\xd3\x0a\x27\xeb\x31\x1b\xb5\x5b\xd1\x54\x33\xad\xaf\x80\x0b\x70\x74\x1c\xa2\x26\x00\x01\x85\x33\xdd\xef\x8f\x40\xf7\xf7\x9b\x82\x0c\xb9\xa3\xf7\x30\x5e\x0e\x83\x1c\xcb\xea\x3e\x50\xf8\xbe\x96\xcc\xc6\xd9\x37\xb2\xab\x96\x4c\xc0\xeb\xa4\x1e\x69\xc4\xfa\xbb\xda\xee\xc6\xac\x6c\xfb\xa9\x88\x81\xa5\xd6\x1a\xf2\x2e\x96\xac\x68\x8c\xd7\x82\x86\x54\x53\x7e\x4e\x14\xba\x76\x5d\xe7\x15\xac\xac\x01\xd9\x67\x35\xb3\x3d\x34\xd6\x84\x5d\x4f\xbe\xa2\xa7\x2a\xe7\xcb\xf5\x15\xf2\xd3\x06\x13\x1a\x0e\x41\x07\x69\x98\x13\x19\x5f\x3c\x52\x65\xe7\xc5\x4c\x38\xeb\xd7\x02\xf6\x44\xca\x34\x37\x6e\x0e\xd9\x3b\x05\x47\x26\x0d\x2a\x5f\x97\x46\x36\x61\xba\x31\x9f\x20\x13\xad\xc5\xd6\xb5\x51\x31\xf9\x87\x7a\xa8\xa9\x81\x9a\x61\x7c\x8c\x50\x21\x2c\x75\xdd\x11\xc4\x56\x2b\x1d\x81\x74\xc1\xd2\x81\x1d\xa7\x61\x72\x11\xec\x4c\x71\x6b\xed\xfe\x26\xa8\x42\x27\xf0\x21\xa5\x4c\x25\x4a\x0c\x20\x81\x94\x49\x12\xa0\x7b\xc3\x29\x53\xda\xba\x59\x67\xcd\x1b\xac\x3b\xa8\x3e\xf0\x57\x5d\xd7\xaa\xbc\x85\xb2\x2d\x70\xb2\x05\x7b\x70\xdc\x1a\x98\x55\x58\xdc\x61\xfb\x72\x3c\x90\xee\xfa\x92\x50\xb7\x64\x98\x0c\xa0\x32\x7a\xe4\x57\x5b\xcd\x32\xc2\xb8\x5d\xdc\xd5\xaa\x3b\x3e\x6a\xa7\xb9\x1c\xaa\x93\xe2\xe4\xba\xb6\x06\xe0\xac\x44\xb3\xfd\xed\x00\xd6\xcf\x96\x92\x47\xa3\xba\x50\xde\x85\xd9\x39\x89\xa2\x8c\x59\x47\xa2\x67\x39\xb6\xcb\xdd\x1d\x0d\xd7\xe2\xd9\xc1\x6e\x63\xdd\xd7\x18\xf2\x9c\x81\xc6\x8a\x3e\x37\xac\xae\xe7\xd7\x96\x04\xe3\xe5\x12\x7b\x97\x51\xe4\x8a\xd5\x1e\x1c\x0f\xc0\x31\x68\xfe\x6a\xd1\x2c\x85\xde\xca\x9c\xdb\x8a\x5e\x7a\x53\xf4\xd3\x08\xf3\x35\x06\xc7\xc7\x88\x2c\xac\x06\xfa\x56\xf8\xed\x64\xbf\x89\x68\x07\x1d\x98\x69\xd1\x9a\x65\x8f\x96\x41\xb6\x1c\x78\x7f\xbf\x2d\xf8\xe6\x52\x8d\x19\x8c\x52\x7f\xa0\xd5\x9c\x3d\x73\x1a\x51\xe9\x2b\x5b\x7f\x29\xd2\x84\xe4\xd1\xac\x90\xcf\x8d\xe0\x31\x95\xda\xbc\xaa\x96\x3c\x37\xaf\xcd\xa8\xaa\xad\xc4\xac\x70\x2f\xe9\x82\x29\x91\x30\x41\x64\x20\xf5\x1f\xc5\x21\xa0\x02\x21\x4d\xf4\xcf\x03\x88\x69\x14\x51\x89\x1e\x67\x3e\x20\x11\xd1\xa2\xb1\xdf\xa6\x70\x06\x35\xe6\x4b\xdb\xcd\xa3\xb4\xb5\x54\x2f\x42\x22\x6e\x33\x2b\xf3\x0b\x7b\xa5\x7e\xa7\x10\x51\x8e\x65\x57\x63\x7d\x5e\x85\x63\x68\xc9\x2d\xaa\xc2\xf6\x42\x63\x7b\x6d\xd4\x50\x81\xc0\x96\x2a\x1a\xc7\xcb\xa8\xa1\xd8\x60\xa5\x2b\x6b\x47\xc0\xdd\xfd\x64\xa1\xb0\xa5\xf4\xd7\x80\xbb\x09\xdd\x6e\xd3\xda\xdc\xde\x55\x8c\x20\x73\xa7\x78\xa9\xb0\x68\xc6\x3b\xbf\x92\xc8\x31\x25\x61\x76\xdf\x07\x81\x41\x3d\xa7\x72\x21\xb5\x0a\x86\x5f\xa5\x5b\xe0\xe8\x96\x8d\x8b\x75\x99\x65\xf9\xb2\x1c\x81\xbd\x30\x9b\x86\xb9\x0f\xa8\x9c\x99\x9d\xdd\x48\x76\x62\x34\xc7\xf2\x2c\x3e\x5f\x1f\x0e\xe0\x13\x06\x11\x7a\xca\x18\xfd\x72\xf0\x9b\xb1\xbc\x2d\x86\x83\xe3\xfe\x8b\x0a\xe2\x76\x5b\x10\x03\x78\x34\x93\x90\xdd\x64\x71\xdb\x51\x16\x39\xdf\xb2\x35\xdf\x03\xd8\x68\xf7\xfa\xf0\x45\x45\x61\x76\x26\xe5\xc2\xa0\x3a\xc2\xee\x60\x12\x06\xc9\xb3\x8d\xbf\xc1\x28\xb6\xc6\x0b\x2f\x20\x87\x5c\x89\xdb\x52\xd8\xd9\x1e\xba\xcb\xa2\xc1\x28\xb6\xf8\xde\x32\x89\xc3\xa3\x17\x35\x09\x9d\x2b\x73\x59\xc4\x85\x7b\x10\x11\x4a\xb8\xbb\xd7\x32\x31\x4b\x0b\x03\x98\x70\x1e\x75\x94\x8f\x46\xdd\x49\x36\x4a\x2c\x5a\xb4\x82\x22\x59\xcf\xd6\x66\xd3\x5a\x8f\x4a\x97\x08\x62\x18\xaf\x19\xe2\xac\x3c\x1a\x75\x41\x69\x24\x37\xde\x9e\x03\x2d\xfc\xb9\x1d\xb2\x4d\x97\xf9\x8f\xe3\xa5\xcb\x90\x24\x89\x16\x4e\x6c\xd6\x8d\x75\x77\x6d\x09\x5c\x9b\xff\x3e\xb1\x78\xf5\x80\xac\x6e\xe3\x5f\x7e\x3d\x81\x47\xcc\x8a\x0a\x0a\xd1\x66\xd5\xbe\x8a\x07\x0d\xfe\x3c\x6a\xdb\x0c\xbd\xea\x57\xe1\x3b\x86\xcc\x19\x7f\xc0\xdc\x43\x5e\xc4\x2f\x2c\xca\x3f\xb7\x67\xd4\x9b\x71\xab\xf9\x1b\x28\x31\x81\xa3\xd1\xa6\x19\xcf\x06\xcb\x25\xcb\xe7\x9b\xf3\xd1\xc9\xef\x67\xce\x9a\x97\xe7\x99\xb3\xa6\xf6\x77\x36\xe7\x2b\x9c\xbf\xa8\x2d\x5f\xe1\xfc\xff\xbc\x21\x17\xdb\x67\x9d\x17\x08\xc8\xff\x6f\xc1\xdb\xd7\x9a\x05\x7f\x34\x9b\x14\xac\x11\x9b\xa9\xac\x6e\x06\x6b\xc1\x9f\x31\x93\x7a\x70\x3c\x80\x84\x08\x89\x97\xac\xa4\x82\xcb\xf6\x50\xbc\x54\xb5\x7a\x23\x30\x21\x02\xb3\x82\x24\xe3\xd9\x3a\xab\x2e\xde\xbb\xb1\xbe\x86\x6c\x87\xb9\x11\xa9\x04\x8c\xf3\x7d\xeb\xae\xfd\xef\x64\xa4\x6d\x09\xa2\x8d\x0f\x6c\xda\xa3\x16\xad\x54\x6d\xec\x71\x5d\x25\x7a\x3c\x21\x95\xc8\x65\xff\x32\xa2\xff\xb8\x2c\x04\xb3\x80\x39\x69\x37\x2b\xb2\x65\x71\x05\xa2\x9d\x65\xbe\x43\x04\xdc\x98\x42\x39\x38\x36\xae\xeb\x34\x8b\xb7\x63\x75\x24\x15\x61\x3a\xaa\xe6\x42\x52\xd6\x46\x75\x2a\xe9\x5a\x15\xe5\xa8\x3a\x89\xa9\x2c\x2a\xe5\x93\x2b\x6b\x22\x03\x9a\x75\xc0\x83\xcd\x97\xcd\x35\x7e\x93\x50\x7a\x3e\x4e\x52\xa3\x5f\xc3\x4b\x07\x15\xe7\xa7\x7a\x2c\x5c\x23\x15\x95\xaf\x4a\xdf\x94\x9c\x5d\x28\x59\xef\x26\x72\xc1\x3c\x10\x29\x73\x72\x11\xd5\x1e\x02\x31\x2b\x65\x30\x2e\xe4\x59\x77\xce\x61\x96\xef\xff\xb8\x33\xc7\x12\xae\xcf\xaf\x4f\x21\x24\x62\x42\x42\x04\x8f\x47\x3a\x7d\x52\xce\x2a\xb9\xba\x22\x57\xd5\x2c\x2f\xcb\xc6\xca\x26\x2c\x8d\xa2\xea\xb7\x4a\xa4\x58\xfd\xd6\xec\x3a\xac\x7e\x6d\xcf\xd5\xd4\x20\x2f\x5d\x54\xac\x6f\x5f\xfe\xb6\xee\xa0\x8d\x59\xa5\x6e\x7b\xe8\x24\x9f\x44\xbe\x9d\xa6\xca\xe7\x73\x06\x63\xcb\x64\xc3\x29\x1d\xb3\x50\x9f\x35\x2c\x6d\xb9\x7a\x38\xa5\xc3\x39\x91\x72\x64\xc3\x21\xdc\x10\x29\xcd\x21\x6a\xc2\x7c\x88\x28\x43\x5d\x54\xa5\x31\x32\x25\xcd\xb9\x6b\x64\x33\x2a\x38\xd3\x0f\x60\x46\x04\x25\x93\x08\x25\x28\x0e\xbf\xe1\xe4\xad\x94\x18\x4f\xa2\x05\x4c\x16\xe6\x1c\x20\x65\x21\xa8\x29\xc6\xfa\xb5\x9a\xa2\x41\x47\x44\xed\xa9\x17\x73\xfe\x31\x08\x24\x6a\x23\x3f\x1a\xfd\x70\x5c\xcb\xb6\x54\xe2\xc6\xc4\x68\x13\x5a\xeb\xfd\x5e\x63\x4e\x4c\x63\x8b\xbf\xda\xdb\x37\x56\xa8\x57\x97\xa3\x2d\xe8\x6a\xc2\x33\x67\xf7\x4c\x74\xdf\xc8\xcd\x3a\x79\xec\x43\xef\x5f\xa3\x5e\x5d\x7c\xcb\x78\xdd\x1f\xaf\xa3\x74\x4e\xe0\x95\x61\x2a\x7f\xf4\xd7\xfa\x09\xea\x6c\xb3\x4e\xa2\xaa\x8e\xc9\xd5\x8a\x91\x88\xd0\xcb\x97\x09\x89\x08\x67\xc5\x5e\xd4\x7a\x98\xd9\x8d\x12\x26\xbc\xd4\x39\x89\xc1\x17\x70\x71\x41\xbc\xa9\xe3\x10\x11\x36\xa8\x29\xc7\x6b\xb7\xb7\x5a\x05\x1b\xb0\xaa\x2d\x44\x65\xa7\xe1\x97\x64\x3e\xe0\x42\x93\x68\xf7\xff\xb8\xfa\xce\xc9\x0f\xc9\xf5\x5d\xc9\x85\xaa\xf2\xda\x75\x3a\x34\x60\x7d\x85\x63\x5a\x14\x7c\x3e\xe0\x62\x17\x3e\xff\xfd\x97\x6f\x0f\xb8\x78\x1a\xff\xe5\x5b\x4e\xe4\xdd\x03\x2e\xee\x9f\xfe\xbd\x23\xf7\xd9\x91\xc1\x3a\x83\x2f\xe8\x28\x68\x4f\x1a\x5d\x49\x7b\xc4\x72\x27\x67\xee\x13\x66\x9f\x4c\xf3\x2e\xd0\x2d\x48\xbb\xa5\xb4\x79\x09\x7e\xe9\x2a\x27\x1d\x85\x31\x9f\xd2\x08\xc1\x31\xf8\x1b\xf7\xd8\x79\xeb\x8b\xc9\x59\x58\xcd\x97\x96\x9d\x6c\xc9\xb9\x45\x91\x51\xb7\x44\xdd\xe2\x24\x61\x7e\x99\xb3\x3e\xcb\x74\xd0\x76\xb8\xb9\x75\x6a\x7d\x42\xfc\x82\xb7\x53\xf8\xc0\x21\x11\x3c\x14\x24\x36\xeb\xdb\x24\x12\x48\xfc\x05\xd8\x3e\xea\x0e\xef\xe6\x57\xfd\xa0\x12\xec\x30\x21\x5f\x88\xcd\x84\x60\x86\xab\x3a\xc9\xc4\xe4\x01\x41\xa6\x02\x81\x2a\xa0\xd2\x16\x41\x53\xc1\x19\x4f\x65\x7d\x11\x56\xb3\x4f\xb3\x86\xe8\x92\x5c\xa8\x2b\x2e\x1d\xf7\xcc\xa4\x41\xd3\xfe\xfe\x0e\xf2\xdf\x79\xb7\x35\x99\x13\xba\x65\x7d\x55\xdb\x26\xb7\x9f\x6c\x3d\x92\x8a\x28\xea\xc1\x17\x2d\xeb\xdc\xfa\xfe\x81\x51\x82\xc2\x6c\xbc\x48\x90\xf9\x94\x85\xf9\x1b\x39\x80\x90\x57\xf1\x96\x1f\xcf\xcb\x76\xa8\x43\xed\xc7\x35\x36\x11\xdb\x00\xf7\x0d\xa8\x7f\x6a\xb6\x2f\x10\x11\xca\xd3\x95\xa2\xa2\x4e\x6f\x21\xaf\xd9\xe3\xf1\x8c\xa2\x7b\x55\x36\x17\x33\x64\x6a\x43\x40\x89\x40\xfd\xf4\x1c\x03\x92\x46\x26\xdd\xf3\xe4\x46\xf0\x84\x84\x44\x0b\xc0\x3e\xb8\x8c\x63\xf4\x29\x51\xb8\xf6\x26\x60\xad\xa5\x68\xfa\xa8\x13\xa5\x36\xbd\x75\x52\x9a\xac\xcf\xb4\x75\xd7\x61\xea\x76\x5b\x55\x9b\xa3\xee\x7b\x83\xed\x76\x9d\x6f\x00\x3d\xaf\xf7\x32\x19\xb7\x27\xa3\x0c\x7a\x37\x7a\x02\x96\x69\xab\x83\xd5\xad\xdf\xed\x6d\xb2\x57\xf7\x8d\x20\xab\x78\xfb\x41\x9a\x95\x82\x0c\xce\xe0\x75\x15\xf7\x79\x73\xa9\x7c\x14\xc2\xee\xb9\x74\x7a\xa9\x24\x21\x9e\x42\xc8\xbf\x7c\x95\x5f\xe6\x44\xc6\x5f\xf0\x11\x3d\xb8\xd3\x3f\x61\x42\x19\x11\x8b\x7b\xb8\x2b\x9c\xf1\xbe\xfa\xa3\x19\x39\x7e\xb3\xdd\xb6\x6c\x8e\xb4\xc4\xcd\xb2\x33\xd8\x3c\xcb\xa1\x1f\x78\x99\xf0\x43\x9e\x7f\xd9\x60\x8d\x63\x69\xa6\x4f\xca\x56\xec\x42\x9e\x7d\xe2\xa0\x20\x89\xcd\xca\x5b\xd9\xaf\x1c\xac\x52\xbe\xdd\x6e\x65\xd8\xe2\xda\x41\xb5\xd2\x16\xe3\x04\xd2\xd5\x29\xf1\x3d\x8d\xec\xd6\xd5\x55\xe2\xee\x0e\xef\xfb\x3a\x5e\xae\x6d\x2e\xef\xbb\x6a\x8a\xcc\x14\x08\xc6\x51\x2b\x93\x7b\x8e\x89\x67\x5f\x48\xe8\x0d\x56\x3f\xc4\xa0\xd3\x63\xfe\x65\x2a\xfd\x56\x6a\x63\x60\xbc\x48\x0c\x3a\x5d\x66\x21\xb6\xd6\x7b\xec\x7e\xf1\xf1\x18\x46\xf0\xb7\xbf\xc1\x77\x99\x38\x9a\xf3\xd7\x70\x08\x3e\x12\x3f\xe2\xde\xc3\xc0\xe6\x69\x53\x2a\x50\x66\xbe\xf9\xc1\x85\x19\x04\x4a\xa5\x29\x51\x82\x78\x58\x9f\xaf\x75\x00\x2f\x19\xf9\x36\xef\x3c\xd7\x80\x5b\xf9\xba\xfb\xe7\x23\x2a\x20\xb2\x50\x1c\x72\x83\xd6\x6a\xcc\x2d\xe6\x5f\x4a\x6c\xbb\xef\x9a\x09\x7f\xc7\x4e\xf8\x57\xaa\xb6\xe6\x63\x25\x9b\xa4\x2c\x09\x7e\xea\xe7\x8c\x9d\x0d\x57\xbf\xe6\x96\x7d\xda\xcd\x1c\xb1\x1a\xf7\x14\x3e\xaa\xe1\x57\x32\x23\xf6\x69\x6f\xf9\xc5\xb7\x22\xa1\x04\xa8\xbc\xe9\x5b\xe6\x5f\xae\x98\x71\x2a\xa2\x01\xac\x99\xe9\x06\xe9\x79\x56\xd2\xb0\xba\x75\x66\xc6\x02\x65\xc2\x99\x2e\x54\xdf\x6c\xb1\x92\xbf\x73\xcd\x11\x9a\x9f\xcc\xd0\xd8\x59\x9f\x97\xcd\xb0\x4c\x16\x0a\x65\x19\x8a\x2a\xaf\x33\x00\x1b\x04\x97\x21\xb6\x2a\x2b\x45\x9d\xbd\x2a\xd4\xb9\x0e\xbe\x54\xc2\x52\x01\x33\x22\xaa\x82\x94\x7e\x15\x73\x33\x07\x53\x22\xdc\x5e\x4c\x28\x33\x5b\xfb\x7b\xdb\xe1\x60\x89\x24\xfb\xce\x15\x67\x11\x37\x47\xfc\xab\xeb\xa8\x98\xfb\x96\xbf\x65\x93\xba\x59\xc1\xcc\x7e\x2b\x0c\x77\xd5\xe2\x9e\x36\x0d\xec\x6c\x68\xbf\x21\x78\xe6\xd3\x19\x50\x7f\x6c\x58\xe9\xbd\x39\x1b\xfa\x74\xf6\x66\x6f\xef\x6c\xc2\xfd\x85\xfe\x3f\x2c\x7e\x98\x6f\x0f\xfe\x6f\x00\x00\x00\xff\xff\x5b\x6e\x14\x5b\x82\x50\x00\x00")

func cmdServerIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_cmdServerIndexHtml,
		"cmd/server/index.html",
	)
}

func cmdServerIndexHtml() (*asset, error) {
	bytes, err := cmdServerIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/server/index.html", size: 20610, mode: os.FileMode(420), modTime: time.Unix(1534060697, 0)}
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
	"cmd/server/index.html": cmdServerIndexHtml,
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
	"cmd": &bintree{nil, map[string]*bintree{
		"server": &bintree{nil, map[string]*bintree{
			"index.html": &bintree{cmdServerIndexHtml, map[string]*bintree{}},
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
