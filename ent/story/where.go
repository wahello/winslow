// Code generated by entc, DO NOT EDIT.

package story

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/satriahrh/winslow/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Story {
	return predicate.Story(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
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
	},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
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
	},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	},
	)
}

// Slug applies equality check predicate on the "slug" field. It's identical to SlugEQ.
func Slug(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSlug), v))
	},
	)
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	},
	)
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	},
	)
}

// SlugEQ applies the EQ predicate on the "slug" field.
func SlugEQ(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSlug), v))
	},
	)
}

// SlugNEQ applies the NEQ predicate on the "slug" field.
func SlugNEQ(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSlug), v))
	},
	)
}

// SlugIn applies the In predicate on the "slug" field.
func SlugIn(vs ...string) predicate.Story {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Story(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSlug), v...))
	},
	)
}

// SlugNotIn applies the NotIn predicate on the "slug" field.
func SlugNotIn(vs ...string) predicate.Story {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Story(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSlug), v...))
	},
	)
}

// SlugGT applies the GT predicate on the "slug" field.
func SlugGT(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSlug), v))
	},
	)
}

// SlugGTE applies the GTE predicate on the "slug" field.
func SlugGTE(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSlug), v))
	},
	)
}

// SlugLT applies the LT predicate on the "slug" field.
func SlugLT(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSlug), v))
	},
	)
}

// SlugLTE applies the LTE predicate on the "slug" field.
func SlugLTE(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSlug), v))
	},
	)
}

// SlugContains applies the Contains predicate on the "slug" field.
func SlugContains(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSlug), v))
	},
	)
}

// SlugHasPrefix applies the HasPrefix predicate on the "slug" field.
func SlugHasPrefix(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSlug), v))
	},
	)
}

// SlugHasSuffix applies the HasSuffix predicate on the "slug" field.
func SlugHasSuffix(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSlug), v))
	},
	)
}

// SlugEqualFold applies the EqualFold predicate on the "slug" field.
func SlugEqualFold(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSlug), v))
	},
	)
}

// SlugContainsFold applies the ContainsFold predicate on the "slug" field.
func SlugContainsFold(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSlug), v))
	},
	)
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	},
	)
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	},
	)
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Story {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Story(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	},
	)
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Story {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Story(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	},
	)
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	},
	)
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	},
	)
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	},
	)
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	},
	)
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	},
	)
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	},
	)
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	},
	)
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	},
	)
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	},
	)
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	},
	)
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	},
	)
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Story {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Story(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	},
	)
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Story {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Story(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	},
	)
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	},
	)
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	},
	)
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	},
	)
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	},
	)
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	},
	)
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	},
	)
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	},
	)
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	},
	)
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	},
	)
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v State) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	},
	)
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v State) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	},
	)
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...State) predicate.Story {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Story(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldState), v...))
	},
	)
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...State) predicate.Story {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Story(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldState), v...))
	},
	)
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ProjectTable, ProjectPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	},
	)
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ProjectTable, ProjectPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	},
	)
}

// HasSprint applies the HasEdge predicate on the "sprint" edge.
func HasSprint() predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SprintTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, SprintTable, SprintPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	},
	)
}

// HasSprintWith applies the HasEdge predicate on the "sprint" edge with a given conditions (other predicates).
func HasSprintWith(preds ...predicate.Sprint) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SprintInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, SprintTable, SprintPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	},
	)
}

// HasActivities applies the HasEdge predicate on the "activities" edge.
func HasActivities() predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActivitiesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ActivitiesTable, ActivitiesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	},
	)
}

// HasActivitiesWith applies the HasEdge predicate on the "activities" edge with a given conditions (other predicates).
func HasActivitiesWith(preds ...predicate.Activity) predicate.Story {
	return predicate.Story(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActivitiesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ActivitiesTable, ActivitiesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Story) predicate.Story {
	return predicate.Story(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Story) predicate.Story {
	return predicate.Story(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Story) predicate.Story {
	return predicate.Story(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
