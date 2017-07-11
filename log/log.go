package log

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

const pathFileLogs = "./data/logs.txt"

// Save logs to file
func Save(link string, statusCode int) {
	file, err := os.OpenFile(pathFileLogs, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	file.WriteString(time.Now().Format("2006-01-02 15:04:05") + " - " + link + " - Status Code: " + strconv.FormatInt(int64(statusCode), 10) + "\n")

	file.Close()
}

// Show logs
func Show() {
	file, err := ioutil.ReadFile(pathFileLogs)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
