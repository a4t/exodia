package exec_script

import (
  "os/exec"
  "time"
  "log"
  "../read_yaml"
)

var debug bool = false

func SetDebugFlag (debug_flag bool) {
  debug = debug_flag
}

func Run (yaml read_yaml.ScriptYaml) bool {
  if (!pre(yaml.Scripts.Pre)) { panic("Pre script is Error") }
  if (!check(yaml)) { panic("Check script is Error") }
  if (!post(yaml.Scripts.Post)) { panic("Post script is Error") }
  return true
}

func exec_run(script string) bool {
  run_logger(script)
  err := exec.Command("sh", "-c", script).Run()
  if err != nil {
    return false
  }
  return true
}

func pre(scripts []string) bool {
  for _, v := range scripts {
    if !exec_run(string(v)) {
      error_logger("Pre", string(v))
      return false
    }
  }
  return true
}

func check(yaml read_yaml.ScriptYaml) bool {
  time.Sleep(time.Duration(yaml.Wait) * time.Second)
  for i := 0; i < yaml.Checkretry; i++ {
    var success_flag bool = true

    for _, v := range yaml.Scripts.Check {
      if !exec_run(string(v)) {
        error_logger("Check", string(v))
        success_flag = false
        time.Sleep(time.Duration(yaml.Interval) * time.Second)
        break
      }
    }

    if success_flag { return true }
  }
  return false
}

func post(scripts []string) bool {
  for _, v := range scripts {
    if !exec_run(string(v)) {
      error_logger("Post", string(v))
      return false
    }
  }
  return true
}

func run_logger(text string) {
  if debug {
    log.Println("Exec:","[\033[0;33m" + text + "\033[0;39m]")
  }
}

func error_logger(group, text string) {
  if debug {
    log.Println("\033[0;31mError:", group + "[" + text + "\033[0;39m]]")
  }
}
