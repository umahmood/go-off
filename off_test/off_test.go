package off_test

import (
	"strings"
	"testing"

	off "github.com/umahmood/go-off"
)

type pair struct {
	key   string
	value interface{}
}

const arrayTest = `;; array.off --- simple off array examples

;; Example array
my_array {1|2|3|Hello, World!|https://www.kernel.org/|200|<p>Hello</p>}

;; Very long array
long_array {Hello|World|10|1|11|12|34|34|12|233|3447|2324}

;; Person
john {doe|11|https://jdoe.foo|+000000000}
`

func TestArray(t *testing.T) {
	config, err := off.LoadConfig(strings.NewReader(arrayTest))
	if err != nil {
		t.Error(err)
	}

	if config.StringCount() != 0 {
		t.Errorf("unexpected string count")
	}

	if config.BoolCount() != 0 {
		t.Errorf("unexpected bool count")
	}

	if config.IntCount() != 0 {
		t.Errorf("unexpected int count")
	}

	if config.ArrayCount() != 3 {
		t.Errorf("unexpected array count")
	}
}

const boolTest = `; boolean.off - boolean examples in off

download_images true
enable_colors true
simple_mode false
disable_http2 true
remove_on_exit false
`

func TestBool(t *testing.T) {
	config, err := off.LoadConfig(strings.NewReader(boolTest))
	if err != nil {
		t.Error(err)
	}

	if config.BoolCount() != 5 {
		t.Errorf("unexpected bool count")
	}

	if config.StringCount() != 0 {
		t.Errorf("unexpected string count")
	}

	if config.IntCount() != 0 {
		t.Errorf("unexpected int count")
	}

	if config.ArrayCount() != 0 {
		t.Errorf("unexpected array count")
	}
}

const commentTest = `;;
; This is an example for comments in off
;;

; I'm a comment
version 1

;; Another one
name awesome-name ; Comment in the end of the line
`

func TestComment(t *testing.T) {
	config, err := off.LoadConfig(strings.NewReader(commentTest))
	if err != nil {
		t.Error(err)
	}

	if config.BoolCount() != 0 {
		t.Errorf("unexpected bool count")
	}

	if config.StringCount() != 1 {
		t.Errorf("unexpected string count")
	}

	if config.IntCount() != 1 {
		t.Errorf("unexpected int count")
	}

	if config.ArrayCount() != 0 {
		t.Errorf("unexpected array count")
	}
}

const integerTest = `;; integer.off: integer examples in off

age 5
id 64
int_array {1|2|3|4|5|6|7|9|10|11}
`

func TestInteger(t *testing.T) {
	config, err := off.LoadConfig(strings.NewReader(integerTest))
	if err != nil {
		t.Error(err)
	}

	if config.BoolCount() != 0 {
		t.Errorf("unexpected bool count")
	}

	if config.StringCount() != 0 {
		t.Errorf("unexpected string count")
	}

	if config.IntCount() != 2 {
		t.Errorf("unexpected int count")
	}

	if config.ArrayCount() != 1 {
		t.Errorf("unexpected array count")
	}
}

const stringTest = `;; string.off: examples of string values in off

name John
last_name Doe
email_address john@foo.bar
pgp 0xDEADBEEF
`

func TestString(t *testing.T) {
	config, err := off.LoadConfig(strings.NewReader(stringTest))
	if err != nil {
		t.Error(err)
	}

	if config.BoolCount() != 0 {
		t.Errorf("unexpected bool count")
	}

	if config.StringCount() != 4 {
		t.Errorf("unexpected string count")
	}

	if config.IntCount() != 0 {
		t.Errorf("unexpected int count")
	}

	if config.ArrayCount() != 0 {
		t.Errorf("unexpected array count")
	}
}

const simpleTest = `;; Example configuration file in off
;; the keys are meaningless and just proof of concept

;; Global 
access_token 1234
base_url https://api.foo.bar

;; Create history file
history true
history_size 10000

;; Locations
config_location /home/user/.config/
save_files false 

;; Example array
user {username|usermail@foo.bar|http://userweb.site/}
`

func TestSimple(t *testing.T) {
	config, err := off.LoadConfig(strings.NewReader(simpleTest))
	if err != nil {
		t.Error(err)
	}

	if config.BoolCount() != 2 {
		t.Errorf("unexpected bool count")
	}

	if config.StringCount() != 2 {
		t.Errorf("unexpected string count")
	}

	if config.IntCount() != 2 {
		t.Errorf("unexpected int count")
	}

	if config.ArrayCount() != 1 {
		t.Errorf("unexpected array count")
	}

	bools := []pair{
		{key: "history", value: true},
		{key: "save_files", value: false},
	}
	for _, p := range bools {
		v, err := config.Bool(p.key)
		if err != nil {
			t.Errorf("bool - key %v not in bools", p.key)
		}
		if p.value.(bool) != v {
			t.Errorf("bool - incorrect value %v for key %v", p.value, p.key)
		}
	}

	ints := []pair{
		{key: "access_token", value: 1234},
		{key: "history_size", value: 10000},
	}
	for _, p := range ints {
		v, err := config.Int(p.key)
		if err != nil {
			t.Errorf("int key %v not in ints", p.key)
		}
		if p.value.(int) != v {
			t.Errorf("int incorrect value %v for key %v", p.value, p.key)
		}
	}

	strs := []pair{
		{key: "base_url", value: "https://api.foo.bar"},
		{key: "config_location", value: "/home/user/.config/"},
	}
	for _, p := range strs {
		v, err := config.String(p.key)
		if err != nil {
			t.Errorf("string key %v not in strings", p.key)
		}
		if p.value.(string) != v {
			t.Errorf("string incorrect value %v for key %v", p.value, p.key)
		}
	}

	arr := []interface{}{"username", "usermail@foo.bar", "http://userweb.site/"}
	user, err := config.Array("user")
	if err != nil {
		t.Errorf("user array not in config")
	}
	if len(arr) != len(user) {
		t.Errorf("array mismatch in length arr: %v user: %v", len(arr), len(user))
		t.FailNow()
	}
	for i := 0; i < len(arr); i++ {
		if arr[i].(string) != user[i].(string) {
			t.Errorf("array mismatch in array items arr: %v user: %v", arr[i], user[i])
		}
	}
}
