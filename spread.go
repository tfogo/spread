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
  app.Usage = "Run scripts and commands in multiple directories at once. \n\nWEBSITE: \n\n   github.com/tfogo/spread\n\nDESCRIPTION:\n\n   Spread is designed to do one thing and do it well. It is a simpler alternative to GNU Parallel. Spread will run commands in sequence on each subdirectory. If any command in the sequence fails, the rest of the commands will not run. Therefore, you can run commands based on the exit code of previous commands. \n\nEXAMPLES: \n\n   spread 'npm test' 'git push origin master'"
  app.Version = "v1.0.1"
  app.UsageText = "spread [global options] [your commands]"
  app.HideHelp = true

  app.Authors = []cli.Author{
    cli.Author{
      Name:  "Tim Fogarty",
      Email: "tim@tfogo.com",
    },
  }

  app.Flags = []cli.Flag {
    cli.StringFlag {
      Name: "exclude, x",
      Value: "",
      Usage: "glob of directories to exclude",
    },
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
