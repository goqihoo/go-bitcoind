package bitcoind

func (b *Bitcoind) BchGetBlockCount() (count uint64, err error) {
	return b.GetBlockCount()
}

