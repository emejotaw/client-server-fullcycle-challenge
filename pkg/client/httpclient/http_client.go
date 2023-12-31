package httpclient

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/emejotaw/client-server-fullcycle-challenge/pkg/constants"
	"github.com/emejotaw/client-server-fullcycle-challenge/pkg/dto"
)

type httpClient struct {
}

func NewHttpClient() *httpClient {
	return &httpClient{}
}

func (hc *httpClient) GetDollarQuotation() (*dto.DollarQuotationDTO, error) {

	httpRequestTimeoutMs, err := strconv.Atoi(os.Getenv(constants.HTTP_REQUEST_TIMEOUT_MS_NAME))

	if err != nil {
		return nil, err
	}

	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	ctx, cancel := context.WithDeadline(context.TODO(), time.Now().Add(time.Millisecond*time.Duration(httpRequestTimeoutMs)))
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		log.Printf("could not generate http request, error: %v", err)
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Printf("could not execute http request, error: %v", err)
		return nil, err
	}

	defer response.Body.Close()
	dataBytes, err := io.ReadAll(response.Body)

	if err != nil {
		log.Printf("could not parse response body, error: %v", err)
		return nil, err
	}

	dollarQuoteDTO := &dto.DollarQuotationDTO{}
	err = json.Unmarshal(dataBytes, dollarQuoteDTO)

	if err != nil {
		log.Printf("could not transform bytes into struct, error: %v", err)
		return nil, err
	}

	return dollarQuoteDTO, nil
}
