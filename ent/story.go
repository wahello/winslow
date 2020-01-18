// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/satriahrh/winslow/ent/story"
)

// Story is the model entity for the Story schema.
type Story struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// State holds the value of the "state" field.
	State story.State `json:"state,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Story) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},
		&sql.NullString{},
		&sql.NullString{},
		&sql.NullString{},
		&sql.NullString{},
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Story fields.
func (s *Story) assignValues(values ...interface{}) error {
	if m, n := len(values), len(story.Columns); m != n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field slug", values[0])
	} else if value.Valid {
		s.Slug = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[1])
	} else if value.Valid {
		s.Name = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field description", values[2])
	} else if value.Valid {
		s.Description = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field state", values[3])
	} else if value.Valid {
		s.State = story.State(value.String)
	}
	return nil
}

// QueryProject queries the project edge of the Story.
func (s *Story) QueryProject() *ProjectQuery {
	return (&StoryClient{s.config}).QueryProject(s)
}

// QuerySprint queries the sprint edge of the Story.
func (s *Story) QuerySprint() *SprintQuery {
	return (&StoryClient{s.config}).QuerySprint(s)
}

// QueryActivities queries the activities edge of the Story.
func (s *Story) QueryActivities() *ActivityQuery {
	return (&StoryClient{s.config}).QueryActivities(s)
}

// Update returns a builder for updating this Story.
// Note that, you need to call Story.Unwrap() before calling this method, if this Story
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Story) Update() *StoryUpdateOne {
	return (&StoryClient{s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Story) Unwrap() *Story {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Story is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Story) String() string {
	var builder strings.Builder
	builder.WriteString("Story(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", slug=")
	builder.WriteString(s.Slug)
	builder.WriteString(", name=")
	builder.WriteString(s.Name)
	builder.WriteString(", description=")
	builder.WriteString(s.Description)
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", s.State))
	builder.WriteByte(')')
	return builder.String()
}

// Stories is a parsable slice of Story.
type Stories []*Story

func (s Stories) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
