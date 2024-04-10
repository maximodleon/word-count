# Word count

Learning Golang by implementing a simple word counter

## Usage

- Count lines from stdin

```bash
echo "line1\nline2" | go run main.go -l

```

- Count bytes from stdin
```bash
echo "line1\nline2" | go run main.go -b
```

- Count words from a file

```bash
go run main.go -f "file"
```

- Count from several files
```bash
go run main.go -file "file1" -file "file2"
```
