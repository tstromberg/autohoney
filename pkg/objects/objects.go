// objects implements standard object types thrown around by autohoney.
package objects

import (
	"strings"
	"time"
)

// Instance data structure
type Instance struct {
	Id                 int
	Name               string
	Image              string
	Recipes            []string
	CreationTime       time.Time
	StartTime          time.Time
	FirstInterceptTime time.Time
	EndTime            time.Time
}

// Structure defining how an Instance object may look flattened into a database.
type SavedInstance struct {
	Id    int
	Name  string
	Image string
	// Recipes is comma delimited.
	Recipes            string
	CreationTime       time.Time
	StartTime          time.Time
	FirstInterceptTime time.Time
	EndTime            time.Time
}

// Instance() converts a SavedInstance into a normal Instance.
func (s SavedInstance) Instance() autohoney.Instance {
	return autohoney.Instance{
		Id:                 s.Id,
		Name:               s.Name,
		Image:              s.Image,
		Recipes:            strings.Split(s.Recipes, ','),
		CreationTime:       s.CreationTime,
		StartTime:          s.StartTime,
		FirstInterceptTime: s.FirstInterceptTime,
		EndTime:            s.EndTime,
	}
}
