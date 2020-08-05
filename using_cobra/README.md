## [Cobra](https://github.com/spf13/cobra)

Cobra is both a library for creating powerful modern CLI applications as well as a program to generate applications and command files.

Many of the most widely used Go projects are built using Cobra, such as: Kubernetes, Hugo, Docker

Content here is taken from an [online tutorial](https://towardsdatascience.com/how-to-create-a-cli-in-golang-with-cobra-d729641c7177).

Cobra is built on a structure of commands, arguments & flags.
-  Commands represent actions
-  Args are things
-  Flags are modifiers for those actions

## Cobra Installation
```bash
go get github.com/spf13/cobra/cobra
```

## Project Setup
```bash
go mod init my-calc
cobra init --pkg-name mycalc
```

## Installing the CLI
Open the terminal inside the my-calc project and run:
```bash
go install my-calc
```

This command will generate the binary or executable file of the project in the $GOPATH/bin folder.
