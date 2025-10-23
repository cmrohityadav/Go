package types

import "time"

type Bhavcopy struct {
	TradDt              time.Time
	BizDt               time.Time
	Sgmt                string
	Src                 string
	FinInstrmTp         string
	FinInstrmId         string
	ISIN                string
	TckrSymb            string
	SctySrs             string
	XpryDt              *time.Time
	FininstrmActlXpryDt *time.Time
	StrkPric            *float64
	OptnTp              string
	FinInstrmNm         string
	OpnPric             *float64
	HghPric             *float64
	LwPric              *float64
	ClsPric             *float64
	LastPric            *float64
	PrvsClsgPric        *float64
	UndrlygPric         *float64
	SttlmPric           *float64
	OpnIntrst           *int64
	ChngInOpnIntrst     *int64
	TtlTradgVol         *int64
	TtlTrfVal           *float64
	TtlNbOfTxsExctd     *int64
	SsnId               string
	NewBrdLotQty        *int64
	Rmks                string
	Rsvd1               string
	Rsvd2               string
	Rsvd3               string
	Rsvd4               string
}

type PriceBand struct {
	Symbol  string
	Series  string
	Name    string
	Band    int
	Remarks string
}