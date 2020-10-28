package dgraphland

import "time"

type Model struct {
	Uid       string    `json:"uid,omitempty"`
	DType     []string  `json:"dgraph.type,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" dgland:"string"`
	UpdatedAt time.Time `json:"updated_at,omitempty" dgland:"string"`
	DeletedAt time.Time `json:"deleted_at,omitempty" dgland:"string"`
}
