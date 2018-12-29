package avgTemp

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil{
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	attrs := []string{"t1", "t2", "t3"}
	attributes := make([]interface{}, 3)

	for i, s := range attrs {
		attributes[i] = s
	}

	//setup attrs
	tc.SetInput("attributes", attributes)
	tc.SetInput("operator", "GT")
	tc.SetInput("threshold", 25.0)

	dt, _ := data.ToTypeEnum("float64")
	//data.GetGlobalScope().AddAttr("d1", dt, 25.6)
	data.GetGlobalScope().AddAttr("t1", dt, 25.6)
	data.GetGlobalScope().AddAttr("t2", dt, 24.3)
	data.GetGlobalScope().AddAttr("t3", dt, 27.8)

	act.Eval(tc)

	result := tc.GetOutput("result").(bool)

	assert.Equal(t, result, true);

	typedVal, _ := data.GetGlobalScope().GetAttr("avgTemp")

	assert.Equal(t, result, typedVal.Value())

	//check result attr
}
