// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCollect = "collect"

// Collect mapped from table <collect>
type Collect struct {
	ID  int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UID int32 `gorm:"column:uid" json:"uid"`
	Vid int32 `gorm:"column:vid" json:"vid"`
}

// TableName Collect's table name
func (*Collect) TableName() string {
	return TableNameCollect
}
