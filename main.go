package main

import "fmt"

func main() {
	
	messageChannel := make(chan int);


	go func(){

		messageChannel <- 10;

		// message := <-messageChannel;

		// fmt.Println(message); //doesn't works on my laptop, because main returns before Println gets a chance

	}();
		
	message := <-messageChannel; // technically always ready to receive

	fmt.Println(message); //doesn't works on my laptop, because main returns before Println gets a chance
}