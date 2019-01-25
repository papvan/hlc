package main

import "flag"

var baseURL = flag.String("url", "http://localhost", "base URL (for loader)")
var nWorkers = flag.Int("w", 8, "number of parallel requests while loading data")
//var dataFileName = flag.String("data", "example.json", "data file name")
var dataFileName = flag.String("data", "test-data/test_accounts.zip", "data file name")

func main() {
	flag.Parse()
	UnzipFile(*dataFileName)
	//LoadData(*baseURL, *dataFileName, *nWorkers)
}
