// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/satriahrh/winslow/ent/sprint"
)

// Sprint is the model entity for the Sprint schema.
type Sprint struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Counter holds the value of the "counter" field.
	Counter uint `json:"counter,omitempty"`
	// SprintGoal holds the value of the "sprint_goal" field.
	SprintGoal string `json:"sprint_goal,omitempty"`
	// State holds the value of the "state" field.
	State sprint.State `json:"state,omitempty"`
	// StartedAt holds the value of the "started_at" field.
	StartedAt *time.Time `json:"started_at,omitempty"`
	// FinishedAt holds the value of the "finished_at" field.
	FinishedAt *time.Time `json:"finished_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Sprint) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},
		&sql.NullInt64{},
		&sql.NullString{},
		&sql.NullString{},
		&sql.NullTime{},
		&sql.NullTime{},
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Sprint fields.
func (s *Sprint) assignValues(values ...interface{}) error {
	if m, n := len(values), len(sprint.Columns); m != n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field counter", values[0])
	} else if value.Valid {
		s.Counter = uint(value.Int64)
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field sprint_goal", values[1])
	} else if value.Valid {
		s.SprintGoal = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field state", values[2])
	} else if value.Valid {
		s.State = sprint.State(value.String)
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field started_at", values[3])
	} else if value.Valid {
		s.StartedAt = new(time.Time)
		*s.StartedAt = value.Time
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field finished_at", values[4])
	} else if value.Valid {
		s.FinishedAt = new(time.Time)
		*s.FinishedAt = value.Time
	}
	return nil
}

// QueryProject queries the project edge of the Sprint.
func (s *Sprint) QueryProject() *ProjectQuery {
	return (&SprintClient{s.config}).QueryProject(s)
}

// QueryStories queries the stories edge of the Sprint.
func (s *Sprint) QueryStories() *StoryQuery {
	return (&SprintClient{s.config}).QueryStories(s)
}

// Update returns a builder for updating this Sprint.
// Note that, you need to call Sprint.Unwrap() before calling this method, if this Sprint
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Sprint) Update() *SprintUpdateOne {
	return (&SprintClient{s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Sprint) Unwrap() *Sprint {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Sprint is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Sprint) String() string {
	var builder strings.Builder
	builder.WriteString("Sprint(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", counter=")
	builder.WriteString(fmt.Sprintf("%v", s.Counter))
	builder.WriteString(", sprint_goal=")
	builder.WriteString(s.SprintGoal)
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", s.State))
	if v := s.StartedAt; v != nil {
		builder.WriteString(", started_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := s.FinishedAt; v != nil {
		builder.WriteString(", finished_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Sprints is a parsable slice of Sprint.
type Sprints []*Sprint

func (s Sprints) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
