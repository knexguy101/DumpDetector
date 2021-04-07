package dumpDetector

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path"
)

type MonitorOptions struct {
	Write bool
	Create bool
	Remove bool
	MaxErrors int
	OtherPaths []string
}

func CheckPanic(err error, max int, curr int) {
	if max > 0 && curr >= max {
		panic(err)
	}
}

func MonitorDumps(options *MonitorOptions) (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		errCount := 0
		fmt.Println("Monitored Started")
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					CheckPanic(err, options.MaxErrors, errCount)
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write && options.Write {
					if path.Ext(event.Name) != ".DMP" {
						continue
					}
					_ = os.RemoveAll(event.Name)
					panic("new .dmp being written to")
				} else if event.Op&fsnotify.Create == fsnotify.Create && options.Create {
					if path.Ext(event.Name) != ".DMP" {
						continue
					}
					_ = os.RemoveAll(event.Name)
					panic("new .dmp being created")
				} else if event.Op&fsnotify.Remove == fsnotify.Remove && options.Remove {
					if path.Ext(event.Name) != ".DMP" {
						continue
					}
					_ = os.RemoveAll(event.Name)
					panic("new .dmp being removed")
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					CheckPanic(err, options.MaxErrors, errCount)
				}
				CheckPanic(err, options.MaxErrors, errCount)
			}
		}
	}()

	ospath, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}

	err = watcher.Add(path.Join(ospath, "Temp"))
	if err != nil {
		panic(err)
	}

	for _, v := range options.OtherPaths {
		err = watcher.Add(v)
		if err != nil {
			panic(err)
		}
	}

	return watcher, nil
}



