package bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // Преобразование int  в ByteCounter
	return len(p), nil
}
