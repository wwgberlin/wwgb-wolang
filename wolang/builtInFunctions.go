package wolang

import (
	"fmt"
)

func plus(terms []interface{}) (result interface{}, err error) {
	result = int(0)

	for _, v := range terms {
		if val, ok := v.(int); !ok {
			return nil, fmt.Errorf("error: '%v' is not a number I can add!", v)
		} else {
			result = result.(int) + val
		}
	}
	return result, nil
}

func strconcat(subs []interface{}) (result interface{}, err error) {
	result = ""

	for _, s := range subs {
		if substr, ok := s.(string); !ok {
			return nil, fmt.Errorf("error: '%v' is not a string I can concat!", s)
		} else {
			result = result.(string) + substr
		}
	}
	return result, nil
}
