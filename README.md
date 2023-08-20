# gfetch
A minimal system fetch written in Go.

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
```shell
gfetch
├── cmd // main file
│   └── gfetch
│       └── main.go
├── internal 
│   ├── config // config library
│   │   └── config.go
│   └── system // system library for getting distro, kernel, memory etc... information
│       ├── color.go
│       ├── desktop.go
│       ├── distro.go
│       ├── kernel.go
│       ├── memory.go
│       └── shell.go
├── pkg
│   ├── cmd // all commands
│   │   └── root
│   │       └── root.go
│   └── user // user library for getting hostname and username
│       ├── hostname.go
│       └── username.go
├── go.mod
├── go.sum
├── README.md
└── Taskfile.yaml // Task runner
```
