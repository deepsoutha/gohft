//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//type Record struct {
//	LineNumber int
//	Timestamp  float64
//	Price      float64
//}
//
//func play() {
//	filenames := []string{"BTCUSDT-2023-01-30.csv", "ETHUSDT-2023-01-30.csv", "ETCUSDT-2023-01-30.csv"}
//	files := make([]*os.File, len(filenames))
//	scanners := make([]*bufio.Scanner, len(filenames))
//	records := make([]*Record, len(filenames))
//
//	for i, filename := range filenames {
//		file, err := os.Open(filename)
//		if err != nil {
//			fmt.Printf("文件打开错误")
//			break
//		}
//		defer file.Close()
//
//		scanner := bufio.NewScanner(file)
//		scanners[i] = scanner
//		files[i] = file
//		scanner.Scan()
//
//		if scanner.Scan() {
//			fields := strings.Split(scanner.Text(), ",")
//			lineNumber := i
//			timestamp, _ := strconv.ParseFloat(fields[1], 64)
//			price, _ := strconv.ParseFloat(fields[2], 64)
//			records[i] = &Record{LineNumber: lineNumber, Timestamp: timestamp, Price: price}
//		}
//	}
//
//	for {//records内部排序
//		sort.Slice(records, func(i, j int) bool {
//			return records[i].Timestamp < records[j].Timestamp
//		})
//		record := records[0]
//		fmt.Printf("%d,%f,%f\n", record.LineNumber, record.Timestamp, record.Price)
//
//
//		records = records[1:]//更新records
//		if len(records) == 0 {
//			break
//		}
//
//		//继续扫描播放过的那个文件
//		scanner := scanners[record.LineNumber]
//		if !scanner.Scan() {
//			continue
//		}
//
//		// 新的数据添加到records
//		fields := strings.Split(scanner.Text(), ",")
//		lineNumber := record.LineNumber
//		timestamp, _ := strconv.ParseFloat(fields[1], 64)
//		price, _ := strconv.ParseFloat(fields[2], 64)
//		newRecord := &Record{LineNumber: lineNumber, Timestamp: timestamp, Price: price}
//		records = append(records, newRecord)
//	}
//
//	for _, file := range files {
//		if file != nil {
//			file.Close()
//		}
//	}
//}
//
//
//
//func main()  {
//	play()
//}






