package handlers

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

// MAP TYPE [json-obj driver]
type Map map[string]interface{}

// typing assertion
func (m Map) M(key string) Map {
	return m[key].(map[string]interface{})
}
func (m Map) S(key string) string {
	return m[key].(string)
}
func (m Map) I(key string) int {
	return m[key].(int)
}
// methods
func (m *Map) to_json() ([]byte, error) {
	return json.Marshal(m)
}
func (m *Map) to_string() (string, error) {
	bytes, err := m.to_json()
	if err != nil {return "", err}
	return string(bytes), nil
}



// JSON STRUCT
type JsonFile struct {
	File	string
	Data	Map
}
func (j *JsonFile) Parse() error {
	// File [open]
	file, err := os.Open(j.File)
	if err != nil { return err }

	// File [read]
	// Bytes -> Map[json-obj driver]
	bytes, _ := ioutil.ReadAll(file)
	if err := json.Unmarshal([]byte(bytes), j.Data); err != nil {
		return err
	}

	// File [close]
	file.Close()

	return nil
}
func (j *JsonFile) Clean() {
	j.Data = nil
}



// func (j *JsonFile) Parse() error {
// 	// File [open]
// 	file, err := os.Open(j.File)
// 	if err != nil {return err}
// 	defer file.Close()

// 	// File [read]
// 	// Bytes -> Map[json-obj driver]
// 	bytes, _ := ioutil.ReadAll(file)
	
// 	var m Map
// 	if err := json.Unmarshal([]byte(bytes), &m); err != nil {
// 		return err
// 	}

// 	j.Data = m
// }