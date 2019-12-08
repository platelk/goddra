package asset

import (
	"io"
	"io/ioutil"
)

// Asset define a default asset behavior
type Asset struct {
	ready bool
	name  string
	data  io.ReadCloser
	bytes []byte
}

func NewAssetFromReader(name string, r io.ReadCloser) *Asset {
	return &Asset{
		ready: true,
		name:  name,
		data:  r,
		bytes: nil,
	}
}

// Name will return the name of the asset
func (a *Asset) Name() string {
	return a.name
}

// IsReady will return true if the asset has been
func (a *Asset) IsReady() bool {
	return a.ready
}

// Stream will return the underlying io.Reader of the data.
// Warning: it can be already close or empty
func (a *Asset) Stream() io.ReadCloser {
	return a.data
}

// DataBytes return the full read data in form of bytes slices
func (a *Asset) DataBytes() []byte {
	if len(a.bytes) == 0 {
		a.readData()
	}
	return a.bytes
}

// DataString return the full read data in form of string
func (a *Asset) DataString() string {
	return string(a.DataBytes())
}

func (a *Asset) readData() {
	b, err := ioutil.ReadAll(a.data)
	if err != nil {
		return
	}
	a.bytes = b
}