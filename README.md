# Terrorist dbf file researcher

A utility for researching a file with an electronic database containing a list of individuals and organizations involved in extremist activities

## Install

Clone project
```bash
git clone git@github.com:will-evil/ter_dbf_research.git
```

Go to project directory
```
cd ter_dbf_research/
```

Install dependencies
```bash
go get ./...
```

Build
```bash
go build -o dbf-research main.go
```

## Usage

As an argument, the utility takes paths to files or the path to a directory with files, which will be used for research.

```bash
./dbf-research spaces /tmp/path/file1.dbf, /tmp/path/file1.dbf
```
or
```bash
./dbf-research spaces /tmp/path/my_directory
```
or
```bash
./dbf-research spaces /tmp/path/file1.dbf, /tmp/path/file1.dbf, /tmp/path/my_directory
```

## Terrorist file format

Information about the file can be read [here](http://www.innovbusiness.ru/pravo/DocumShow_DocumID_107162.html)
