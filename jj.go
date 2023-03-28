////package jj
//
//import (
//	"awesomeProject2/src"
//	"encoding/csv"
//	"log"
//	"math"
//	"os"
//	"strconv"
//	"strings"
//)
//
//var pos  = 0
//var entry_time float64
//var entry_price float64
//const period int = 80
//
////func run_st(delay float64,fileName string) {
////	file, err := os.Open(fileName)
////	if err != nil {
////		log.Fatal(err)
////	}
////	defer file.Close()
////	reader := csv.NewReader(file)
////	rows, err := reader.ReadAll()
////	if err != nil {
////		log.Fatal(err)
////	}
////	var priceSlice []float64
////	var tsSlice []float64
////	var result [][]string
////
////	for _, row := range rows {
////		pp, _ := strconv.ParseFloat(row[2], 64)
////		tt, _ := strconv.ParseFloat(row[1], 64)
////		length := len(tsSlice)
////
////			if length>0 { // 这是直接切数据
////				if tt - tsSlice[length-1] >1000 { //更新价格和时间戳
////					priceSlice = append(priceSlice,pp)
////					tsSlice = append(tsSlice,tt)
////				}else if tt ==tsSlice[length-1]{
////					priceSlice[length-1] = pp //如果时间戳相同，更新pp
////				}else {}//其他情况不更新
////			}else {//还没数据时，更新第一个
////				priceSlice = append(priceSlice,pp)
////				tsSlice = append(tsSlice,tt)
////			}
////
////		if length > period+1 {
////			priceSlice = append(priceSlice[:0], priceSlice[1:]...) //删掉，保持长度在80
////			tsSlice = append(tsSlice[:0], tsSlice[1:]...)
////		}
////
////
////		//以下是信号和数据统计
////		if length > period+1 {
////			signal := src.St(priceSlice,period)
////			if signal == 1 {
////				pos = 1
////				entry_price = pp
////				entry_time = tt
////
////			} else if signal == -1 {
////				pos = -1
////				entry_price = pp
////				entry_time = tt
////
////			} else {}
////
////			if pos == 1 { //当上一个row带来的pos变成1以后，新进来的row启动时间戳检查
////				if tt-entry_time > delay {
////					var res []string
////					pnl := 100 * (pp - entry_price) / entry_price
////					res = append(res, "buy", strconv.FormatFloat(entry_time, 'f', 2, 64), strconv.FormatFloat(tt, 'f', 2, 64), strconv.FormatFloat(pnl, 'f', 2, 64))
////					result = append(result, res)
////					//fmt.Println("buy ", tt-entry_time, strconv.FormatFloat(tt, 'f', 2, 64), strconv.FormatFloat(entry_time, 'f', 2, 64), pnl)
////					pos = 0
////					//入场价格和时间不需要重置，在新的pos出现后会重新赋值
////				}
////			} else if pos == -1 {
////				if tt-entry_time > delay {
////					var res []string
////					pnl := 100 * (entry_price - pp) / entry_price
////					res = append(res, "sell", strconv.FormatFloat(entry_time, 'f', 2, 64), strconv.FormatFloat(tt, 'f', 2, 64), strconv.FormatFloat(pnl, 'f', 2, 64))
////					result = append(result, res)
////					pos = 0
////				} else {}
////			} else {}
////		}
////
////	}
////
////	src.WriteToCSV(strings.Split(fileName,".")[0]+"_pnl.csv", result)
////
////}
//


