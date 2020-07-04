# clap

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
