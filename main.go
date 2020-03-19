package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	file := flag.String("f", "", "File name")
	n := flag.Int("p", 2, "Number of parts")
	flag.Parse()
	f, err := os.Open(*file)
	if err != nil {
		log.Fatalln(err)
	}

	stat, err := f.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	size := stat.Size()/int64(*n) + 1
	baseName := filepath.Base(*file)
	for i := 1; ; i++ {
		b := make([]byte, size)
		n, err := f.Read(b)
		if err == io.EOF {
			fmt.Println("done creating files")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		b = b[:n]
		partf, err := os.Create(fmt.Sprintf("%s.%d", baseName, i))
		if err != nil {
			log.Fatalln(err)
		}
		n, err = partf.Write(b)
		if err != nil {
			log.Fatalln(err)
		}
		if n < len(b) {
			fmt.Println("done creating files")
			partf.Close()
			break
		}
		partf.Close()
	}
}
