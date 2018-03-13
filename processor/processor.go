package processor

import "github.com/seefan/tablekv/tables"

type ProcessStart interface {
	Start(tm *tables.TableManager, host string, port int) error
}
