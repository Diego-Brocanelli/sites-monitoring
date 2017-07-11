# Sites monitoring in golang
Pet project for study of go lang.

# Required
   - [Goland](https://golang.org/) V.1.8.3

# Features
    - Monitoring 
    - Logs
        - Save
        - Show
    - Exit
## Run system

Open your terminal.

Access root your project.

Run command:
```go
go run main.go
```

#### Result
```
------------------------------------------------------------------
Welcome to the site monitoring system, choose the option you want.
------------------------------------------------------------------
Option: 0 - Exit System
Option: 1 - Start Monitoring
Option: 2 - Show logs
```

# Option: 0 - Exit System
#### Result
```
Thank you for use the monitoring system...goodbye :)
```

# Option: 1 - Start monitoring
#### Result
```
------------------------------------------------------------------
Start monitoring
------------------------------------------------------------------
```
After, choose number of monitoring.
```
Choose the amount of analysis you want.
------------------------------------------------------------------
```

```
------------------------------------------------------------------
Preparing analysis...
Runing...

Monitoring: http://www.diegobrocanelli.com.br | Status Code: 200
Monitoring: http://www.google.com.br | Status Code: 200
Monitoring: http://www.github.com | Status Code: 200
------------------------------------------------------------------
Finished monitoring
```

# Option: 2 - Show logs
#### Result
```
------------------------------------------------------------------
2017-07-10 22:30:07 - http://www.diegobrocanelli.com.br - Status Code: 200
2017-07-10 22:30:07 - http://www.google.com.br - Status Code: 200
2017-07-10 22:30:10 - http://www.github.com - Status Code: 200
```
