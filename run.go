package validate

import "errors"

func Run(rules ...Rule) error {
	var errs []error
	for _, rule := range rules {
		if err := rule(); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}
