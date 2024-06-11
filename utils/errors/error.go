package errors

import "fmt"

type JsonError struct {
	error
	Code	int
}

func ToJsonError(e error, code int) JsonError {
	return JsonError{error: e, Code: code}
}

func (e JsonError) MarshalJSON() ([]byte, error) {
	return []byte(`{"error": "` + e.Error() +`",` + `"code": ` + fmt.Sprint(e.Code) + `}`), nil
}


func Catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}