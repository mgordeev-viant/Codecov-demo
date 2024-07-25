package bidder

import "testing"

func TestAuction(t *testing.T) {
    auction := NewAuction()
    auction.PlaceBid(Bid{Bidder: "Alice", Amount: 100})
    auction.PlaceBid(Bid{Bidder: "Bob", Amount: 200})

    highestBid, err := auction.GetHighestBid()
    if err != nil {
        t.Error("Expected no error, got", err)
    }

    if highestBid.Bidder != "Bob" || highestBid.Amount != 200 {
        t.Errorf("Expected highest bid to be Bob with amount 200, got %s with amount %d", highestBid.Bidder, highestBid.Amount)
    }
}
