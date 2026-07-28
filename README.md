# Golang

- the following folders contain only small practice scripts and are personal note for learning Go.
## install go

- install go

```sh
sudo dnf install golang
```

- use alternative `goenv`

```sh
git clone https://github.com/syndbg/goenv.git ~/.goenv
```

- add environment variables in bashrc file

```sh
# goenv for multiple versions
export GOENV_ROOT="$HOME/.goenv"
export PATH="$GOENV_ROOT/bin:$PATH"
export GOENV_AUTOMATICALLY_DETECT_VERSION=1
eval "$(goenv init -)"
```

```sh
source ~/.bashrc
```

```sh
goenv install 1.26.5
goenv global 1.26.5
```

```sh
exec $SHELL
```

- verify instalation

```sh
go version
which go
```

- use the following command to use  a specific version of GO in your personal project

```sh
cd folder_project
goenv local 1.17.2
```

## datatype

- `uint` stores positive data types

`unit`
* uint8   unsigned 8-bit   integers (0 to 255)
* uint16  unsigned 16-bit  integers (0 to 65535)
* uint32  unsigned 32-bit  integers (0 to 4294967295)
* uint64  unsigned 64-bit  integers (0 to 18446744073709551615)

- `INT` stores both positive and negative data types

`int`
* int8   signed 8-bit  integers (-128 to 127)
* int16  signed 16-bit integers (-32768 to 32767)
* int32  signed 32-bit integers (-2147483648 to 2147483647)
* int64  signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

`byte` // alias for uint8

`rune` // alias for int32
       // represents a Unicode code point

`float32` `float64`

## references

- check out [the unicode character table](https://en.wikipedia.org/wiki/List_of_Unicode_characters).
- check out [Blank identifier](https://www.naukri.com/code360/library/blank-identifier-in-go)

### goenv

- check out [manage diferents versions of go](https://medium.com/@tinchoram/c%C3%B3mo-administrar-diferentes-versiones-de-golang-con-goenv-ca4f2cbb84c5)
- check out [goenv github](https://github.com/go-nv/goenv/tree/master)
- check out [goenv installation](https://github.com/go-nv/goenv/blob/master/INSTALL.md)

