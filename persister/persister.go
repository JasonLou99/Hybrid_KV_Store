package persister

import (
	"hybrid_kv_store/util"

	"github.com/syndtr/goleveldb/leveldb"
)

type Persister struct {
	// path string
	db *leveldb.DB
}

func (p *Persister) Init(path string) {
	var err error
	//数据存储路径和一些初始文件
	p.db, err = leveldb.OpenFile(path, nil)
	if err != nil {
		util.EPrintf("Open db failed, err: %s", err)
	}
}

func (p *Persister) Put(key string, value string) {
	err := p.db.Put([]byte(key), []byte(value), nil)
	if err != nil {
		util.EPrintf("Put key %s value %s failed, err: %s", key, value, err)
	}
}

func (p *Persister) Get(key string) []byte {
	value, err := p.db.Get([]byte(key), nil)
	if err != nil {
		util.EPrintf("Get key %s failed, err: %s", key, err)
		return nil
	}
	return value
}
