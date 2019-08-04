// Code generated by vfsgen; DO NOT EDIT.

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// assets statically implements the virtual filesystem provided to vfsgen.
var assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Time{},
		},
		"/source": &vfsgen۰DirInfo{
			name:    "source",
			modTime: time.Time{},
		},
		"/source/flag-types.json": &vfsgen۰CompressedFileInfo{
			name:             "flag-types.json",
			modTime:          time.Time{},
			uncompressedSize: 2559,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x94\xc1\x6e\x9c\x30\x10\x86\xef\x3c\xc5\xc8\xbd\x40\xb5\x82\x1c\x56\x7b\xd8\x63\x55\xb5\xda\x5b\xa5\x34\xbd\x24\x51\xe4\x80\x21\x56\x1d\x1b\xd9\x43\xd4\x28\xca\xbb\x57\xf6\x2e\xbb\x60\x0c\x62\x5b\x2d\xb9\xfe\x23\x98\xef\x9b\xb1\xe6\x36\x02\x78\x8b\x00\x00\x88\xa4\xcf\x8c\x6c\x81\x7c\x51\x4a\x90\xd5\x3e\xc3\xd7\xda\x65\x8f\x9d\xec\x85\x8a\xc6\x86\x25\x15\x86\x1d\xb2\x82\x19\x24\x5b\x40\xdd\xb4\x49\xae\x24\xb2\x3f\xf8\x50\xb0\x92\x36\xc2\x16\x89\xfb\xa0\xfd\x4b\x4d\xb5\x61\xda\xc6\x06\x75\xae\xe4\x4b\xfa\xc3\x26\xb6\x79\x5c\xa6\xbf\x6c\x8f\xf4\x1a\x35\x97\x55\x9c\x24\x24\x02\x78\x5f\x85\x51\x7f\xfe\x3f\x6b\xa1\x72\xa4\x5c\xd8\x8f\x01\x9f\x28\x02\x37\xae\x0c\x8f\xaf\xd0\xf2\x2f\xa9\xf5\xb5\xd1\x14\xb9\x92\xbe\x19\xf2\x67\x96\xfa\xc5\x56\xb1\xeb\x33\x69\x18\x1b\xc6\xe0\x09\xb1\x36\xdb\x2c\xab\x94\xa0\xb2\x4a\x95\xae\xb2\xfa\x77\x95\xd9\x0e\xd9\x27\x87\xdc\xf6\x49\x26\xd4\xaf\x86\xda\x8e\xb1\xf7\x83\xf9\xde\xdf\x84\xa2\xb8\x59\xfb\xda\x65\x3f\x9e\x25\x3c\x0f\xb6\xb7\x23\xd7\x7d\x00\xbb\x82\xcd\x7a\x14\xf8\x3b\x93\x4c\xf3\xdc\x07\xf6\xe2\x71\xe0\xee\xb3\x0c\x10\x4b\x2e\xfc\xe1\xb7\x2d\xb8\x44\xa6\x4b\x9a\xb3\xb7\xf7\x31\xb8\x9d\x0c\xcc\x92\xcb\x05\x26\xb9\x93\xa1\x39\x5e\x4d\x8e\x72\x27\x31\xc0\xfa\xa1\xa4\xbd\x4f\x1f\x72\xea\xfa\x59\xaa\xd8\x45\xc5\x94\xcb\xb5\xe0\x39\xf3\x85\x3e\xfb\x85\x4b\xbc\x8c\xdb\xfb\xce\xdc\x4e\xda\x47\xcd\xf8\x48\x91\x24\xfb\xc8\x2a\x33\xad\x95\x8e\x25\x17\x53\x52\x9b\xf5\xa8\x96\x57\xba\x98\xd8\xe9\xf1\x8e\xa9\x1d\x48\xce\x93\xdb\xef\xde\x17\x33\xbd\xf4\x5f\xdf\xe0\x1d\xb9\x23\x43\xe8\xe1\xab\x9b\x4b\x19\xde\x41\xa0\x76\x99\x25\xf4\x87\x12\xdc\x42\x87\xe5\xbc\x35\xdc\xf0\xd0\xcd\x6a\x16\x39\x5a\xb6\xf7\xd9\x57\xeb\x86\x0f\xcf\x56\xb3\xc0\xdd\x9a\x86\x0d\x1f\xae\xc6\xbf\x5c\xd1\x7d\xf4\x37\x00\x00\xff\xff\x66\x52\x85\x44\xff\x09\x00\x00"),
		},
		"/templates": &vfsgen۰DirInfo{
			name:    "templates",
			modTime: time.Time{},
		},
		"/templates/altsrc_flags_generated.gotpl": &vfsgen۰CompressedFileInfo{
			name:             "altsrc_flags_generated.gotpl",
			modTime:          time.Time{},
			uncompressedSize: 1063,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x92\x41\x8b\xdb\x30\x10\x85\xcf\xab\x5f\xf1\x08\x4b\x89\x97\x60\xdf\x53\xf6\x50\xba\x5b\xe8\x25\x5b\xd8\x85\x9e\x15\x7b\x64\x8b\x2a\xb6\x91\xc6\x09\x41\xf8\xbf\x97\x91\x97\x34\xa5\x4a\xe9\xa5\x37\x59\xa3\xf7\xde\x7c\x33\xae\x2a\x7c\x1e\x1a\x42\x4b\x3d\x79\xcd\xd4\x60\x7f\x86\x69\x3f\xe2\xe9\x05\xbb\x97\x37\x3c\x3f\x7d\x7d\x2b\x95\x1a\x75\xfd\x43\xb7\x84\x18\x51\x7e\x5b\xce\x3b\x7d\x20\xcc\xb3\x52\xf6\x30\x0e\x9e\xb1\x56\x77\x2b\xe3\x74\xbb\x52\x00\xb0\x6a\x2d\x77\xd3\xbe\xac\x87\x43\x35\x79\xa3\x8f\x54\xd5\xce\xae\x54\xa1\x54\x8c\xf0\xba\x6f\x09\xf7\x76\x83\x7b\x91\x60\xfb\x88\xf2\x8b\xd3\x6d\x10\xc3\xaa\x92\x98\x54\x28\xdf\x43\xa4\x06\x1b\xc0\x1d\x21\x09\xf8\x3c\x12\xb8\xd3\x8c\x93\xd7\x63\x40\xed\x6c\x99\x15\xf1\x00\xed\xdc\x70\x12\x57\x33\x78\x0c\xdc\x91\xc7\x51\xbb\x89\x82\x14\xf7\x84\x30\x52\x6d\x8d\xa5\x46\x25\xd7\xac\x4d\x60\x3f\xd5\x8c\x98\xd8\x6e\x85\xa5\x62\x20\xc6\x43\xaa\xc8\xd5\x2b\xb1\x9a\x95\x84\xef\xe8\x94\x75\xae\x3d\x69\xa6\x00\x8d\x9e\x4e\xd9\x70\x65\xa6\xbe\xbe\xa5\x5f\x1b\x77\xb3\x9f\x02\x0f\xd9\xc8\x85\xc2\x13\x4f\xbe\xc7\x87\xdc\x93\x98\x6d\x64\x0b\xe3\x36\x02\xb8\x45\x6f\x1d\xe6\x77\xb2\x4f\xe3\xe8\xce\x08\xfa\x48\xbf\x16\xf4\x4a\x9c\xc6\xed\x34\x93\xc7\x14\xe4\xdf\xa9\xb5\x73\x61\x23\x4f\xfa\xe5\x2c\x62\x11\xc8\x0a\x47\x6a\xb2\x99\x65\x72\x5f\x46\xb0\x36\x79\xa0\x62\x69\x61\xfd\xc7\xe8\x0b\x44\x75\x67\x4a\xb9\x7f\x94\xbe\xe5\xe3\x76\x88\xe8\x8b\x6b\xa6\xef\x96\xbb\x67\xef\x07\xff\x7f\xe1\x2e\x31\xff\x44\x79\x79\x9d\xc3\xa5\xd4\xed\xb2\xdf\x6b\xee\xab\x7d\xff\x6d\x02\xbf\x79\xcb\x28\x62\x04\xf5\x0d\xe6\xf9\x67\x00\x00\x00\xff\xff\x1d\x0d\x69\x06\x27\x04\x00\x00"),
		},
		"/templates/cli_flags_generated.gotpl": &vfsgen۰CompressedFileInfo{
			name:             "cli_flags_generated.gotpl",
			modTime:          time.Time{},
			uncompressedSize: 2230,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x55\x41\x4f\xeb\x38\x10\x3e\x37\xbf\x62\xb6\xe2\x90\xa0\x6e\x7b\x5f\xc4\x89\xc2\x2e\xd2\x0a\xd0\x83\xc7\xdd\x4d\xc6\xa9\x85\xb1\x8b\xed\xc0\x43\x51\xff\xfb\xd3\x8c\x9d\xb4\x4d\x0b\xf4\xc2\xeb\xc9\x9e\xf9\xfc\xcd\x37\xdf\x8c\x9a\xd9\x0c\x2e\x6c\x85\x50\xa3\x41\x27\x02\x56\xb0\x78\x07\x59\x9f\xc1\xfc\x16\x6e\x6e\x1f\xe0\x72\x7e\xfd\x30\xcd\xb2\x95\x28\x9f\x44\x8d\xd0\xb6\x30\xbd\x8b\xe7\x1b\xf1\x8c\xb0\x5e\x67\x99\x7a\x5e\x59\x17\x20\xcf\x46\x63\xa9\x45\x3d\xce\x46\x63\x1f\x5c\x69\xcd\x2b\x1d\x83\x7a\xc6\x71\x56\x64\x6d\x0b\x4e\x98\x1a\xe1\x44\x4d\xe0\x84\x80\xf0\xcf\x39\x4c\xaf\xb4\xa8\x3d\xd1\xcc\x66\x44\xce\x89\x69\xa2\xa6\x1c\x28\x0f\x02\x18\xfe\xa6\xc2\x12\xc2\xfb\x0a\x37\xc0\x07\xba\xad\xd7\xfd\x7d\x6e\xcb\x20\x94\x26\xbe\x5d\xe0\x36\xa3\x0f\xae\x29\x03\xb4\x19\x00\x00\x27\xd2\xcf\x07\xa7\x4c\xcd\xe1\x9f\x9e\xba\xdd\x0b\x5f\x9a\xd7\x47\xe1\xf6\xc2\x57\x4a\xe3\x9d\x08\xcb\x41\xf8\x07\xbe\x34\xca\x61\x45\xe7\x85\xb5\x9a\x83\xff\xa9\xaa\x42\x13\x29\xfa\x60\xdb\xfe\x0d\x4a\x02\xbe\x24\xb9\x8f\x42\x37\x08\xc1\x35\xec\xf0\x88\xaf\xa3\xd1\xb0\xed\x6c\x44\xef\xd0\x54\xfd\x79\x9b\x63\x8e\x3e\x28\x23\x82\xb2\x66\xc3\xb4\x15\x1c\x9d\x7e\xca\xb7\xce\x68\x24\xf7\xdc\x0d\x38\x0c\x8d\x33\x34\x09\x87\xa2\x12\x0b\x8d\xe0\x70\xe5\xd0\xa3\x09\xb1\x82\x95\x10\x96\xca\xc3\x2b\x49\xa5\x97\xb9\xb4\x0e\x1a\xb6\xb1\x42\x29\x1a\x1d\x7c\x91\xc9\xc6\x94\x90\xcb\x83\x73\x29\x52\xb1\xbc\x48\x1e\xa6\x09\xc5\xda\x40\x90\x08\x40\x97\xcb\x22\xe9\xfb\x17\x03\x53\x74\x02\xc3\x12\xc1\x50\x80\xf5\x20\x2f\xce\x17\x55\x13\xc5\x07\x65\x25\x83\x53\xb5\x6b\xdf\x8f\xb4\x2b\xf8\xb6\xc4\xb0\x44\x07\xd6\x81\xb1\xa1\xaf\x49\x6b\xeb\x12\xf6\x8b\xfa\x1b\xd2\xbc\xe0\x8d\x18\x0a\xe8\xd2\x49\xc4\x90\x05\xb4\xb5\x4f\x1e\x9a\x15\x17\x67\xff\xa9\x7b\x01\xda\x96\x42\x1f\x2c\x3a\xe9\xe4\xef\xf0\x5d\x58\x13\xf0\x57\x98\xc7\x69\x11\xb3\x92\xdc\x94\xb4\x8d\xe9\xba\x28\xe1\x34\xe1\x8a\x3d\xea\x9c\x9d\x8f\x2e\x16\x6d\xcb\xcf\x11\x3a\x5e\x5e\xb2\xf1\x98\x78\x87\x25\xd3\xfe\x51\x1c\xb5\xc7\x1d\xc8\x26\xd7\xad\xe6\xae\x3f\xd4\x7d\xb3\x3a\x28\x65\x02\xe5\x94\x82\xf7\x18\xfa\x7d\xd1\x76\x21\xf4\xd1\x16\xd6\x0c\xff\x4e\x0f\x0f\x0b\xfa\xb3\x4e\x2a\x09\xd2\xd3\xbf\x71\xf4\x32\x4a\xba\x8a\xc6\x75\x46\x16\x67\x84\xf9\xeb\x1c\x8c\xea\x16\xf4\xd8\x21\x48\x5f\x30\x7e\xbd\x3d\xb6\x8f\x1d\xa3\x49\xb1\x4f\x9f\x90\x26\x6b\x26\xe0\x31\xc0\x29\xa7\x93\xde\xef\x77\x4b\x92\x53\x1e\xc3\xf4\x7f\xd6\xc7\x7a\x8a\xde\xc7\x7d\x8b\x36\x82\xee\x84\xf3\xe8\xa2\x96\x15\x9d\xab\x09\xa0\x73\xc4\xd7\x97\x4d\x18\xfe\xa6\x25\x4d\x03\xa8\x8c\x5f\x07\xbe\x5b\x97\x1b\xa5\xa9\xe7\xf4\xaf\xdd\x15\xa5\x0f\x81\x73\xfb\x62\x8e\x9b\x40\x87\x5c\x7f\xd8\xc4\x85\xf0\x21\x36\x32\x64\xdb\xca\x6f\x37\x91\x60\xb1\x97\x5d\xbd\xc7\xef\x45\xff\xee\x77\x00\x00\x00\xff\xff\x4d\xaf\x90\xd0\xb6\x08\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/source"].(os.FileInfo),
		fs["/templates"].(os.FileInfo),
	}
	fs["/source"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/source/flag-types.json"].(os.FileInfo),
	}
	fs["/templates"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/templates/altsrc_flags_generated.gotpl"].(os.FileInfo),
		fs["/templates/cli_flags_generated.gotpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
