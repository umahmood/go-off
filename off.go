package off

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	offComment    = ";"
	offArrayStart = "{"
	offArrayEnd   = "}"
	offArraySep   = "|"
)

const (
	offArray = iota
	offBool
	offInt
	offString
)

// Config represents a single off configuration instance
type Config struct {
	strings map[string]string
	bools   map[string]bool
	ints    map[string]int
	arrays  map[string][]interface{}
}

// LoadConfig opens and reads in, an off configuration file
func LoadConfig(r io.Reader) (*Config, error) {
	off := &Config{
		strings: make(map[string]string),
		arrays:  make(map[string][]interface{}),
		bools:   make(map[string]bool),
		ints:    make(map[string]int),
	}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		err := off.parseLine(scanner.Text())
		if err != nil {
			return nil, err
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return off, nil
}

// parseArray parses an off array type
func parseArray(value string) []interface{} {
	var q []interface{}
	a := strings.Trim(value, offArrayStart+offArrayEnd)
	for _, b := range strings.Split(a, offArraySep) {
		_, val := parseValue(b)
		q = append(q, val)
	}
	return q
}

// parseValue parses the value from a key/value pair
func parseValue(value string) (int, interface{}) {
	if string(value[0]) == offArrayStart {
		return offArray, nil
	} else if value == "true" {
		return offBool, true
	} else if value == "false" {
		return offBool, false
	} else if byte(value[0]) >= 48 && byte(value[0]) <= 57 {
		// we want numbers like '+00000000' to be treated as a string and not
		// an int, if we use Atoi it becomes '0'.
		if v, err := strconv.Atoi(value); err == nil {
			return offInt, v
		}
	}
	return offString, value
}

// parseLine parses a single line in an off config file
func (c *Config) parseLine(line string) error {
	if line == "" {
		return nil
	} else if strings.Contains(line, offComment) {
		// single line comment
		if line[:1] == offComment {
			return nil
		}
		// comment on the same line as key value pair i.e.:
		// server_id 42 ; this is a comment
		idx := strings.Index(line, offComment)
		line = strings.Trim(line[:idx], " ")
	}
	line += " "
	idx := strings.Index(line, " ")
	key := line[:idx]
	value := strings.Trim(line[idx:], " ")
	if len(value) == 0 {
		return fmt.Errorf("off: key %q has no value", key)
	}
	t, v := parseValue(value)
	switch t {
	case offArray:
		a := parseArray(value)
		c.arrays[key] = append(c.arrays[key], a...)
	case offBool:
		c.bools[key] = v.(bool)
	case offInt:
		c.ints[key] = v.(int)
	case offString:
		c.strings[key] = v.(string)
	}
	return nil
}

// String retrieves a string value from configuration
func (c *Config) String(key string) (string, error) {
	if _, ok := c.strings[key]; !ok {
		return "", fmt.Errorf("off: key %q not found in configuration", key)
	}
	return c.strings[key], nil
}

// StringCount returns the number of string key/value pairs in the configuration
func (c *Config) StringCount() int {
	return len(c.strings)
}

// Bool retrieves a boolean value from configuration
func (c *Config) Bool(key string) (bool, error) {
	if _, ok := c.bools[key]; !ok {
		return false, fmt.Errorf("off: key %q not found in configuration", key)
	}
	return c.bools[key], nil
}

// BoolCount returns the number of boolean key/value pairs in the configuration
func (c *Config) BoolCount() int {
	return len(c.bools)
}

// Int retrieves a integer value from configuration
func (c *Config) Int(key string) (int, error) {
	if _, ok := c.ints[key]; !ok {
		return 0, fmt.Errorf("off: key %q not found in configuration", key)
	}
	return c.ints[key], nil
}

// IntCount returns the number of integer key/value pairs in the configuration
func (c *Config) IntCount() int {
	return len(c.ints)
}

// Array retrieves an offArray value from configuration
func (c *Config) Array(key string) ([]interface{}, error) {
	if _, ok := c.arrays[key]; !ok {
		return nil, fmt.Errorf("off: key %q not found in configuration", key)
	}
	return c.arrays[key], nil
}

// ArrayCount returns the number of offArray key/value pairs in the configuration
func (c *Config) ArrayCount() int {
	return len(c.arrays)
}
