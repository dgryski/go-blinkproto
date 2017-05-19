package blinkproto

const epochOffset = 730425

func ToDays(yy, mm, dd int32) int32 {
	m := (int64(mm) + 9) % 12
	y := int64(yy) - m/10
	edays := 365*y + y/4 - y/100 + y/400 + (m*306+5)/10 + (int64(dd) - 1)
	return int32(edays - epochOffset)
}

func ToDate(days int32) (year, month, day int32) {

	edays := int64(days) + epochOffset

	y := (10000*edays + 14780) / 3652425
	ddd := edays - (365*y + y/4 - y/100 + y/400)
	if ddd < 0 {
		y--
		ddd = edays - (365*y + y/4 - y/100 + y/400)
	}
	mi := (100*ddd + 52) / 3060
	mm := (mi+2)%12 + 1
	y = y + (mi+2)/12
	dd := ddd - (mi*306+5)/10 + 1
	return int32(y), int32(mm), int32(dd)
}
