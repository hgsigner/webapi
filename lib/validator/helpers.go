package validator

import (
	"reflect"
	"regexp"
	"strings"
)

func getValidationTags(s interface{}) []validationData {

	strc := reflect.ValueOf(s)
	slen := strc.NumField()

	vTagsList := []validationData{}

	for i := 0; i < slen; i++ {

		typeF := strc.Type().Field(i)
		fname := strings.ToLower(typeF.Name)
		ftag := typeF.Tag.Get("validates")

		fvalue := strc.Field(i).Interface()

		if ftag != "" {
			v := regexp.MustCompile(",").Split(ftag, -1)
			if len(v) != 1 {
				for _, val := range v {
					vTagsList = append(vTagsList, validationData{fname, val, fvalue})
				}
			} else {
				vTagsList = append(vTagsList, validationData{fname, ftag, fvalue})
			}
		}
	}

	//fmt.Println(vTagsList)

	return vTagsList

}

func filterValidationName(name string) string {
	l := regexp.MustCompile("=").Split(name, -1)
	if len(l) > 1 {
		return l[0]
	}

	return name
}
