package kv

type Kv[Key, Value any] struct {
	sub map[uint]*entity[Key, Value]
	len uint
}

func (k *Kv[Key, Value]) Load(key Key) (value Value, exist bool) {
	index := hash(key) % k.len
	return k.sub[index].query(key)
}

func (k *Kv[Key, Value]) Store(key Key, value Value) {
	index := hash(key) % k.len
	if k.sub[index].has(key) {
		k.sub[index].update(key, value)
	} else {
		k.sub[index].create(key, value)
	}
}

func (k *Kv[Key, Value]) Delete(key Key) {
	index := hash(key) % k.len
	if k.sub[index].has(key) {
		k.sub[index].delete(key)
	}
}

func (k *Kv[Key, Value]) Range(f func(Key, Value) bool) {
	for _, s := range k.sub {
		if !s.r(f) {
			return
		}
	}
}

func (k *Kv[Key, Value]) LoadOrStore(key Key, value Value) (load Value, exist bool) {
	index := hash(key) % k.len
	load, exist = k.sub[index].query(key)
	if !exist {
		k.sub[index].create(key, value)
	}
	return load, exist
}

func (k *Kv[Key, Value]) LoadAndDelete(key Key) (load Value, exist bool) {
	index := hash(key) % k.len
	load, exist = k.sub[index].query(key)
	if exist {
		k.sub[index].delete(key)
	}
	return load, exist
}

func NewKv[Key, Value any](length ...uint) IMap[Key, Value] {
	kv := &Kv[Key, Value]{
		sub: map[uint]*entity[Key, Value]{},
	}

	if len(length) == 0 {
		kv.len = 1007
	} else {
		kv.len = length[0]
	}
	for i := uint(0); i < kv.len; i++ {
		kv.sub[i] = newEntity[Key, Value]()
	}

	return kv
}
