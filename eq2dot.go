package main

import (
  "os"
  "flag"
  "fmt"
)

func main() {
  var inputFile string
  var outputFile string

  consoleOut := flag.Bool("print", false, "Print to the console")

  flag.Parse()

  if flag.NArg() < 1 {
    usage();
  } else {
    inputFile = flag.Arg(0)
    fmt.Println(inputFile)
    if flag.NArg() == 1 {
      outputFile = inputFile + ".dot"
    } else {
      outputFile = flag.Arg(1)
    }
  }

  fmt.Printf("Cmd: %s %s %s\n", os.Args[0], inputFile, outputFile)
  fmt.Println(consoleOut)
}

func usage() {
  fmt.Printf("Usage: %s inputFile [outputFile]\n", os.Args[0])
  flag.PrintDefaults()
}

