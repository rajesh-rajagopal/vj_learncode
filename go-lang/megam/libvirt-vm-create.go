package main

import (
   "fmt"
   "github.com/libvirt/libvirt-go"
)

// # https://libvirt.org/uri.html
// http://www.thegeekstuff.com/2014/10/linux-kvm-create-guest-vm/?utm_source=feedburner&utm_medium=feed&utm_campaign=Feed%3A+TheGeekStuff+(The+Geek+Stuff)
// https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/6/html/Virtualization_Administration_Guide/section-libvirt-dom-xml-example.html

func buildConnection(url string) *libvirt.Connect {
	conn, err := libvirt.NewConnect(url)
	if err != nil {
		fmt.Println("********** NewConnect ****** error :",err)
           panic("NewConnect problem")
	}
	return conn
}

func buildTestConnection() *libvirt.Connect {
  return buildConnection("test:///default")
}


func buildRemoteConnection(url string) *libvirt.Connect {
   return buildConnection(url)
}

func buildLocalConnection() *libvirt.Connect {
   return buildConnection("qemu:///system")
}

func defineDomain(conn *libvirt.Connect) (*libvirt.Domain, error){
return conn.DomainDefineXML(`<domain type='kvm'>
  <name>demo2</name>
  <uuid>4dea24b3-1d52-d8f3-2516-782e98a23fa0</uuid>
  <memory>131072</memory>
  <vcpu>1</vcpu>
  <os>
    <type arch='x86_64' >hvm</type>
  </os>
  <clock sync="localtime"/>
  <devices>
    <emulator>/usr/bin/qemu-system-x86_64</emulator>
    <disk type='file' device='disk'>
      <source file='/home/gabriel/kvm/images/one-app-9-ttylinux-kvm.img'/>
      <target dev='hda'/>
    </disk>
    <interface type='network'>
      <source network='default'/>
      <mac address='24:42:53:21:52:45'/>
    </interface>
    <graphics type='vnc' port='6001' autoport='no' listen='0.0.0.0'/>
  </devices>
</domain>`)
}

// 'unsupported configuration: vnc port must be in range [5900,65535]'




  /*  <disk type='file' device='cdrom'>
     <source file='/home/ranjitha/kvm/iso/ubuntu-16.04-kvm.iso'/>
     <target dev='hdc'/>
     <readonly/>
    </disk> */
// "qemu+ssh://ranjitha@192.168.1.22/system"
func main() {
        
       //  conn := buildTestConnection()
         remote_qemu_url := "qemu+ssh://gabriel@188.240.231.85/system"
        conn := buildRemoteConnection(remote_qemu_url)
        defer conn.Close()
        fmt.Println("successfully connected to remote server : ", remote_qemu_url)
        ninfo, err := conn.GetNodeInfo()
	if err != nil {
           fmt.Println("********** GetNodeInfo ****** error :",err)
	}
        fmt.Printf("\n\n Node Info :  %#v",ninfo)
        nd, err := conn.NumOfDomains()
        fmt.Println("number of domains :", nd)

        host, err := conn.GetHostname()
        fmt.Println("remote server hostname :", host)
        doms, err := conn.ListDefinedDomains()
	if err != nil {
           fmt.Println("********** ListDefinedDomains ****** error :",err)
	}
        fmt.Printf("\n\n List of Domains that defined :  %v",doms)

// /*
	dom, err := defineDomain(conn)
	if err != nil {
           fmt.Println("********** DomainDefineXML ****** error :",err)
           panic("DomainDefineXML problem")
		
	}
       fmt.Printf("***** DomainDefineXML ******%#v",dom)
       name, err := dom.GetName()
	if err != nil {
		fmt.Println("***** Getname *****",err)
	}
       fmt.Println(name)
       err = dom.Create()
	if err != nil {
		fmt.Println("***** Create Domain*****error ",err)
	}
	fmt.Println("***** Domain Created Successfull ******")
        nd, err = conn.NumOfDomains()
        fmt.Println("after create domain : number of domains :", nd)


// */

}