//
//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//type Record struct {
//	LineNumber int
//	Timestamp  float64
//	Price      float64
//}
//
//type MovingAverage struct {
//	sum   float64
//	queue []float64
//	size  int
//}
//
//func (m *MovingAverage) push(val float64) {
//	m.sum += val
//	m.queue = append(m.queue, val)
//	if len(m.queue) > m.size {
//		m.sum -= m.queue[0]
//		m.queue = m.queue[1:]
//	}
//}
//
//func (m *MovingAverage) value() float64 {
//	return m.sum / float64(len(m.queue))
//}
//
//func play() {
//	filenames := []string{"BTCUSDT-2023-01-30.csv", "ETHUSDT-2023-01-30.csv", "ETCUSDT-2023-01-30.csv"}
//	files := make([]*os.File, len(filenames))
//	scanners := make([]*bufio.Scanner, len(filenames))
//	records := make([]*Record, len(filenames))
//
//	for i, filename := range filenames {
//		file, err := os.Open(filename)
//		if err != nil {
//			fmt.Printf("文件打开错误")
//			break
//		}
//		defer file.Close()
//
//		scanner := bufio.NewScanner(file)
//		scanners[i] = scanner
//		files[i] = file
//		scanner.Scan()
//
//		if scanner.Scan() {
//			fields := strings.Split(scanner.Text(), ",")
//			lineNumber := i
//			timestamp, _ := strconv.ParseFloat(fields[1], 64)
//			price, _ := strconv.ParseFloat(fields[2], 64)
//			records[i] = &Record{LineNumber: lineNumber, Timestamp: timestamp, Price: price}
//		}
//	}
//
//	// 初始化均线对象，movingAvg1 计算5日均线，movingAvg2 计算10日均线，size为均线窗口大小
//	movingAvg1 := MovingAverage{size: 5}
//	movingAvg2 := MovingAverage{size: 10}
//
//	// 定义一个数组，用于存放计算均线的结果，最多存储60个
//	movingAvgResults := make([]float64, 60)
//
//	signalCount := 0 // 统计信号计算次数，当信号计算次数超过60次以后，就每接收一个新数据，就计算一次信号
//
//	for {
//		sort.Slice(records, func(i, j int) bool {
//			return records[i].Timestamp < records[j].Timestamp
//		})
//
//		record := records[0]
//		fmt.Printf("%d,%f,%f\n", record.LineNumber, record.Timestamp, record.Price)
//
//		// 计算5日和10日均线
//		movingAvg1.push(record.Price)
//		movingAvg2.push(record.Price)
//
//		if len(movingAvgResults) >= 60 {
//			// 当 movingAvgResults 数组被填满，就每接收一个新数据，就计算一次信号
//			signalCount++
//
//			if signalCount >= 60 {
//				movingAvg1Value := movingAvg1.value()
//				movingAvg2Value := movingAvg2.value()
//
//				// 判断是金叉还是死叉
//				if movingAvg1Value > movingAvg2Value {
//					fmt.Println("产生金叉信号！")
//				} else {
//					fmt.Println("产生死叉信号！")
//				}
//			}
//		}
//
//		records = records[1:] // 更新 records
//		if len(records) == 0 {
//			break
//		}
//
//		// 继续扫描播放过的那个文件
//		scanner := scanners[record.LineNumber]
//		if !scanner.Scan() {
//			continue
//		}
//
//		// 新的数据添加到 records
//		fields := strings.Split(scanner.Text(), ",")
//		lineNumber := record.LineNumber
//		timestamp, _ := strconv.ParseFloat(fields[1], 64)
//		price, _ := strconv.ParseFloat(fields[2], 64)
//		newRecord := &Record{LineNumber: lineNumber, Timestamp: timestamp, Price: price}
//		records = append(records, newRecord)
//
//		// 将当前时间段的收盘价存入 movingAvgResults 数组，用于计算均线
//		movingAvgResults = append(movingAvgResults, price)
//		if len(movingAvgResults) > 60 {
//			movingAvgResults = movingAvgResults[1:]
//		}
//	}
//
//	for _, file := range files {
//		if file != nil {
//			file.Close()
//		}
//	}
//}
//
//func main() {
//	play()
//}
//






//
//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//type Record struct {
//	LineNumber int
//	Timestamp  float64
//	Price      float64
//}
//
//// used for storing latest Record
//
//// St is used to store the latest Record in a stack
//func St(record *Record) {
//	stack:=make( []*Record,60)
//	stack = append(stack, record)
//
//}
//
//func play() {
//	filenames := []string{"BTCUSDT-2023-01-30.csv", "ETHUSDT-2023-01-30.csv", "ETCUSDT-2023-01-30.csv"}
//	files := make([]*os.File, len(filenames))
//	scanners := make([]*bufio.Scanner, len(filenames))
//	records := make([]*Record, len(filenames))
//
//	for i, filename := range filenames {
//		file, err := os.Open(filename)
//		if err != nil {
//			fmt.Printf("文件打开错误")
//			break
//		}
//		defer file.Close()
//
//		scanner := bufio.NewScanner(file)
//		scanners[i] = scanner
//		files[i] = file
//		scanner.Scan()
//
//		if scanner.Scan() {
//			fields := strings.Split(scanner.Text(), ",")
//			lineNumber := i
//			timestamp, _ := strconv.ParseFloat(fields[1], 64)
//			price, _ := strconv.ParseFloat(fields[2], 64)
//			records[i] = &Record{LineNumber: lineNumber, Timestamp: timestamp, Price: price}
//		}
//	}
//
//	for {
//		//records内部排序
//		sort.Slice(records, func(i, j int) bool {
//			return records[i].Timestamp < records[j].Timestamp
//		})
//		record := records[0]
//		//fmt.Printf("%d,%f,%f\n", record.LineNumber, record.Timestamp, record.Price)
//
//		St(record) // add latest Record to stack
//
//
//		records = records[1:] //更新records
//		if len(records) == 0 {
//			break
//		}
//
//		//继续扫描播放过的那个文件
//		scanner := scanners[record.LineNumber]
//		if !scanner.Scan() {
//			continue
//		}
//
//		// 新的数据添加到records
//		fields := strings.Split(scanner.Text(), ",")
//		lineNumber := record.LineNumber
//		timestamp, _ := strconv.ParseFloat(fields[1], 64)
//		price, _ := strconv.ParseFloat(fields[2], 64)
//		newRecord := &Record{LineNumber: lineNumber, Timestamp: timestamp, Price: price}
//		records = append(records, newRecord)
//	}
//
//	for _, file := range files {
//		if file != nil {
//			file.Close()
//		}
//	}
//
//}
//
//func main() {
//	play()
//}






