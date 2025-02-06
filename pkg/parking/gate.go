package parking

import "sync"

type Gate struct {
	Lock sync.Mutex
}

func (g *Gate) VehicleEntry(v *Vehicle) {
	g.Lock.Lock()
	defer g.Lock.Unlock()

}
