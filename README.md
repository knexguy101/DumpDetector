# DumpDetector

Uses Windows pathing to detect if new .DMP files are being created using the Task Manager

## Installation

```bash
go get -u github.com/knexguy101/DumpDetector
```

## Usage

```go

package main

import (
	"errors"
	"github.com/knexguy101/DumpDetector"
)

func main(){
	done := make(chan bool)
	watcher, _ := dumpDetector.MonitorDumps(&dumpDetector.MonitorOptions{
		Write: true,
		Create: true,
		Remove: false,
		MaxErrors: 10,
		OtherPaths: []string {
			"add/new/paths/here",
			"we/monitor/them/all",
		},
		OnDetectedFile: func(){
			//panic will stop the file
			panic(errors.New("detected file"))
		},
		OnErrorCountExceeded: func(){
			//panic will stop the program
			panic(errors.New("error count exceeded"))
		},
	})
	defer watcher.Close()
	<-done
}


```

## Contributing
Made for fun, fork and make better

## To Do
want to find registry address for .DMP location instead of trusting no one will change it

## License
[MIT](https://choosealicense.com/licenses/mit/)