////读出所有文件
////遍历所有文件，逐个按时间戳播放，记录下当前播放的索引
////
//func run_st(delay float64,fileName string) {
//		file, err := os.Open(fileName)
//		if err != nil {
//			log.Fatal(err)
//		}
//		defer file.Close()
//		reader := csv.NewReader(file)
//		rows, err := reader.ReadAll()
//		if err != nil {
//			log.Fatal(err)
//		}
//		var priceSlice []float64
//		var tsSlice []float64
//		var result [][]string
//
//		for _, row := range rows {
//			pp, _ := strconv.ParseFloat(row[2], 64)
//			tt, _ := strconv.ParseFloat(row[1], 64)
//			length := len(tsSlice)
//
//				if length>1 {
//					last_t := tsSlice[length-1]
//					last_p := priceSlice[length-1]
//
//					if tt - tsSlice[length-1] >500 && tt - tsSlice[length-1] <1000  {
//						priceSlice = append(priceSlice,pp)
//						tsSlice = append(tsSlice,last_t + 500)
//					}else if tt - tsSlice[length-1] >1000 {
//
//
//						cishu := math.Floor((tt - tsSlice[length-1])/500) //
//
//
//						for  i:=1; i<int(cishu); i++ {
//							priceSlice = append(priceSlice,last_p)
//							tsSlice = append(tsSlice,last_t + float64(500*i))
//
//						}
//
//						priceSlice = append(priceSlice,pp)
//						tsSlice = append(tsSlice,last_t + float64(cishu * 500))
//					}else {}//小于500不更新
//
//				}else {//还没数据时，更新第一个
//					priceSlice = append(priceSlice,pp)
//					tsSlice = append(tsSlice,tt)
//
//				}
//
//
//			if length > period+1 {
//				priceSlice = append(priceSlice[:0], priceSlice[1:]...) //删掉，保持长度在80
//				tsSlice = append(tsSlice[:0], tsSlice[1:]...)
//			}
//
//			//以下是信号和数据统计
//			if length > period+1 {
//				//fmt.Println(tsSlice[80])
//				signal := src.St(priceSlice,period)
//				if signal == 1 {
//					pos = 1
//					entry_price = pp
//					entry_time = tt
//
//				} else if signal == -1 {
//					pos = -1
//					entry_price = pp
//					entry_time = tt
//
//				} else {}
//
//				if pos == 1 { //当上一个row带来的pos变成1以后，新进来的row启动时间戳检查
//					if tt-entry_time > delay {
//						var res []string
//						pnl := 100 * (pp - entry_price) / entry_price
//						res = append(res, "buy", strconv.FormatFloat(entry_time, 'f', 2, 64), strconv.FormatFloat(tt, 'f', 2, 64), strconv.FormatFloat(pnl, 'f', 2, 64))
//						result = append(result, res)
//						//fmt.Println("buy ", tt-entry_time, strconv.FormatFloat(tt, 'f', 2, 64), strconv.FormatFloat(entry_time, 'f', 2, 64), pnl)
//						pos = 0
//						//入场价格和时间不需要重置，在新的pos出现后会重新赋值
//					}
//				} else if pos == -1 {
//					if tt-entry_time > delay {
//						var res []string
//						pnl := 100 * (entry_price - pp) / entry_price
//						res = append(res, "sell", strconv.FormatFloat(entry_time, 'f', 2, 64), strconv.FormatFloat(tt, 'f', 2, 64), strconv.FormatFloat(pnl, 'f', 2, 64))
//						result = append(result, res)
//						pos = 0
//					} else {}
//				} else {}
//			}
//
//		}
//
//		src.WriteToCSV(strings.Split(fileName,".")[0]+"_pnl.csv", result)
//
//	}
//
//
//
//
//
////func main() {
////	run_st(2000,"BTCUSDT-2023-01-30.csv")
////}





//以上是单文件


