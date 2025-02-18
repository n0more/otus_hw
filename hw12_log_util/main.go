package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// logLevelStatistic counts the occurrences of a specified log level in the log data.
func logLevelStatistic(data []byte, logLevel string) int {
	count := 0
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if strings.Contains(line, logLevel) {
			count++
		}
	}

	return count
}

func parseArguments() (filename, logLevel, outputFile string) {
	filename = os.Getenv("LOG_ANALYZER_FILE")

	if os.Getenv("LOG_ANALYZER_LEVEL") == "" {
		logLevel = "info"
	} else {
		logLevel = os.Getenv("LOG_ANALYZER_LEVEL")
	}

	outputFile = os.Getenv("LOG_ANALYZER_OUTPUT")

	flag.StringVar(&filename, "file", filename, "Path to logfile")
	flag.StringVar(&logLevel, "level", logLevel, "LogLevel to analyze logfile")
	flag.StringVar(&outputFile, "output", outputFile, "Path to output file")
	flag.Parse()

	return filename, logLevel, outputFile
}

func writeStatistics(logLevelStatisticResult string, outputFile string) {
	if outputFile == "" {
		fmt.Println(logLevelStatisticResult)
	} else {
		err := os.WriteFile(outputFile, []byte(logLevelStatisticResult), 0o600)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
		fmt.Printf("Statistics written to %s\n", outputFile)
	}
}

func main() {
	os.Setenv("LOG_ANALYZER_FILE", "./main.old.log")

	filename, logLevel, outputFile := parseArguments()

	if filename == "" {
		log.Fatal("Log file path must be specified either via -file flag or LOG_ANALYZER_FILE environment variable.")
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	logLevelStatisticResult := fmt.Sprintf("logLevel %s=%d", logLevel, logLevelStatistic(data, logLevel))
	writeStatistics(logLevelStatisticResult, outputFile)
}
