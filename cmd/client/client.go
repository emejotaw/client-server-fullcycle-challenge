package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/emejotaw/client-server-fullcycle-challenge/pkg/constants"
	"github.com/emejotaw/client-server-fullcycle-challenge/pkg/dto"
)

const HTTP_CLIENT_REQUEST_TIMEOUT_MS = 3000

func main() {

	os.Setenv(constants.HTTP_CLIENT_REQUEST_TIMEOUT_MS_NAME, fmt.Sprintf("%d", HTTP_CLIENT_REQUEST_TIMEOUT_MS))

	dollarQuotationDTO, err := GetDollarQuotation()

	if err != nil {
		panic(err)
	}

	currentDollarAmount := dollarQuotationDTO.Bid
	quotation := fmt.Sprintf("Dolar: %v", currentDollarAmount)
	filename := "cotacao.txt"
	WriteToFile(filename, quotation)
	fmt.Printf("Everything is fine, look at the file %s", filename)
}

func WriteToFile(filename string, quotation string) error {

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, os.ModeAppend)

	if err != nil {
		return err
	}

	defer file.Close()

	numberOfBytesWritten, err := file.WriteString(fmt.Sprintf("%s\n", quotation))

	if err != nil {
		return err
	}

	fmt.Printf("%d bytes written to the file\n", numberOfBytesWritten)
	return nil
}

func GetDollarQuotation() (*dto.DollarQuotationDTO, error) {

	httpClientTimeoutMs, err := strconv.Atoi(os.Getenv(constants.HTTP_CLIENT_REQUEST_TIMEOUT_MS_NAME))

	if err != nil {
		return nil, err
	}

	url := "http://localhost:8080/cotacao"
	ctx, cancel := context.WithDeadline(context.TODO(), time.Now().Add(time.Millisecond*time.Duration(httpClientTimeoutMs)))
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		log.Printf("client could not generate http request, error: %v", err)
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Printf("client could not execute http request, error: %v", err)
		return nil, err
	}

	defer response.Body.Close()

	dataBytes, err := io.ReadAll(response.Body)

	if err != nil {
		log.Printf("could not parse response body, error: %v", err)
		return nil, err
	}

	if response.StatusCode == http.StatusOK {

		dollarQuotationDTO := &dto.DollarQuotationDTO{}
		err = json.Unmarshal(dataBytes, dollarQuotationDTO)

		if err != nil {
			log.Printf("could not transform bytes into dollar quotation dto, error: %v", err)
			return nil, err
		}

		log.Printf("Got dollar quotation successfully with response body %v", dollarQuotationDTO)

		return dollarQuotationDTO, nil
	}

	return nil, fmt.Errorf("request fail with status %d, error: %v", response.StatusCode, string(dataBytes))
}
