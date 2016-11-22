package validator_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/go-jstmpl/go-jsvalidator"
	"github.com/gocraft/dbr"
)

func TestNewRequiredValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition validator.RequiredValidatorDefinition
		Error      error
	}
	cases := []Case{
		{
			Message:    "single element",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"foo"}},
			Error:      nil,
		},
		{
			Message:    "multi elements",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"foo", "bar"}},
			Error:      nil,
		},
		{
			Message:    "empty slice",
			Definition: validator.RequiredValidatorDefinition{Required: []string{}},
			Error:      validator.RequiredDefinitionEmptyError,
		},
		{
			Message:    "duplicated elements",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"foo", "foo"}},
			Error:      validator.RequiredDefinitionDuplicationError,
		},
		{
			Message:    "duplicated elements at first and second",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"foo", "foo", "bar"}},
			Error:      validator.RequiredDefinitionDuplicationError,
		},
		{
			Message:    "duplicated elements at first and end",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"foo", "bar", "foo"}},
			Error:      validator.RequiredDefinitionDuplicationError,
		},
		{
			Message:    "duplicated elements at second end end",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"bar", "foo", "foo"}},
			Error:      validator.RequiredDefinitionDuplicationError,
		},
		{
			Message:    "duplicated all elements",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"foo", "foo", "foo"}},
			Error:      validator.RequiredDefinitionDuplicationError,
		},
	}
	for _, c := range cases {
		_, err := validator.NewRequiredValidator(c.Definition)
		if err != c.Error {
			t.Errorf("Test with %s: fail to NewPatternValidator with error %v", c.Message, err)
		}
	}
}

