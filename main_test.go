package main

import (
	"bytes"
	"testing"
)

func TestCountBytes(t *testing.T) {
  b := bytes.NewBufferString("word1word2word3word4word5word6")

  expected := 4

  res := count(b, false, true)

  if res != expected {
    t.Errorf("Expected %d, got %d instead", expected, res)
  }
}

func TestCountLines(t *testing.T) {
  b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")

  expected := 3

  res := count(b, true,  false)

  if res != expected {
    t.Errorf("Expected %d, got %d instead", expected, res)
  }
}

func TestCountWords (t *testing.T) {

  b := bytes.NewBufferString("word1 word2 word3 word4\n")

  expected := 4

  res := count(b, false, false)

  if res != expected {
    t.Errorf("Expected %d, got %d instead", expected, res)
  }
}
