package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	file := flag.String("file", "", "enter file")
	url := flag.String("url", "", "enter url")
	token := flag.String("token", "", "enter token")
	flag.Parse()
	csv := readCsvFile(*file)
	client := influxdb2.NewClient(*url, *token)
	defer client.Close()
	for _, v := range csv {
		v[8] = strings.ReplaceAll(v[8], " ", "T")
		v[8] = fmt.Sprintf("%sZ", v[8])
		dt, err := time.Parse(time.RFC3339, v[8])
		if err != nil {
			fmt.Println(err)
			return
		}
		price, err := strconv.ParseFloat(v[7], 64)
		if err != nil {
			return
		}
		coin, err := strconv.Atoi(v[6])
		if err != nil {
			return
		}
		writeAPI := client.WriteAPI("dekd", "wallet")
		p := influxdb2.NewPointWithMeasurement("novel_transaction").AddField("coin", coin).AddField("price", price).AddTag("novel_id", v[2]).AddTag("owner_id", v[5]).AddTag("product_type", v[3]).SetTime(dt)
		fmt.Println("datetime :", dt)
		writeAPI.WritePoint(p)
		writeAPI.Flush()
	}
	fmt.Println("success :)")
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
