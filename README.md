# gfetch
A minimal system fetch written in Go.

# Table of Contents
- [Usage](#usage)
    - [Examples](#examples)
    - [Flags](#flags)
- [Configuration](#configuration)
- [Building](#building)
- [File Architecture](#file-architecture)
- [License](#license)

# Usage
```shell
gfetch
```

## Examples
```shell
gfetch
```

```shell
gfetch --help
```

## Flags
```
-h, --help      help for gfetch
-v, --version   version for gfetch
```

# Configuration
You can customize title and value text styles with this simple command:
```shell
// Command will added soon...
```

# Building
- Install [Go](https://go.dev/) programming language.
    - Arch Linux:
        ```shell
        pacman -S go
        ```
    - Debian / Ubuntu or based distros:
        ```shell
        apt install golang
        ```
    - MacOS or All Linux Distros:
        ```shell
        brew install go
        ```
      
- Clone Repository
    - With [Git](https://git-scm.com/):
        ```shell
        git clone https://github.com/lnxwizard/gfetch.git
        ```
    - With [GitHub CLI](https://cli.github.com/):
        ```shell
        gh repo clone lnxwizard/gfetch
        ```

- Change Directory
    ```shell
    cd gfetch
    ```

- Build the Program with Go
    ```shell
    go build ./cmd/gfetch
    ```

- Run Program
	```
	./gfetch
	```
  
# File Architecture
```
gfetch
├── cmd
│   └── gfetch
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   └── system
│       ├── color.go
│       ├── desktop.go
│       ├── distro.go
│       ├── kernel.go
│       ├── memory.go
│       └── shell.go
├── LICENSE
├── pkg
│   ├── cmd
│   │   ├── bug
│   │   │   └── bug.go
│   │   └── root
│   │       └── root.go
│   └── user
│       ├── hostname.go
│       └── username.go
├── README.md
└── Taskfile.yaml

11 directories, 17 files
```

# License
MIT License
