package app

import (
	"sort"

	"github.com/jameycribbs/hare"
	"github.com/jameycribbs/hare/datastores/disk"
)

const quotesTable = "quotes"

type Hare struct {
	ds *disk.Disk
	db *hare.Database
}

func NewHare() (*Hare, error) {
	ds, err := disk.New("./data", ".json")
	if err != nil {
		return nil, err
	}

	if !ds.TableExists(quotesTable) {
		err = ds.CreateTable(quotesTable)
		if err != nil {
			return nil, err
		}
	}

	db, err := hare.New(ds)
	if err != nil {
		return nil, err
	}

	dbs := Hare{
		ds: ds,
		db: db,
	}

	return &dbs, nil

}

func (hare Hare) InsertQuote(quote *Quote) (int, error) {
	i, err := hare.db.Insert(quotesTable, quote)
	return i, err
}

func (hare Hare) FetchQuote() error {
	return nil
}

func (hare Hare) GetLastQuote() (*Quote, error) {
	ids, err := hare.db.IDs(quotesTable)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return nil, nil
	}

	sort.Ints(ids)
	var quote Quote

	err = hare.db.Find(quotesTable, ids[len(ids)-1], &quote)
	if err != nil {
		return nil, err
	}

	return &quote, nil
}
