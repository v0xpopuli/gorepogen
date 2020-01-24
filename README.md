[![Build Status](https://travis-ci.com/v0xpopuli/gorepogen.svg?branch=master)](https://travis-ci.com/v0xpopuli/gorepogen) [![Coverage Status](https://coveralls.io/repos/github/v0xpopuli/gorepogen/badge.svg?branch=master)](https://coveralls.io/github/v0xpopuli/gorepogen?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/v0xpopuli/gorepogen)](https://goreportcard.com/report/github.com/v0xpopuli/gorepogen)
[![codebeat badge](https://codebeat.co/badges/8f05934f-566f-45fb-abdc-7df276f03c7b)](https://codebeat.co/projects/github-com-v0xpopuli-gorepogen-master)
    
# GOREPOGEN
Simple Go tool, for repositories auto generation. 


#### **Installation**
```go get -u -v github.com/v0xpopuli/gorepogen/.../```, then from downloaded ```go build -o $GOPATH/bin/gorepogen```


#### **Usage**
Use ```gorepogen -h``` for help
```
NAME:
   gorepogen - tool for repositories auto generation

USAGE:
   gorepogen [global options]

VERSION:
   1.0.0

AUTHOR:
   v0xpopuli <vadim.rozhkalns@gmail.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --name value, -n value  Entity name
   --root value, -r value  Project root
   --help, -h              show help (default: false)
   --version, -v           print the version (default: false)
```
If arg ```-r``` or ```--root``` were not define, program take current directory path

#### **Note**
[GORM](https://github.com/jinzhu/gorm) are used as ORM

