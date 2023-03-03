package parser

import (
	"bufio"
	"reflect"
	"strings"
)

// Result stores parsed values.
// TODO: pass result struct as a pointer to the Parse to make parser more flexible.
type Result struct {
	Commands      []string `set:"gt"`
	WalletAddress string   `set:"wallet,address"`
	Reward        string   `set:"reward"`
}

var resultType = reflect.TypeOf(Result{})

func Parse(body string) Result {
	resultValue := reflect.ValueOf(&Result{}).Elem()

	scanner := bufio.NewScanner(strings.NewReader(body))
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\t\n")
		if len(line) < len(actionPrefix) || line[:len(actionPrefix)] != actionPrefix {
			continue
		}

		ss := strings.FieldsFunc(line[len(actionPrefix):], func(r rune) bool {
			return r == ' ' || r == ':'
		})

		if len(ss) < 2 {
			continue
		}

		action := strings.ToLower(ss[0])
		value := strings.Join(ss[1:], " ")

	SETTER:
		for i := 0; i < resultType.NumField(); i++ {
			field := resultType.Field(i)
			val, isSettable := field.Tag.Lookup("set")
			if !isSettable {
				continue
			}

			for _, v := range strings.Split(val, ",") {
				if v == action {
					fieldValue := resultValue.Field(i)

					switch fieldValue.Interface().(type) {
					case []string:
						fieldValue.Set(reflect.Append(fieldValue, reflect.ValueOf(value)))
					case string:
						fieldValue.SetString(value)
					}

					break SETTER
				}
			}
		}
	}

	return resultValue.Interface().(Result)
}

func SearchLabel(target LabelText, labels []string) bool {
	for _, v := range labels {
		if v == string(target) {
			return true
		}
	}

	return false
}
