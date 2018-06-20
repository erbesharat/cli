# cli
Cli is the command line tool designed for resumic users to easily create an empty and new resume.json file via the command line.

 [![state](https://img.shields.io/badge/state-unstable-red.svg)]() [![release](https://img.shields.io/github/release/resumic/cli.svg)](https://github.com/resumic/cli/releases) [![license](https://img.shields.io/github/license/resumic/cli.svg)](LICENSE) [![Build Status](https://travis-ci.org/resumic/cli.svg?branch=master)](https://travis-ci.org/resumic/cli) [![Go Report Card](https://goreportcard.com/badge/github.com/resumic/cli)](https://goreportcard.com/report/github.com/resumic/cli)


# Getting started
```
$ go get github.com/resumic/cli
```
## Using commands
### help 
```
$ resumic help
```
Lists the available commands and flags.
### init
```
$ resumic init
```
Creates an empty resume.json file with examples to give a better understanding of the json file.

### theme
The theme command has the following sub-commands:
#### list
```
$ resumic theme list
```
Lists the avvailable theme for download.

#### get
```
$ resumic theme get [theme-name]
```
Downloads the theme locally for use.

### test 
```
$ resumic test filename
```
### serve
```
$ resumic serve
```
##### flags usage --theme:
```
$ resumic serve --theme [themeName]
```
###  render
```
$ resumic render --theme [themename] --format html
```
Downloads the page in HTML format.
