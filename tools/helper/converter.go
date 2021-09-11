package helper

import "github.com/ulule/deepcopier"

func ConvertStruct(source interface{}, dest interface{}) {
	deepcopier.Copy(source).To(dest)
}
