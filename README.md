# pwgen

cli password generator by Golang

## Introduction

only type `pw`

	% pwgen
	ZIqUwU0znI
	kLmDwR70LM
	84ZlsZKQWX
	kzwtaGJynq
	WVQ89MpG7s
	c6DPUJy6Fx
	vznrIwJVcJ
	8EpfBIoHRk
	JFTX3OhitP
	4igsXqCiKf

## Installation

### step1. Install pwgen

* go get (for developer)
* Homebrew
* download binaries (for end-user)

If you are developer (case: Linux, MacOS)

	$ go get github.com/jniltinho/go-pwgen

	$ cd $GOPATH/src/github.com/jniltinho/go-pwgen/cmd/pw

	$ go build -ldflags "-s -w" -v -o pwgen cmd/pw/*.go
	

or Homebrew (case: MacOS)

	$ brew tap jniltinho/go-pwgen

	$ brew install go-pwgen

If you are end-user, you can download release binaries.

<https://github.com/girigiribauer/go-pwgen/releases>

### step2. You can read manual

	$ pwgen -h


## Options

See `pw -h`

    GLOBAL OPTIONS:
       --length value, -l value  password length (range is [8...64]) (default: 10)
       --count value, -c value   number of counts (default: 10)
       --digit                   add password characters 0-9
       --alphabetlarge           add password characters A-Z
       --alphabetsmall           add password characters a-z
       --underscore, -u          add password character Underscore(\_)
       --specialchars, -s        add password characters special characters, exclude from Space, Backslash, Underscore and Delete
       --help, -h                show help
       --version, -v             print the version

