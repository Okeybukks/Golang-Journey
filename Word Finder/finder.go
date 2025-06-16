package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Log struct {
	Text      string
	LineNum   int
	Timestamp time.Time
}

type Logs []Log

func (logs *Logs) AddLog(line string, lineNum int) {
	stringLine := strings.SplitN(line, "] ", 2)
	timestamp := logs.ParseTimestamp(stringLine[0][1:])
	title := stringLine[1]

	log := Log{
		Text:      title,
		LineNum:   lineNum,
		Timestamp: timestamp,
	}

	*logs = append(*logs, log)

}

func (logs Logs) ParseTimestamp(timestamp string) time.Time {
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, timestamp)
	if err != nil {
		fmt.Printf("%v", err)
		log.Fatal()
	}

	return parsedTime
}

func (logs Logs) TimestampChecker(beforeTime, afterTime, checkTime time.Time) bool {
	if (checkTime.After(beforeTime) && checkTime.Before(afterTime)) || (checkTime.Equal(beforeTime) || checkTime.Equal(afterTime)) {
		return true
	}
	return false
}

func (logs *Logs) fileScanner(filePath, keyword string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("%v", err)
		log.Fatal()
	}
	defer file.Close()

	lineNumber := 1

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.Contains(strings.ToLower(line), strings.ToLower(keyword)) {
			logs.AddLog(line, lineNumber)
		}
		lineNumber++
	}

}

func mainPrint(index int, log Log) {
	fmt.Printf("📌 Match # %v\n", index+1)
	fmt.Printf("🕒 Time: %v\n", log.Timestamp)
	fmt.Printf(">> [Line %v] %v\n", log.LineNum, log.Text)
	fmt.Println("_______________________________________________________________________________")
	fmt.Println()
}

func filteredLogs(logs Logs, head, tail int) Logs {
	logsLen := len(logs)
	if head > -1 && tail == -1 {
		return logs[:head]
	}
	if head == -1 && tail > -1 {
		return logs[logsLen-tail:]
	}

	return logs
}

func lengthLogs(logs Logs) {
	fmt.Println("Search Result")
	fmt.Printf("🔄 Total Query Found: %d\n", len(logs))
}

func (logs Logs) Print(fromTime, toTime string, head, tail int) {

	if fromTime == "" && toTime == "" {
		filtered := filteredLogs(logs, head, tail)
		for index, log := range filtered {
			mainPrint(index, log)
		}

		lengthLogs(filtered)

		return
	}

	if fromTime != "" && toTime == "" {
		localLogs := Logs{}
		fTime := logs.ParseTimestamp(fromTime)
		for _, log := range logs {
			if fTime.Equal(log.Timestamp) || fTime.Before(log.Timestamp) {
				localLogs = append(localLogs, log)
			}
		}
		filtered := filteredLogs(localLogs, head, tail)
		for index, log := range filtered {
			mainPrint(index, log)
		}
		lengthLogs(filtered)
		return
	}

	if fromTime == "" && toTime != "" {
		localLogs := Logs{}
		tTime := logs.ParseTimestamp(toTime)
		for _, log := range logs {
			if tTime.Equal(log.Timestamp) || tTime.After(log.Timestamp) {
				localLogs = append(localLogs, log)
			}
		}
		filtered := filteredLogs(localLogs, head, tail)
		for index, log := range filtered {
			mainPrint(index, log)
		}
		lengthLogs(filtered)
		return
	}

	if fromTime != "" && toTime != "" {
		localLogs := Logs{}
		tTime := logs.ParseTimestamp(toTime)
		fTime := logs.ParseTimestamp(fromTime)
		for _, log := range logs {
			if (fTime.Equal(log.Timestamp) || fTime.Before(log.Timestamp)) && (tTime.Equal(log.Timestamp) || tTime.After(log.Timestamp)) {
				localLogs = append(localLogs, log)
			}
		}
		filtered := filteredLogs(localLogs, head, tail)
		for index, log := range filtered {
			mainPrint(index, log)
		}
		lengthLogs(filtered)
		return
	}
}
