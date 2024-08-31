package main

type SortedOrderList interface {
	Slice() []Order
	Len() int
	Get(index int) Order

	RemoveFirst()
	RemoveLast()
	Insert(o Order)
}

type Order struct {
	IsBuy	bool
	Shares	int
	Price	int
}

func MatchOrInsert(asks SortedOrderList, bids SortedOrderList, order Order) {
	n := 0
	filledShares := 0
	if order.IsBuy {
		// While not all shares of the order have been filled and there
		// there are still asks to match
		for filledShares < order.Shares && asks.Len() > 0 && asks.Get(0).Price <= order.Price {
			n++
			// If the remaining number of shares to fill is equal to or larger than
			// the ask's shares, remove the order
			// Else the remaining number of shares to fill is less than the ask's
			// shares, subtract the remaining number of shares to fill from the ask's
			if order.Shares - filledShares >= asks.Get(0).Shares {
				filledShares = filledShares + asks.Get(0).Shares
				asks.RemoveFirst()
			} else {
				firstOrder := asks.Get(0)
				firstOrder.Shares = firstOrder.Shares - (order.Shares - filledShares)
				filledShares = order.Shares
			}
		}
		// If not all shares were filled, add the order with remaining shares to fill to
		// the bids
		if filledShares < order.Shares {
			order.Shares = order.Shares - filledShares
			bids.Insert(order)
		}
	} else {
		for filledShares < order.Shares && bids.Len() > 0 && bids.Get(bids.Len() - 1).Price >= order.Price {
			n++
			if order.Shares - filledShares >= bids.Get(bids.Len() - 1).Shares {
				filledShares = filledShares + bids.Get(bids.Len() - 1).Shares
				bids.RemoveLast()
			} else {
				lastOrder := bids.Get(bids.Len() - 1)
				lastOrder.Shares = lastOrder.Shares - (order.Shares - filledShares)
				filledShares = order.Shares
			}
		}
		if filledShares < order.Shares {
			order.Shares = order.Shares - filledShares
			asks.Insert(order)
		}
	}
}
