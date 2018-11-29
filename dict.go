package opencc

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/go-ego/cedar"
)

// dict contains the Trie and dict values
type dict struct {
	Trie   *cedar.Cedar
	Values [][]string
}

// buildFromFile builds the da dict from fileName
func buildFromFile(fileName string) (*dict, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	trie := cedar.New()
	values := [][]string{}
	br := bufio.NewReader(file)
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		items := strings.Split(strings.TrimSpace(line), "\t")
		if len(items) < 2 {
			continue
		}
		err = trie.Insert([]byte(items[0]), len(values))
		if err != nil {
			return nil, err
		}

		if len(items) > 2 {
			values = append(values, items[1:])
		} else {
			values = append(values, strings.Fields(items[1]))
		}
	}
	return &dict{Trie: trie, Values: values}, nil
}

// prefixMatch str by Dict, returns the matched string and its according values
func (d *dict) prefixMatch(str string) (map[string][]string, error) {
	if d.Trie == nil {
		return nil, fmt.Errorf("Trie is nil")
	}
	ret := make(map[string][]string)
	for _, id := range d.Trie.PrefixMatch([]byte(str), 0) {
		key, err := d.Trie.Key(id)
		if err != nil {
			return nil, err
		}
		value, err := d.Trie.Value(id)
		if err != nil {
			return nil, err
		}
		ret[string(key)] = d.Values[value]
	}
	return ret, nil
}
