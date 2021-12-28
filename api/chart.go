package api

type Chart struct {
	Ticker    string
	Exchange  string
	Currency  string
	Start     *datetime.Datetime
	End       *datetime.Datetime
	Length    int
	High      decimal.Decimal
	Low       decimal.Decimal
	Open      decimal.Decimal
	Close     decimal.Decimal
	Interval  datetime.Interval
	Bars      []*Bar
	Change    decimal.Decimal
	ChangeVal decimal.Decimal
	Prev      decimal.Decimal
}

type Bar struct {
	Timestamp *datetime.Datetime
	Current   decimal.Decimal
	Y         int
	Char      string
}