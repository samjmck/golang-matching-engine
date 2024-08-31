package main

import "fmt"

func Test(asks SortedOrderList, bids SortedOrderList) {
	order := Order{ IsBuy: true, Shares: 100, Price: 100 }
	MatchOrInsert(asks, bids, order)

	order = Order{ IsBuy: true, Shares: 50, Price: 110 }
	MatchOrInsert(asks, bids, order)

	order = Order{ IsBuy: true, Shares: 25, Price: 105 }
	MatchOrInsert(asks, bids, order)

	// Orders should be bids and should be sorted
	fmt.Printf("asks=%v bids=%v\n", asks.Slice(), bids.Slice())

	// Should not get matched
	order = Order{ IsBuy: false, Shares: 30, Price: 120 }
	MatchOrInsert(asks, bids, order)

	fmt.Printf("asks=%v bids=%v\n", asks.Slice(), bids.Slice())

	// Should match top 2 bids
	order = Order{ IsBuy: false, Shares: 60, Price: 105 }
	MatchOrInsert(asks, bids, order)

	fmt.Printf("asks=%v bids=%v\n", asks.Slice(), bids.Slice())
}

func main() {
	var asks SortedOrderList
	var bids SortedOrderList

	fmt.Println("=== Testing slice order list implementation ===")
	asks = NewSliceSortedOrderList()
	bids = NewSliceSortedOrderList()
	Test(asks, bids)

	fmt.Println("=== Testing linked list order list implementation ===")
	asks = NewOrderedLinkedOrderList()
	bids = NewOrderedLinkedOrderList()
	Test(asks, bids)
}
