package httpclient

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestResponseGetStringAndInt(t *testing.T) {
	const (
		text        = "test response text."
		keyCode     = "code"
		code        = 100
		keyMsg      = "msg"
		msg         = "this is a message"
		keyNotExist = "not-exist"
	)

	resp := newResponse([]byte(text))
	if resp.PlainText() != text {
		t.Errorf("plain text: (is) %s != (want) %s", resp.PlainText(), text)
	}

	resp = newResponse([]byte(fmt.Sprintf(`{"%s": %d, "%s": "%s"}`, keyCode, code, keyMsg, msg)))

	c1, err := resp.GetInt(keyCode)
	if err != nil || c1 != code {
		t.Errorf("GetInt: (is) %d != (want) %d, error: %v", c1, code, err)
	}

	_, err = resp.GetInt(keyNotExist)
	if err == nil {
		t.Errorf("GetInt: when a key does not exist should return error, but nil")
	} else {
		t.Log(err)
	}

	s1, err := resp.GetString(keyMsg)
	if err != nil || s1 != msg {
		t.Errorf("GetString: (is) %s != (want) %s, error: %v", s1, msg, err)
	}
	_, err = resp.GetString(keyCode)
	if err == nil {
		t.Errorf("GetString: type mismatch should return error, but nil")
	}
	_, err = resp.GetInt(keyMsg)
	if err == nil {
		t.Errorf("GetString: type mismatch should return error, but nil")
	}
}

func TestResponseGetMapAndArray(t *testing.T) {
	const (
		keyMap   = "map"
		mStr     = `{"name": "ls", "age": 18}`
		keyArray = "array"
		arrayStr = `["lg", "dell", "aoc"]`
	)
	var (
		m     = make(map[string]any)
		array = make([]any, 0)
	)

	if err := json.Unmarshal([]byte(mStr), &m); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(arrayStr), &array); err != nil {
		t.Fatal(err)
	}

	resp := newResponse([]byte(fmt.Sprintf(`{"%s": %s, "%s": %s}`, keyMap, mStr, keyArray, arrayStr)))
	m1, err := resp.GetMap(keyMap)
	if err != nil {
		t.Errorf("GetMap: %v", err)
	}
	if equal := reflect.DeepEqual(m, m1); !equal {
		t.Errorf("GetMap: not equal: is %v want %v", m1, m)
	}

	array1, err := resp.GetArray(keyArray)
	if err != nil {
		t.Errorf("GetArray: %v", err)
	}
	if equal := reflect.DeepEqual(m, m1); !equal {
		t.Errorf("GetArray: not equal: is %v want %v", array1, array)
	}
}
