package main

import (
  "encoding/csv"
  "fmt"
  "net/http"
  //"io/ioutil"
  "github.com/gojektech/heimdall/httpclient"
  "time"
  "os"
  "log"
  "io"

)
func main() {
	
	for testNumber := 0; testNumber<=1000; testNumber+=100{
		result:= retrieveURLS(testNumber)
		start := time.Now()
		for i:=0; i<testNumber; i++ {
			//native("https://"+result[i])
			heimdall("https://"+result[i])
		}
		elapsed := time.Since(start)
	    fmt.Println("Caching,",testNumber, ",",elapsed.Seconds())

	    start = time.Now()
	    for i:=0; i<testNumber; i++ {
	    	native("https://"+result[i])
		}
		elapsed = time.Since(start)
	    fmt.Println("Native,", testNumber, ",",elapsed.Seconds())

	    start = time.Now()
	    for i:=0; i<testNumber; i++ {
			heimdall("https://"+result[i])
		}
		elapsed = time.Since(start)
	    fmt.Println("Heimdall,",testNumber, ",",elapsed.Seconds())
	}
	


	

}
func retrieveURLS(num int) (result []string){
	
	n := num
	i := 1
	// Open the file
	csvfile, err := os.Open("top-1m.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	for i<=n {
		i += 1
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		result = append(result,record[1])
		
	}
	return
}
func native(url string) {

  method := "GET"

  client := &http.Client {
  	Timeout: 1000 * time.Millisecond,
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
  }
  res, err := client.Do(req)
  if res != nil {
  	defer res.Body.Close()
  }
  
  //body, err := ioutil.ReadAll(res.Body)

  //fmt.Println(string(body))
}
func heimdall(url string) {
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	_, err := client.Get(url, nil)
	if err != nil{
		//fmt.Println(err)
	}

	//body, err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	
}