func TestValidateOfRequiredValidator(t *testing.T) {
	type Types struct {
		NullableStringValue dbr.NullString
		NullableIntValue    dbr.NullInt64
		NullableFloatValue  dbr.NullFloat64
		NullableBoolValue   dbr.NullBool
		NullableTimeValue   dbr.NullTime
		StringValue         string
		IntValue            int
		FloatValue          float64
		BoolValue           bool
		TimeValue           time.Time
	}

	type Input struct {
		Types      Types
		Definition validator.RequiredValidatorDefinition
	}

	type Case struct {
		Message  string
		Input    Input
		Expected error
	}

	cases := []Case{
		{
			Message: "complete struct against required",
			Input: Input{
				Types: Types{
					NullableStringValue: dbr.NewNullString("value"),
					NullableIntValue:    dbr.NewNullInt64(1),
					NullableFloatValue:  dbr.NewNullFloat64(1.1),
					NullableBoolValue:   dbr.NewNullBool(true),
					NullableTimeValue:   dbr.NewNullTime("2009-11-10 23:00:00"),
					StringValue:         "value",
					IntValue:            1,
					FloatValue:          1.1,
					BoolValue:           true,
					TimeValue:           time.Date(2009, 11, 10, 23, 00, 0, 0, time.UTC),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{
						"NullableStringValue",
						"NullableIntValue",
						"NullableFloatValue",
						"NullableBoolValue",
						"NullableTimeValue",
						"StringValue",
						"IntValue",
						"FloatValue",
						"BoolValue",
						"TimeValue",
					},
				},
			},
			Expected: nil,
		},
		{
			Message: "NullableStringValue is missing",
			Input: Input{
				Types: Types{
					NullableStringValue: dbr.NewNullString(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableStringValue"},
				},
			},
			Expected: &validator.RequiredValidationError{
				Input: Types{
					NullableStringValue: dbr.NewNullString(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableStringValue"},
				},
			},
		},
		{
			Message: "NullableIntValue is missing",
			Input: Input{
				Types: Types{
					NullableIntValue: dbr.NewNullInt64(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableIntValue"},
				},
			},
			Expected: &validator.RequiredValidationError{
				Input: Types{
					NullableIntValue: dbr.NewNullInt64(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableIntValue"},
				},
			},
		},
		{
			Message: "NullableFloatValue is missing",
			Input: Input{
				Types: Types{
					NullableFloatValue: dbr.NewNullFloat64(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableFloatValue"},
				},
			},
			Expected: &validator.RequiredValidationError{
				Input: Types{
					NullableFloatValue: dbr.NewNullFloat64(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableFloatValue"},
				},
			},
		},
		{
			Message: "NullableBoolValue is missing",
			Input: Input{
				Types: Types{
					NullableBoolValue: dbr.NewNullBool(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableBoolValue"},
				},
			},
			Expected: &validator.RequiredValidationError{
				Input: Types{
					NullableBoolValue: dbr.NewNullBool(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableBoolValue"},
				},
			},
		},
		{
			Message: "NullableTimeValue is missing",
			Input: Input{
				Types: Types{
					NullableTimeValue: dbr.NewNullTime(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableTimeValue"},
				},
			},
			Expected: &validator.RequiredValidationError{
				Input: Types{
					NullableTimeValue: dbr.NewNullTime(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableTimeValue"},
				},
			},
		},

		{
			Message: "NullableTimeValue is missing",
			Input: Input{
				Types: Types{
					NullableTimeValue: dbr.NewNullTime(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableTimeValue"},
				},
			},
			Expected: &validator.RequiredValidationError{
				Input: Types{
					NullableTimeValue: dbr.NewNullTime(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableTimeValue"},
				},
			},
		},
	}

	for _, c := range cases {
		definition := c.Input.Definition
		va, err := validator.NewRequiredValidator(definition)
		if err != nil {
			t.Errorf("test with %s: fail to create new required validator: %s", c.Message, err)
			continue
		}

		err = va.Validate(c.Input.Types)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("test with %s: expected %+v, but actual %+v", c.Message, c.Expected, err)
		}
	}
}

func TestConvertToConcreteValue(t *testing.T) {
	// Output expected is always Kind of non Ptr
	type Case struct {
		Message string
		Input   reflect.Value
	}
	var (
		stringValue = "string"
		intValue    = 1
		floatValue  = 1.1
		boolValue   = true
		structValue = time.Now()
	)

	cases := []Case{
		{
			Message: "kind of string",
			Input:   reflect.ValueOf(stringValue),
		},
		{
			Message: "kind of int",
			Input:   reflect.ValueOf(intValue),
		},
		{
			Message: "kind of float",
			Input:   reflect.ValueOf(floatValue),
		},
		{
			Message: "kind of bool",
			Input:   reflect.ValueOf(boolValue),
		},
		{
			Message: "kind of struct",
			Input:   reflect.ValueOf(structValue),
		},
		{
			Message: "kind of pointer of string",
			Input:   reflect.ValueOf(&stringValue),
		},
		{
			Message: "kind of pointer of int",
			Input:   reflect.ValueOf(&intValue),
		},
		{
			Message: "kind of pointer of float",
			Input:   reflect.ValueOf(&floatValue),
		},
		{
			Message: "kind of pointer of bool",
			Input:   reflect.ValueOf(&boolValue),
		},
		{
			Message: "kind of pointer of struct",
			Input:   reflect.ValueOf(&structValue),
		},
	}

	for _, c := range cases {
		v, ok := validator.ConvertToConcreteValue(c.Input)
		if !ok {
			t.Errorf("test with %s: fail to convert to concrete value %v", c.Message, c.Input)
		}
		if v.Kind() == reflect.Ptr {
			t.Errorf("test with  %s: expected non Ptr but not", c.Message)
		}
	}
}

func TestGetFieldByName(t *testing.T) {
	type Sample struct {
		Hoge string
		Foo  string
		Bar  string
	}

	v, ok := validator.GetFieldByName(
		reflect.ValueOf(
			Sample{
				Hoge: "hoge",
				Foo:  "foo",
				Bar:  "bar",
			},
		),
		"Foo",
	)
	if !ok {
		t.Fatal("test with existing field key: expected true but not")
	}
	i := v.Interface()
	if i.(string) != "foo" {
		t.Errorf("test with existing field: expected `foo` but not")
	}

	_, ok = validator.GetFieldByName(
		reflect.ValueOf(
			Sample{
				Hoge: "hoge",
				Foo:  "foo",
				Bar:  "bar",
			},
		),
		"Piyo",
	)
	if ok {
		t.Errorf("test with not existing field: expected false but not")
	}
}

func TestIsPresentString(t *testing.T) {
	type Case struct {
		Description       string
		Input             string
		ExpectedIsPresent bool
	}

	cases := []Case{
		{
			Description:       "presented value",
			Input:             "value",
			ExpectedIsPresent: true,
		},
		{
			Description:       "empty",
			Input:             "",
			ExpectedIsPresent: false,
		},
		{
			Description:       "single space",
			Input:             " ",
			ExpectedIsPresent: false,
		},
		{
			Description:       "double space",
			Input:             "  ",
			ExpectedIsPresent: false,
		},
		{
			Description:       "many space",
			Input:             "                         ",
			ExpectedIsPresent: false,
		},
		{
			Description:       "horizontal tab",
			Input:             "\t",
			ExpectedIsPresent: false,
		},
		{
			Description:       "newline",
			Input:             "\n",
			ExpectedIsPresent: false,
		},
		{
			Description:       "vertical tab character",
			Input:             "\v",
			ExpectedIsPresent: false,
		},
		{
			Description:       "form feed",
			Input:             "\f",
			ExpectedIsPresent: false,
		},
		{
			Description:       "carriage return",
			Input:             "\r",
			ExpectedIsPresent: false,
		},
		{
			Description:       "whitespace",
			Input:             "\t\n\v\f\r ",
			ExpectedIsPresent: false,
		},
		{
			Description:       "whitespace with value",
			Input:             "\t\n\v\f\r value",
			ExpectedIsPresent: true,
		},
	}

	for _, c := range cases {
		ok := validator.IsPresentString(c.Input)
		if ok != c.ExpectedIsPresent {
			t.Errorf("test with %s: expected %t but not", c.Description, c.ExpectedIsPresent)
		}
	}
}

func TestIsPresentArray(t *testing.T) {
	type Case struct {
		Description       string
		Input             reflect.Value
		ExpectedIsPresent bool
	}

	cases := []Case{
		{
			Description:       "empty array of int",
			Input:             reflect.ValueOf([]int{}),
			ExpectedIsPresent: false,
		},
		{
			Description:       "empty array of string",
			Input:             reflect.ValueOf([]string{}),
			ExpectedIsPresent: false,
		},
		{
			Description:       "empty array of float64",
			Input:             reflect.ValueOf([]float64{}),
			ExpectedIsPresent: false,
		},
		{
			Description:       "empty array of bool",
			Input:             reflect.ValueOf([]bool{}),
			ExpectedIsPresent: false,
		},
		{
			Description:       "empty array of struct",
			Input:             reflect.ValueOf([]time.Time{}),
			ExpectedIsPresent: false,
		},
		{
			Description:       "array of int with one element",
			Input:             reflect.ValueOf([]int{1}),
			ExpectedIsPresent: true,
		},
		{
			Description:       "array of string with one element",
			Input:             reflect.ValueOf([]string{"value"}),
			ExpectedIsPresent: true,
		},
		{
			Description:       "array of float64 with one element",
			Input:             reflect.ValueOf([]float64{1.1}),
			ExpectedIsPresent: true,
		},
		{
			Description:       "array of bool with one element",
			Input:             reflect.ValueOf([]bool{true}),
			ExpectedIsPresent: true,
		},
		{
			Description:       "array of struct with one element",
			Input:             reflect.ValueOf([]time.Time{time.Now()}),
			ExpectedIsPresent: true,
		},
		{
			Description:       "array of int with many element",
			Input:             reflect.ValueOf([]int{1, 2, 3, 4, 5}),
			ExpectedIsPresent: true,
		},
	}

	for _, c := range cases {
		ok := validator.IsPresentArray(c.Input)
		if ok != c.ExpectedIsPresent {
			t.Errorf("test with %s: expected %t but not", c.Description, c.ExpectedIsPresent)
		}
	}
}

func IsPresentStruct(t *testing.T) {
	type Case struct {
		Description       string
		Input             reflect.Value
		ExpectedIsPresent bool
	}
	type emptyStruct struct{}
	type nonEmptyStruct struct {
		field string
	}

	cases := []Case{
		{
			Description:       "empty struct",
			Input:             reflect.ValueOf(emptyStruct{}),
			ExpectedIsPresent: false,
		},
		{
			Description:       "non empty struct",
			Input:             reflect.ValueOf(nonEmptyStruct{field: "value"}),
			ExpectedIsPresent: true,
		},
	}

	for _, c := range cases {
		ok := validator.IsPresentStruct(c.Input)
		if ok != c.ExpectedIsPresent {
			t.Errorf("test with %s: expected %t but not", c.Description, c.ExpectedIsPresent)
		}
	}
}

func TestIsValid(t *testing.T) {
	type Case struct {
		Description     string
		Input           interface{}
		ExpectedIsValid bool
	}

	cases := []Case{
		{
			Description:     "primitive value",
			Input:           "value",
			ExpectedIsValid: true,
		},
		{
			Description:     "valid nullable string",
			Input:           dbr.NewNullString("value"),
			ExpectedIsValid: true,
		},
		{
			Description:     "valid nullable int",
			Input:           dbr.NewNullInt64(1),
			ExpectedIsValid: true,
		},
		{
			Description:     "valid nullable float",
			Input:           dbr.NewNullFloat64(1.1),
			ExpectedIsValid: true,
		},
		{
			Description:     "valid nullable bool",
			Input:           dbr.NewNullBool(true),
			ExpectedIsValid: true,
		},
		{
			Description:     "valid nullable time",
			Input:           dbr.NewNullTime(time.Now()),
			ExpectedIsValid: true,
		},
		{
			Description:     "invalid nullable string",
			Input:           dbr.NewNullString(nil),
			ExpectedIsValid: false,
		},
		{
			Description:     "invalid nullable int",
			Input:           dbr.NewNullInt64(nil),
			ExpectedIsValid: false,
		},
		{
			Description:     "invalid nullable float",
			Input:           dbr.NewNullFloat64(nil),
			ExpectedIsValid: false,
		},
		{
			Description:     "invalid nullable bool",
			Input:           dbr.NewNullBool(nil),
			ExpectedIsValid: false,
		},
		{
			Description:     "invalid nullable time",
			Input:           dbr.NewNullTime(nil),
			ExpectedIsValid: false,
		},
	}

	for _, c := range cases {
		ok := validator.IsValid(c.Input)
		if ok != c.ExpectedIsValid {
			t.Errorf("test with %s: expected %t but not", c.Description, c.ExpectedIsValid)
		}
	}
}
