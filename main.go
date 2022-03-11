package main

import (
	tm "github.com/buger/goterm"
	"math/rand"
	"sync"
	"time"
)

var wgPrinter sync.WaitGroup
var scrMtx = make(chan bool, 1)

func goPrint(row int, text string) {
	scrMtx <- true
	tm.MoveCursor(1, row)
	tm.Println(text)
	tm.Flush()
	time.Sleep(time.Second / 10)
	<-scrMtx
}

func main() {
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Println("\033[?25l") //  hide cursor
	tm.Flush()
	tasks := 5
	for i := 0; i < tasks; i++ {
		wgPrinter.Add(1)
		go func(row int) {
			var n int64
			//var pb Pbar
			//pb.New(50, 50)
			pb := NewPbar(50, 50)
			defer wgPrinter.Done()
			for n = 0; n < pb.Total; n++ {
				goPrint(row, pb.String())
				//goPrint(row, fmt.Sprintf("hi %d!, Current Time: %s", n, time.Now().Format(time.RFC1123)))
				time.Sleep(time.Second/10 + time.Duration(rand.Intn(10000000)))
				pb.Inc(1)
			}
			pb.End()
			goPrint(row, pb.String())
		}(i + 1)
	}
	wgPrinter.Wait()
	goPrint(tasks+1, "Done!\033[?25h") // Done + show cursor
}
