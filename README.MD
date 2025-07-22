# Web Crawler Go

A simple and efficient web crawler written in Go. This project allows you to crawl websites, extract links, and gather data for analysis or indexing.

## Features

- Fast concurrent crawling
- Configurable depth and domain restrictions
- Extracts links and page metadata
- Simple CLI interface

## Installation

```bash
git clone https://github.com/toyeafo/web-crawler-go.git
cd web-crawler-go
go build -o webcrawler
```

## Usage

```bash
./webcrawler https://example.com 2 10
```

### Options

- `https://example.com` : The starting URL to crawl.
- `2` : Maximum go routines.
- `10` : Maximum number of pages.

## Example

```bash
./webcrawler https://golang.org 3 20
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/foo`)
3. Commit your changes
4. Push to the branch
5. Open a pull request

## License

MIT License