package main
//func getNextTick()  {
//	files := []string{"BTCUSDT-2023-01-30.csv", "ETHUSDT-2023-01-30.csv", "ETCUSDT-2023-01-30.csv"}
//	fh := []*os.File{}
//	for _, f := range files {
//		fileHandle, err := os.Open(f)
//		if err != nil {
//			fmt.Println("Error opening file: ", err)
//		}
//		defer fileHandle.Close()
//		fh = append(fh, fileHandle)
//	}
//
//	scanners := make([]*bufio.Scanner, 3)
//	for i, f := range fh {
//		scanner := bufio.NewScanner(f)
//		scanner.Scan() // skip header
//		scanners[i] = scanner
//	}
//
//	records := make([]*record, 3)
//	for i, s := range scanners {
//		if s.Scan() {
//			line := s.Text()
//			recordFields, _ := csv.NewReader(bytes.NewBufferString(line)).Read()
//			timestamp, _ := strconv.ParseFloat(recordFields[1], 64)
//			price, _ := strconv.ParseFloat(recordFields[2], 64)
//			records[i] = &record{Filename: i, Timestamp: timestamp, Price: price, Line: recordFields}
//		} else {
//			records[i] = nil
//		}
//	}
//
//	for { //找到时间戳最小的一行
//		sort.Slice(records, func(i, j int) bool {
//			return records[i].Timestamp < records[j].Timestamp
//		})
//		smallestTimestamp := records[0].Timestamp
//
//		//逐行播放
//		var wg sync.WaitGroup
//		for _, r := range records {
//			if r != nil && r.Timestamp == smallestTimestamp {
//				wg.Add(1)
//				go func(record *record) {
//
//					if scanners[record.Filename].Scan() {
//						line := scanners[record.Filename].Text()
//
//						recordFields, _ := csv.NewReader((bytes.NewBufferString(line))).Read()
//						timestamp, _ := strconv.ParseFloat(recordFields[1], 64)
//						price, _ := strconv.ParseFloat(recordFields[2], 64)
//						record.Timestamp = timestamp
//						record.Price = price
//						record.Line = recordFields
//					} else {
//						record = nil
//					}
//
//					fmt.Println(strings.Split(files[record.Filename], "-")[0], record.Timestamp)
//
//					wg.Done()
//				}(r)
//			}
//		}
//		wg.Wait()
//
//	}
//
//	for _, f := range fh {
//		f.Close()
//	}
//}

//func getNextTick()  {
//	files := []string{"BTCUSDT-2023-01-30.csv", "ETHUSDT-2023-01-30.csv", "ETCUSDT-2023-01-30.csv"}
//	fh := []*os.File{}
//	for _, f := range files {
//		fileHandle, err := os.Open(f)
//		if err != nil {
//			fmt.Println("Error opening file: ", err)
//		}
//		defer fileHandle.Close()
//		fh = append(fh, fileHandle)
//	}
//
//	scanners := make([]*bufio.Scanner, 3)
//	for i, f := range fh {
//		scanner := bufio.NewScanner(f)
//		scanner.Scan() // skip header
//		scanners[i] = scanner
//	}
//	records := make([]*record, 3)
//	for i, s := range scanners {
//		if s.Scan() {
//			line := s.Text()
//			recordFields, _ := csv.NewReader(bytes.NewBufferString(line)).Read()
//			timestamp, _ := strconv.ParseFloat(recordFields[1], 64)
//			price, _ := strconv.ParseFloat(recordFields[2], 64)
//			records[i] = &record{Filename: i, Timestamp: timestamp, Price: price, Line: recordFields}
//		} else {
//			records[i] = nil
//		}
//	}
//
//	var currentRecord *record
//	leftRecords := make([]*record, 0)
//
//	for _, r := range records {
//		if r != nil {
//			leftRecords = append(leftRecords, r)
//		}
//	}
//
//	for len(leftRecords) > 0 {
//		sort.Slice(leftRecords, func(i, j int) bool {
//			return leftRecords[i].Timestamp < leftRecords[j].Timestamp
//		})
//
//		// Get the next available record
//		currentRecord = leftRecords[0]
//
//		if scanners[currentRecord.Filename].Scan() {
//			line := scanners[currentRecord.Filename].Text()
//
//			recordFields, _ := csv.NewReader((bytes.NewBufferString(line))).Read()
//			timestamp, _ := strconv.ParseFloat(recordFields[1], 64)
//			price, _ := strconv.ParseFloat(recordFields[2], 64)
//
//			nextRecord := &record{Filename: currentRecord.Filename, Timestamp: timestamp, Price: price, Line: recordFields}
//
//			leftRecords[0] = nextRecord
//
//			fmt.Println(strings.Split(files[nextRecord.Filename], "-")[0], nextRecord.Timestamp)
//
//		} else {
//			leftRecords = leftRecords[1:]
//		}
//	}
//
//	fmt.Println("All records have been played.")
//	for _, f := range fh {
//		f.Close()
//	}
//}
//
//
//func main()  {
//	getNextTick()
//
//}

//单币种，维护一个tick时间序列
//然后数据接收，逐个播放

