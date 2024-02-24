package app

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

type QuoteCoinInfo struct {
	Code       string
	CodeIn     string
	Name       string
	High       string
	Low        string
	Bid        string
	Ask        string
	CreateDate string `json:"create_date"`
}

type QuoteResponse struct {
	USDBRL QuoteCoinInfo `json:"USDBRL"`
}

func FetchQuote() (*Quote, error) {
	url := "https://economia.awesomeapi.com.br/last/USD-BRL"
	body, err := doRequest(url)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return parseBody(*body)
}

func doRequest(url string) (*[]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer resp.Body.Close()

	return &body, nil
}

func parseBody(body []byte) (*Quote, error) {
	var quoteResponse QuoteResponse
	err := json.Unmarshal(body, &quoteResponse)
	if err != nil {
		return nil, err
	}

	return buildQuoteFromResponse(quoteResponse)
}

func buildQuoteFromResponse(response QuoteResponse) (*Quote, error) {
	var quote Quote
	var err error

	quote.FromCoin = response.USDBRL.Code
	quote.ToCoin = response.USDBRL.CodeIn
	quote.Name = response.USDBRL.Name

	quote.High, err = parseMoney(response.USDBRL.High)
	if err != nil {
		return nil, err
	}

	quote.Low, err = parseMoney(response.USDBRL.Low)
	if err != nil {
		return nil, err
	}

	quote.Bid, err = parseMoney(response.USDBRL.Bid)
	if err != nil {
		return nil, err
	}

	quote.Ask, err = parseMoney(response.USDBRL.Ask)
	if err != nil {
		return nil, err
	}

	quote.CreateDate, err = parseTime(response.USDBRL.CreateDate)
	if err != nil {
		return nil, err
	}

	return &quote, nil
}

func parseMoney(s string) (int64, error) {
	dec, err := decimal.NewFromString(s)
	if err != nil {
		return 0, err
	}

	three := decimal.NewFromInt(1000)
	parsed := dec.Mul(three).IntPart()
	return parsed, nil
}

func parseTime(s string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, s)
	if err != nil {
		return time.Now(), err
	}

	return t, nil
}
