package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func getDescription() string {
	resp, err := http.Get("https://workspace-rho.vercel.app/api/description")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body)
}

func getJobList() string {
	resp, err := http.Get("https://workspace-rho.vercel.app/api/jobs")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body)
}

type Cache struct {
	DescString string
	JobList    string
	IsFilled   bool
}

func (c *Cache) aggregate() {
	wg := sync.WaitGroup{}

	wg.Add(2)

	if !c.IsFilled {

		go func() {
			c.DescString = getDescription()
			wg.Done()
		}()

		go func() {
			c.JobList = getJobList()
			wg.Done()
		}()
		wg.Wait()

		c.IsFilled = true
	}

	fmt.Println(c.DescString)
	fmt.Println(c.JobList)
}

func main() {
	start1 := time.Now()
	fmt.Println("\ndari calculate ", start1)

	chc := Cache{}
	chc.aggregate()

	fmt.Printf("took %v\n\n", time.Since(start1))

	start2 := time.Now()
	fmt.Println("dari calculate ", start2)

	chc.aggregate()

	fmt.Printf("took %v\n\n", time.Since(start2))
}
