package ginmdw

// var report = TimeReporter(1)

// func main() {
// 	go Read()
// 	go Write()
// 	time.Sleep(1 * time.Minute)
// }

// func Read() {
// 	var totalTime time.Duration
// 	for i := 0; i < 100000; i++ {
// 		t := time.Now()
// 		report.TimeReport.Get("me")
// 		totalTime += time.Since(t)
// 	}

// 	fmt.Println(totalTime)
// }

// func Write() {
// 	var totalTime time.Duration
// 	for i := 0; i < 100000; i++ {
// 		t := time.Now()
// 		report.TimeReport.Set("me", TimeReport{
// 			Hits: 1,
// 		})

// 		totalTime += time.Since(t)
// 	}

// 	fmt.Println(totalTime)
// }
