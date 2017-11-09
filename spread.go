package main

import (
  "fmt"
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
    cli.StringFlag {
      Name: "setup, s",
      Value: "",
      Usage: "Setup command to run before running commands in subdirectories",
    },
  }

  app.Action = action

  app.Run(os.Args)
}

func action(c *cli.Context) error {
  command := c.Args().Get(0)
  exclude := c.String("exclude")
  setup := c.String("setup")

  setupCmd := exec.Command("sh", "-c", setup)
  setupCmd.Stdout = os.Stdout
  setupCmd.Stderr = os.Stderr
  err := setupCmd.Run()
  if err != nil {
    log.Fatal(err)
  }

  files, err := ioutil.ReadDir(".")
  if err != nil {
    log.Fatal(err)
  }

  for _, file := range files {

    excluded, _ := filepath.Match(exclude, file.Name())

    if file.IsDir() && !excluded {
      fmt.Println(file.Name())
      cmd := exec.Command("sh", "-c", command)
      cmd.Dir = file.Name()
      cmd.Stdout = os.Stdout
      cmd.Stderr = os.Stderr
      err := cmd.Run()
      if err != nil {
        log.Fatal(err)
      }
    }
  }

  return nil
}
