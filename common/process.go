package common

type Processor interface {
	// Parameters:
	//  - Key
	Get(key []byte) (r []byte, err error)
	// Parameters:
	//  - Key
	//  - Value
	Set(key []byte, value []byte) (err error)
	// Parameters:
	//  - Key
	Exists(key []byte) (r bool, err error)
	// Parameters:
	//  - Key
	Delete(key []byte) (err error)
	// Parameters:
	//  - Keys
	//  - Values
	BatchSet(keys [][]byte, values [][]byte) (err error)
}

type GetProcessor interface {
	GetProcessor(name string) (Processor, error)
}
