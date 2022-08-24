package kv

type IMap[Key, Value any] interface {
	Load(key Key) (value Value, exist bool)
	Store(key Key, value Value)
	Delete(key Key)
	Range(f func(Key, Value) bool)
	LoadOrStore(key Key, value Value) (load Value, exist bool)
	LoadAndDelete(key Key) (load Value, exist bool)
}
