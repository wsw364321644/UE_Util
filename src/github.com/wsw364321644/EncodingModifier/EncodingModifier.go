package main

import (
	"flag"
	"github.com/gogs/chardet"
	"github.com/wsw364321644/go-botil"
	"github.com/wsw364321644/go-botil/log"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)
var path = flag.String("path", "", "path")
var exts = flag.String("exts", ".h .c .cpp .hpp   .cs", "exts")
var extlist []string


func main() {
	flag.Parse()
	extlist=strings.Fields(*exts)
	ModPath(*path)

}

func ModPath(path string){

	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Panicln(err)
	}
	// IsDir is short for fileInfo.Mode().IsDir()
	if fileInfo.IsDir() {
		ModDir(path)
	} else {
		ModFile(path)
	}
}

func ModDir(path string) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		log.Errorln("ReadDir error")
		return
	}

	for _, entry := range entries {
		entryPath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			ModDir(entryPath)
		} else {
			if(botil.StringInSlice(filepath.Ext(entryPath),extlist,true)) {
				ModFile(entryPath)
			}
		}
	}
}
func ModFile(path string) {

	file, err := os.OpenFile(path,os.O_RDWR, 0666)
	if err != nil {
		log.Errorln("path error")
		return
	}

	defer file.Close()
	b,err:=ioutil.ReadAll(file)
	if err != nil {
		log.Errorln("ReadAll error")
		return
	}

	det:=chardet.NewTextDetector()
	result,err:=det.DetectBest(b)
	if err != nil {
		log.Errorln("Detect error")
		return
	}
	encoder,err:=ianaindex.IANA.Encoding(result.Charset)
	if err != nil {
		log.Errorln("Charset error")
		return
	}
	rb,_,err:=transform.Bytes(encoder.NewDecoder(),b)
	if err != nil {
		log.Errorln("transform error")
		return
	}

	err=file.Truncate(0);
	if err != nil {
		log.Errorln(err)
		return
	}

	_,err=file.WriteAt(rb,0);
	if err != nil {
		log.Errorln(err)
		return
	}
}