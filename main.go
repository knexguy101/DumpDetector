package dumpDetector

import (
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
	OnDetectedFile func()
	OnErrorCountExceeded func()
}

func checkPanic(options *MonitorOptions, err error, curr int) {
	if options.MaxErrors > 0 && curr >= options.MaxErrors {
		options.OnErrorCountExceeded()
	}
}

func MonitorDumps(options *MonitorOptions) (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		errCount := 0
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					checkPanic(options, err, errCount)
				}
				if event.Op&fsnotify.Write == fsnotify.Write && options.Write {
					if path.Ext(event.Name) != ".DMP" {
						continue
					}
					_ = os.RemoveAll(event.Name)
					options.OnDetectedFile()
					panic("new .dmp being written to")
				} else if event.Op&fsnotify.Create == fsnotify.Create && options.Create {
					if path.Ext(event.Name) != ".DMP" {
						continue
					}
					_ = os.RemoveAll(event.Name)
					options.OnDetectedFile()
					panic("new .dmp being created")
				} else if event.Op&fsnotify.Remove == fsnotify.Remove && options.Remove {
					if path.Ext(event.Name) != ".DMP" {
						continue
					}
					_ = os.RemoveAll(event.Name)
					options.OnDetectedFile()
					panic("new .dmp being removed")
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					checkPanic(options, err, errCount)
				}
				checkPanic(options, err, errCount)
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



