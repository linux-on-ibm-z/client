// Code generated by "esc -o fixtures.go -pkg client _fixtures/final.yml _fixtures/m2.yml _fixtures/m3.yml"; DO NOT EDIT.

package client

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/_fixtures/final.yml": {
		local:   "_fixtures/final.yml",
		size:    770,
		modtime: 1510323992,
		compressed: `
H4sIAAAAAAAC/7SRvY7bMBCEez3FwI6bADR1hpEclDKPcEhtUNRKIsw/cEk7fvuAtHzVtakI7M7Ofjvc
46OMzjCb4DEbS5hDAmnCuX/HrKzFqX/72SVlhg64UarKAf3x1AHGqYUG6HsklTh4WbWzupCmc/9+cX89
5csU9JXSoNz04yyWWIRVmTh3iTiUpImrsY6lPoBKejWZdC6JBrShDli+bkfFWtnW0KH4POCtAzzle0jX
oeIzdTo4p/zU1ozF2OnpJCBH4+WoeIXQ2OkIkW6QnLT8DtmUO+yhQ3wgh60CDsgr4YkOFdvlML5VS7RB
TV/a11g56arc0mmLfmEKqKtv+FbbsmXWIEKkpHJIUhfOwclj1Xr6ZNoQ6qdxBWyTyIloA/BGE4Q/9XDq
ShC/N3fsnzk8JzZxNBHGc64/LkRhShD0womPvAYPYP+p2Uqj8ZPxC79uLpza3dk4gpixO/xpXocPfnAm
hwORVZFp2r0sJN2UFayTiZnlbLyyx/jYUhKrWdb/Zm7DvfsXAAD//0SrpRcCAwAA
`,
	},

	"/_fixtures/m2.yml": {
		local:   "_fixtures/m2.yml",
		size:    780,
		modtime: 1510324978,
		compressed: `
H4sIAAAAAAAC/7SRvY7bMBCEez3FwI6bADR1hpEclPJw6VIFqQ2KWlmE+QcuacdvH5CWr7o2FYHd2dlv
h1v8MpY4B084gMvoDLMJHrOxhDkkvL+949i/4qeyFof+5XuXlBk64EqpKgf0+0MHGKfONEDfIqnEwcuq
ndWJNB3715P76ymfpqAvlAblpm9HoWMRVmXi3CXiUJImrsY6lvoAKunFZNK5JBrQhjrg/Hk7KtbKtoYO
xecBfQd4yreQLgNmZZk6HZxTfmprxmLs9HASkKPxclS8QGhsdIRIV0hOWn6FbMoNttAh3pHDWgEH5IXw
QIeK7XIY36ol2qCmT+1rrJx0Va7ptEU/MAXU1Vd8qW3ZMmsQIVJSOSSpC+fg5L5qPX0wrQj107gCtknk
RLQCeKMJwh96OHUhiLfVHdtHDo+JVRxNhPGc648LUZgSBD1x4j0vwQPYfmjW0mj8ZPyZnzcXTu3ubBxB
zNjs/jSv3W++cyaHHZFVkWnaPC0kXZUVrJOJmaU77F/28b6GJBZzXv6Xtw237l8AAAD//6jENYwMAwAA
`,
	},

	"/_fixtures/m3.yml": {
		local:   "_fixtures/m3.yml",
		size:    780,
		modtime: 1510324971,
		compressed: `
H4sIAAAAAAAC/7SRP48aMRDF+/0UTxCaSMbAoeS0KU+XLlWUGhnv7K6F/8ljQ/j2kc1y1bWpLM28efOb
5zV+GUucgye8gMvZGWYTPEZjCWNIeH97x3H3ip/KWhx2++9dUqbvgCulquyx2x46wDg1UQ99i6QSBy+r
dlQn0nTcvZ7cX0/5NAR9odQrN3w7iikWYVUmzl0iDiVp4mqsY6kPoJKeTSadS6IebagDps/bUbFWtjV0
KD732HeAp3wL6dJjVJap08E55Ye25lyMHR5OAvJsvDwrniE0VjpCpCskJy2/QjblCmvoEO/IYamAA/JM
eKBDxXY5jG/VEm1Qw6f2NVZOuiqXdNqiHxgC6uorvtS2bJk1iBApqRyS1IVzcHJbtZ4+mBaE+mlcAdsk
ciJaALzRBOEPOzh1IYi3xR3rRw6PiUUcTYTxnOuPC1GYEgQ9ceI9z8EDWH9oltLZ+MH4iZ83F07t7mwc
QYxYbf40r81vvnMmhw2RVZFpWD0tJF2VFayTiZmle9nut/G+hCRmM83/y9uGW/cvAAD//4J/SdsMAwAA
`,
	},

	"/": {
		isDir: true,
		local: "",
	},

	"/_fixtures": {
		isDir: true,
		local: "_fixtures",
	},
}
