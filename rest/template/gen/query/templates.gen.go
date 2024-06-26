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

	"hermes/rest/domain"
)

func newTemplate(db *gorm.DB, opts ...gen.DOOption) template {
	_template := template{}

	_template.templateDo.UseDB(db, opts...)
	_template.templateDo.UseModel(&domain.Template{})

	tableName := _template.templateDo.TableName()
	_template.ALL = field.NewAsterisk(tableName)
	_template.ID = field.NewUint(tableName, "id")
	_template.CreatedAt = field.NewTime(tableName, "created_at")
	_template.UpdatedAt = field.NewTime(tableName, "updated_at")
	_template.DeletedAt = field.NewField(tableName, "deleted_at")
	_template.TemplateId = field.NewString(tableName, "template_id")
	_template.Name = field.NewString(tableName, "name")
	_template.Channel = field.NewString(tableName, "channel")
	_template.Content = field.NewString(tableName, "content")
	_template.SendAccount = field.NewString(tableName, "send_account")
	_template.ClientId = field.NewString(tableName, "client_id")
	_template.Enabled = field.NewBool(tableName, "enabled")
	_template.Comment = field.NewString(tableName, "comment")

	_template.fillFieldMap()

	return _template
}

type template struct {
	templateDo

	ALL         field.Asterisk
	ID          field.Uint
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	TemplateId  field.String
	Name        field.String
	Channel     field.String
	Content     field.String
	SendAccount field.String
	ClientId    field.String
	Enabled     field.Bool
	Comment     field.String

	fieldMap map[string]field.Expr
}

func (t template) Table(newTableName string) *template {
	t.templateDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t template) As(alias string) *template {
	t.templateDo.DO = *(t.templateDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *template) updateTableName(table string) *template {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewUint(table, "id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.TemplateId = field.NewString(table, "template_id")
	t.Name = field.NewString(table, "name")
	t.Channel = field.NewString(table, "channel")
	t.Content = field.NewString(table, "content")
	t.SendAccount = field.NewString(table, "send_account")
	t.ClientId = field.NewString(table, "client_id")
	t.Enabled = field.NewBool(table, "enabled")
	t.Comment = field.NewString(table, "comment")

	t.fillFieldMap()

	return t
}

func (t *template) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *template) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 12)
	t.fieldMap["id"] = t.ID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["template_id"] = t.TemplateId
	t.fieldMap["name"] = t.Name
	t.fieldMap["channel"] = t.Channel
	t.fieldMap["content"] = t.Content
	t.fieldMap["send_account"] = t.SendAccount
	t.fieldMap["client_id"] = t.ClientId
	t.fieldMap["enabled"] = t.Enabled
	t.fieldMap["comment"] = t.Comment
}

func (t template) clone(db *gorm.DB) template {
	t.templateDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t template) replaceDB(db *gorm.DB) template {
	t.templateDo.ReplaceDB(db)
	return t
}

type templateDo struct{ gen.DO }

