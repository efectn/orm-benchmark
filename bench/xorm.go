package bench

import (
	"github.com/efectn/go-orm-benchmarks/helper"
	"sync"
	"testing"
	xormdb "xorm.io/xorm"
)

type Xorm struct {
	helper.ORMInterface
	mu         sync.Mutex
	conn       *xormdb.Session
	iterations int // Same as b.N, just to customize it
}

func CreateXorm(iterations int) helper.ORMInterface {
	xorm := &Xorm{
		iterations: iterations,
	}

	return xorm
}

func (xorm *Xorm) Name() string {
	return "xorm"
}

func (xorm *Xorm) Init() error {
	engine, err := xormdb.NewEngine("postgres", helper.OrmSource)
	if err != nil {
		return err
	}

	xorm.conn = engine.NewSession()

	return nil
}

func (xorm *Xorm) Close() error {
	return xorm.conn.Close()
}

func (xorm *Xorm) Insert(b *testing.B) {
	m := NewModel5()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.ID = 0
		_, err := xorm.conn.Insert(m)
		if err != nil {
			helper.SetError(b, xorm.Name(), "insert", err.Error())
		}
	}
}

func (xorm *Xorm) InsertMulti(b *testing.B) {
	ms := make([]*Model5, 0, 100)
	for i := 0; i < 100; i++ {
		ms = append(ms, NewModel5())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := xorm.conn.Insert(&ms)
		if err != nil {
			helper.SetError(b, xorm.Name(), "insert_multi", err.Error())
		}
	}
}

func (xorm *Xorm) Update(b *testing.B) {
	m := NewModel5()

	_, err := xorm.conn.Insert(m)
	if err != nil {
		helper.SetError(b, xorm.Name(), "update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := xorm.conn.ID(m.ID).Update(m)
		if err != nil {
			helper.SetError(b, xorm.Name(), "update", err.Error())
		}
	}
}

func (xorm *Xorm) Read(b *testing.B) {
	m := NewModel5()

	_, err := xorm.conn.Insert(m)
	if err != nil {
		helper.SetError(b, xorm.Name(), "read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := xorm.conn.ID(1).NoAutoCondition().Get(m)
		if err != nil {
			helper.SetError(b, xorm.Name(), "read", err.Error())
		}
	}
}

func (xorm *Xorm) ReadSlice(b *testing.B) {
	m := NewModel5()

	for i := 0; i < 100; i++ {
		m.ID = 0
		_, err := xorm.conn.Insert(m)
		if err != nil {
			helper.SetError(b, xorm.Name(), "read_slice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var models []*Model5
		err := xorm.conn.Where("id > 0").Limit(100).Find(&models)
		if err != nil {
			helper.SetError(b, xorm.Name(), "read_slice", err.Error())
		}
	}
}