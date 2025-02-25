// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rodrigo-brito/ninjabot/pkg/ent/order"
	"github.com/rodrigo-brito/ninjabot/pkg/ent/predicate"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// Where adds a new predicate for the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.predicates = append(ou.mutation.predicates, ps...)
	return ou
}

// SetExchangeID sets the "exchange_id" field.
func (ou *OrderUpdate) SetExchangeID(i int64) *OrderUpdate {
	ou.mutation.ResetExchangeID()
	ou.mutation.SetExchangeID(i)
	return ou
}

// AddExchangeID adds i to the "exchange_id" field.
func (ou *OrderUpdate) AddExchangeID(i int64) *OrderUpdate {
	ou.mutation.AddExchangeID(i)
	return ou
}

// SetDate sets the "date" field.
func (ou *OrderUpdate) SetDate(t time.Time) *OrderUpdate {
	ou.mutation.SetDate(t)
	return ou
}

// SetSymbol sets the "symbol" field.
func (ou *OrderUpdate) SetSymbol(s string) *OrderUpdate {
	ou.mutation.SetSymbol(s)
	return ou
}

// SetSide sets the "side" field.
func (ou *OrderUpdate) SetSide(s string) *OrderUpdate {
	ou.mutation.SetSide(s)
	return ou
}

// SetType sets the "type" field.
func (ou *OrderUpdate) SetType(s string) *OrderUpdate {
	ou.mutation.SetType(s)
	return ou
}

// SetStatus sets the "status" field.
func (ou *OrderUpdate) SetStatus(s string) *OrderUpdate {
	ou.mutation.SetStatus(s)
	return ou
}

// SetPrice sets the "price" field.
func (ou *OrderUpdate) SetPrice(f float64) *OrderUpdate {
	ou.mutation.ResetPrice()
	ou.mutation.SetPrice(f)
	return ou
}

// AddPrice adds f to the "price" field.
func (ou *OrderUpdate) AddPrice(f float64) *OrderUpdate {
	ou.mutation.AddPrice(f)
	return ou
}

// SetQuantity sets the "quantity" field.
func (ou *OrderUpdate) SetQuantity(f float64) *OrderUpdate {
	ou.mutation.ResetQuantity()
	ou.mutation.SetQuantity(f)
	return ou
}

// AddQuantity adds f to the "quantity" field.
func (ou *OrderUpdate) AddQuantity(f float64) *OrderUpdate {
	ou.mutation.AddQuantity(f)
	return ou
}

// SetGroupID sets the "group_id" field.
func (ou *OrderUpdate) SetGroupID(i int64) *OrderUpdate {
	ou.mutation.ResetGroupID()
	ou.mutation.SetGroupID(i)
	return ou
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableGroupID(i *int64) *OrderUpdate {
	if i != nil {
		ou.SetGroupID(*i)
	}
	return ou
}

// AddGroupID adds i to the "group_id" field.
func (ou *OrderUpdate) AddGroupID(i int64) *OrderUpdate {
	ou.mutation.AddGroupID(i)
	return ou
}

// ClearGroupID clears the value of the "group_id" field.
func (ou *OrderUpdate) ClearGroupID() *OrderUpdate {
	ou.mutation.ClearGroupID()
	return ou
}

// SetStop sets the "stop" field.
func (ou *OrderUpdate) SetStop(f float64) *OrderUpdate {
	ou.mutation.ResetStop()
	ou.mutation.SetStop(f)
	return ou
}

// SetNillableStop sets the "stop" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableStop(f *float64) *OrderUpdate {
	if f != nil {
		ou.SetStop(*f)
	}
	return ou
}

// AddStop adds f to the "stop" field.
func (ou *OrderUpdate) AddStop(f float64) *OrderUpdate {
	ou.mutation.AddStop(f)
	return ou
}

// ClearStop clears the value of the "stop" field.
func (ou *OrderUpdate) ClearStop() *OrderUpdate {
	ou.mutation.ClearStop()
	return ou
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ou.hooks) == 0 {
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: order.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.ExchangeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: order.FieldExchangeID,
		})
	}
	if value, ok := ou.mutation.AddedExchangeID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: order.FieldExchangeID,
		})
	}
	if value, ok := ou.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldDate,
		})
	}
	if value, ok := ou.mutation.Symbol(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldSymbol,
		})
	}
	if value, ok := ou.mutation.Side(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldSide,
		})
	}
	if value, ok := ou.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldType,
		})
	}
	if value, ok := ou.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldStatus,
		})
	}
	if value, ok := ou.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldPrice,
		})
	}
	if value, ok := ou.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldPrice,
		})
	}
	if value, ok := ou.mutation.Quantity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldQuantity,
		})
	}
	if value, ok := ou.mutation.AddedQuantity(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldQuantity,
		})
	}
	if value, ok := ou.mutation.GroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: order.FieldGroupID,
		})
	}
	if value, ok := ou.mutation.AddedGroupID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: order.FieldGroupID,
		})
	}
	if ou.mutation.GroupIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: order.FieldGroupID,
		})
	}
	if value, ok := ou.mutation.Stop(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldStop,
		})
	}
	if value, ok := ou.mutation.AddedStop(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldStop,
		})
	}
	if ou.mutation.StopCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: order.FieldStop,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderMutation
}

