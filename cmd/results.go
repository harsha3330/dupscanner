package main

import (
	"sync"
)

type Result struct {
	Hash string
	File string
}

func ExtractResults(config *Config) ([]Result, []AppError) {
	files, fileerrs := FileProducer(config.Dir)

	resultsCh := make(chan Result, config.Workers)
	errsCh := make(chan AppError, config.Workers)

	var wg sync.WaitGroup

	for i := 0; i < config.Workers; i++ {
		wg.Add(1)
		go HashWorker(config, files, resultsCh, errsCh, &wg)
	}

	go func() {
		wg.Wait()
		close(resultsCh)
		close(errsCh)
	}()

	var results []Result
	var errorsList []AppError

	for resultsCh != nil || errsCh != nil || fileerrs != nil {
		select {
		case r, ok := <-resultsCh:
			if !ok {
				resultsCh = nil
				continue
			}
			results = append(results, r)

		case err, ok := <-errsCh:
			if !ok {
				errsCh = nil
				continue
			}
			errorsList = append(errorsList, err)

		case err, ok := <-fileerrs:
			if !ok {
				fileerrs = nil
				continue
			}
			errorsList = append(errorsList, err)
		}
	}

	return results, errorsList
}
