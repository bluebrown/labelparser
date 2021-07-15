package labelparser

import (
	"strings"
)

func walk(parts []string, val string, conf map[string]interface{}) {

	// if its the last part
	// assign the value to the key in the map
	if len(parts) == 1 {
		// if the value already exists
		// something must be going on
		if oldval, ok := conf[parts[0]]; ok {
			switch v := oldval.(type) {
			case map[string]interface{}:
				v[parts[0]] = val
			}
			// otherwise just assign the value
		} else {
			conf[parts[0]] = val
		}
		return
	}

	// create map at part if not exists
	if _, ok := conf[parts[0]]; !ok {
		c := map[string]interface{}{}
		conf[parts[0]] = c
	}

	// it can be that c is not a map here but a string
	// in that case i should wrap c in a new map
	switch v := conf[parts[0]].(type) {
	case string:
		c := map[string]interface{}{}
		conf[parts[0]] = c
		conf[parts[0]].(map[string]interface{})[parts[0]] = v
	}

	// use current part conf map as next conf
	c := conf[parts[0]].(map[string]interface{})

	walk(parts[1:], val, c)
}

func Parse(labels map[string]string, conf *map[string]interface{}) {
	for k, v := range labels {
		walk(strings.Split(k, "."), v, *conf)
	}
}
