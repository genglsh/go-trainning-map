// Code generated by entc, DO NOT EDIT.

package people

import (
	"entgo.io/ent/dialect/sql"
	"github.com/genglsh/go-trainning-map/app/people/service/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Idnum applies equality check predicate on the "idnum" field. It's identical to IdnumEQ.
func Idnum(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIdnum), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Telephone applies equality check predicate on the "telephone" field. It's identical to TelephoneEQ.
func Telephone(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTelephone), v))
	})
}

// IdnumEQ applies the EQ predicate on the "idnum" field.
func IdnumEQ(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIdnum), v))
	})
}

// IdnumNEQ applies the NEQ predicate on the "idnum" field.
func IdnumNEQ(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIdnum), v))
	})
}

// IdnumIn applies the In predicate on the "idnum" field.
func IdnumIn(vs ...string) predicate.People {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.People(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIdnum), v...))
	})
}

// IdnumNotIn applies the NotIn predicate on the "idnum" field.
func IdnumNotIn(vs ...string) predicate.People {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.People(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIdnum), v...))
	})
}

// IdnumGT applies the GT predicate on the "idnum" field.
func IdnumGT(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIdnum), v))
	})
}

// IdnumGTE applies the GTE predicate on the "idnum" field.
func IdnumGTE(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIdnum), v))
	})
}

// IdnumLT applies the LT predicate on the "idnum" field.
func IdnumLT(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIdnum), v))
	})
}

// IdnumLTE applies the LTE predicate on the "idnum" field.
func IdnumLTE(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIdnum), v))
	})
}

// IdnumContains applies the Contains predicate on the "idnum" field.
func IdnumContains(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIdnum), v))
	})
}

// IdnumHasPrefix applies the HasPrefix predicate on the "idnum" field.
func IdnumHasPrefix(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIdnum), v))
	})
}

// IdnumHasSuffix applies the HasSuffix predicate on the "idnum" field.
func IdnumHasSuffix(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIdnum), v))
	})
}

// IdnumEqualFold applies the EqualFold predicate on the "idnum" field.
func IdnumEqualFold(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIdnum), v))
	})
}

// IdnumContainsFold applies the ContainsFold predicate on the "idnum" field.
func IdnumContainsFold(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIdnum), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.People {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.People(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.People {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.People(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// TelephoneEQ applies the EQ predicate on the "telephone" field.
func TelephoneEQ(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTelephone), v))
	})
}

// TelephoneNEQ applies the NEQ predicate on the "telephone" field.
func TelephoneNEQ(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTelephone), v))
	})
}

// TelephoneIn applies the In predicate on the "telephone" field.
func TelephoneIn(vs ...string) predicate.People {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.People(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTelephone), v...))
	})
}

// TelephoneNotIn applies the NotIn predicate on the "telephone" field.
func TelephoneNotIn(vs ...string) predicate.People {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.People(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTelephone), v...))
	})
}

// TelephoneGT applies the GT predicate on the "telephone" field.
func TelephoneGT(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTelephone), v))
	})
}

// TelephoneGTE applies the GTE predicate on the "telephone" field.
func TelephoneGTE(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTelephone), v))
	})
}

// TelephoneLT applies the LT predicate on the "telephone" field.
func TelephoneLT(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTelephone), v))
	})
}

// TelephoneLTE applies the LTE predicate on the "telephone" field.
func TelephoneLTE(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTelephone), v))
	})
}

// TelephoneContains applies the Contains predicate on the "telephone" field.
func TelephoneContains(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTelephone), v))
	})
}

// TelephoneHasPrefix applies the HasPrefix predicate on the "telephone" field.
func TelephoneHasPrefix(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTelephone), v))
	})
}

// TelephoneHasSuffix applies the HasSuffix predicate on the "telephone" field.
func TelephoneHasSuffix(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTelephone), v))
	})
}

// TelephoneEqualFold applies the EqualFold predicate on the "telephone" field.
func TelephoneEqualFold(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTelephone), v))
	})
}

// TelephoneContainsFold applies the ContainsFold predicate on the "telephone" field.
func TelephoneContainsFold(v string) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTelephone), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.People) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.People) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.People) predicate.People {
	return predicate.People(func(s *sql.Selector) {
		p(s.Not())
	})
}
