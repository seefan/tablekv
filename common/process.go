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
	QGet() (r []byte, err error)
	// Parameters:
	//  - Value
	QSet(value []byte) (err error)
	// Parameters:
	//  - Value
	BatchQSet(value [][]byte) (err error)
	// Parameters:
	//  - Size
	BatchQGet(size int) (r [][]byte, err error)
	// Parameters:
	//  - KeyStart
	//  - KeyEnd
	//  - Limit
	Scan(key_start []byte, key_end []byte, limit int) (ks [][]byte, vs [][]byte, err error)
}

type GetProcessor interface {
	GetProcessor(name string) (Processor, error)
}
