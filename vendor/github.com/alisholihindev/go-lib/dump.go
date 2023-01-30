package lib

import (
	"encoding/json"
	"fmt"

	"github.com/TylerBrock/colorjson"
)

func BeautifyStruct(m struct{}) {
	var obj map[string]interface{}

	j, _ := json.Marshal(m)
	json.Unmarshal([]byte(j), &obj)
	f := colorjson.NewFormatter()
	f.Indent = 4

	s, _ := f.Marshal(obj)
	fmt.Println(string(s))
}

func BeautifyJson(j string) {
	var obj map[string]interface{}

	json.Unmarshal([]byte(j), &obj)
	f := colorjson.NewFormatter()
	f.Indent = 4

	s, _ := f.Marshal(obj)
	fmt.Println(string(s))
}
