package main

import (
    "fmt"
    "flag"
    "os"
    "io/ioutil"
    "crypto/md5"
    "crypto/sha1"
)

func main() {
    // Get the name of the file.
    flag.Parse()
    fmt.Fprintf(os.Stdout, "Filename: %s\n", flag.Arg(0))

    // Slurp the file into a string.
    file, err := ioutil.ReadFile(flag.Arg(0))
    if (err != nil) {
        fmt.Fprintf(os.Stdout, "Error reading the file: %s", err.String())
        os.Exit(1)
    }

    // Create the MD5 hash.
    digest1 := md5.New()
    _, err = digest1.Write(file)
    if (err != nil) {
        fmt.Fprintf(os.Stdout, "Error creating the first hash: %s\n", err.String())
        os.Exit(1)
    }
    fmt.Fprintf(os.Stdout, "MD5: %x\n", digest1.Sum())

    // Create the SHA-1 hash.
    digest2 := sha1.New()
    _, err = digest2.Write(file)
    if (err != nil) {
        fmt.Fprintf(os.Stdout, "Error creating the first hash: %s\n", err.String())
        os.Exit(1)
    }
    fmt.Fprintf(os.Stdout, "SHA-1: %x\n", digest2.Sum())
}
