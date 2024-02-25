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

func (this *FileSystemBuilder) Build() func(additionalCleanupPaths ...string) {
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

	return func(additionalCleanupPaths ...string) {
		for _, symlink := range this.symlinks {
			err := os.RemoveAll(symlink[1])
			if err != nil {
				panic(err)
			}
		}

		for _, file := range this.files {
			err := os.RemoveAll(file)
			if err != nil {
				panic(err)
			}
		}

		for _, dir := range this.directories {
			err := os.RemoveAll(dir)
			if err != nil {
				panic(err)
			}
		}

		for _, additionalPath := range additionalCleanupPaths {
			err := os.RemoveAll(additionalPath)
			if err != nil {
				panic(err)
			}
		}
	}
}
