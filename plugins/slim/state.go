package slim

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/snple/slim"
	"github.com/snple/types/cache"
	"go.etcd.io/bbolt"
)

type Store struct {
	ss *SlimSlot

	cache *cache.Cache[slim.Object]
	ttl   time.Duration
}

const bucket = "store"

func NewStore(ss *SlimSlot, ttl time.Duration) (*Store, error) {
	err := ss.db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	s := &Store{
		ss:    ss,
		cache: cache.NewCache[slim.Object](nil),
		ttl:   ttl,
	}

	return s, nil
}

func (s *Store) Cache() *Cache {
	return &Cache{store: s}
}

func (s *Store) Module() map[string]slim.Object {
	return map[string]slim.Object{
		"cache": s.Cache(),
		"set":   &slim.UserFunction{Name: "set", Value: s.buildtinSet},
		"get":   &slim.UserFunction{Name: "get", Value: s.buildtinGet},
		"save":  &slim.UserFunction{Name: "save", Value: s.buildtinSave},
		"load":  &slim.UserFunction{Name: "load", Value: s.buildtinLoad},

		"clear_cache": &slim.UserFunction{Name: "clear_cache", Value: s.buildtinClearCache},
		"clear_db":    &slim.UserFunction{Name: "clear_db", Value: s.buildtinClearDB},
	}
}

func (s *Store) AutoGC(duration time.Duration) chan<- struct{} {
	return s.cache.AutoGC(duration)
}

func (s *Store) buildtinSet(args ...slim.Object) (res slim.Object, err error) {
	argsLen := len(args)
	if argsLen != 2 {
		return nil, slim.ErrWrongNumArguments
	}

	key, ok := slim.ToString(args[0])
	if !ok {
		err = slim.ErrInvalidIndexType
		return
	}

	s.cache.Set(key, args[1], s.ttl)
	return slim.TrueValue, nil
}

func (s *Store) buildtinGet(args ...slim.Object) (res slim.Object, err error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, slim.ErrWrongNumArguments
	}

	key, ok := slim.ToString(args[0])
	if !ok {
		err = slim.ErrInvalidIndexType
		return
	}

	if option := s.cache.Get(key); option.IsSome() {
		res = option.Unwrap()
		return
	}

	if argsLen == 2 {
		return args[1], nil
	}

	res = slim.UndefinedValue
	return
}

func (s *Store) buildtinClearCache(args ...slim.Object) (res slim.Object, err error) {
	argsLen := len(args)
	if argsLen != 0 {
		return nil, slim.ErrWrongNumArguments
	}

	s.cache.DeleteAll()

	res = slim.TrueValue
	return
}

func (s *Store) buildtinSave(args ...slim.Object) (res slim.Object, err error) {
	argsLen := len(args)
	if argsLen != 2 {
		return nil, slim.ErrWrongNumArguments
	}

	key, ok := slim.ToString(args[0])
	if !ok {
		err = slim.ErrInvalidIndexType
		return
	}

	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err = enc.Encode([]slim.Object{args[1]})
	if err != nil {
		return nil, err
	}

	// put
	if err := s.ss.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))

		return b.Put([]byte(key), buffer.Bytes())
	}); err != nil {
		return nil, err
	}

	return slim.TrueValue, nil
}

func (s *Store) buildtinLoad(args ...slim.Object) (res slim.Object, err error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, slim.ErrWrongNumArguments
	}

	key, ok := slim.ToString(args[0])
	if !ok {
		err = slim.ErrInvalidIndexType
		return
	}

	// get
	var value []byte
	if err := s.ss.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))

		value = b.Get([]byte(key))
		return nil
	}); err != nil {
		return nil, err
	}

	if len(value) == 0 {
		if argsLen == 2 {
			return args[1], nil
		}

		return slim.UndefinedValue, nil
	}

	buffer := bytes.NewBuffer(value)
	dec := gob.NewDecoder(buffer)
	var decode []slim.Object
	err = dec.Decode(&decode)
	if err != nil {
		return nil, err
	}

	if len(decode) != 1 {
		return slim.UndefinedValue, nil
	}

	return decode[0], nil
}

func (s *Store) buildtinClearDB(args ...slim.Object) (res slim.Object, err error) {
	argsLen := len(args)
	if argsLen != 0 {
		return nil, slim.ErrWrongNumArguments
	}

	// put
	if err := s.ss.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))

		cursor := b.Cursor()
		for key, _ := cursor.First(); key != nil; key, _ = cursor.Next() {
			err := b.Delete(key)
			if err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	res = slim.TrueValue
	return
}

type Cache struct {
	slim.ObjectImpl
	store *Store
}

func (o *Cache) TypeName() string {
	return "cache"
}

func (o *Cache) String() string {
	return "cache"
}

func (o *Cache) Copy() slim.Object {
	return &Cache{store: o.store}
}

func (o *Cache) IsFalsy() bool {
	return o.store.cache.Size() == 0
}

func (o *Cache) Equals(x slim.Object) bool {
	return false
}

// IndexGet returns the value for the given key.
func (o *Cache) IndexGet(index slim.Object) (res slim.Object, err error) {
	// fmt.Printf("get: %v\n", index)
	strIdx, ok := slim.ToString(index)
	if !ok {
		err = slim.ErrInvalidIndexType
		return
	}

	if option := o.store.cache.Get(strIdx); option.IsSome() {
		res = option.Unwrap()
		return
	}

	res = slim.UndefinedValue
	return
}

// IndexSet sets the value for the given key.
func (o *Cache) IndexSet(index, value slim.Object) (err error) {
	// fmt.Printf("set: %v, %v\n", index, value)
	strIdx, ok := slim.ToString(index)
	if !ok {
		err = slim.ErrInvalidIndexType
		return
	}

	o.store.cache.Set(strIdx, value, o.store.ttl)
	return nil
}
