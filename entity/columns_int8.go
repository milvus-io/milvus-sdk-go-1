// Code generated by go generate; DO NOT EDIT
// This file is generated by go genrated at 2021-07-08 10:54:30.984859088 &#43;0800 CST m=&#43;0.002222031

//Package entity defines entities used in sdk
package entity 

import "github.com/milvus-io/milvus-sdk-go/internal/proto/schema"

// ColumnInt8 generated columns type for Int8
type ColumnInt8 struct {
	name   string
	values []int8
}

// Name returns column name
func (c *ColumnInt8) Name() string {
	return c.name
}

// Type returns column FieldType
func (c *ColumnInt8) Type() FieldType {
	return FieldTypeInt8
}

// Len returns column values length
func (c *ColumnInt8) Len() int {
	return len(c.values)
}

// FieldData return column data mapped to schema.FieldData
func (c *ColumnInt8) FieldData() *schema.FieldData {
	fd := &schema.FieldData{
		Type: schema.DataType_Int8,
		FieldName: c.name,
	}
	fd.Field = &schema.FieldData_Scalars{
		Scalars: &schema.ScalarField{
			Data: &schema.ScalarField_IntData{
				IntData: &schema.IntArray{
					Data: []int32{},
				},
			},
		},
	}
	return fd
}

// NewColumnInt8 auto generated constructor
func NewColumnInt8(name string, values []int8) *ColumnInt8 {
	return &ColumnInt8 {
		name: name,
		values: values,
	}
}
