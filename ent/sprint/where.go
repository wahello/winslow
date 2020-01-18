// Code generated by entc, DO NOT EDIT.

package sprint

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/satriahrh/winslow/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Sprint {
	return predicate.Sprint(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	},
	)
}

// Counter applies equality check predicate on the "counter" field. It's identical to CounterEQ.
func Counter(v uint) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCounter), v))
	},
	)
}

// SprintGoal applies equality check predicate on the "sprint_goal" field. It's identical to SprintGoalEQ.
func SprintGoal(v string) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSprintGoal), v))
	},
	)
}

// StartedAt applies equality check predicate on the "started_at" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartedAt), v))
	},
	)
}

// FinishedAt applies equality check predicate on the "finished_at" field. It's identical to FinishedAtEQ.
func FinishedAt(v time.Time) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFinishedAt), v))
	},
	)
}

// CounterEQ applies the EQ predicate on the "counter" field.
func CounterEQ(v uint) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCounter), v))
	},
	)
}

// CounterNEQ applies the NEQ predicate on the "counter" field.
func CounterNEQ(v uint) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCounter), v))
	},
	)
}

// CounterIn applies the In predicate on the "counter" field.
func CounterIn(vs ...uint) predicate.Sprint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sprint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCounter), v...))
	},
	)
}

// CounterNotIn applies the NotIn predicate on the "counter" field.
func CounterNotIn(vs ...uint) predicate.Sprint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sprint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCounter), v...))
	},
	)
}

// CounterGT applies the GT predicate on the "counter" field.
func CounterGT(v uint) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCounter), v))
	},
	)
}

// CounterGTE applies the GTE predicate on the "counter" field.
func CounterGTE(v uint) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCounter), v))
	},
	)
}

// CounterLT applies the LT predicate on the "counter" field.
func CounterLT(v uint) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCounter), v))
	},
	)
}

// CounterLTE applies the LTE predicate on the "counter" field.
func CounterLTE(v uint) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCounter), v))
	},
	)
}

// SprintGoalEQ applies the EQ predicate on the "sprint_goal" field.
func SprintGoalEQ(v string) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSprintGoal), v))
	},
	)
}

// SprintGoalNEQ applies the NEQ predicate on the "sprint_goal" field.
func SprintGoalNEQ(v string) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSprintGoal), v))
	},
	)
}

// SprintGoalIn applies the In predicate on the "sprint_goal" field.
func SprintGoalIn(vs ...string) predicate.Sprint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sprint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSprintGoal), v...))
	},
	)
}

// SprintGoalNotIn applies the NotIn predicate on the "sprint_goal" field.
func SprintGoalNotIn(vs ...string) predicate.Sprint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sprint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSprintGoal), v...))
	},
	)
}

// SprintGoalGT applies the GT predicate on the "sprint_goal" field.
func SprintGoalGT(v string) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSprintGoal), v))
	},
	)
}

// SprintGoalGTE applies the GTE predicate on the "sprint_goal" field.
func SprintGoalGTE(v string) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSprintGoal), v))
	},
	)
}

// SprintGoalLT applies the LT predicate on the "sprint_goal" field.
func SprintGoalLT(v string) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSprintGoal), v))
	},
	)
}

// SprintGoalLTE applies the LTE predicate on the "sprint_goal" field.
func SprintGoalLTE(v string) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSprintGoal), v))
	},
	)
}

// SprintGoalContains applies the Contains predicate on the "sprint_goal" field.
func SprintGoalContains(v string) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSprintGoal), v))
	},
	)
}

// SprintGoalHasPrefix applies the HasPrefix predicate on the "sprint_goal" field.
func SprintGoalHasPrefix(v string) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSprintGoal), v))
	},
	)
}

// SprintGoalHasSuffix applies the HasSuffix predicate on the "sprint_goal" field.
func SprintGoalHasSuffix(v string) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSprintGoal), v))
	},
	)
}

// SprintGoalIsNil applies the IsNil predicate on the "sprint_goal" field.
func SprintGoalIsNil() predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSprintGoal)))
	},
	)
}

// SprintGoalNotNil applies the NotNil predicate on the "sprint_goal" field.
func SprintGoalNotNil() predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSprintGoal)))
	},
	)
}

// SprintGoalEqualFold applies the EqualFold predicate on the "sprint_goal" field.
func SprintGoalEqualFold(v string) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSprintGoal), v))
	},
	)
}

