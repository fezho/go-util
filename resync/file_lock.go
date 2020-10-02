package gpu

import (
	"errors"
	"os"
	"syscall"
)

https://github.com/etcd-io/bbolt/blob/master/bolt_unix.go

type fileLock struct {
	filename string
	fd       int
}

func newFileLock(filename string) *fileLock {
	return &fileLock{filename: filename}
}

func (l *fileLock) GetFilename() string {
	return l.filename
}

func (l *fileLock) open() error {
	fd, err := syscall.Open(l.filename, syscall.O_CREAT|syscall.O_RDONLY, 0600)
	if err != nil {
		return err
	}
	l.fd = fd
	return nil
}

func (l *fileLock) TryLock() error {
	if err := l.open(); err != nil {
		return err
	}
	err := syscall.Flock(l.fd, syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		syscall.Close(l.fd) // nolint
	}
	if err == syscall.EWOULDBLOCK {
		return errors.New("locked")
	}
	return err
}

func (l *fileLock) Unlock() error {
	return syscall.Close(l.fd)
}

func (l *fileLock) Remove() error {
	return os.Remove(l.filename)
}
