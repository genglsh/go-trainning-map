// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/genglsh/go-trainning-map/app/people/service/internal/data/ent/people"
)

// People is the model entity for the People schema.
type People struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Idnum holds the value of the "idnum" field.
	Idnum string `json:"idnum,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Telephone holds the value of the "telephone" field.
	Telephone string `json:"telephone,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*People) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case people.FieldID:
			values[i] = new(sql.NullInt64)
		case people.FieldIdnum, people.FieldName, people.FieldTelephone:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type People", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the People fields.
func (pe *People) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case people.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pe.ID = int(value.Int64)
		case people.FieldIdnum:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field idnum", values[i])
			} else if value.Valid {
				pe.Idnum = value.String
			}
		case people.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pe.Name = value.String
			}
		case people.FieldTelephone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field telephone", values[i])
			} else if value.Valid {
				pe.Telephone = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this People.
// Note that you need to call People.Unwrap() before calling this method if this People
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *People) Update() *PeopleUpdateOne {
	return (&PeopleClient{config: pe.config}).UpdateOne(pe)
}

// Unwrap unwraps the People entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *People) Unwrap() *People {
	tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: People is not a transactional entity")
	}
	pe.config.driver = tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *People) String() string {
	var builder strings.Builder
	builder.WriteString("People(")
	builder.WriteString(fmt.Sprintf("id=%v", pe.ID))
	builder.WriteString(", idnum=")
	builder.WriteString(pe.Idnum)
	builder.WriteString(", name=")
	builder.WriteString(pe.Name)
	builder.WriteString(", telephone=")
	builder.WriteString(pe.Telephone)
	builder.WriteByte(')')
	return builder.String()
}

// Peoples is a parsable slice of People.
type Peoples []*People

func (pe Peoples) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}
