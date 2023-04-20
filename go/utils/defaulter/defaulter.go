// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package defaulter

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	yaml "github.com/ghodss/yaml"
)

func readYaml(filename string) (content map[string]interface{}, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)
	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}

	jsonDoc, err := yaml.YAMLToJSON(buffer)
	if err != nil {
		return nil, err
	}

	var defaultsFileStruct map[string]interface{}
	err = json.Unmarshal(jsonDoc, &defaultsFileStruct)
	if err != nil {
		return nil, err
	}

	return defaultsFileStruct, nil
}

func SetDefaults(ptr interface{}) error {
	buff, err := readYaml("./modules/gke-cluster/defaults.yaml")
	if err != nil {
		return err
	}

	fmt.Println(buff)

	v := reflect.ValueOf(ptr).Elem()
	//t := v.Type()

	// Read Yaml File into Struct
	// Read 'comment tags' into struct

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() == reflect.Struct {
			t := v.Type()
			fmt.Println(t.Field(i).Name)
			SetDefaults(f.Addr().Interface())
		} else {
			t := v.Type()
			fmt.Println(t.Field(i).Name)
			fmt.Println(t.Field(i).Tag.Get("default"))
			fmt.Println(f.Kind())
			//if f.Kind() != reflect.Int {
			//fmt.Println(f.Elem().Interface())
			//}
		}

	}
	return nil
}
