package main

import "fmt"
import "runtime"
import "io/ioutil"
import osr "github.com/dominodatalab/os-release"

func main() {
	const OS string = runtime.GOOS
	const ARCH string = runtime.GOARCH

	if OS == "linux" {
		data, err := ioutil.ReadFile("/etc/os-release")
		if err != nil {
			fmt.Println(OS, "/", ARCH)
			return
		}

		info := osr.Parse(string(data))

		fmt.Println(OS, "/", ARCH)
		fmt.Println(info.ID, "/", info.VersionID)
		
		return
	}
	
	fmt.Println(OS, ARCH)
}