type ITemplateDo interface {
	gen.SubQuery
	Debug() ITemplateDo
	WithContext(ctx context.Context) ITemplateDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITemplateDo
	WriteDB() ITemplateDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITemplateDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITemplateDo
	Not(conds ...gen.Condition) ITemplateDo
	Or(conds ...gen.Condition) ITemplateDo
	Select(conds ...field.Expr) ITemplateDo
	Where(conds ...gen.Condition) ITemplateDo
	Order(conds ...field.Expr) ITemplateDo
	Distinct(cols ...field.Expr) ITemplateDo
	Omit(cols ...field.Expr) ITemplateDo
	Join(table schema.Tabler, on ...field.Expr) ITemplateDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITemplateDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITemplateDo
	Group(cols ...field.Expr) ITemplateDo
	Having(conds ...gen.Condition) ITemplateDo
	Limit(limit int) ITemplateDo
	Offset(offset int) ITemplateDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITemplateDo
	Unscoped() ITemplateDo
	Create(values ...*domain.Template) error
	CreateInBatches(values []*domain.Template, batchSize int) error
	Save(values ...*domain.Template) error
	First() (*domain.Template, error)
	Take() (*domain.Template, error)
	Last() (*domain.Template, error)
	Find() ([]*domain.Template, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.Template, err error)
	FindInBatches(result *[]*domain.Template, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*domain.Template) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITemplateDo
	Assign(attrs ...field.AssignExpr) ITemplateDo
	Joins(fields ...field.RelationField) ITemplateDo
	Preload(fields ...field.RelationField) ITemplateDo
	FirstOrInit() (*domain.Template, error)
	FirstOrCreate() (*domain.Template, error)
	FindByPage(offset int, limit int) (result []*domain.Template, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITemplateDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t templateDo) Debug() ITemplateDo {
	return t.withDO(t.DO.Debug())
}

func (t templateDo) WithContext(ctx context.Context) ITemplateDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t templateDo) ReadDB() ITemplateDo {
	return t.Clauses(dbresolver.Read)
}

func (t templateDo) WriteDB() ITemplateDo {
	return t.Clauses(dbresolver.Write)
}

func (t templateDo) Session(config *gorm.Session) ITemplateDo {
	return t.withDO(t.DO.Session(config))
}

func (t templateDo) Clauses(conds ...clause.Expression) ITemplateDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t templateDo) Returning(value interface{}, columns ...string) ITemplateDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t templateDo) Not(conds ...gen.Condition) ITemplateDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t templateDo) Or(conds ...gen.Condition) ITemplateDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t templateDo) Select(conds ...field.Expr) ITemplateDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t templateDo) Where(conds ...gen.Condition) ITemplateDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t templateDo) Order(conds ...field.Expr) ITemplateDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t templateDo) Distinct(cols ...field.Expr) ITemplateDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t templateDo) Omit(cols ...field.Expr) ITemplateDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t templateDo) Join(table schema.Tabler, on ...field.Expr) ITemplateDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t templateDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITemplateDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t templateDo) RightJoin(table schema.Tabler, on ...field.Expr) ITemplateDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t templateDo) Group(cols ...field.Expr) ITemplateDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t templateDo) Having(conds ...gen.Condition) ITemplateDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t templateDo) Limit(limit int) ITemplateDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t templateDo) Offset(offset int) ITemplateDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t templateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITemplateDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t templateDo) Unscoped() ITemplateDo {
	return t.withDO(t.DO.Unscoped())
}

func (t templateDo) Create(values ...*domain.Template) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t templateDo) CreateInBatches(values []*domain.Template, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t templateDo) Save(values ...*domain.Template) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t templateDo) First() (*domain.Template, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Template), nil
	}
}

func (t templateDo) Take() (*domain.Template, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Template), nil
	}
}

func (t templateDo) Last() (*domain.Template, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Template), nil
	}
}

func (t templateDo) Find() ([]*domain.Template, error) {
	result, err := t.DO.Find()
	return result.([]*domain.Template), err
}

func (t templateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.Template, err error) {
	buf := make([]*domain.Template, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t templateDo) FindInBatches(result *[]*domain.Template, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t templateDo) Attrs(attrs ...field.AssignExpr) ITemplateDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t templateDo) Assign(attrs ...field.AssignExpr) ITemplateDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t templateDo) Joins(fields ...field.RelationField) ITemplateDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t templateDo) Preload(fields ...field.RelationField) ITemplateDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t templateDo) FirstOrInit() (*domain.Template, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Template), nil
	}
}

func (t templateDo) FirstOrCreate() (*domain.Template, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Template), nil
	}
}

func (t templateDo) FindByPage(offset int, limit int) (result []*domain.Template, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t templateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t templateDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t templateDo) Delete(models ...*domain.Template) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *templateDo) withDO(do gen.Dao) *templateDo {
	t.DO = *do.(*gen.DO)
	return t
}