//
//package main
//
//import (
//	"bufio"
//	"bytes"
//	"encoding/csv"
//	"errors"
//	"fmt"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//type record struct {
//	Filename  int
//	Timestamp float64
//	Price     float64
//	Line      []string
//}
//
//func getNextTick() ([]string, error) {
//	files := []string{"BTCUSDT-2023-01-30.csv", "ETHUSDT-2023-01-30.csv", "ETCUSDT-2023-01-30.csv"}
//	fh := []*os.File{}
//	for _, f := range files {
//		fileHandle, err := os.Open(f)
//		if err != nil {
//			fmt.Println("Error opening file: ", err)
//			return nil, err
//		}
//		defer fileHandle.Close()
//		fh = append(fh, fileHandle)
//	}
//
//	scanners := make([]*bufio.Scanner, 3)
//	for i, f := range fh {
//		scanner := bufio.NewScanner(f)
//		scanner.Scan() // skip header
//		scanners[i] = scanner
//	}
//
//	records := make([]*record, 3)
//	for i, s := range scanners {
//		if s.Scan() {
//			line := s.Text()
//			recordFields, _ := csv.NewReader(bytes.NewBufferString(line)).Read()
//			timestamp, _ := strconv.ParseFloat(recordFields[1], 64)
//			price, _ := strconv.ParseFloat(recordFields[2], 64)
//			records[i] = &record{Filename: i, Timestamp: timestamp, Price: price, Line: recordFields}
//		} else {
//			records[i] = nil
//		}
//	}
//
//	var currentRecord *record
//	leftRecords := make([]*record, 0)
//
//	for _, r := range records {
//		if r != nil {
//			leftRecords = append(leftRecords, r)
//		}
//	}
//
//	for len(leftRecords) > 0 {
//		sort.Slice(leftRecords, func(i, j int) bool {
//			return leftRecords[i].Timestamp < leftRecords[j].Timestamp
//		})
//
//		// Get the next available record
//		currentRecord = leftRecords[0]
//
//		if scanners[currentRecord.Filename].Scan() {
//			line := scanners[currentRecord.Filename].Text()
//
//			recordFields, _ := csv.NewReader((bytes.NewBufferString(line))).Read()
//			timestamp, _ := strconv.ParseFloat(recordFields[1], 64)
//			price, _ := strconv.ParseFloat(recordFields[2], 64)
//
//			nextRecord := &record{Filename: currentRecord.Filename, Timestamp: timestamp, Price: price, Line: recordFields}
//
//			leftRecords[0] = nextRecord
//
//			return []string{strings.Split(files[nextRecord.Filename], "-")[0], strconv.FormatFloat(nextRecord.Timestamp, 'f', -1, 64)}, nil
//
//		} else {
//			leftRecords = leftRecords[1:]
//		}
//	}
//
//	for _, f := range fh {
//		f.Close()
//	}
//
//	return nil, errors.New("all records have been played")
//}
//
//func main() {
//	for {
//		line, err := getNextTick()
//		if err != nil {
//			fmt.Println(err.Error())
//			break
//
//
//
//		}
//
//		fmt.Println(line)
//	}
//}




