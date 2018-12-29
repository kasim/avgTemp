package avgTemp

import (
	"fmt"
	//"reflect"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
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
	input_attributes := context.GetInput("attributes")
	operator := context.GetInput("operator").(string)
	threshold := context.GetInput("threshold").(float64)

	attributes := input_attributes.([]interface {})
	avg := 0.0
	sum := 0.0 
	completed := true
	breakpoint := -1
	result := false

	//log.Infof("%-v", attributes)
	//typeAttr := reflect.TypeOf(attributes).Kind()
	//log.Infof("%T: %s", typeAttr, typeAttr)
	//log.Infof("attributes: %v", attributes)
	for n, a := range attributes{
		typedVal, ok := data.GetGlobalScope().GetAttr(a.(string))
		if !ok {
			errorMsg := fmt.Sprintf("Attribute not defined: '%s'", attributes[n])
			log.Error(errorMsg)
			return false, activity.NewError(errorMsg, "", nil)
		}
		//log.Infof("typedVal.Value(): %v", typedVal.Value())
		if (typedVal.Value().(float64) == 0.0) {
			breakpoint = n 
			completed = false
			break;
		}
		sum += typedVal.Value().(float64) 
	}
	if (completed) {
		avg = sum / float64(len(attributes))
		//log.Infof("float64 of attributes: %v", float64(len(attributes)))
		//log.Infof("sum temperature: %v", sum)
		switch operator {
			case "GE":
				if (avg >= threshold) { result = true }
			case "GT":
				if (avg > threshold) { result = true }
			case "EQ":
				if (avg == threshold) { result = true }
			case "LT":
				if (avg < threshold) { result = true }
			case "LE":
				if (avg <= threshold) { result = true }
			default:
				errorMsg := fmt.Sprintf("Unsupported operator: '%s' ", operator)
				log.Error(errorMsg)
				return false, activity.NewError(errorMsg, "", nil)
		}
		log.Infof("avg temperature: %v", avg)
		log.Infof("operator: %v", operator)
		log.Infof("threshold : %v", threshold)
	} else {
		errorMsg := fmt.Sprintf("Not all temperatures are recorded, missed [%s]!", attributes[breakpoint])
		log.Error(errorMsg)
		return false, activity.NewError(errorMsg, "", nil)
	}
	dtype, _ := data.ToTypeEnum("bool")
	data.GetGlobalScope().AddAttr("avgTemp", dtype, result)
	context.SetOutput("result", result)

	return true, nil
}