//
//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//type Record struct {
//	LineNumber int
//	Timestamp  float64
//	Price      float64
//}
//
//type LineStack struct {
//	LineNumber  int
//	LineRecords []*Record
//}
//
//var stack []*LineStack // used for storing latest Reƒcord
//
//// St is used to store the latest Record in a stack for each LineNumber
//func St(record *Record) {
//
//	filenames := []string{"BTCUSDT-2023-01-30.csv", "ETHUSDT-2023-01-30.csv", "ETCUSDT-2023-01-30.csv"}
//
//	var allData map[]
//	for i,filename := range filenames{
//		dataSlice := []Record
//		allData[filenames[record.LineNumber] = dataSlice
//
//
//	}
//}
//
//
//func play() {
//	filenames := []string{"BTCUSDT-2023-01-30.csv", "ETHUSDT-2023-01-30.csv", "ETCUSDT-2023-01-30.csv"}
//	files := make([]*os.File, len(filenames))
//	scanners := make([]*bufio.Scanner, len(filenames))
//	records := make([]*Record, len(filenames))
//
//	for i, filename := range filenames {
//		file, err := os.Open(filename)
//		if err != nil {
//			fmt.Printf("文件打开错误")
//			break
//		}
//		defer file.Close()
//
//		scanner := bufio.NewScanner(file)
//		scanners[i] = scanner
//		files[i] = file
//		scanner.Scan()
//
//		if scanner.Scan() {
//			fields := strings.Split(scanner.Text(), ",")
//			lineNumber := i
//			timestamp, _ := strconv.ParseFloat(fields[1], 64)
//			price, _ := strconv.ParseFloat(fields[2], 64)
//			records[i] = &Record{LineNumber: lineNumber, Timestamp: timestamp, Price: price}
//		}
//	}
//
//	for {
//		//records内部排序
//		sort.Slice(records, func(i, j int) bool {
//			return records[i].Timestamp < records[j].Timestamp
//		})
//		record := records[0]
//		fmt.Printf("%d,%f,%f\n", record.LineNumber, record.Timestamp, record.Price)
//
//		St(record) // add latest Record to stack
//
//		records = records[1:] //更新records
//		if len(records) == 0 {
//			break
//		}
//
//		//继续扫描播放过的那个文件
//		scanner := scanners[record.LineNumber]
//		if !scanner.Scan() {
//			continue
//		}
//
//		// 新的数据添加到records
//		fields := strings.Split(scanner.Text(), ",")
//		lineNumber := record.LineNumber
//		timestamp, _ := strconv.ParseFloat(fields[1], 64)
//		price, _ := strconv.ParseFloat(fields[2], 64)
//		newRecord := &Record{LineNumber: lineNumber, Timestamp: timestamp, Price: price}
//		records = append(records, newRecord)
//	}
//
//	for _, file := range files {
//		if file != nil {
//			file.Close()
//		}
//	}
//
//	// print the stack values
//	fmt.Println("--- Stack values ---")
//
//}
//
//func main() {
//	play()
//}
//


