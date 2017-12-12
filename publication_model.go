package m

type Publication struct {
	ID              int64
	Operation       string //Sale, Rent
	Property        *Property
	Price           int64
	Timestamp       string
	Owner           *User
	MinContractTime int
	MaxContractTime int
}
