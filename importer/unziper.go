package main

import (
	"archive/zip"
	"log"
	"sync"
	"time"
)

type unzipper struct {
	wg sync.WaitGroup
	file string
}

func UnzipFile(filename string)  {
	uz := &unzipper{
		file: filename,
	}

	uz.unzipFile()
}

func (uz *unzipper) unzipFile() {
	r, err := zip.OpenReader(uz.file)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	log.Printf("Unziper is working...")

	tStart := time.Now()

	for _, f := range r.File {
		uz.wg.Add(1)
		// TODO: обработка данных файла
	}

	uz.wg.Wait()

	tEnd := time.Now()

	log.Printf("Unzipper has finished ( Worktime: %v)", tEnd.Sub(tStart))
}
