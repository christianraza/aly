# Aly
A cross-platform command-line assistant, at their core Aly is a parser that handles Google's organic SERP results and processes them all for you.  

Inspired by [Tuxi](https://github.com/Bugswriter/tuxi), written in [Go](https://golang.org/), and with the help of [goquery](https://github.com/PuerkitoBio/goquery) for parsing, Aly answers all those burning questions directly from the command line!  

![aly](assets/aly.gif)

## Installation
##### Go
If you don't already have Go installed, follow this link to get started: https://golang.org/doc/install  

Alternatively, you can use your package manager of choice, for Windows I recommend [Scoop](https://github.com/lukesampson/scoop), for Linux check your distro, and for Mac I recommend [Homebrew](https://github.com/Homebrew/brew).  

Then just install Go using your package manager.
##### Aly
After installing Go just run the command below, Go will automatically download any dependencies and install Aly to the [GOBIN](https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies)!
```
$ go install github.com/christianraza/aly
```
If you don't want to download Go, don't worry! You can get the executable from the [releases](https://github.com/christianraza/aly/releases). Just find your operating system and architecture, download the zip file, unzip it, and add the executable's directory to your path or execute it manually if you just wanna test everything out :smiley:
## Usage
```
$ aly what is Git?
```
General usage takes the form of `aly [flags] <query>` where `[flags]` are optional and `<query>` is required.

For a list of flags and any additional information, run Aly with the help flag:
```
$ aly -h
```
## Notes
##### Special characters
It's important to note that depending on your shell there may be special characters, such as `'` and `"` that usually delimit strings. In these situations where you must explicitly use special characters, wrap your query in quotes and everything will work!
