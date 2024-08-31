package main

import "fmt"

func main() {
	asks := SortedOrderList{}
	bids := SortedOrderList{}

	order := Order{ IsBuy: true, Shares: 100, Price: 100 }
	asks, bids = MatchOrInsert(asks, bids, order)

	order = Order{ IsBuy: true, Shares: 50, Price: 110 }
	asks, bids = MatchOrInsert(asks, bids, order)

	order = Order{ IsBuy: true, Shares: 25, Price: 105 }
	asks, bids = MatchOrInsert(asks, bids, order)

	// Orders should be bids and should be sorted
	fmt.Printf("asks=%v bids=%v\n", asks, bids)

	// Should not get matched
	order = Order{ IsBuy: false, Shares: 30, Price: 120 }
	asks, bids = MatchOrInsert(asks, bids, order)

	fmt.Printf("asks=%v bids=%v\n", asks, bids)

	// Should match top 2 bids
	order = Order{ IsBuy: false, Shares: 60, Price: 105 }
	asks, bids = MatchOrInsert(asks, bids, order)

	fmt.Printf("asks=%v bids=%v\n", asks, bids)
}
