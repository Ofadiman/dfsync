package main

import (
	"os"
)

type FileSystemBuilder struct {
	files       []string
	directories []string
	symlinks    [][2]string
}

func (this *FileSystemBuilder) File(path string) *FileSystemBuilder {
	this.files = append(this.files, path)
	return this
}

func (this *FileSystemBuilder) Directory(path string) *FileSystemBuilder {
	this.directories = append(this.directories, path)
	return this
}

func (this *FileSystemBuilder) Symlink(source, target string) *FileSystemBuilder {
	this.symlinks = append(this.symlinks, [2]string{source, target})
	return this
}

func (this *FileSystemBuilder) Build() func() {
	for _, dir := range this.directories {
		err := os.MkdirAll(dir, 0700)
		if err != nil {
			panic(err)
		}
	}

	for _, file := range this.files {
		err := os.WriteFile(file, []byte{}, 0700)
		if err != nil {
			panic(err)
		}
	}

	for _, symlink := range this.symlinks {
		err := os.Symlink(symlink[0], symlink[1])
		if err != nil {
			panic(err)
		}
	}

	// Cleanup function
	return func() {
		for _, symlink := range this.symlinks {
			err := os.Remove(symlink[1])
			if err != nil {
				panic(err)
			}
		}

		for _, file := range this.files {
			err := os.Remove(file)
			if err != nil {
				panic(err)
			}
		}

		for _, dir := range this.directories {
			err := os.Remove(dir)
			if err != nil {
				panic(err)
			}
		}
	}
}
