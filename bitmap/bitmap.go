package bitmap

type BitMap []byte

func NewBitmap(maxID int) BitMap {
	size := (maxID + 7) / 8
	return make([]byte, size)
}

func (bitmap BitMap) Add(userID int) {
	index := userID / 8
	offset := uint(userID % 8)
	bitmap[index] |= 1 << offset
}

func (bitmap BitMap) Contains(userID int) bool {
	index := userID / 8
	offset := uint(userID % 8)
	return bitmap[index]&(1<<offset) != 0
}
