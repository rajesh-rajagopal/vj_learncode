
yum update -y
yum install -y libguestfs-tools
yum install -y genisoimage

virsh shutdown $1


#Get backup of Solus VM 

dd if=/dev/vg-one-0/$1_img | gzip | dd of=/home/$1_backup.gz bs=4096

#Customize VM 

wget https://github.com/OpenNebula/addon-context-linux/releases/download/v4.14.4/one-context_4.14.4.deb

mkdir packages
mv one-context_4.14.4.deb packages

genisoimage -o packages.iso -R -J -V PACKAGES packages/

qemu-img create -f qcow2 -b $1_backup $1.qcow2


cat >script.sh<<"EOF"
#!/bin/bash
mount LABEL=PACKAGES /mnt

# Install opennebula context package
dpkg -i /mnt/one-context*deb


# Script to add a user to Linux system
if [ $(id -u) -eq 0 ]; then
        username=vertice
        password=vertice
        egrep "^$username" /etc/passwd >/dev/null
        if [ $? -eq 0 ]; then
                echo "$username exists!"
                exit 1
        else
            	pass=$(perl -e 'print crypt($ARGV[0], "password")' $password)
                useradd -m -p $pass $username
                [ $? -eq 0 ] && echo "User has been added to system!" || echo "Failed to add a user!"
        fi
else
    	echo "Only root may add a user to the system"
        exit 2
fi


echo 'vertice ALL = (root) NOPASSWD:ALL' >/etc/sudoers.d/vertice 


echo "deb [arch=amd64] http://get.megam.io/0.9/ubuntu/14.04/ trusty testing" >/etc/apt/sources.list.d/vertice.list

apt-key adv --keyserver keyserver.ubuntu.com --recv B3E0C1B7

sudo apt-get update -y

sudo apt-get install -y open-vm-tools git curl

sudo apt-get install -y megamcadvisor

sudo apt-get install -y verticegulpd

if [ -f /etc/lsb-release ]; then
    . /etc/lsb-release
    DISTRO=$DISTRIB_ID
elif [ -f /etc/debian_version ]; then
    DISTRO=Debian
    # XXX or Ubuntu
elif [ -f /etc/redhat-release ]; then
    DISTRO="Red Hat"
    # XXX or CentOS or Fedora
else
    DISTRO=$(uname -s)
fi



if [ "$DISTRO" = "Red Hat" ]  || [ "$DISTRO" = "Ubuntu" ] || [ "$DISTRO" = "Debian" ] then
CONF='//usr/share/megam/verticegulpd/conf/gulpd.conf'
else if [ "$DISTRO" = "CoreOS" ]; then
  CONF='//var/lib/megam/verticegulpd/conf/gulpd.conf'
fi
fi
cat >$CONF  << 'END'

### Welcome to the Gulpd configuration file.

  ###
  ### [meta]
  ###
  ### Controls the parameters for the Raft consensus group that stores metadata
  ### about the gulp.
  ###

  [meta]
  [meta]
      user = "root"
      nsqd = ["192.168.1.100:4150"]
      scylla = ["192.168.1.100"]
      scylla_keyspace = "vertice"

  ###
  ### [gulpd]
  ###
  ### Controls which assembly to be deployed into machine
  ###

  [gulpd]
    enabled =true
    name_gulp = "hostname"
    cats_id = "AMS1259077729232486400"
    cat_id = "ASM1260230009767985152"
	provider = "chefsolo"
	cookbook = "megam_run"
	repository = "github"
	repository_path = "https://github.com/megamsys/chef-repo.git"
  repository_tar_path = "https://github.com/megamsys/chef-repo/archive/0.96.tar.gz"

  ###
  ### [http]
  ###
  ### Controls how the HTTP endpoints are configured. This a frill
  ### mechanism for pinging gulpd (ping)
  ###

  [http]
    enabled = false
    bind-address = "127.0.0.1:6666"

END

sed -i "s/^[ \t]*name_gulp.*/    name = \"kvm101\"/" $CONF
sed -i "s/^[ \t]*cats_id.*/    cats_id = \"AMS000000000000001\"/" $CONF
sed -i "s/^[ \t]*cat_id.*/    cat_id = \"ASM100000000000000\"/" $CONF


EOF

chmod 755 script.sh

export LIBGUESTFS_ATTACH_METHOD=appliance

virt-customize -v -x --attach packages.iso --format qcow2 -a $1.qcow2 --run script.sh --root-password password:pass1234

qemu-img convert -f qcow2 -O qcow2 -o compat=0.10 $1.qcow2 $1-final.qcow2

#lvcreate -n kvm101_img --size 10G /dev/vps   
qemu-img convert $1-final.qcow2 $1-final.raw

#from opennebula 
#create an ttylinux image update lvm template  used for temprary image
#resize the ttylinux image 

lvextend -L10G /dev/vg-one-0/lv-one-21

dd if=$1-final.raw of=/dev/vg-one-0/$1_img bs=4096

virsh start $1

#on node because  error: Cannot check QEMU binary /usr/bin/qemu-system-x86_64: No such file or directory

ln -s /usr/libexec/qemu-kvm /usr/bin/qemu-system-x86_64

#for permission denied error
cat /etc/libvirt/qemu.conf 
user  = "oneadmin"
group = "oneadmin"
dynamic_ownership = 0




#get source code 

wget http://libguestfs.org/download/1.33-development/libguestfs-1.33.16.tar.gz^C

http://rpm.pbone.net/index.php3/stat/4/idpl/31494289/dir/mageia_cauldron/com/golang-libguestfs-1.31.20-1.mga6.x86_64.rpm.html

	
https://www.rpmfind.net/linux/rpm2html/search.php?query=libguestfs



#dependencies error 

	http://superuser.com/questions/850276/how-to-resolve-rpm-dependencies

http://libguestfs.org/guestfs-faq.1.html

