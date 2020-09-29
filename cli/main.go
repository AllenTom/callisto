package main

import (
	"fmt"
	"github.com/allentom/callisto"
	"github.com/allentom/callisto/parser"
	"github.com/fsnotify/fsnotify"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var opts struct {
	// build from source path
	Source string `short:"s" long:"source" description:"source file path" required:"true"`
	// auto rebuild when file changed
	Watch bool `short:"w" long:"watch" description:"watch file change" required:"false"`
}

func Generation() error {
	renderEngine := callisto.RenderEngine{}
	err := parser.BuildTree(&renderEngine, opts.Source)
	if err != nil {
		return err
	}
	renderEngine.RenderImage()

	filename := filepath.Base(opts.Source)
	fileExt := filepath.Ext(filename)
	err = renderEngine.SaveAsPNG(strings.ReplaceAll(opts.Source, fileExt, ".png"))
	if err != nil {
		return err
	}
	return nil
}

func RunWatch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					err := Generation()
					if err != nil {
						fmt.Println(err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(opts.Source)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
func main() {
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		log.Fatal(err)
	}

	if opts.Watch {
		RunWatch()
	} else {
		err = Generation()
		if err != nil {
			log.Fatal(err)
		}
	}
}