////////////以下是一次性读
////package main
////////////
////////////import (
////////////	"bufio"
////////////	"encoding/csv"
////////////	"fmt"
////////////	"os"
////////////)
////////////
//////////////import (
//////////////	"bufio"
//////////////	"encoding/csv"
//////////////	"fmt"
//////////////	"os"
//////////////	"sort"
//////////////	"strconv"
//////////////	"strings"
//////////////)
//////////////type Data struct {
//////////////	symbol string
//////////////	timestamp float64
//////////////	price     float64
//////////////}
//////////////func readCSV(filePath string) []Data {
//////////////	var data []Data
//////////////	file, _ := os.Open(filePath)
//////////////	symbol := strings.Split(filePath,"-")[0]
//////////////	defer file.Close()
//////////////	reader := csv.NewReader(bufio.NewReader(file))
//////////////	records, _ := reader.ReadAll()
//////////////	for _, record := range records {
//////////////		timestamp, _ := strconv.ParseFloat(record[1],64)
//////////////		price, _ := strconv.ParseFloat(record[2], 64)
//////////////		data = append(data, Data{symbol,timestamp, price})
//////////////	}
//////////////	return data
//////////////}
//////////////
//////////////func main() {
//////////////	var allData []Data
//////////////	fileList := []string{"BTCUSDT-2023-01-30.csv", "ETHUSDT-2023-01-30.csv", "ETCUSDT-2023-01-30.csv"}
//////////////	for _, file := range fileList {
//////////////		dd := readCSV(file)
//////////////		for _, row := range dd {
//////////////			allData = append(allData, row)
//////////////		}
//////////////		sort.Slice(allData, func(i, j int) bool {
//////////////			return allData[i].timestamp < allData[j].timestamp
//////////////		})
//////////////	}
//////////////	fmt.Println("ok")
//////////////
//////////////}
////////////////一次读完，逐个读
////////////
////////////
////////////func main() {
////////////	// 打开 CSV 文件
////////////	fileList := []string{"BTCUSDT-2023-01-30.csv", "ETHUSDT-2023-01-30.csv", "ETCUSDT-2023-01-30.csv"}
////////////	for _, file := range fileList {
////////////		csvFile, err := os.Open(file)
////////////		if err != nil {
////////////			fmt.Println(err)
////////////			return
////////////		}
////////////		defer csvFile.Close()
////////////
////////////		// 创建 CSV 读取器
////////////		reader := csv.NewReader(bufio.NewReader(csvFile))
////////////
////////////		// 逐行读取 CSV 中的数据
////////////		for {
////////////			line, _ := reader.Read()
////////////			// 获取每一行的文本内容
////////////			fmt.Println(line)
////////////
////////////		}
////////////
////////////	}
////////////}
//////////package main
//////////
//////////import (
//////////	"bufio"
//////////	"encoding/csv"
//////////	"fmt"
//////////	"os"
//////////	"sort"
//////////	"strconv"
//////////	"strings"
//////////)
//////////
//////////type Record struct {
//////////	Timestamp float64
//////////	Price     float64
//////////	Symbol  string
//////////}
//////////
//////////type ByTimestamp []Record
//////////
//////////func (a ByTimestamp) Len() int           { return len(a) }
//////////func (a ByTimestamp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
//////////func (a ByTimestamp) Less(i, j int) bool { return a[i].Timestamp < a[j].Timestamp }
//////////
//////////func main() {
//////////	files := []string{"BTCUSDT-2023-01-30.csv", "ETHUSDT-2023-01-30.csv", "ETCUSDT-2023-01-30.csv"}
//////////
//////////	for _, file := range files {
//////////		fmt.Printf("Reading file: %s\n", file)
//////////
//////////		// Open the file
//////////		f, err := os.Open(file)
//////////		if err != nil {
//////////			panic(err)
//////////		}
//////////		defer f.Close()
//////////
//////////		// Read the file with CSV scanner
//////////		scanner := bufio.NewScanner(f)
//////////		var records []Record
//////////		for scanner.Scan() {
//////////			line := scanner.Text()
//////////			fields, err := csv.NewReader(strings.NewReader(line)).Read()
//////////			if err != nil {
//////////				fmt.Printf("Error parsing line: %s\n", line)
//////////				continue
//////////			}
//////////
//////////			// Parse the timestamp and price from the CSV record
//////////			ts, err := strconv.ParseFloat(fields[0], 64)
//////////			if err != nil {
//////////				fmt.Printf("Error parsing timestamp: %s\n", fields[0])
//////////				continue
//////////			}
//////////			price, err := strconv.ParseFloat(fields[1], 64)
//////////			if err != nil {
//////////				fmt.Printf("Error parsing price: %s\n", fields[1])
//////////				continue
//////////			}
//////////
//////////			// Add the record to the slice
//////////			records = append(records, Record{
//////////				Timestamp: ts,
//////////				Price:     price,
//////////				Symbol: strings.Split(file,"-")[0],
//////////			})
//////////		}
//////////
//////////		// Sort the records by timestamp
//////////		sort.Sort(ByTimestamp(records))
//////////
//////////		// Play the records
//////////		for _, r := range records {
//////////			fmt.Printf("%.1f,%.2f\n", r.Timestamp, r.Price,r.Symbol)
//////////		}
//////////	}
//////////}
////////
////////






