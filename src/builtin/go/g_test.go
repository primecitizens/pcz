package stdgo

import "github.com/primecitizens/std/core/alloc"

var _ G = (*GHead)(nil)

func (gp *GHead) Rand32() uint32           { return gp.G().Rand32() }
func (gp *GHead) Rand64() uint64           { return gp.G().Rand64() }
func (gp *GHead) Status() Status           { return gp.G().Status() }
func (gp *GHead) DefaultAlloc() alloc.M    { return gp.G().DefaultAlloc() }
func (gp *GHead) PersistantAlloc() alloc.P { return gp.G().PersistantAlloc() }
