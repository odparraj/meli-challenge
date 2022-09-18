package services

import (
	"bufio"
	"fmt"
	m "goravel/app/models"
	s "goravel/contracts/services"
	"goravel/facades"
	"gorm.io/gorm/clause"
	"os"
	"strconv"
	"sync"
	"time"
)

type Job struct {
	id     int
	itemId uint64
	site   string
	line   string
}

type Result struct {
	job  Job
	data *s.DataItem
	err  error
}

type FileImportService struct {
}

var workersNumber int
var jobs chan Job
var results chan Result
var responseErrors []error
var headers = []string{"site", "id"}

func (importer *FileImportService) ProcessFile(filePath string) (string, []error) {

	responseErrors = nil
	
	readFile, err := os.Open("storage/" + filePath)

	if err != nil {
		appendError(err)
		return "fails reading file", responseErrors
	}

	startTime := time.Now()
	fileScanner := bufio.NewScanner(readFile)

	workersNumber = 100
	jobs = make(chan Job, 10)
	results = make(chan Result, 10)

	decoder := facades.Decoders[facades.Config.GetString("decoder.type")]

	if fileScanner.Scan() {
		fileScanner.Text()
	}

	go enqueuJobs(fileScanner, decoder)

	done := make(chan bool)
	go result(done)

	createWorkerPool()
	<-done

	readFile.Close()

	endTime := time.Now()
	diff := endTime.Sub(startTime)

	message := fmt.Sprintf("execution time: %.2f seconds", diff.Seconds())
	return message, responseErrors
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		dataItem, err := facades.MeliClient.GetDataItem(job.line)
		output := Result{job, dataItem, err}
		results <- output
	}
	wg.Done()
}

func createWorkerPool() {
	var wg sync.WaitGroup
	for i := 0; i < workersNumber; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func enqueuJobs(fileScanner *bufio.Scanner, decoder s.Decoder) {
	i := 0
	for fileScanner.Scan() {
		i++
		values, err := decoder.Decode(headers, fileScanner.Text(), facades.Config.GetString("decoder.delimiter"))
		if err != nil {
			appendLineError(i, fmt.Errorf("decoding error [%s]", err.Error()))
			continue
		}
		itemId, err := strconv.ParseUint(values["id"], 10, 64)
		if err != nil {
			appendLineError(i, fmt.Errorf("invalid field [id] with value [%s]", values["id"]))
			continue
		}
		job := Job{
			i,
			itemId,
			values["site"],
			values["site"] + values["id"],
		}
		jobs <- job
	}
	close(jobs)
}

func saveModelFromResult(r Result) error {
	res := facades.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns(m.Fillable),
	}).Create(&m.Item{
		ID:          r.job.itemId,
		Name:        r.data.Name,
		Description: r.data.Description,
		Nickname:    r.data.Nickname,
		StartTime:   r.data.StartTime,
		Price:       r.data.Price,
		Site:        r.data.Site,
	})
	return res.Error
}

func result(done chan bool) {
	for result := range results {
		if result.err != nil {
			appendLineError(result.job.id, result.err)
			continue
		}
		if err := saveModelFromResult(result); err != nil {
			panic(fmt.Sprintf("database error: %s", err.Error()))
		}
	}
	done <- true
}

func appendError(err error) {
	responseErrors = append(responseErrors, err)
}

func appendLineError(line int, err error) {
	responseErrors = append(responseErrors, fmt.Errorf("error proccessing line %d, message: %s", line, err.Error()))
}
