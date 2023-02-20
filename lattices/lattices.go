package lattices

/*
	Lattice结构，kvs间传输的数据结构
*/

import (
	config "hybrid_kv_store/config"
	"hybrid_kv_store/util"
	"sync"
)

// Lattice接口，用来实现Merge方法的多态
type Lattice interface {

	// 返回lattice结构包裹的原始数据
	Reveal() config.Log

	// Lattice的Merge操作
	Merge(other Lattice)
}

// 基于VectorClock实现的ValueLattice
type ValueLattice struct {
	value config.Log

	// Vector Clock
	vectorClock sync.Map
}

type MapLattice struct {
	key string
	vl  ValueLattice
}

func (vl ValueLattice) Reveal() config.Log {
	return vl.value
}

func (vl *ValueLattice) Merge(other ValueLattice) {
	// 根据vectorClock决定合并策略
	if util.IsUpper(other.vectorClock, vl.vectorClock) {
		// other >= vl
		// vl.value = other.value
		vl.vectorClock = other.vectorClock
	}
}

func (ml MapLattice) Reveal() config.Log {
	return ml.vl.value
}

func (ml *MapLattice) Merge(other MapLattice) {
	if util.IsUpper(other.vl.vectorClock, ml.vl.vectorClock) {
		// other >= vl
		// ml.vl.value = other.vl.value
		ml.vl.vectorClock = other.vl.vectorClock
	}
}
