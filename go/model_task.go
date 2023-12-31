/*
 * todos
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Task struct {
	Id string `json:"id"`

	Title string `json:"title"`

	Description string `json:"description,omitempty"`
}

// AssertTaskRequired checks if the required fields are not zero-ed
func AssertTaskRequired(obj Task) error {
	elements := map[string]interface{}{
		"id":    obj.Id,
		"title": obj.Title,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseTaskRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Task (e.g. [][]Task), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTaskRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTask, ok := obj.(Task)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTaskRequired(aTask)
	})
}
