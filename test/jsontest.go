package main

 import (
         "fmt"
         "time"
		 "os"
 )
 
 func main(){
	filename := "urltest.go"
	
	file, err := os.Stat(filename)
	
	if err != nil {
        fmt.Println(err)
     }
	 
	modifiedtime := file.ModTime()

	now := time.Now()
	minus30 := time.Minute*time.Duration(-30)
	past := now.Add(minus30)
	
	if modifiedtime.Before(past){
		fmt.Println("too old")
	}else{
		fmt.Println("use it")
	}
 }