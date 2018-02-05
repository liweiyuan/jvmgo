package main

import (
	"jvmgo/ch03/classpath"
	"strings"
	"fmt"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}
func startJVM(cmd *Cmd) {
	//fmt.Printf("classpath: %s class: %s args: %v", cmd.cpOption, cmd.class, cmd.args)
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("conld not find or load main class %s\n", cmd.class)
		return
	}

	fmt.Printf("class data: %v\n", classData)
}

