# clap

[![Build Status](https://github.com/suzuki-shunsuke/clap/workflows/CI/badge.svg)](https://github.com/suzuki-shunsuke/clap/actions)
[![codecov](https://codecov.io/gh/suzuki-shunsuke/clap/branch/master/graph/badge.svg)](https://codecov.io/gh/suzuki-shunsuke/clap)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/clap)](https://goreportcard.com/report/github.com/suzuki-shunsuke/clap)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/clap.svg)](https://github.com/suzuki-shunsuke/clap)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/clap/master/LICENSE)

Simple installer

## What's clap? Why do we develop clap?

clap is a simple CLI tool to install tools.
clap does the following things.

1. download an archved file such as tarball
2. extract files from the downloaded file and install them to specified paths

We can do the same thing by shell scripts with some commands like `mktemp`, `curl`, `tar`, and `zip`.
But we don't want to write such shell scripts for each tools.
We don't want to search how to use these commands everytime we write the install script.

It's bothersome.

With `clap` we can install tools by simply specifying required parameters.

For example, to install [conftest](https://www.conftest.dev/), we have to write the following shell script.

```sh
#!/usr/bin/env bash

set -eu

CONFTEST_VERSION=0.18.2

dirpath=$(mktemp -d)
pushd "$dirpath"
TARFILE=conftest_${CONFTEST_VERSION}_Linux_x86_64.tar.gz
curl -OL https://github.com/instrumenta/conftest/releases/download/v${CONFTEST_VERSION}/${TARFILE}
tar xvzf $TARFILE
mv conftest /usr/local/bin/conftest
chmod a+x /usr/local/bin/conftest
popd
rm -R "$dirpath"
```

On the other hand, we can do the same thing with clap more simply.

```sh
#!/usr/bin/env bash

set -eu

CONFTEST_VERSION=0.18.2
clap install https://github.com/instrumenta/conftest/releases/download/v${CONFTEST_VERSION}/conftest_${CONFTEST_VERSION}_Linux_x86_64.tar.gz conftest:/usr/local/bin/conftest
chmod a+x /usr/local/bin/conftest
```

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

## Blog (written in Japanese)

https://techblog.szksh.cloud/clap/

## License

[MIT](LICENSE)
