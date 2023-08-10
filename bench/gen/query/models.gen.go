// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/efectn/go-orm-benchmarks/bench/gen/models"
)

func newModel(db *gorm.DB, opts ...gen.DOOption) model {
	_model := model{}

	_model.modelDo.UseDB(db, opts...)
	_model.modelDo.UseModel(&models.Model{})

	tableName := _model.modelDo.TableName()
	_model.ALL = field.NewAsterisk(tableName)
	_model.Id = field.NewInt(tableName, "id")
	_model.Name = field.NewString(tableName, "name")
	_model.Title = field.NewString(tableName, "title")
	_model.Fax = field.NewString(tableName, "fax")
	_model.Web = field.NewString(tableName, "web")
	_model.Age = field.NewInt(tableName, "age")
	_model.Right = field.NewBool(tableName, "right")
	_model.Counter = field.NewInt64(tableName, "counter")

	_model.fillFieldMap()

	return _model
}

type model struct {
	modelDo modelDo

	ALL     field.Asterisk
	Id      field.Int
	Name    field.String
	Title   field.String
	Fax     field.String
	Web     field.String
	Age     field.Int
	Right   field.Bool
	Counter field.Int64

	fieldMap map[string]field.Expr
}

func (m model) Table(newTableName string) *model {
	m.modelDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m model) As(alias string) *model {
	m.modelDo.DO = *(m.modelDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *model) updateTableName(table string) *model {
	m.ALL = field.NewAsterisk(table)
	m.Id = field.NewInt(table, "id")
	m.Name = field.NewString(table, "name")
	m.Title = field.NewString(table, "title")
	m.Fax = field.NewString(table, "fax")
	m.Web = field.NewString(table, "web")
	m.Age = field.NewInt(table, "age")
	m.Right = field.NewBool(table, "right")
	m.Counter = field.NewInt64(table, "counter")

	m.fillFieldMap()

	return m
}

func (m *model) WithContext(ctx context.Context) IModelDo { return m.modelDo.WithContext(ctx) }

func (m model) TableName() string { return m.modelDo.TableName() }

func (m model) Alias() string { return m.modelDo.Alias() }

func (m model) Columns(cols ...field.Expr) gen.Columns { return m.modelDo.Columns(cols...) }

func (m *model) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *model) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 8)
	m.fieldMap["id"] = m.Id
	m.fieldMap["name"] = m.Name
	m.fieldMap["title"] = m.Title
	m.fieldMap["fax"] = m.Fax
	m.fieldMap["web"] = m.Web
	m.fieldMap["age"] = m.Age
	m.fieldMap["right"] = m.Right
	m.fieldMap["counter"] = m.Counter
}

func (m model) clone(db *gorm.DB) model {
	m.modelDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m model) replaceDB(db *gorm.DB) model {
	m.modelDo.ReplaceDB(db)
	return m
}

type modelDo struct{ gen.DO }

type IModelDo interface {
	gen.SubQuery
	Debug() IModelDo
	WithContext(ctx context.Context) IModelDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IModelDo
	WriteDB() IModelDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IModelDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IModelDo
	Not(conds ...gen.Condition) IModelDo
	Or(conds ...gen.Condition) IModelDo
	Select(conds ...field.Expr) IModelDo
	Where(conds ...gen.Condition) IModelDo
	Order(conds ...field.Expr) IModelDo
	Distinct(cols ...field.Expr) IModelDo
	Omit(cols ...field.Expr) IModelDo
	Join(table schema.Tabler, on ...field.Expr) IModelDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IModelDo
	RightJoin(table schema.Tabler, on ...field.Expr) IModelDo
	Group(cols ...field.Expr) IModelDo
	Having(conds ...gen.Condition) IModelDo
	Limit(limit int) IModelDo
	Offset(offset int) IModelDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IModelDo
	Unscoped() IModelDo
	Create(values ...*models.Model) error
	CreateInBatches(values []*models.Model, batchSize int) error
	Save(values ...*models.Model) error
	First() (*models.Model, error)
	Take() (*models.Model, error)
	Last() (*models.Model, error)
	Find() ([]*models.Model, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Model, err error)
	FindInBatches(result *[]*models.Model, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Model) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IModelDo
	Assign(attrs ...field.AssignExpr) IModelDo
	Joins(fields ...field.RelationField) IModelDo
	Preload(fields ...field.RelationField) IModelDo
	FirstOrInit() (*models.Model, error)
	FirstOrCreate() (*models.Model, error)
	FindByPage(offset int, limit int) (result []*models.Model, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IModelDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m modelDo) Debug() IModelDo {
	return m.withDO(m.DO.Debug())
}

func (m modelDo) WithContext(ctx context.Context) IModelDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m modelDo) ReadDB() IModelDo {
	return m.Clauses(dbresolver.Read)
}

func (m modelDo) WriteDB() IModelDo {
	return m.Clauses(dbresolver.Write)
}

func (m modelDo) Session(config *gorm.Session) IModelDo {
	return m.withDO(m.DO.Session(config))
}

func (m modelDo) Clauses(conds ...clause.Expression) IModelDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m modelDo) Returning(value interface{}, columns ...string) IModelDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m modelDo) Not(conds ...gen.Condition) IModelDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m modelDo) Or(conds ...gen.Condition) IModelDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m modelDo) Select(conds ...field.Expr) IModelDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m modelDo) Where(conds ...gen.Condition) IModelDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m modelDo) Order(conds ...field.Expr) IModelDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m modelDo) Distinct(cols ...field.Expr) IModelDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m modelDo) Omit(cols ...field.Expr) IModelDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m modelDo) Join(table schema.Tabler, on ...field.Expr) IModelDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m modelDo) LeftJoin(table schema.Tabler, on ...field.Expr) IModelDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m modelDo) RightJoin(table schema.Tabler, on ...field.Expr) IModelDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m modelDo) Group(cols ...field.Expr) IModelDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m modelDo) Having(conds ...gen.Condition) IModelDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m modelDo) Limit(limit int) IModelDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m modelDo) Offset(offset int) IModelDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m modelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IModelDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m modelDo) Unscoped() IModelDo {
	return m.withDO(m.DO.Unscoped())
}

func (m modelDo) Create(values ...*models.Model) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m modelDo) CreateInBatches(values []*models.Model, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m modelDo) Save(values ...*models.Model) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m modelDo) First() (*models.Model, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Model), nil
	}
}

func (m modelDo) Take() (*models.Model, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Model), nil
	}
}

func (m modelDo) Last() (*models.Model, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Model), nil
	}
}

func (m modelDo) Find() ([]*models.Model, error) {
	result, err := m.DO.Find()
	return result.([]*models.Model), err
}

func (m modelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Model, err error) {
	buf := make([]*models.Model, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m modelDo) FindInBatches(result *[]*models.Model, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m modelDo) Attrs(attrs ...field.AssignExpr) IModelDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m modelDo) Assign(attrs ...field.AssignExpr) IModelDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m modelDo) Joins(fields ...field.RelationField) IModelDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m modelDo) Preload(fields ...field.RelationField) IModelDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m modelDo) FirstOrInit() (*models.Model, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Model), nil
	}
}

func (m modelDo) FirstOrCreate() (*models.Model, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Model), nil
	}
}

func (m modelDo) FindByPage(offset int, limit int) (result []*models.Model, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m modelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m modelDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m modelDo) Delete(models ...*models.Model) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *modelDo) withDO(do gen.Dao) *modelDo {
	m.DO = *do.(*gen.DO)
	return m
}