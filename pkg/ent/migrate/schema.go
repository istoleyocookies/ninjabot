// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "exchange_id", Type: field.TypeInt64},
		{Name: "date", Type: field.TypeTime},
		{Name: "symbol", Type: field.TypeString},
		{Name: "side", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "status", Type: field.TypeString},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "quantity", Type: field.TypeFloat64},
		{Name: "group_id", Type: field.TypeInt64, Nullable: true},
		{Name: "stop", Type: field.TypeFloat64, Nullable: true},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:        "orders",
		Columns:     OrdersColumns,
		PrimaryKey:  []*schema.Column{OrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OrdersTable,
	}
)

func init() {
}
