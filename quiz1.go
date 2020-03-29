package quiz22

import "fmt"

func Quiz1(){
	var n int
	p:=false
	sk:=false
	f:=false
	for true{
		fmt.Printf("1 – Print Covid19 cases in Pakistan\n2 – Print Covid19 cases in South Korea\n3 – Print Covid19 cases in France\n4 – Print my personalized message to Coronavirus\n0 – Exit\n>> ")
		fmt.Scan(&n)
	if n==1{
		fmt.Println("Covid19 cases in Pakistan: 1500")
		p=true
		
	}else if n==2{
		fmt.Println("Covid19 cases in South Korea: 9332")
		sk=true
		
	}else if n==3{
		fmt.Println("Covid19 cases in France: 25233")
		f=true
		
	}else if n==4{
		fmt.Println("My Personalized message: Stay at HOME")
		
	}else if n==0 {
		if p && f && sk{
			break
			}else {
			fmt.Println("Please view information of all the countries before exiting")
			}
	}else{
		fmt.Println("Wrong input")
		}
	}
		
	}