// SetExchangeID sets the "exchange_id" field.
func (ouo *OrderUpdateOne) SetExchangeID(i int64) *OrderUpdateOne {
	ouo.mutation.ResetExchangeID()
	ouo.mutation.SetExchangeID(i)
	return ouo
}

// AddExchangeID adds i to the "exchange_id" field.
func (ouo *OrderUpdateOne) AddExchangeID(i int64) *OrderUpdateOne {
	ouo.mutation.AddExchangeID(i)
	return ouo
}

// SetDate sets the "date" field.
func (ouo *OrderUpdateOne) SetDate(t time.Time) *OrderUpdateOne {
	ouo.mutation.SetDate(t)
	return ouo
}

// SetSymbol sets the "symbol" field.
func (ouo *OrderUpdateOne) SetSymbol(s string) *OrderUpdateOne {
	ouo.mutation.SetSymbol(s)
	return ouo
}

// SetSide sets the "side" field.
func (ouo *OrderUpdateOne) SetSide(s string) *OrderUpdateOne {
	ouo.mutation.SetSide(s)
	return ouo
}

// SetType sets the "type" field.
func (ouo *OrderUpdateOne) SetType(s string) *OrderUpdateOne {
	ouo.mutation.SetType(s)
	return ouo
}

// SetStatus sets the "status" field.
func (ouo *OrderUpdateOne) SetStatus(s string) *OrderUpdateOne {
	ouo.mutation.SetStatus(s)
	return ouo
}

// SetPrice sets the "price" field.
func (ouo *OrderUpdateOne) SetPrice(f float64) *OrderUpdateOne {
	ouo.mutation.ResetPrice()
	ouo.mutation.SetPrice(f)
	return ouo
}

// AddPrice adds f to the "price" field.
func (ouo *OrderUpdateOne) AddPrice(f float64) *OrderUpdateOne {
	ouo.mutation.AddPrice(f)
	return ouo
}

// SetQuantity sets the "quantity" field.
func (ouo *OrderUpdateOne) SetQuantity(f float64) *OrderUpdateOne {
	ouo.mutation.ResetQuantity()
	ouo.mutation.SetQuantity(f)
	return ouo
}

// AddQuantity adds f to the "quantity" field.
func (ouo *OrderUpdateOne) AddQuantity(f float64) *OrderUpdateOne {
	ouo.mutation.AddQuantity(f)
	return ouo
}

// SetGroupID sets the "group_id" field.
func (ouo *OrderUpdateOne) SetGroupID(i int64) *OrderUpdateOne {
	ouo.mutation.ResetGroupID()
	ouo.mutation.SetGroupID(i)
	return ouo
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableGroupID(i *int64) *OrderUpdateOne {
	if i != nil {
		ouo.SetGroupID(*i)
	}
	return ouo
}

// AddGroupID adds i to the "group_id" field.
func (ouo *OrderUpdateOne) AddGroupID(i int64) *OrderUpdateOne {
	ouo.mutation.AddGroupID(i)
	return ouo
}

// ClearGroupID clears the value of the "group_id" field.
func (ouo *OrderUpdateOne) ClearGroupID() *OrderUpdateOne {
	ouo.mutation.ClearGroupID()
	return ouo
}

// SetStop sets the "stop" field.
func (ouo *OrderUpdateOne) SetStop(f float64) *OrderUpdateOne {
	ouo.mutation.ResetStop()
	ouo.mutation.SetStop(f)
	return ouo
}

// SetNillableStop sets the "stop" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableStop(f *float64) *OrderUpdateOne {
	if f != nil {
		ouo.SetStop(*f)
	}
	return ouo
}

// AddStop adds f to the "stop" field.
func (ouo *OrderUpdateOne) AddStop(f float64) *OrderUpdateOne {
	ouo.mutation.AddStop(f)
	return ouo
}

// ClearStop clears the value of the "stop" field.
func (ouo *OrderUpdateOne) ClearStop() *OrderUpdateOne {
	ouo.mutation.ClearStop()
	return ouo
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	var (
		err  error
		node *Order
	)
	if len(ouo.hooks) == 0 {
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			mut = ouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: order.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Order.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.ExchangeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: order.FieldExchangeID,
		})
	}
	if value, ok := ouo.mutation.AddedExchangeID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: order.FieldExchangeID,
		})
	}
	if value, ok := ouo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldDate,
		})
	}
	if value, ok := ouo.mutation.Symbol(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldSymbol,
		})
	}
	if value, ok := ouo.mutation.Side(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldSide,
		})
	}
	if value, ok := ouo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldType,
		})
	}
	if value, ok := ouo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldStatus,
		})
	}
	if value, ok := ouo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldPrice,
		})
	}
	if value, ok := ouo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldPrice,
		})
	}
	if value, ok := ouo.mutation.Quantity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldQuantity,
		})
	}
	if value, ok := ouo.mutation.AddedQuantity(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldQuantity,
		})
	}
	if value, ok := ouo.mutation.GroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: order.FieldGroupID,
		})
	}
	if value, ok := ouo.mutation.AddedGroupID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: order.FieldGroupID,
		})
	}
	if ouo.mutation.GroupIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: order.FieldGroupID,
		})
	}
	if value, ok := ouo.mutation.Stop(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldStop,
		})
	}
	if value, ok := ouo.mutation.AddedStop(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldStop,
		})
	}
	if ouo.mutation.StopCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: order.FieldStop,
		})
	}
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