//
//package main
//
//import (
//	"awesomeProject2/src"
//	"bufio"
//	"fmt"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//const period = 80
//var pos  = 0
//var entry_time float64
//var entry_price float64
//
//
//type Record struct {
//	LineNumber int
//	Timestamp  float64
//	Price      float64
//}
//
//type Res struct {
//	Symbol int
//	Signal string
//	EntryTime float64
//	OutTime float64
//	Pnl float64
//
//}
//
//type Pos struct {
//	Symbol int
//	Pos int
//	EntryTime float64
//	EntryPrice float64
//	OutTime float64
//
//}
//var recordsMap map[int][]*Record
//var allpnl map[int][]*Res
//var allpos map[int]*Pos
//
//
//func St(record *Record,delay float64) { //接收一个新的record，分别按symbol存成slice。然后看新接收的是哪个symbol，就只触发那个symbol的信号计算
//	//以下是按symbol存record
//	if recordsMap[record.LineNumber] == nil {
//		recordsMap[record.LineNumber] = []*Record{record}
//	} else {
//		recordsMap[record.LineNumber] = append(recordsMap[record.LineNumber], record)
//		if len(recordsMap[record.LineNumber]) > period+1{
//			recordsMap[record.LineNumber] = append(recordsMap[record.LineNumber][:0], recordsMap[record.LineNumber][1:]...)
//		}
//	}
//	//以下是按symbol存pos
//
//
//
//
//
//	rMap :=  recordsMap[record.LineNumber]
//
//	if len(rMap) > period {
//			PriceSlice := []float64{}
//			for _, r := range rMap {
//				PriceSlice = append(PriceSlice, r.Price)
//			}
//
//			//以下是信号生成
//			//
//			sig := src.Strategy(PriceSlice, period)
//			if sig == 1 {
//				//fmt.Println(record.LineNumber,record.Timestamp, "----buy")
//				allpos[record.LineNumber].Pos =1
//				allpos[record.LineNumber].EntryTime = record.Timestamp
//				allpos[record.LineNumber].EntryPrice = record.Price
//
//
//			} else if sig == -1 {
//				allpos[record.LineNumber].Pos =1
//				allpos[record.LineNumber].EntryTime = record.Timestamp
//				allpos[record.LineNumber].EntryPrice = record.Price
//				//fmt.Println(record.LineNumber,record.Timestamp, "----sell")
//			}else {}
//
//
//			fmt.Println(allpos[record.LineNumber])
//
//
//			//以下是收益统计
//
//		//if pos == 1 {
//		//	if record.Timestamp-entry_time > delay {
//		//		pnl := 100 * (record.Price - entry_price) / entry_price
//		//		rr := &Res{Symbol: record.LineNumber, Signal: "buy", EntryTime: entry_time,OutTime:record.Timestamp,Pnl:pnl}
//		//		pos = 0
//		//		fmt.Println(rr)
//		//						//入场价格和时间不需要重置，在新的pos出现后会重新赋值
//		//	}
//		//} else if pos == -1 {
//		//	if record.Timestamp-entry_time > delay {
//		//		pnl := 100 * ( entry_price-record.Price) / entry_price
//		//		rr := &Res{Symbol: record.LineNumber, Signal: "sell", EntryTime: entry_time,OutTime:record.Timestamp,Pnl:pnl}
//		//		pos = 0
//		//		fmt.Println(rr)
//		//					} else {}
//		//				} else {}
//		//			}
//
//		}}
//
//
//func play() {
//	filenames := []string{"BTCUSDT-2023-01-30.csv", "ETHUSDT-2023-01-30.csv", "ETCUSDT-2023-01-30.csv"}
//	files := make([]*os.File, len(filenames))
//	scanners := make([]*bufio.Scanner, len(filenames))
//	records := make([]*Record, len(filenames))
//
//	recordsMap = make(map[int][]*Record)
//
//	for i, filename := range filenames {
//		file, err := os.Open(filename)
//		if err != nil {
//			fmt.Printf("文件打开错误")
//			break
//		}
//		defer file.Close()
//		scanner := bufio.NewScanner(file)
//		scanners[i] = scanner
//		files[i] = file
//		scanner.Scan()
//
//		if scanner.Scan() {//每个文件扫描一行，加入records
//			fields := strings.Split(scanner.Text(), ",")
//			lineNumber := i
//			timestamp, _ := strconv.ParseFloat(fields[1], 64)
//			price, _ := strconv.ParseFloat(fields[2], 64)
//			records[i] = &Record{LineNumber: lineNumber, Timestamp: timestamp, Price: price}
//		}
//	}
//
//	for {//比较时间戳
//
//		sort.Slice(records, func(i, j int) bool {
//			return records[i].Timestamp < records[j].Timestamp
//		})
//		record := records[0]
//		St(record,2000)
//
//		records = records[1:]
//		if len(records) == 0 {
//			break
//		}
//		//继续扫描播放过的那个文件
//		scanner := scanners[record.LineNumber]
//		if !scanner.Scan() {
//			continue
//		}
//		// 新的数据添加到records
//		fields := strings.Split(scanner.Text(), ",")
//		lineNumber := record.LineNumber
//		timestamp, _ := strconv.ParseFloat(fields[1], 64)
//		price, _ := strconv.ParseFloat(fields[2], 64)
//		newRecord := &Record{LineNumber: lineNumber, Timestamp: timestamp, Price: price}
//		records = append(records, newRecord)
//	}
//
//	for _, file := range files {
//		if file != nil {
//			file.Close()
//		}
//	}
//
//}
//
//func main() {
//	play()
//
//}


