package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// People holds the schema definition for the People entity.
type People struct {
	ent.Schema
}

// Fields of the People.
func (People) Fields() []ent.Field {

	return []ent.Field{
		field.String("idnum"),
		field.String("name"),
		field.String("telephone"),
	}
}

// Edges of the People.
func (People) Edges() []ent.Edge {
	return nil
}
