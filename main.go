package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

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

func main() {
  // Define a boolean flag called l
  // with deafult value false
  lines := flag.Bool("l", false, "count lines")
  bytes := flag.Bool("b", false, "count bytes")

  // parse the flags
  flag.Parse()

  fmt.Println(count(os.Stdin, *lines, *bytes))
}