//
//
//
//package main
//
//import (
//	"awesomeProject2/src"
//	"bufio"
//	"fmt"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//	"time"
//)
//
//
//const period = 80
//var pos = 0
//var entry_time float64
//var entry_price float64
//
//type Record struct {
//	LineNumber int
//	Timestamp  float64
//	Price      float64
//}
//
//
//var recordsMap map[int][]*Record
//
//
//type Trade struct {
//	EntryTime     float64
//	EntryPrice    float64
//	CloseTime     float64
//	ClosePrice    float64
//	ProfitPercent float64
//}
//
//var tradesMap map[int][]*Trade
//
//func St(record *Record, delay float64) {
//	lineNumber := record.LineNumber
//	//以下是按symbol存record
//	if recordsMap[lineNumber] == nil {
//		recordsMap[lineNumber] = []*Record{record}
//		tradesMap[lineNumber] = []*Trade{}
//	} else {
//		recordsMap[lineNumber] = append(recordsMap[lineNumber], record)
//		if len(recordsMap[lineNumber]) > period+1 {
//			recordsMap[lineNumber] = append(recordsMap[lineNumber][:0], recordsMap[lineNumber][1:]...)
//		}
//	}
//
//	rMap := recordsMap[lineNumber]
//
//	if len(rMap) > period {
//		PriceSlice := []float64{}
//		for _, r := range rMap {
//			PriceSlice = append(PriceSlice, r.Price)
//		}
//		trades := tradesMap[lineNumber]
//
//		if len(trades) == 0 { // no trades open for this symbol
//			sig := src.Strategy(PriceSlice, period)
//			if sig == 1 {
//				tradesMap[lineNumber] = append(trades, &Trade{
//					EntryTime:  record.Timestamp,
//					EntryPrice: record.Price,
//				})
//			} else if sig == -1 {
//				tradesMap[lineNumber] = append(trades, &Trade{
//					EntryTime:  record.Timestamp,
//					EntryPrice: record.Price,
//				})
//			}
//		} else {}
//
//		//// check for closing signals
//		//	trade := trades[len(trades)-1] // assume only one trade open at a time
//		//	sig := src.Strategy(PriceSlice, period)
//		//	if sig == 1 {
//		//		// do nothing - keep long trade open
//		//	} else if sig == -1 {
//		//		// close long trade and open new short trade
//		//		closeTrade(record, trade, delay)
//		//		tradesMap[lineNumber] = append(trades, &Trade{
//		//			EntryTime:  record.Timestamp,
//		//			EntryPrice: record.Price,
//		//		})
//		//	} else { // sig == 0 or other value
//		//		// close open trade
//		//		closeTrade(record, trade, delay)
//		//		tradesMap[lineNumber] = trades[:len(trades)-1] // remove closed trade from list
//		//	}
//		fmt.Println(record.LineNumber,trades)
//		}
//
//	}
//
//
//func closeTrade(record *Record, trade *Trade, delay float64) {
//	closeTime := record.Timestamp
//	closePrice := record.Price
//	profitPercent := ((closePrice - trade.EntryPrice) / trade.EntryPrice) * 100.0
//	time.Sleep(time.Duration(delay) * time.Millisecond)
//	fmt.Printf("开仓时间: %.2f, 开仓价格: %.2f, 平仓时间: %.2f, 平仓价格: %.2f, 收益百分比: %.2f%%\n",
//		trade.EntryTime, trade.EntryPrice, closeTime, closePrice, profitPercent)
//	trade.CloseTime = closeTime
//	trade.ClosePrice = closePrice
//	trade.ProfitPercent = profitPercent
//}
//
//func play() {
//	filenames := []string{"BTCUSDT-2023-01-30.csv", "ETHUSDT-2023-01-30.csv", "ETCUSDT-2023-01-30.csv"}
//	files := make([]*os.File, len(filenames))
//	scanners := make([]*bufio.Scanner, len(filenames))
//	records := make([]*Record, len(filenames))
//
//	recordsMap = make(map[int][]*Record)
//
//	for i, filename := range filenames {
//		file, err := os.Open(filename)
//		if err != nil {
//			fmt.Printf("文件打开错误")
//			break
//		}
//		defer file.Close()
//		scanner := bufio.NewScanner(file)
//		scanners[i] = scanner
//		files[i] = file
//		scanner.Scan()
//
//		if scanner.Scan() {//每个文件扫描一行，加入records
//			fields := strings.Split(scanner.Text(), ",")
//			lineNumber := i
//			timestamp, _ := strconv.ParseFloat(fields[1], 64)
//			price, _ := strconv.ParseFloat(fields[2], 64)
//			records[i] = &Record{LineNumber: lineNumber, Timestamp: timestamp, Price: price}
//		}
//	}
//
//	for {//比较时间戳
//
//		sort.Slice(records, func(i, j int) bool {
//			return records[i].Timestamp < records[j].Timestamp
//		})
//		record := records[0]
//		St(record,2000)
//
//		records = records[1:]
//		if len(records) == 0 {
//			break
//		}
//		//继续扫描播放过的那个文件
//		scanner := scanners[record.LineNumber]
//		if !scanner.Scan() {
//			continue
//		}
//		// 新的数据添加到records
//		fields := strings.Split(scanner.Text(), ",")
//		lineNumber := record.LineNumber
//		timestamp, _ := strconv.ParseFloat(fields[1], 64)
//		price, _ := strconv.ParseFloat(fields[2], 64)
//		newRecord := &Record{LineNumber: lineNumber, Timestamp: timestamp, Price: price}
//		records = append(records, newRecord)
//	}
//
//	for _, file := range files {
//		if file != nil {
//			file.Close()
//		}
//	}
//
//}
//
//func main() {
//	play()
//
//}



