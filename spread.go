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
  command := c.Args().Get(0)
  exclude := c.String("exclude")


  files, err := ioutil.ReadDir(".")
  if err != nil {
    log.Fatal(err)
  }

  errors := make([]error, 0, len(files))
  log.Print(len(files))

  for _, file := range files {

    excluded, _ := filepath.Match(exclude, file.Name())

    if file.IsDir() && !excluded {
      // fmt.Println(file.Name())

      cmd := exec.Command("sh", "-c", command)
      cmd.Dir = file.Name()
      cmd.Stdout = os.Stdout
      cmd.Stderr = os.Stderr

      err := cmd.Run()

      if err != nil {
        errors = append(errors, err)
      }
    }
  }

  result := cli.NewMultiError(errors...)
  log.Print(result)
  return result
}
