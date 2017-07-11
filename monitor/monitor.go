package monitor

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"sites-monitoring/log"
	"strings"
	"time"
)

const lineSeparator = "------------------------------------------------------------------"
const pathFileSites = "./data/site_list.txt"
const delaySeconds = 5

// Start sites monitoring
func Start(quantity int) {
	sitesList := getSiteList()

	for i := 0; i < quantity; i++ {

		fmt.Println(lineSeparator)
		fmt.Println("Preparing analysis...")
		fmt.Println("Runing...")
		fmt.Println("")

		for _, site := range sitesList {

			statusCode := requestMonitor(site)

			fmt.Println("Monitoring:", site, "| Status Code:", statusCode)
			log.Save(site, statusCode)
		}

		time.Sleep(delaySeconds * time.Second)
	}
}

// Send request for site
func requestMonitor(site string) int {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Request error for site:", site)
		os.Exit(-1)
	}

	return response.StatusCode
}

// Get the sites list.
func getSiteList() []string {
	var siteList []string

	file, err := os.Open(pathFileSites)

	if err != nil {
		fmt.Println("open file fail")
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		siteList = append(siteList, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return siteList
}
