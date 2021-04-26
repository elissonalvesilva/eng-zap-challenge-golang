package coleta

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
var tmp_point_to_path_coleta *string
var tmp_point_to_path_dados *string

func InitColeta() {
	log.Println("routine.coleta - Starting ...")
	dir, _ := os.Getwd()
	path := createFolderIfNotExists(dir + "/" + os.Getenv("PATH_COLETA_FILES"))
	path_dados := createFolderIfNotExists(dir + "/" + os.Getenv("PATH_DADOS"))
	tmp_point_to_path_coleta = &path
	tmp_point_to_path_dados = &path_dados
	createFileIfNotExist(*tmp_point_to_path_dados + os.Getenv("FILENAME_CATALOG"))
}

func Run() {
	log.Println("routine.coleta - Running Coleta ...")
	log.Println("routine.coleta - Download file from: http://grupozap-code-challenge.s3-website-us-east-1.amazonaws.com/sources/source-2.json")
	res, _ := http.Head("http://grupozap-code-challenge.s3-website-us-east-1.amazonaws.com/sources/source-2.json") // 187 MB file of random numbers per line
	maps := res.Header
	length, _ := strconv.Atoi(maps["Content-Length"][0]) // Get the content length from the header request
	limit := 10                                          // 10 Go-routines for the process so each downloads 18.7MB
	len_sub := length / limit                            // Bytes for each Go-routine
	diff := length % limit                               // Get the remaining for the last request
	body := make([]string, 11)                           // Make up a temporary array to hold the data to be written to the file
	for i := 0; i < limit; i++ {
		wg.Add(1)

		min := len_sub * i       // Min range
		max := len_sub * (i + 1) // Max range

		if i == limit-1 {
			max += diff // Add the remaining bytes in the last request
		}

		go func(min int, max int, i int) {
			client := &http.Client{}
			req, _ := http.NewRequest("GET", "http://grupozap-code-challenge.s3-website-us-east-1.amazonaws.com/sources/source-2.json", nil)
			range_header := "bytes=" + strconv.Itoa(min) + "-" + strconv.Itoa(max-1)
			req.Header.Add("Range", range_header)
			resp, _ := client.Do(req)
			defer resp.Body.Close()
			reader, _ := ioutil.ReadAll(resp.Body)
			body[i] = string(reader)
			ioutil.WriteFile(*tmp_point_to_path_coleta+strconv.Itoa(i), []byte(string(body[i])), 0x777)
			wg.Done()
		}(min, max, i)
	}
	wg.Wait()
	log.Println("routine.coleta - Combine files ...")
	log.Println("routine.coleta - Create full.json ...")
	combineFile(limit, *tmp_point_to_path_dados+os.Getenv("FILENAME_CATALOG"))
	log.Println("routine.coleta - Finished coleta ...")

}

func combineFile(workes int, filename string) {

	f, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < workes; i++ {
		content, _ := ioutil.ReadFile(*tmp_point_to_path_coleta + fmt.Sprint(i))
		if _, err := f.Write([]byte(content)); err != nil {
			f.Close()
			log.Fatal(err)
		}
		e := os.Remove(*tmp_point_to_path_coleta + fmt.Sprint(i))
		if e != nil {
			log.Fatal(e)
		}
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func createFileIfNotExist(filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		var file, err = os.Create(filename)
		if err != nil {
			return
		}
		defer file.Close()
	}
}

func createFolderIfNotExists(path string) string {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.Mkdir(path, 0700)
		return path
	}
	return path
}
