package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"simple-go-server/pkg/constant"

	"github.com/go-playground/validator"
)

func ValidateBody(r *http.Request, reqData interface{}) error {
	// Decode the JSON request body into the provided struct
	if err := json.NewDecoder(r.Body).Decode(reqData); err != nil {
		return err
	}

	if fmt.Sprintf("%v", reqData) == "&{  }" {
		return constant.ErrBodyCannotBeEmpty
	}
	// Validate the decoded data using the validator
	validate := validator.New()
	if err := validate.Struct(reqData); err != nil {
		return err
	}

	// Return nil if both decoding and validation succeed
	return nil
}

func ValidateQueryParams(jsonData []byte, reqData interface{}) error {
	// Decode the JSON request body into the provided struct
	err := json.Unmarshal(jsonData, reqData)

	if err != nil {
		return err
	}

	// Validate the decoded data using the validator
	validate := validator.New()
	if err := validate.Struct(reqData); err != nil {
		return err
	}

	// Return nil if both decoding and validation succeed
	return nil
}

func QueryToJSON(query url.Values) (string, error) {
	// Create a map to store our processed query parameters
	queryMap := make(map[string]interface{})

	// Iterate through all query parameters
	for key, values := range query {
		// If there's only one value, store it directly
		if len(values) == 1 {
			queryMap[key] = values[0]
		} else {
			// If there are multiple values, store them as an array
			queryMap[key] = values
		}
	}

	// Convert map to JSON
	jsonBytes, err := json.Marshal(queryMap)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}