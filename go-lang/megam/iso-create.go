package main

import (
  "fmt"
  "os"
  "os/exec"
)
func main() {
	vmdir := "/var/lib/rioos/vms/" + "1234562"
	err := mkdir(vmdir)
	if (err != nil) {
	fmt.Printf("************mkdir*****err******* %s",err.Error())
	}
  isoPath, err := createISO(vmdir, "1234562")
	if err != nil {
		fmt.Printf("************createISO*****err******* %s",err.Error())
	}
   fmt.Println("**************************",isoPath)
}

func  cpContext(vmDir,assemblyId  string) error {
        app1 := "cp"
        src := "/var/lib/rioos/context/context.sh"
        dest := vmDir + "/context.sh"

cmd := exec.Command(app1, src, dest)
out, err := cmd.Output()
if (err != nil) {
    fmt.Println("context copy error",err)
    fmt.Fprintln(os.Stderr, err.Error())
	 return err
}
fmt.Printf("************cpContext*****command**out***** %v",string(out))
	app2 := "sed"
        arg := "-i"
	arg0 := fmt.Sprintf(`"s/^[ \t]*ASSEMBLY_ID.*/ASSEMBLY_ID= \"%s\"/"`,assemblyId)
  	arg1 := dest
	fmt.Printf("************sed Context**exec***command******* %s %s %s",app2,arg0, arg1)
cmd2 := exec.Command(app2, arg, arg0, arg1)
out2, err := cmd2.Output()
if (err != nil) {
    fmt.Println(" context error  ",err)
    fmt.Fprintln(os.Stderr, err.Error())
	 return err
}
fmt.Printf("************cpContext**exec***command**out***** %v",string(out2))
return nil
}

func  mkdir(path string) error {
	app := "mkdir"
	arg0 := "-p"
	arg1 := path
cmd := exec.Command(app, arg0,arg1)
fmt.Printf("************mkdir**exec***command******* %s %s %s",app,arg0, arg1)
out, err := cmd.Output()
if (err != nil) {
			fmt.Printf("************cpContext**exec***err******* %s",err.Error())
    fmt.Fprintln(os.Stderr, err.Error())
	 return err

}
    fmt.Println(string(out))
return nil
}



func  createISO(vmdir, assemblyID string ) (string, error) {
  err := cpContext(vmdir,assemblyID)
	if (err != nil) {
			fmt.Printf("************cpContext*****err******* %s",err.Error())
	}
	outPath := vmdir + "/context.iso"
	contextPath := vmdir + "/context.sh"
	initPath := "/var/lib/rioos/context/init.sh"
	gulpd := "/var/lib/rioos/context/gulpd"
	versionPath := "/var/lib/rioos/context/VERSION"
	// genisoimage -V CONTEXT -o iso_outputname -r context.sh init.sh gulpd VERSION
	app := "genisoimage"
arg0 := "-V"
arg0a := "CONTEXT"
arg1 := "-o" 
arg1a := outPath
arg2 := "-r"
fmt.Printf("************iso**exec***command******* %s %s %s %s %s ",app,arg0, arg1,arg2,initPath)
cmd := exec.Command(app, arg0,arg0a,arg1,arg1a,arg2,contextPath,initPath,gulpd,versionPath)
out, err := cmd.Output()
if (err != nil) {
fmt.Fprintln(os.Stderr, err.Error())
	fmt.Printf("************createISO**exec***err******* %s",err.Error())
   fmt.Fprintln(os.Stderr, err.Error())
	 return "", err
}
fmt.Println(string(out))
return  outPath, nil
}

func  cloneImage(vmpath, srcimage string ) (string, error) {
	// genisoimage -V CONTEXT -o iso_outputname -r context.sh init.sh gulpd VERSION
 app := "cp"
cmd := exec.Command(app, srcimage,vmpath)
_, err := cmd.Output()
if (err != nil) {
	 fmt.Fprintln(os.Stderr, err.Error())
	 return "", err
}
return vmpath, nil
}

