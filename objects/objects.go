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

// Filters for an instance query
type InstanceQuery struct {
	Id int
}

// Structure defining how an Instance object may look flattened into a database.
type FlatInstance struct {
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

// Instance() converts a FlatInstance into a normal Instance.
func (s FlatInstance) Instance() Instance {
	return Instance{
		Id:                 s.Id,
		Name:               s.Name,
		Image:              s.Image,
		Recipes:            strings.Split(s.Recipes, ","),
		CreationTime:       s.CreationTime,
		StartTime:          s.StartTime,
		FirstInterceptTime: s.FirstInterceptTime,
		EndTime:            s.EndTime,
	}
}
