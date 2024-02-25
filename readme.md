# Dfsync

`dfsync` is a CLI application that allows you to synchronize dot files to your home directory. After running the `dfsync` command, symlinks will be created from all dot files located in source directory to the home directory keeping the correct paths in the file system.

## Installation

There are 2 options to install `dfsync`.

### Golang installation

If you use golang then you have the option to install `dfsync` as a go module.

#### 1. Install package using golang cli

```shell
go install github.com/ofadiman/dfsync@latest
```

#### 2. Verify installation

```shell
dfsync --help
```

### Binary installation

If you do not use golang then you can also install `dfsync` as a standalone binary. Check the [releases](https://github.com/Ofadiman/dfsync/releases) page to find out all possible binary versions.

#### 1. Download the archive containing binary version of your choice:

```shell
wget https://github.com/Ofadiman/dfsync/releases/download/v1.0.0/dfsync-v1.0.0-linux-amd64.tar.gz
```

#### 2. Unzip the archive to the current directory:

```shell
tar -xzvf dfsync-v1.0.0-linux-amd64.tar.gz
```

#### 3. Move the binary to the path that is available under the `$PATH` environment variable:

```shell
sudo mv dfsync /usr/bin/
```

#### 4. Verify installation

```shell
dfsync --help
```

## Options

| Option                  | Default   | Description                                                                                                                                                                                                                                                           |
| ----------------------- | --------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `--source-directory`    | -         | Directory to be used as a source of dot files                                                                                                                                                                                                                         |
| `--dry`                 | false     | Decide whether to execute the command or just display the logs of the operations that would be performed.                                                                                                                                                             |
| `--conflict-resolution` | no-action | Decide what to do when the file in the symlink destination already exists. `no-action` means skip the file, `backup` means change the extension of the existing file to `*.bak` and create a symlink, `override` means delete the existing file and create a symlink. |
