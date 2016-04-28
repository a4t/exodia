package read_yaml

import (
  "io/ioutil"
  "gopkg.in/yaml.v2"
)

type ScriptYaml struct {
  Wait int
  Interval int
  Checkretry int
  Scripts struct {
    Pre []string
    Check []string
    Post []string
  }
}

func Read(yaml_file string) ScriptYaml {
  buf, err := ioutil.ReadFile(yaml_file)
  if err != nil {
    panic(err)
  }

  var m ScriptYaml
  err = yaml.Unmarshal(buf, &m)
  if err != nil {
    panic(err)
  }

  return m
}