/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 19:11:42
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-14 13:20:12
 */
package errors

import "github.com/beego/beego/v2/core/validation"

type Errors struct {
	errors map[string][]string
}

func (e *Errors) Add(key, err string) {
	if _, ok := e.errors[key]; !ok {
		e.errors[key] = make([]string, 0, 5)
	}
	e.errors[key] = append(e.errors[key], err)
}

func (e *Errors) AddValization(valid *validation.Validation) {
	if valid.HasErrors() {
		for key, errs := range valid.ErrorsMap {
			for _, err := range errs {
				e.Add(key, err.Message)
			}
		}
	}
}

func (e *Errors) Errors() map[string][]string {
	return e.errors
}

func (e *Errors) ErrorsByKey(key string) []string {
	return e.errors[key]
}

func (e *Errors) HasErrorByKey() bool {
	return len(e.errors) != 0
}

func (e *Errors) HasErrors(key string) []string {
	return e.errors[key]
}

func New() *Errors {
	return &Errors{
		errors: make(map[string][]string),
	}
}
