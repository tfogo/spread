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
  app.Usage = "Run scripts and commands in multiple directories. \n\nWEBSITE: \n\n   github.com/tfogo/spread\n\nDESCRIPTION: \n\n   The first argument to spread is the command that will be run. By default, spread will run the command in all the subdirectories of the present working directory. Add a list of files as arguments after the command to specify the directories in which the command should run.\n\nEXAMPLES: \n\n   spread 'npm test && git push origin master'\n   spread ls dir1 dir2 dir3"
  app.Version = "v2.0.0"
  app.UsageText = "spread [your command] [directories ...]"
  app.HideHelp = true

  app.Authors = []cli.Author{
    cli.Author{
      Name:  "Tim Fogarty",
      Email: "tim@tfogo.com",
    },
  }

  app.Flags = []cli.Flag {
    cli.BoolFlag {
      Name: "help, h",
      Usage: "show help",
    },
  }

  app.Action = action

  app.Run(os.Args)
}

func action(c *cli.Context) error {
  if c.Bool("help") {
    cli.ShowAppHelp(c)
    return nil
  }

  if len(c.Args()) == 0 {
    cli.ShowAppHelp(c)
    return nil
  }

  command, fileArgs := c.Args()[0], c.Args()[1:]
  files := make([]string, 0, len(fileArgs))

  if len(fileArgs) == 0 {
    dirs, err := ioutil.ReadDir(".")
    if err != nil {
      log.Fatal(err)
    }
    for _, fileInfo := range dirs {
      file, err := filepath.Abs(fileInfo.Name())
      if err != nil {
        log.Fatal(err)
      }

      if fileInfo.IsDir() {
        files = append(files, file)
      }
    }
  } else {
    for _, fileString := range fileArgs {
      file, err := filepath.Abs(fileString)
      if err != nil {
        log.Fatal(err)
      }
      fileInfo, err := os.Stat(file)
      if err != nil {
        log.Fatal(err)
      }

      if fileInfo.IsDir() {
        files = append(files, file)
      } else {
        log.Fatal(file, " is not a directory")
      }

    }
  }

  errors := make([]error, 0, len(files))

  for _, file := range files {
    err := runCmd(command, file)
    if err != nil {
      errors = append(errors, err)
    }
  }

  if len(errors) > 0 {
    return cli.NewMultiError(errors...)
  } else {
    return nil
  }
}

func runCmd(command string, file string) error {

  cmd := exec.Command("sh", "-c", command)
  cmd.Dir = file
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  err := cmd.Run()

  return err
}
