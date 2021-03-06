package schema

import "time"

// schema 模块不要引入官方库以外的其它模块或内部模块

// TableLock is a table name in db.
const TableLock = "urbs_lock"

// Lock 详见 ./sql/schema.sql table `urbs_lock`
// 内部锁
type Lock struct {
	ID       int64     `db:"id" goqu:"skipinsert"`
	Name     string    `db:"name"` // varchar(255) 锁键，表内唯一
	ExpireAt time.Time `db:"expire_at"`
}

// TableName retuns table name
func (Lock) TableName() string {
	return "urbs_lock"
}
