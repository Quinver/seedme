<h1 align="center">seedme</h1>

<p align="center"><strong><i>A magnet scraper in your CLI</i></strong></p>

**seedme** Is a command line tool that let's you watch video's using magnet links. It streams directly to your mpv. 
I know there are a lot of other tools like this but this one is mine, and I have the superior name.

## Table of Contents
- [Features](#features)
- [Installation](#installation)
  - [Build From Source](#build-from-source)
    - [Dependencies](#dependencies)
- [Usage](#usage)

## Features

* User-friendly
* Watch the video while it is being downloaded
* Sorted results from 2 different providers at once (Might add more later) 

## Installation
### Dependencies
* mpv
* fzf
* [webtorrent-cli](https://github.com/webtorrent/webtorrent-cli)

### Build From Source

```bash
$ git clone https://github.com/Quinver/seedme
$ cd seedme
$ go build cmd/seedme/
```

## Usage
```bash
$ ./seedme <name>
```
