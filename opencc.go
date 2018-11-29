package opencc

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var (
	// Dir .
	Dir               string
	configPath        = "config"
	dictionaryPath    = "dictionary"
	defaultDir        = `./`
	defaultConversion = "s2twp"
)

// Group holds a sequence of dicts
type Group struct {
	Files []string
	Dicts []*dict
}

func (g *Group) String() string {
	return fmt.Sprintf("%+v", g.Files)
}

// OpenCC contains the converter
type openCC struct {
	Conversion  string
	Description string
	DictChains  []*Group
}

var conversions = map[string]*openCC{
	"hk2s": {}, "s2hk": {}, "s2t": {}, "s2tw": {}, "s2twp": {},
	"t2hk": {}, "t2s": {}, "t2tw": {}, "tw2s": {}, "tw2sp": {},
}

// Init construct an instance of OpenCC.
func Init() {
	if Dir == "" {
		Dir = defaultDir
	}
	for k, v := range conversions {
		if err := v.dict(k); err != nil {
			panic("init dict")
		}
	}
}

// Convert .
func Convert(ctx context.Context, in string) (string, error) {
	return convert(in, defaultConversion)
}

// convert string from Simplified Chinese to Traditional Chinese or vice versa
func convert(in, conversion string) (string, error) {
	cc := conversions["s2t"]
	for _, group := range cc.DictChains {
		r := []rune(in)
		var tokens []string
		for i := 0; i < len(r); {
			s := r[i:]
			var token string
			max := 0
			for _, dict := range group.Dicts {
				ret, err := dict.prefixMatch(string(s))
				if err != nil {
					return "", err
				}
				if len(ret) > 0 {
					o := ""
					for k, v := range ret {
						if len(k) > max {
							max = len(k)
							token = v[0]
							o = k
						}
					}
					i += len([]rune(o))
					break
				}
			}
			if max == 0 { //no match
				token = string(r[i])
				i++
			}
			tokens = append(tokens, token)
		}
		in = strings.Join(tokens, "")
	}
	return in, nil
}

func (cc *openCC) dict(conversion string) error {
	cc.Conversion = conversion
	configFile := filepath.Join(Dir, configPath, conversion+".json")
	body, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	var m interface{}
	if err = json.Unmarshal(body, &m); err != nil {
		return err
	}
	config := m.(map[string]interface{})
	name, ok := config["name"]
	if !ok {
		return fmt.Errorf("name not found in %s", configFile)
	}
	cc.Description = name.(string)
	dictChain, ok := config["conversion_chain"].([]interface{})
	if !ok {
		return fmt.Errorf("conversion_chain not found in %s", configFile)
	}
	for _, v := range dictChain {
		d, ok := v.(map[string]interface{})
		if !ok {
			return fmt.Errorf("should be map inside conversion_chain")
		}
		dictMap, ok := d["dict"]
		if !ok {
			return fmt.Errorf("should have dict inside conversion_chain")
		}
		if dict, ok := dictMap.(map[string]interface{}); ok {
			group, err := cc.group(dict)
			if err != nil {
				return err
			}
			cc.DictChains = append(cc.DictChains, group)
		}
	}
	return nil
}

func (cc *openCC) group(d map[string]interface{}) (*Group, error) {
	t, ok := d["type"]
	if !ok {
		return nil, fmt.Errorf("type not found in %+v", d)
	}
	if typ, ok := t.(string); ok {
		ret := &Group{}
		switch typ {
		case "group":
			ds, ok := d["dicts"]
			if !ok {
				return nil, fmt.Errorf("no dicts field found")
			}
			dicts, ok := ds.([]interface{})
			if !ok {
				return nil, fmt.Errorf("dicts field invalid")
			}
			for _, dict := range dicts {
				d, ok := dict.(map[string]interface{})
				if !ok {
					return nil, fmt.Errorf("dicts items invalid")
				}
				group, err := cc.group(d)
				if err != nil {
					return nil, err
				}
				ret.Files = append(ret.Files, group.Files...)
				ret.Dicts = append(ret.Dicts, group.Dicts...)
			}
		case "txt":
			file, ok := d["file"]
			if !ok {
				return nil, fmt.Errorf("no file field found")
			}
			daDict, err := buildFromFile(filepath.Join(Dir, dictionaryPath, file.(string)))
			if err != nil {
				return nil, err
			}
			ret.Files = append(ret.Files, file.(string))
			ret.Dicts = append(ret.Dicts, daDict)
		default:
			return nil, fmt.Errorf("type should be txt or group")
		}
		return ret, nil
	}
	return nil, fmt.Errorf("type should be string")
}
