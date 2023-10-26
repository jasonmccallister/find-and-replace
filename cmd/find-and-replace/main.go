package main

import (
	"bytes"
	"flag"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	var fileName = flag.String("filename", "", "File name to search")
	var find = flag.String("find", "", "String to find")
	var replace = flag.String("replace", "", "String to replace")
	flag.Parse()

	var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

	if *fileName == "" || *find == "" || *replace == "" {
		flag.PrintDefaults()
		return
	}

	f, err := os.Open(*fileName)
	if err != nil {
		logger.Error("Error opening file", err)
		return
	}
	defer f.Close()

	buffer := new(bytes.Buffer)
	if _, err := buffer.ReadFrom(f); err != nil {
		logger.Error("Error reading file", err)
		return
	}

	output := bytes.Replace(buffer.Bytes(), []byte(*find), []byte(*replace), -1)

	if err = os.WriteFile("modified-"+*fileName, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Done. Output file to modified-" + *fileName + ".")
}
