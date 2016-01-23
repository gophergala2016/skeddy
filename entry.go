package main

import(
  "github.com/pborman/uuid"
  "encoding/json"
  "net/http"
)

type Entry struct {
  ID          string
  Expression  string
  Endpoint    string
  Payload     string
}

func NewEntry(exp string, ep string, p string) *Entry {
	return &Entry{ID: uuid.New(), Expression: exp, Endpoint: ep, Payload: p}
}

func (e *Entry) Bytes() []byte {
  bt, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return bt
}

func NewEntryFromBytes(b []byte) (*Entry, error) {
  var e Entry
  err := json.Unmarshal(b, &e)
  if err != nil {
    return nil, err
  }
  return &e, nil
}

func NewEntryFromReq(r *http.Request) *Entry {
	return &Entry{ID: r.FormValue("id"), Expression: r.FormValue("expression"), Endpoint: r.FormValue("endpoint"), Payload: r.FormValue("payload")}
}
