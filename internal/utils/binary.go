package utils

type Binary = uint64

func Toggle(b Binary, index int) Binary {
	return b ^ (1 << index)
}

func Set(b Binary, index int) Binary {
	return b | (1 << index)
}

func Clear(b Binary, index int) Binary {
	return b & ^(1 << index)
}

func IsSet(b Binary, index int) bool {
	return (b & (1 << index)) != 0
}

func ShiftLeft(b Binary, n int) Binary {
	return b << n
}

func ShiftRight(b Binary, n int) Binary {
	return b >> n
}

func SetLSB(b Binary) Binary {
	return b | 1
}

func ClearLSB(b Binary) Binary {
	return b & ^Binary(1)
}
