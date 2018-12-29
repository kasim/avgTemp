package avgTemp

import (
	"reflect"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activity-avgTemp")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {
	attributes := context.GetInput("attributes")

	log.Infof("%-v", attributes)
	typeAttr := reflect.TypeOf(attributes).Kind()
	log.Infof("%T: %s\n", typeAttr, typeAttr)

	context.SetOutput("avgTemp", 0.0)

	return true, nil
}
