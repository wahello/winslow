// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/satriahrh/winslow/ent/project"
	"github.com/satriahrh/winslow/ent/sprint"
	"github.com/satriahrh/winslow/ent/story"
)

// SprintCreate is the builder for creating a Sprint entity.
type SprintCreate struct {
	config
	counter     *uint
	sprint_goal *string
	state       *sprint.State
	started_at  *time.Time
	finished_at *time.Time
	project     map[int]struct{}
	stories     map[int]struct{}
}

// SetCounter sets the counter field.
func (sc *SprintCreate) SetCounter(u uint) *SprintCreate {
	sc.counter = &u
	return sc
}

// SetSprintGoal sets the sprint_goal field.
func (sc *SprintCreate) SetSprintGoal(s string) *SprintCreate {
	sc.sprint_goal = &s
	return sc
}

// SetNillableSprintGoal sets the sprint_goal field if the given value is not nil.
func (sc *SprintCreate) SetNillableSprintGoal(s *string) *SprintCreate {
	if s != nil {
		sc.SetSprintGoal(*s)
	}
	return sc
}

// SetState sets the state field.
func (sc *SprintCreate) SetState(s sprint.State) *SprintCreate {
	sc.state = &s
	return sc
}

// SetStartedAt sets the started_at field.
func (sc *SprintCreate) SetStartedAt(t time.Time) *SprintCreate {
	sc.started_at = &t
	return sc
}

// SetNillableStartedAt sets the started_at field if the given value is not nil.
func (sc *SprintCreate) SetNillableStartedAt(t *time.Time) *SprintCreate {
	if t != nil {
		sc.SetStartedAt(*t)
	}
	return sc
}

// SetFinishedAt sets the finished_at field.
func (sc *SprintCreate) SetFinishedAt(t time.Time) *SprintCreate {
	sc.finished_at = &t
	return sc
}

// SetNillableFinishedAt sets the finished_at field if the given value is not nil.
func (sc *SprintCreate) SetNillableFinishedAt(t *time.Time) *SprintCreate {
	if t != nil {
		sc.SetFinishedAt(*t)
	}
	return sc
}

// AddProjectIDs adds the project edge to Project by ids.
func (sc *SprintCreate) AddProjectIDs(ids ...int) *SprintCreate {
	if sc.project == nil {
		sc.project = make(map[int]struct{})
	}
	for i := range ids {
		sc.project[ids[i]] = struct{}{}
	}
	return sc
}

// AddProject adds the project edges to Project.
func (sc *SprintCreate) AddProject(p ...*Project) *SprintCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sc.AddProjectIDs(ids...)
}

// AddStoryIDs adds the stories edge to Story by ids.
func (sc *SprintCreate) AddStoryIDs(ids ...int) *SprintCreate {
	if sc.stories == nil {
		sc.stories = make(map[int]struct{})
	}
	for i := range ids {
		sc.stories[ids[i]] = struct{}{}
	}
	return sc
}

// AddStories adds the stories edges to Story.
func (sc *SprintCreate) AddStories(s ...*Story) *SprintCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddStoryIDs(ids...)
}

// Save creates the Sprint in the database.
func (sc *SprintCreate) Save(ctx context.Context) (*Sprint, error) {
	if sc.counter == nil {
		return nil, errors.New("ent: missing required field \"counter\"")
	}
	if sc.state == nil {
		return nil, errors.New("ent: missing required field \"state\"")
	}
	if err := sprint.StateValidator(*sc.state); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"state\": %v", err)
	}
	return sc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SprintCreate) SaveX(ctx context.Context) *Sprint {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *SprintCreate) sqlSave(ctx context.Context) (*Sprint, error) {
	var (
		s    = &Sprint{config: sc.config}
		spec = &sqlgraph.CreateSpec{
			Table: sprint.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sprint.FieldID,
			},
		}
	)
	if value := sc.counter; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  *value,
			Column: sprint.FieldCounter,
		})
		s.Counter = *value
	}
	if value := sc.sprint_goal; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: sprint.FieldSprintGoal,
		})
		s.SprintGoal = *value
	}
	if value := sc.state; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: sprint.FieldState,
		})
		s.State = *value
	}
	if value := sc.started_at; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: sprint.FieldStartedAt,
		})
		s.StartedAt = value
	}
	if value := sc.finished_at; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: sprint.FieldFinishedAt,
		})
		s.FinishedAt = value
	}
	if nodes := sc.project; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sprint.ProjectTable,
			Columns: sprint.ProjectPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if nodes := sc.stories; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sprint.StoriesTable,
			Columns: sprint.StoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: story.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, sc.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}
