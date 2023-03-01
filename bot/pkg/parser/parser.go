package parser

import (
	"bufio"
	"reflect"
	"strings"
)

type Result struct {
	Commands      []string
	WalletAddress string `set:"wallet,address"`
	Reward        string `set:"reward"`
}

var resultType = reflect.TypeOf(Result{})

func ParseBody(body string) Result {
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

		action := ss[0]
		value := ss[1]

	SETTER:
		for i := 0; i < resultType.NumField(); i++ {
			field := resultType.Field(i)
			val, found := field.Tag.Lookup("set")
			if !found {
				continue
			}

			for _, v := range strings.Split(val, ",") {
				if v == action {
					resultValue.Field(i).SetString(value)
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
