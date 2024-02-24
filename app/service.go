package app

import (
	"log"
	"time"
)

func UpdateQuotes(hare Hare) {
	lastQuote, err := hare.GetLastQuote()
	if err != nil {
		log.Fatal(err)
		return
	}
	if shouldFetchNew(lastQuote) {
		log.Println("Fetching new Quote")
		_, err := fetchNew(hare)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Println("Ignoring fetch")
	}
}

func shouldFetchNew(lastQuote *Quote) bool {

	if lastQuote == nil {
		return true
	}

	weekday := time.Now().Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		return false
	}

	diff := time.Since(lastQuote.CreateDate)
	return diff.Hours() > 2
}

func fetchNew(hare Hare) (*Quote, error) {
	quote, err := FetchQuote()
	if err != nil {
		return nil, err
	}

	id, err := hare.InsertQuote(quote)
	if err != nil {
		return nil, err
	}

	quote.ID = id
	return quote, nil
}

// notification := toast.Notification{
// 	AppID:   "Cotação",
// 	Title:   "Dólar abaixou para R$ 4,90",
// 	Message: "Some message about how important something is...",
// 	Actions: []toast.Action{
// 		{Type: "protocol", Label: "I'm a button", Arguments: ""},
// 		{Type: "protocol", Label: "I'm another button", Arguments: ""},
// 	},
// }
// err := notification.Push()
// if err != nil {
// 	print(err.Error())
// 	log.Fatalln(err)
// }
