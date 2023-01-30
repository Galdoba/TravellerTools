package t5

type star struct {
	starType        string // буква определяющая спектр OBAFGKM
	spectralDecimal int    //число определяющее близость к типу 0123456789
	sizeClass       string //римское число определяющее размер Ia Ib II III IV V VI D (D - определяет белого карлика) BD (определяет коричнегого карлика)
}
