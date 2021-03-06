// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PeoplesColumns holds the columns for the "peoples" table.
	PeoplesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "idnum", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "telephone", Type: field.TypeString},
	}
	// PeoplesTable holds the schema information for the "peoples" table.
	PeoplesTable = &schema.Table{
		Name:       "peoples",
		Columns:    PeoplesColumns,
		PrimaryKey: []*schema.Column{PeoplesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PeoplesTable,
	}
)

func init() {
}
