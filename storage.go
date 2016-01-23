package main

import(
  "github.com/syndtr/goleveldb/leveldb"
)

type Storage struct {
  DB *leveldb.DB
}

func NewStorage(dbname string) (*Storage, error) {
  db, err := leveldb.OpenFile(dbname, nil)
  if err != nil {
    return nil, err
  }
  return &Storage{DB: db}, nil
}

func (s *Storage) Close() {
  s.DB.Close()
}

func (s *Storage) SaveEntry(e *Entry) error {
  err := s.DB.Put([]byte(e.ID), e.Bytes(), nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) AllEntries() []*Entry {
  result := make([]*Entry, 0)
	iter := s.DB.NewIterator(nil, nil)
	for iter.Next() {
		entry, _ := NewEntryFromBytes(iter.Value())
		result = append(result, entry)
	}
	return result
}
