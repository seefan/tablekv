package process

import (
	"github.com/seefan/tablekv/common"
	"github.com/seefan/tablekv/tables"
	"sync"
)

type ProcessorManager struct {
	tm   *tables.TableManager
	pm   map[string]common.Processor
	lock sync.RWMutex
}

func NewProcessorManager(tm *tables.TableManager) *ProcessorManager {
	return &ProcessorManager{
		tm: tm,
		pm: make(map[string]common.Processor),
	}
}
func (p *ProcessorManager) GetProcessor(name string) (common.Processor, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if pt, ok := p.pm[name]; ok {
		return pt, nil
	}
	if tb, err := p.tm.GetTable(name); err == nil {
		cp := NewLevelDbProcess(name, tb)//TODO multi-processor
		return cp, nil
	} else {
		return nil, err
	}
}