///勉强还行

//
//package main
//
//import (
//	"awesomeProject2/src"
//	"bufio"
//	"fmt"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//const period = 80
//type Record struct {
//	LineNumber int
//	Timestamp  float64
//	Price      float64
//}
//
//
//type Trade struct {
//	Symbol int
//	EntryTime     float64
//	EntryPrice    float64
//	Pos int
//}
//
//
//
//var recordsMap map[int][]*Record
//
//func St(record *Record) { //接收一个新的record，分别按symbol存成slice。然后看新接收的是哪个symbol，就只触发那个symbol的信号计算
//	if recordsMap[record.LineNumber] == nil {
//		recordsMap[record.LineNumber] = []*Record{record}
//	} else {
//		recordsMap[record.LineNumber] = append(recordsMap[record.LineNumber], record)
//		if len(recordsMap[record.LineNumber]) > period+1{
//			recordsMap[record.LineNumber] = append(recordsMap[record.LineNumber][:0], recordsMap[record.LineNumber][1:]...)
//		}
//	}
//	rMap :=  recordsMap[record.LineNumber]
//
//	tradesMap:=make(map[int][]*Trade)
//	if tradesMap[record.LineNumber] ==nil{
//		trade := &Trade{
//			Symbol:     record.LineNumber,
//			EntryTime:  record.Timestamp,
//			EntryPrice: record.Price,
//			Pos:        0,
//		}
//		tradesMap[record.LineNumber] = append(tradesMap[record.LineNumber], trade)
//	}else {
//		tradesMap[record.LineNumber] = append(tradesMap[record.LineNumber], tradesMap[record.LineNumber][len(tradesMap[record.LineNumber])-1])
//	}
//
//	if len(rMap) > period {
//		PriceSlice := []float64{}
//		for _, r := range rMap {
//			PriceSlice = append(PriceSlice, r.Price)
//		}
//		sig := src.Strategy(PriceSlice, period)
//		if sig == 1 {
//
//			trade := &Trade{
//				Symbol:     record.LineNumber,
//				EntryTime:  record.Timestamp,
//				EntryPrice: record.Price,
//				Pos:        1,
//			}
//			tradesMap[record.LineNumber] = append(tradesMap[record.LineNumber], trade)
//
//		} else if sig == -1 {
//			trade := &Trade{
//				Symbol:     record.LineNumber,
//				EntryTime:  record.Timestamp,
//				EntryPrice: record.Price,
//				Pos:        -1,
//			}
//			tradesMap[record.LineNumber] = append(tradesMap[record.LineNumber], trade)
//		} else {
//			tradesMap[record.LineNumber] = append(tradesMap[record.LineNumber], tradesMap[record.LineNumber][len(tradesMap[record.LineNumber])-1])
//		}
//		for _,r :=range tradesMap[record.LineNumber]{
//			fmt.Println(r)
//		}
//
//		//for _,r :=range (trades){
//		//	//fmt.Println(r.Symbol,r.Pos,r.EntryTime,len(trades))
//		//	if record.Timestamp- r.EntryTime > 2000{
//		//		if r.Pos ==1{
//		//			pnl :=100* (record.Price-r.EntryPrice)/r.EntryPrice
//		//			fmt.Println(r.Symbol,pnl)
//		//
//		//		}else if r.Pos ==-1{
//		//			pnl :=100* (record.Price-r.EntryPrice)/r.EntryPrice
//		//			fmt.Println(r.Symbol,pnl)
//		//		} else {}
//		//	}
//		//}
//
//	}
//}
//
//func play() {
//	filenames := []string{"BTCUSDT-2023-01-30.csv", "ETHUSDT-2023-01-30.csv", "ETCUSDT-2023-01-30.csv"}
//	files := make([]*os.File, len(filenames))
//	scanners := make([]*bufio.Scanner, len(filenames))
//	records := make([]*Record, len(filenames))
//
//	recordsMap = make(map[int][]*Record)
//
//	for i, filename := range filenames {
//		file, err := os.Open(filename)
//		if err != nil {
//			fmt.Printf("文件打开错误")
//			break
//		}
//		defer file.Close()
//		scanner := bufio.NewScanner(file)
//		scanners[i] = scanner
//		files[i] = file
//		scanner.Scan()
//
//		if scanner.Scan() {//每个文件扫描一行，加入records
//			fields := strings.Split(scanner.Text(), ",")
//			lineNumber := i
//			timestamp, _ := strconv.ParseFloat(fields[1], 64)
//			price, _ := strconv.ParseFloat(fields[2], 64)
//			records[i] = &Record{LineNumber: lineNumber, Timestamp: timestamp, Price: price}
//		}
//	}
//
//	for {//比较时间戳
//
//		sort.Slice(records, func(i, j int) bool {
//			return records[i].Timestamp < records[j].Timestamp
//		})
//		record := records[0]
//		St(record)
//
//		records = records[1:]
//		if len(records) == 0 {
//			break
//		}
//		//继续扫描播放过的那个文件
//		scanner := scanners[record.LineNumber]
//		if !scanner.Scan() {
//			continue
//		}
//		// 新的数据添加到records
//		fields := strings.Split(scanner.Text(), ",")
//		lineNumber := record.LineNumber
//		timestamp, _ := strconv.ParseFloat(fields[1], 64)
//		price, _ := strconv.ParseFloat(fields[2], 64)
//		newRecord := &Record{LineNumber: lineNumber, Timestamp: timestamp, Price: price}
//		records = append(records, newRecord)
//	}
//
//	for _, file := range files {
//		if file != nil {
//			file.Close()
//		}
//	}
//
//}
//
//func main() {
//	play()
//
//}



