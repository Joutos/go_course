package main

import "fmt"

func main() {
	jobs := make(chan int, 200)
	results := make(chan int, 200)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := range 82 {
		jobs <- i
	}

	close(jobs)

	for result := range results {
		fmt.Println(result)
	}
	close(results)
}

func worker(jobs <-chan int, results chan<- int) {
	for numero := range jobs {
		results <- fibonnacci(numero)
	}
}

func fibonnacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonnacci(posicao-2) + fibonnacci(posicao-1)
}
