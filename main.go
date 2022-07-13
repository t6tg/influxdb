package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	csv := readCsvFile("data.csv")
	client := influxdb2.NewClient("http://localhost:8086", "AoYJomd8FtlAfG1jUE_h0TRxabhRa15HlZcgzEoHFK-CPgtszef9fgRjxlMHWbXXPDDBCoQSdA0IhQ6qoLW-OQ==")
	defer client.Close()
	for _, v := range csv {
		v[8] = strings.ReplaceAll(v[8], " ", "T")
		v[8] = fmt.Sprintf("%sZ", v[8])
		dt, err := time.Parse(time.RFC3339, v[8])
		if err != nil {
			fmt.Println(err)
			return
		}
		novel_id, err := strconv.Atoi(v[2])
		if err != nil {
			return
		}
		buyer_id, err := strconv.Atoi(v[1])
		if err != nil {
			return
		}
		coin, err := strconv.Atoi(v[6])
		if err != nil {
			return
		}
		writeAPI := client.WriteAPI("dek_d", "writer_transaction")
		p := influxdb2.NewPointWithMeasurement("transaction_6").AddField("novel_id", novel_id).AddField("buyer_id", buyer_id).AddField("coin", coin).SetTime(dt).AddTag("novel_id", v[2])
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
