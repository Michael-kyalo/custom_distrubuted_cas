package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log"
	"os"
	"strings"
)

func PathTransformation(key string) PathData {
	hash := sha1.Sum([]byte(key))

	hashStr := hex.EncodeToString(hash[:])

	blockSize := 5
	sliceSize := len(hashStr) / blockSize

	path := make([]string, sliceSize)

	for i := 0; i < sliceSize; i++ {
		from, to := i*blockSize, (i*blockSize)+blockSize
		path[i] = hashStr[from:to]
	}

	return PathData{
		Path: strings.Join(path, "/"),
		Key:  hashStr,
	}

}

type TransformPathFunc func(string) PathData

var NoTransformPathFunc = func(key string) PathData {
	return PathData{
		Path: key,
		Key:  key,
	}
}

type PathData struct {
	Path string
	Key  string
}

func (p PathData) FullPath() string {
	return p.Path + "/" + p.Key
}

type StoreConfig struct {
	TransformPathFunc TransformPathFunc
}
type Store struct {
	StoreConfig
}

func NewStore(storeConfig StoreConfig) *Store {
	return &Store{storeConfig}
}

func (s *Store) readStream(key string) (io.Reader, error) {

	pathName := s.TransformPathFunc(key)

	file, err := os.Open(pathName.FullPath())
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (s *Store) writeStream(key string, reader io.Reader) error {
	pathName := s.TransformPathFunc(key)

	if err := os.MkdirAll(pathName.Path, 0755); err != nil {
		return err
	}

	buf := new((bytes.Buffer))
	io.Copy(buf, reader)
	filenameBytes := md5.Sum(buf.Bytes())
	filename := hex.EncodeToString(filenameBytes[:])
	file, err := os.Create(pathName.Path + "/" + filename)

	if err != nil {
		return err
	}

	n, err := io.Copy(file, buf)

	if err != nil {
		return err
	}

	log.Printf("written %d bytes ", n)

	return nil
}
