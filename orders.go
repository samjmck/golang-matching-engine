package main

import "slices"

type Order struct {
	IsBuy	bool
	Shares	int
	Price	int
}

type SortedOrderList []Order

func (l SortedOrderList) RemoveFirst() SortedOrderList {
	return l[1:]
}

func (l SortedOrderList) RemoveLast() SortedOrderList {
	return l[:len(l) - 1]
}

func (l SortedOrderList) Insert(o Order) SortedOrderList {
	i := 0
	for i < len(l) && o.Price > l[i].Price {
		i++
	}
	return slices.Insert(l, i, o)
}

// Returns the new asks and new bids
func MatchOrInsert(asks SortedOrderList, bids SortedOrderList, order Order) (SortedOrderList, SortedOrderList) {
	n := 0
	filledShares := 0
	if order.IsBuy {
		// While not all shares of the order have been filled and there
		// there are still asks to match
		for filledShares < order.Shares && len(asks) > 0 && asks[0].Price <= order.Price {
			n++
			// If the remaining number of shares to fill is equal to or larger than
			// the ask's shares, remove the order
			// Else the remaining number of shares to fill is less than the ask's
			// shares, subtract the remaining number of shares to fill from the ask's
			if order.Shares - filledShares >= asks[0].Shares {
				filledShares = filledShares + asks[0].Shares
				asks = asks.RemoveFirst()
			} else {
				asks[0].Shares = asks[0].Shares - (order.Shares - filledShares)
				filledShares = order.Shares
			}
		}
		// If not all shares were filled, add the order with remaining shares to fill to
		// the bids
		if filledShares < order.Shares {
			order.Shares = order.Shares - filledShares
			bids = bids.Insert(order)
		}
	} else {
		for filledShares < order.Shares && len(bids) > 0 && bids[len(bids) - 1].Price >= order.Price {
			n++
			if order.Shares - filledShares >= bids[len(bids) - 1].Shares {
				filledShares = filledShares + bids[len(bids) - 1].Shares
				bids = bids.RemoveLast()
			} else {
				bids[len(bids) - 1].Shares = bids[len(bids) - 1].Shares - (order.Shares - filledShares)
				filledShares = order.Shares
			}
		}
		if filledShares < order.Shares {
			order.Shares = order.Shares - filledShares
			asks = asks.Insert(order)
		}
	}
	return asks, bids
}
