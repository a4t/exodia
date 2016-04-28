package main

import (
  "log"
  "flag"

  "./read_yaml"
  "./exec_script"
)

type ParseFlag struct {
  script_yml string
  debug_flag bool
}

func parse_flag() ParseFlag {
  var script_yml *string = flag.String("f", "./script.yml", "Setting yml file")
  var debug_flag *bool = flag.Bool("d", false, "Debug")
  flag.Parse()

  var p ParseFlag
  p.script_yml = string(*script_yml)
  p.debug_flag = bool(*debug_flag)

  return p
}

func main() {
  var p ParseFlag = parse_flag()

  exec_script.SetDebugFlag(p.debug_flag)
  yaml := read_yaml.Read(p.script_yml)

  if exec_script.Run(yaml) { log.Println("Success") }
}