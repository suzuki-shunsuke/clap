# clap

[![Build Status](https://github.com/suzuki-shunsuke/clap/workflows/CI/badge.svg)](https://github.com/suzuki-shunsuke/clap/actions)
[![codecov](https://codecov.io/gh/suzuki-shunsuke/clap/branch/master/graph/badge.svg)](https://codecov.io/gh/suzuki-shunsuke/clap)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/clap)](https://goreportcard.com/report/github.com/suzuki-shunsuke/clap)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/clap.svg)](https://github.com/suzuki-shunsuke/clap)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/clap/master/LICENSE)

Simple installer

## Install

Download from [GitHub Releases](https://github.com/suzuki-shunsuke/clap/releases)

## Usage

install - Download a file and extract files from downloaded file and install them

```
$ clap install <URL> <file path in archive>:<install path> [<file path in archive>:<install path> ...]
```

ex. Install [cmdx](https://github.com/suzuki-shunsuke/cmdx) to /usr/local/bin/cmdx

```
$ clap install https://github.com/suzuki-shunsuke/cmdx/releases/download/v1.6.0/cmdx_1.6.0_darwin_amd64.tar.gz cmdx:/usr/local/bin/cmdx
```

## License

[MIT](LICENSE)
