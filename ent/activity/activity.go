// Code generated by entc, DO NOT EDIT.

package activity

import (
	"time"

	"github.com/satriahrh/winslow/ent/schema"
)

const (
	// Label holds the string label denoting the activity type in the database.
	Label = "activity"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldActivity holds the string denoting the activity vertex property in the database.
	FieldActivity = "activity"
	// FieldLoggedAt holds the string denoting the logged_at vertex property in the database.
	FieldLoggedAt = "logged_at"

	// Table holds the table name of the activity in the database.
	Table = "activities"
	// StoryTable is the table the holds the story relation/edge. The primary key declared below.
	StoryTable = "story_activities"
	// StoryInverseTable is the table name for the Story entity.
	// It exists in this package in order to avoid circular dependency with the "story" package.
	StoryInverseTable = "stories"
)

// Columns holds all SQL columns are activity fields.
var Columns = []string{
	FieldID,
	FieldActivity,
	FieldLoggedAt,
}

var (
	// StoryPrimaryKey and StoryColumn2 are the table columns denoting the
	// primary key for the story relation (M2M).
	StoryPrimaryKey = []string{"story_id", "activity_id"}
)

var (
	fields = schema.Activity{}.Fields()

	// descLoggedAt is the schema descriptor for logged_at field.
	descLoggedAt = fields[1].Descriptor()
	// DefaultLoggedAt holds the default value on creation for the logged_at field.
	DefaultLoggedAt = descLoggedAt.Default.(func() time.Time)
)
