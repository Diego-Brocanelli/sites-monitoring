package main

import ( 
	"fmt"
	"net/http"
	"os"
	"bufio"
	"strings"
	"io"
	"io/ioutil"
	"time"
	"strconv"
)

const line_separetor = "------------------------------------------------------------------"
const path_file_sites = "site_list.txt"
const path_file_logs = "logs.txt"
const delay_seconds = 5

func main() {
	showOptions()

	command := getCommand()
	
	switch command{
		case 0:
			fmt.Println("Thank you for use the monitoring system...goodbye :)")
			os.Exit(0)
		break
		case 1:
			fmt.Println()
			fmt.Println(line_separetor)
			fmt.Println("Start monitoring")
			
			quantity := getQuantity()
			startMonitor(quantity)

			fmt.Println(line_separetor)
			fmt.Println("Finished monitoring")
			fmt.Println("")
			os.Exit(0)
		break
		case 2:
		 	fmt.Println("Show logs")
			fmt.Println(line_separetor)
			
			showLogs()

			fmt.Println(line_separetor)
			os.Exit(0)
		break
		default:
			fmt.Println("Command not suported...")
			os.Exit(-1)
		break
	}
}

func showOptions(){
	fmt.Println(line_separetor)
	fmt.Println("Welcome to the site monitoring system, choose the option you want.")
	fmt.Println(line_separetor)
	fmt.Println("Option: 0 - Exit System")
	fmt.Println("Option: 1 - Start Monitoring")
	fmt.Println("Option: 2 - Show logs")
}

func chooeseQtd(){
	fmt.Println(line_separetor)
	fmt.Println("Choose the amount of analysis you want.")
	fmt.Println(line_separetor)
}

func getCommand() int{
	var command int
	fmt.Scan(&command)

	return command
}

func getQuantity() int{

	chooeseQtd()

	var quantity int
	fmt.Scan(&quantity)

	return quantity
}

func startMonitor(quantity int){
	sites_list := getSiteList()

	for i := 0; i < quantity; i++{
		
		fmt.Println(line_separetor)
		fmt.Println("Preparing analysis...")
		fmt.Println("Runing...")
		fmt.Println("")
		
		for _, site := range sites_list {
			
			status_code := requestMonitor(site)
			
			fmt.Println("Monitoring:", site, "| Status Code:", status_code)
			saveLog(site, status_code)
		}

		time.Sleep(delay_seconds * time.Second)
	}
}

func requestMonitor(site string) int {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Request error for site:", site)
		os.Exit(-1)
	}

	return response.StatusCode
}

func getSiteList() []string {
	var site_list []string

	file, err := os.Open(path_file_sites)

	if err != nil{
		fmt.Println("open file fail")
	}

	reader := bufio.NewReader(file)

	for{
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		
		site_list = append(site_list, line)

		if err == io.EOF{
			break
		}
	}

	file.Close()

	return site_list
}

func saveLog(link string, status_code int){
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil{
		fmt.Println(err)
		os.Exit(-1)
	}
	
	file.WriteString(time.Now().Format("2006-01-02 15:04:05") + " - " + link + " - Status Code: " + strconv.FormatInt(int64(status_code), 10) + "\n")

	file.Close()
}

func showLogs(){
	file, err := ioutil.ReadFile(path_file_logs)

	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(string(file))
}