package validator

import (
	"regexp"
	"strconv"
)

type validationData struct {
	fieldName string
	validaton string
	value     interface{}
}

func validatePresence(data validationData) {
	if data.value == "" {
		ErrosListHolder.AppendError(ErrorMessage{data.fieldName, "Can not be blank"})
		return
	}
}

func validateFormat(data validationData) {
	// ErrosListHolder.AppendError(ErrorMessage{data.fieldName, "Must e valid"})
	// return
}

func validateMinMax(data validationData) {

	vl := regexp.MustCompile("=").Split(data.validaton, -1)
	refValue, err := strconv.Atoi(vl[1])
	if err != nil {
		ErrosListHolder.AppendError(ErrorMessage{data.fieldName, "Must be a number"})
		return
	}

	switch vl[0] {
	case "min":
		if data.value.(int) < refValue {
			ErrosListHolder.AppendError(ErrorMessage{data.fieldName, "Must be greater than " + vl[1]})
			return
		}
	case "max":
		if data.value.(int) > refValue {
			ErrosListHolder.AppendError(ErrorMessage{data.fieldName, "Must be less than " + vl[1]})
			return
		}
	}

}
