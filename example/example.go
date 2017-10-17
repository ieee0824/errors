package main

import (
	"github.com/ieee0824/errors"
	"fmt"
	"encoding/json"
	"strings"
)

func main() {
	var err = errors.Error{}
	fmt.Println(errors.New("hoge").String())
	json.NewDecoder(strings.NewReader(errors.New("hoge").String())).Decode(&err)

	err.Level = errors.Err
	fmt.Println(err.Level)
}