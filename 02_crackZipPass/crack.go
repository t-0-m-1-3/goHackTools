package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/alexmullins/zip"
)

var (
	zipfile    string
	dictionary string
)

func init() {
	flag.StringVar(&zipfile, "f", "", "Open zipfile")
	flag.StringVar(&dictionary, "d", "", "Open pass dictionary")
}

func main() {
	// parse flags
	flag.Parse()

	// output the help
	if zipfile == "" || dictionary == "" {
		println("Please " + os.Args[0] + " -h")
		os.Exit(0)
	}

	// open the password protected archive
	zipr, err := zip.OpenReader(zipfile)
	if err != nil {
		log.Fatal(err)
	}

	// open the dictionary file
	// ../00_addMaterials/dict.txt
	dictFile, err := os.Open(dictionary)
	if err != nil {
		log.Fatalln(err)
	}
	defer dictFile.Close()

	scanner := bufio.NewScanner(dictFile)
	for scanner.Scan() {
		pass := scanner.Text()
		for _, z := range zipr.File {
			z.SetPassword(pass)
			_, err := z.Open()
			// if  ок
			if err == nil {
				println("[+] Found password")
				println("[+] Password = " + pass)
				os.Exit(0)
			}
		}
	}
}
