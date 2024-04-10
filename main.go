package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type files []string

func (f *files) Set(value string) error {
  *f = append(*f, value)
  return nil
}

func (f *files) String() string {
  return strings.Join(*f, ",")
}

func count(r io.Reader, countLines bool, countBytes bool) int {
  // This is used to read text from a reader
  scanner := bufio.NewScanner(r)

  // Split by words
  // only if the flag is not set
  if !countLines {
    scanner.Split(bufio.ScanWords)
  } else if countBytes {
    scanner.Split(bufio.ScanBytes)
  }

  // init the word counter
  wc := 0

  // count the words
  for scanner.Scan() {
    if countBytes {
      wc = wc + len(scanner.Bytes())
    } else {
      wc++
    }
  }

  // Return the word count
  return wc
}

func getFile(filename string) (*os.File) {
    f, err := os.Open(filename)

    if err != nil {
      panic(err)
    }

  return f
}

func main() {
  // Define a boolean flag called l
  // with deafult value false
  var filesArg files
  lines := flag.Bool("l", false, "count lines")
  bytes := flag.Bool("b", false, "count bytes")
  file := flag.String("f", "", "File to read from")
  flag.Var(&filesArg, "file", "files to read from")

  // parse the flags
  flag.Parse()

  reader := os.Stdin

  // If we need to read more than one file
  if filesArg.String() != "" {
    counter := 0
    // Go through each file
    for _, file := range filesArg {
      f := getFile(file)
      defer f.Close()
      reader = f
      // sum line or byte count
      count := count(reader, *lines, *bytes)
      counter = counter + count
      fmt.Fprintf(os.Stdout, "File: %v, count: %v\n", f.Name(), counter)
    }

    // print the total count
    fmt.Println(counter)
    // we don't need to execute the remainig code
    return
  }

  // If reaeding from file
  if *file != "" {
    f := getFile(*file)
    defer f.Close()

    reader = f
  }

  fmt.Println(count(reader, *lines, *bytes))
}
