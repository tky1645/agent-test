package main

type Stringer interface{ String() string }

type S string

func (s S) String() string {
	return string(s)
}

type error interface{ Error() string }
type typeError struct {
	message string
}

func (e typeError) Error() string {
	return e.message + " | typeError"
}

func cast2Stringer(v interface{}) (Stringer, error) {
	if n, ok := v.(Stringer); ok {
		return n, nil
	}
	return nil, typeError{message: "Stringerへのキャスト"}
}

func main() {
	// s := S("aaa") 成功パターン
	s := 10 //失敗パターン
	v, err := cast2Stringer(s)

	if err != nil {
		println(err.Error())
	} else {
		println(v.String())
	}
}