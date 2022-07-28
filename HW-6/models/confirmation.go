package models

import "encoding/json"

type Confirmation struct {
	UserName string `json:"username"`
	Code     string `json:"code"`
}

func (c *Confirmation) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}
func (c *Confirmation) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}