package main

import (
	"awesomeProject2/src"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const period = 80
type Record struct {
	LineNumber int
	Timestamp  float64
	Price      float64
	Side int
	EntryTime float64
	EntryPrice float64
	Profit2s  float64
}

var recordsMap map[int][]*Record


type Data struct {
	symbol int
	sig       int
	Timestamp float64
	Price     float64
	Profit2s float64
}

var dataMap map[int][]Data

func St(record *Record) {
	if recordsMap[record.LineNumber] == nil {
		recordsMap[record.LineNumber] = []*Record{record}
	} else {
		recordsMap[record.LineNumber] = append(recordsMap[record.LineNumber], record)
		if len(recordsMap[record.LineNumber]) > period+1 {
			recordsMap[record.LineNumber] = append(recordsMap[record.LineNumber][:0], recordsMap[record.LineNumber][1:]...)
		}
	}
	rMap := recordsMap[record.LineNumber]

	if dataMap == nil {
		dataMap = make(map[int][]Data)
	}

	if len(rMap) > period {
		PriceSlice := []float64{}
		for _, r := range rMap {
			PriceSlice = append(PriceSlice, r.Price)
		}

		sig := src.Strategy(PriceSlice, period)


		if sig != 0 {
			newData := Data{
				symbol: record.LineNumber,
				sig:       sig,
				Timestamp: rMap[len(rMap)-1].Timestamp,
				Price:     rMap[len(rMap)-1].Price,

			}
			if dataMap[record.LineNumber] == nil {
				dataMap[record.LineNumber] = []Data{newData}
			} else {
				dataMap[record.LineNumber] = append(dataMap[record.LineNumber], newData)

			}
		} else {
			// 如果sig值为0，Data[LineNumbers]等于上一个值
			if dataMap[record.LineNumber] == nil {
				newData := Data{
					symbol: record.LineNumber,
					sig:       0,
					Timestamp: 0,
					Price:     0,
				}
				dataMap[record.LineNumber] = []Data{newData}
			} else {
				prevData := dataMap[record.LineNumber][len(dataMap[record.LineNumber])-1]
				dataMap[record.LineNumber] = append(dataMap[record.LineNumber], prevData)
			}
		}


		if len(dataMap[record.LineNumber]) >1 {//有数据时

			lastData:=dataMap[record.LineNumber][len(dataMap[record.LineNumber])-1]
			preData:=dataMap[record.LineNumber][len(dataMap[record.LineNumber])-2]

			if preData.Profit2s ==0 {
				//还没开始计算收益或者刚生成一个新信号
				//要生成第一个信号

				if record.Timestamp - preData.Timestamp>2000{
					if lastData.sig ==1 {
						dataMap[record.LineNumber][len(dataMap[record.LineNumber])-1].Profit2s = 100 * (record.Price - lastData.Price) / lastData.Price


					}else if lastData.sig ==-1{
						dataMap[record.LineNumber][len(dataMap[record.LineNumber])-1].Profit2s = 100 * (lastData.Price-record.Price ) / lastData.Price
					}

				}
			}else {//此时收益列是有数据的
				if preData.sig == lastData.sig{ //信号没变，时间在增加，此时收益沿用前一个数据的收益
					dataMap[record.LineNumber][len(dataMap[record.LineNumber])-1].Profit2s =preData.Profit2s
				}else {
					//信号变了
					dataMap[record.LineNumber][len(dataMap[record.LineNumber])-1].Profit2s = 0
				}

			}

		}

		//dataMap保留一定长度

		if len(dataMap[record.LineNumber])>period+2 {
			dataMap[record.LineNumber] = append(dataMap[record.LineNumber][:0], dataMap[record.LineNumber][1:]...)
		}
		//fmt.Println((dataMap[2]))
	}
}

func play() {
	filenames := []string{"BTCUSDT-2023-01-30.csv", "ETHUSDT-2023-01-30.csv", "ETCUSDT-2023-01-30.csv"}
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

		if scanner.Scan() {//每个文件扫描一行，加入records
			fields := strings.Split(scanner.Text(), ",")
			lineNumber := i
			timestamp, _ := strconv.ParseFloat(fields[1], 64)
			price, _ := strconv.ParseFloat(fields[2], 64)
			records[i] = &Record{LineNumber: lineNumber, Timestamp: timestamp, Price: price}
		}
	}

	for {//比较时间戳
		sort.Slice(records, func(i, j int) bool {
			return records[i].Timestamp < records[j].Timestamp
		})
		record := records[0]
		St(record)
		//fmt.Println(record)

		records = records[1:]
		if len(records) == 0 {
			break
		}
		//继续扫描播放过的那个文件
		scanner := scanners[record.LineNumber]
		if !scanner.Scan() {
			continue
		}
		// 新的数据添加到records
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

func main() {
	play()

}