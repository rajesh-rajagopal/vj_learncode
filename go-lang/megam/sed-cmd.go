package main

import (
  "fmt"
  "os"
  "os/exec"
)
func main() {
	assemblyID := "1234562"
        vmdir := "/var/lib/rioos/vms/" + assemblyID
        err := cpContext(vmdir,assemblyID)
	if (err != nil) {
			fmt.Printf("************cpContext*****err******* %s",err.Error())
	}
}
func  cpContext(vmDir,assemblyId  string) error {
        dest :=  vmDir + "/context.sh"
	cmd := fmt.Sprintf(`sed -i "s/^[ \t]*ASSEMBLY_ID.*/ ASSEMBLY_ID=\"%s\"/" %s`,assemblyId,dest)
	out2, err := exec.Command("/bin/sh", "-c", cmd).Output()
if (err != nil) {
    fmt.Println(" context error  ",err)
    fmt.Fprintln(os.Stderr, err.Error())
	 return err
}
fmt.Println(string(out2))
return nil
}
