package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const period = 180 //信号计算的tick长度

type Record struct {
	LineNumber int
	Timestamp  float64
	Price      float64
}
type Data struct {
	symbol    int
	sig       int
	Timestamp float64
	Price     float64
}

var recordsMap map[int][]*Record
var dataMap map[int][]Data

var rows [][]string
func Strategy(priceSlice []float64,period int) int {

	ma1 := MA(priceSlice, period/2)
	ma2 := MA(priceSlice, period)
	if ma1[len(ma1)-1] > ma2[len(ma2)-1] && ma1[len(ma1)-2] < ma2[len(ma2)-2] {
		return 1
	} else if ma1[len(ma1)-1] < ma2[len(ma2)-1] && ma1[len(ma1)-2] > ma2[len(ma2)-2] {
		return -1
	} else {
		return 0
	}
}

func MA(data []float64, windowSize int) []float64 {
	ma := make([]float64, len(data)-windowSize+1)
	sum := 0.0
	for i := 0; i < windowSize; i++ {
		sum += data[i]
	}
	ma[0] = sum / float64(windowSize)
	for i := windowSize; i < len(data); i++ {
		sum += data[i] - data[i-windowSize]
		ma[i-windowSize+1] = sum / float64(windowSize)
	}
	return ma
}


func Run(record *Record, delay float64,files []string) {

	//接收最新的一个数据，根据不同symbol也就是LineNumer，存到固定长度的切片下
	if recordsMap[record.LineNumber] == nil {
		recordsMap[record.LineNumber] = []*Record{record}
	} else {
		recordsMap[record.LineNumber] = append(recordsMap[record.LineNumber], record)
		if len(recordsMap[record.LineNumber]) > period+1 {
			recordsMap[record.LineNumber] = append(recordsMap[record.LineNumber][:0], recordsMap[record.LineNumber][1:]...)
		}
	}
	rMap := recordsMap[record.LineNumber] //接收到的当前一个数据，属于哪个symbol，就触发哪个的信号计算

	if dataMap == nil { //dataMap用来存信号
		dataMap = make(map[int][]Data)
	}

	if len(rMap) > period { //数据长度达到后，触发信号计算
		PriceSlice := []float64{}
		for _, r := range rMap {
			PriceSlice = append(PriceSlice, r.Price)
		}
		sig := Strategy(PriceSlice, period) //策略出信号，1，-1，,0

		if sig != 0 {
			newData := Data{
				symbol:    record.LineNumber,
				sig:       sig,
				Timestamp: rMap[len(rMap)-1].Timestamp,
				Price:     rMap[len(rMap)-1].Price,
			}
			dataMap[record.LineNumber] = append(dataMap[record.LineNumber], newData)
			//如果有非0的新信号，存到当前dataMap下当前symbol里
		} else {
		}

		if len(dataMap[record.LineNumber]) > 0 { //如果信号长度大于0,那么每进来一个新的数据，可以对比最老的一个信号，是否达到离场条件，
			// 如果达到，则输出收益，删掉最老的一个信号
			if rMap[len(rMap)-1].Timestamp-dataMap[record.LineNumber][0].Timestamp > delay {

				signal:=dataMap[record.LineNumber][0].sig
				outTime:= rMap[len(rMap)-1].Timestamp
				entryTime:=dataMap[record.LineNumber][0].Timestamp
				var pnl float64
				if signal ==1{
					pnl = 100*(rMap[len(rMap)-1].Price-dataMap[record.LineNumber][0].Price)/dataMap[record.LineNumber][0].Price
				}else if signal ==-1{
					pnl = -100*(rMap[len(rMap)-1].Price-dataMap[record.LineNumber][0].Price)/dataMap[record.LineNumber][0].Price

				}
				row := []string{
					strings.Split(files[record.LineNumber],"-")[0],
					strconv.Itoa(signal),
					strconv.FormatFloat(entryTime, 'f', 1, 64),
					strconv.FormatFloat(outTime, 'f', 1, 64),
					strconv.FormatFloat(pnl, 'f', 5, 64),

				}
				rows = append(rows, row)
				dataMap[record.LineNumber] = append(dataMap[record.LineNumber][:0], dataMap[record.LineNumber][1:]...)
			}
		}
	}


	//fmt.Println("输出完毕")
}

func Play(delay float64,filenames []string) {

	files := make([]*os.File, len(filenames))
	scanners := make([]*bufio.Scanner, len(filenames))
	records := make([]*Record, len(filenames))
	recordsMap = make(map[int][]*Record)

	for i, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("文件打开错误")
			break
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanners[i] = scanner
		files[i] = file
		scanner.Scan()

		if scanner.Scan() { //每个文件扫描一行，加入records
			fields := strings.Split(scanner.Text(), ",")
			lineNumber := i
			timestamp, _ := strconv.ParseFloat(fields[1], 64)
			price, _ := strconv.ParseFloat(fields[2], 64)
			records[i] = &Record{LineNumber: lineNumber, Timestamp: timestamp, Price: price}
		}
	}

	for { //比较时间戳
		sort.Slice(records, func(i, j int) bool {
			return records[i].Timestamp < records[j].Timestamp
		})
		record := records[0]
		Run(record, delay,filenames) //此处触发信号计算
		//fmt.Println(record) //此处逐个播放数据

		records = records[1:]
		if len(records) == 0 {
			break
		}
		//继续扫描播放过的那个文件
		scanner := scanners[record.LineNumber]
		if !scanner.Scan() {
			continue
		}
		// 新的数据添加到剩余的records
		fields := strings.Split(scanner.Text(), ",")
		lineNumber := record.LineNumber
		timestamp, _ := strconv.ParseFloat(fields[1], 64)
		price, _ := strconv.ParseFloat(fields[2], 64)
		newRecord := &Record{LineNumber: lineNumber, Timestamp: timestamp, Price: price}
		records = append(records, newRecord)
	}

	for _, file := range files {
		if file != nil {
			file.Close()
		}
	}
}

func WriteCsv(name string)  {
	file, _ := os.Create(name)
	defer file.Close()
	header := []string{"symbol", "signal", "entryTime", "outTime", "pnl"}
	writer := csv.NewWriter(file)
	writer.Write(header)
	writer.WriteAll(rows)
	writer.Flush()

}
func main() {
	filenames := []string{"BTCUSDT-2023-01-30.csv", "ETHUSDT-2023-01-30.csv", "ETCUSDT-2023-01-30.csv"}
	Play(10000,filenames)
	WriteCsv("pnlDelay10.csv")


}
