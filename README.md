# Waitress

A simple little static file server. It can be used to serve a custom browser start page.

### Installation

Navigate to the project folder and run `go build` (need to have go installed).
Then copy or symlink built file to local bin:

link
```bash
ln waitress ~/.local/bin/waitress
```

copy

```bash
cp waitress ~/.local/bin/
```

### Usage

```bash
waitress --path <path_to_file> --port <port_number>
```