// SprintGoalContainsFold applies the ContainsFold predicate on the "sprint_goal" field.
func SprintGoalContainsFold(v string) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSprintGoal), v))
	},
	)
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v State) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	},
	)
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v State) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	},
	)
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...State) predicate.Sprint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sprint(func(s *sql.Selector) {
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
func StateNotIn(vs ...State) predicate.Sprint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sprint(func(s *sql.Selector) {
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

// StartedAtEQ applies the EQ predicate on the "started_at" field.
func StartedAtEQ(v time.Time) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartedAt), v))
	},
	)
}

// StartedAtNEQ applies the NEQ predicate on the "started_at" field.
func StartedAtNEQ(v time.Time) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartedAt), v))
	},
	)
}

// StartedAtIn applies the In predicate on the "started_at" field.
func StartedAtIn(vs ...time.Time) predicate.Sprint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sprint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartedAt), v...))
	},
	)
}

// StartedAtNotIn applies the NotIn predicate on the "started_at" field.
func StartedAtNotIn(vs ...time.Time) predicate.Sprint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sprint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartedAt), v...))
	},
	)
}

// StartedAtGT applies the GT predicate on the "started_at" field.
func StartedAtGT(v time.Time) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartedAt), v))
	},
	)
}

// StartedAtGTE applies the GTE predicate on the "started_at" field.
func StartedAtGTE(v time.Time) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartedAt), v))
	},
	)
}

// StartedAtLT applies the LT predicate on the "started_at" field.
func StartedAtLT(v time.Time) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartedAt), v))
	},
	)
}

// StartedAtLTE applies the LTE predicate on the "started_at" field.
func StartedAtLTE(v time.Time) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartedAt), v))
	},
	)
}

// StartedAtIsNil applies the IsNil predicate on the "started_at" field.
func StartedAtIsNil() predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartedAt)))
	},
	)
}

// StartedAtNotNil applies the NotNil predicate on the "started_at" field.
func StartedAtNotNil() predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartedAt)))
	},
	)
}

// FinishedAtEQ applies the EQ predicate on the "finished_at" field.
func FinishedAtEQ(v time.Time) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFinishedAt), v))
	},
	)
}

// FinishedAtNEQ applies the NEQ predicate on the "finished_at" field.
func FinishedAtNEQ(v time.Time) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFinishedAt), v))
	},
	)
}

// FinishedAtIn applies the In predicate on the "finished_at" field.
func FinishedAtIn(vs ...time.Time) predicate.Sprint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sprint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFinishedAt), v...))
	},
	)
}

// FinishedAtNotIn applies the NotIn predicate on the "finished_at" field.
func FinishedAtNotIn(vs ...time.Time) predicate.Sprint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sprint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFinishedAt), v...))
	},
	)
}

// FinishedAtGT applies the GT predicate on the "finished_at" field.
func FinishedAtGT(v time.Time) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFinishedAt), v))
	},
	)
}

// FinishedAtGTE applies the GTE predicate on the "finished_at" field.
func FinishedAtGTE(v time.Time) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFinishedAt), v))
	},
	)
}

// FinishedAtLT applies the LT predicate on the "finished_at" field.
func FinishedAtLT(v time.Time) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFinishedAt), v))
	},
	)
}

// FinishedAtLTE applies the LTE predicate on the "finished_at" field.
func FinishedAtLTE(v time.Time) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFinishedAt), v))
	},
	)
}

// FinishedAtIsNil applies the IsNil predicate on the "finished_at" field.
func FinishedAtIsNil() predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFinishedAt)))
	},
	)
}

// FinishedAtNotNil applies the NotNil predicate on the "finished_at" field.
func FinishedAtNotNil() predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFinishedAt)))
	},
	)
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
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
func HasProjectWith(preds ...predicate.Project) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
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

// HasStories applies the HasEdge predicate on the "stories" edge.
func HasStories() predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StoriesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, StoriesTable, StoriesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	},
	)
}

// HasStoriesWith applies the HasEdge predicate on the "stories" edge with a given conditions (other predicates).
func HasStoriesWith(preds ...predicate.Story) predicate.Sprint {
	return predicate.Sprint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StoriesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, StoriesTable, StoriesPrimaryKey...),
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
func And(predicates ...predicate.Sprint) predicate.Sprint {
	return predicate.Sprint(
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
func Or(predicates ...predicate.Sprint) predicate.Sprint {
	return predicate.Sprint(
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
func Not(p predicate.Sprint) predicate.Sprint {
	return predicate.Sprint(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
