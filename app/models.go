package app

import (
	"time"

	"github.com/jameycribbs/hare"
)

type Quote struct {
	ID         int       `json:"id"`
	FromCoin   string    `json:"fromCoin"`
	ToCoin     string    `json:"toCoin"`
	Name       string    `json:"name"`
	High       int64     `json:"high"`
	Low        int64     `json:"low"`
	Bid        int64     `json:"bid"`
	Ask        int64     `json:"ask"`
	CreateDate time.Time `json:"createDate"`
}

func (c *Quote) GetID() int {
	return c.ID
}

func (c *Quote) SetID(id int) {
	c.ID = id
}

func (c *Quote) AfterFind(db *hare.Database) error {
	// IMPORTANT!!!  These two lines of code are necessary in your AfterFind
	//               in order for the Find method to work correctly!
	*c = Quote(*c)

	return nil
}
