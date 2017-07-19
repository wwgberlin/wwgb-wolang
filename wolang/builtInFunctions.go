package wolang

import (
	"fmt"
)

func plus(terms []DataType) (result Addable, err error) {
	result = Integer{0}

	for _, v := range terms {
		if t, ok := v.(Addable); !ok {
			return nil, fmt.Errorf("error: '%v' is not a number I can add!", v.GetValue())
		} else {
			if result, err = result.Add(t); err != nil {
				return nil, fmt.Errorf("error: '%v' is not a numeric I can add!", t)
			}
		}
	}
	return result, nil
}

func strconcat(subs []DataType) (result Addable, err error) {
	result = String{""}

	for _, s := range subs {
		if t, ok := s.(Addable); !ok {
			return nil, fmt.Errorf("error: '%v' is not a string I can concat!", s.GetValue())
		} else {
			if result, err = result.Add(t); err != nil {
				return nil, fmt.Errorf("error: '%v' is not a string I can concat!", t)
			}
		}
	}
	return result, nil
}
