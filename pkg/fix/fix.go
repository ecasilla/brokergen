package fixgen

type FixOrder interface {
	SingleOrder() string
	ExecutionReport() string
}

type Fix struct {
	clordid   string
	orderid   string
	orderqty  int
	leavesqty int
	symbol    string
}

var (
	pairDelimiter = "\001"
	kvDelimiter   = "="
	stocks        []Stock
)

func init() {
	byt, err := ioutil.ReadFile("./symbols.json")
	if err != nil {
		panic("Can not read symbols data")
	}

	if err := json.Unmarshal(byt, &stocks); err != nil {
		panic(err)
	}
}

func NewFix() Fix {
	return &Fix{
		clordid:   uuid.v4(),
		orderid:   uuid.v4(),
		orderqty:  rand.Int(10000),
		leavesqty: orderqty,
		// choose a random stock symbol between zero and the symbol slice len
		symbol: randomStockSymbol(),
	}
}

type Stock struct {
	Price   float64 `json:"price"`
	Company string  `json:"company"`
	Symbol  string  `json:"symbol"`
}

func randomStockSymbol() string {
	return stocks[rand.Int(len(s))].Symbol
}
