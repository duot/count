# count

Count file metrics: bytes (default is bytes), lines, words, chars.

## Usage

```bash
❯ ./count --help
NAME:
   count all - Count file metric: bytes (default is bytes), lines, words, chars.

USAGE:
   count all [global options] [command [command options]] [arguments...]

VERSION:
   v0.0.1

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --characters, -c  count characters (default: true)
   --words, -w       count words (default: true)
   --lines, -l       count lines (default: true)
   --bytes, -b       count bytes (default: true)
   --help, -h        show help (default: false)
   --version, -v     print the version (default: false)
```

## Build

Run `go build` and `do install` to install on the local go bin.

## todo

- [ ] add test
- [ ] fix default to bytes
- [ ] extract functions to separate lib
- [ ] support reading from std input

## License

See [`LICENSE`](./LICENSE)
