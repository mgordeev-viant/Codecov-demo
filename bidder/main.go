package bidder


import (
    "errors"
    "sort"
)

type Bid struct {
    Bidder string
    Amount int
}

type Auction struct {
    Bids []Bid
}

func NewAuction() *Auction {
    return &Auction{Bids: []Bid{}}
}

func (a *Auction) PlaceBid(bid Bid) {
    a.Bids = append(a.Bids, bid)
}

func (a *Auction) GetHighestBid() (Bid, error) {
    if len(a.Bids) == 0 {
        return Bid{}, errors.New("no bids placed")
    }
    sort.SliceStable(a.Bids, func(i, j int) bool {
        return a.Bids[i].Amount > a.Bids[j].Amount
    })
    return a.Bids[0], nil

}
