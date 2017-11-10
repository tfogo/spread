package main

import (
  "os"
  "os/exec"
  "io/ioutil"
  "log"
  "path/filepath"

  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "spread"
  app.Usage = "Run commands in subdirectories"

  app.Flags = []cli.Flag {
    cli.StringFlag {
      Name: "exclude, x",
      Value: "",
      Usage: "Glob of directories to exclude",
    },
  }

  app.Action = action

  app.Run(os.Args)
}

func action(c *cli.Context) error {
  commands := c.Args()
  exclude := c.String("exclude")

  files, err := ioutil.ReadDir(".")
  if err != nil {
    log.Fatal(err)
  }

  errors := make([]error, 0, len(files))

  for _, file := range files {

    excluded, _ := filepath.Match(exclude, file.Name())

    if file.IsDir() && !excluded {
      // fmt.Println(file.Name())
      for _, command := range commands {
        err := runCmd(command, file)
        if err != nil {
          errors = append(errors, err)
          break
        }
      }
    }
  }

  if len(errors) > 0 {
    return cli.NewMultiError(errors...)
  } else {
    return nil
  }

}

func runCmd(command string, file os.FileInfo) error {

  cmd := exec.Command("sh", "-c", command)
  cmd.Dir = file.Name()
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  err := cmd.Run()

  return err
}
