package data

import (
	"bytes"
	"sync"
)

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

//func Test_() {
//	p := new(SyncedBuffer) // type *SyncedBuffer
//	var v SyncedBuffer     // type SyncedBuffer
//}

//func NewFile(fd int, name string) *File {
//	if fd < 0 {
//		return nil
//	}
//	f := new(File)
//	f.fd = fd
//	f.name = name
//	f.dirinfo = nil
//	f.nepipe = 0
//	return f
//}
//
//func NewFile(fd int, name string) *File {
//	if fd < 0 {
//		return nil
//	}
//	f := File{fd, name, nil, 0}
//	return &f
//}
