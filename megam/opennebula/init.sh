#!/bin/bash

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

dist=$(grep PRETTY_NAME /etc/*-release | awk -F '="' '{print $2}')
OS=$(echo $dist | awk '{print $1;}')

#DISTRO=$OS
#ulimit  -n 65535

if [ "$DISTRO" = "Red Hat" ]  || [ "$DISTRO" = "Ubuntu" ] || [ "$DISTRO" = "Debian" ]
then
CONF='//var/lib/megam/gulp/gulpd.conf'
else 
  CONF='//var/lib/megam/verticegulpd/conf/gulpd.conf'
fi
cat >$CONF  <<'EOF'
### Welcome to the Gulpd configuration file.

  ###
  ### [meta]
  ###
  ### Controls the parameters for the Raft consensus group that stores metadata
  ### about the gulp.
  ###

  [meta]
    user = "root"
    nsqd = ["88.198.139.80:4150"]
    scylla = ["88.198.139.81"]
    scylla_keyspace = "vertice" 
  ###
  ### [gulpd]
  ###
  ### Controls which assembly to be deployed into machine
  ###

  [gulpd]
    enabled = true
    name_gulp = "hostname"
    assemblies_id = "AMS1259077729232486400"
    assembly_id = "ASM1260230009767985152"
	provider = "chefsolo"
	cookbook = "megam_run"
	chefrepo = "https://github.com/megamsys/chef-repo.git"
        chefrepo_tarball = "https://github.com/megamsys/chef-repo/archive/0.96.tar.gz"

  ###
  ### [http]
  ###
  ### Controls how the HTTP endpoints are configured. This a frill
  ### mechanism for pinging gulpd (ping)
  ###

  [http]
    enabled = false
    bind-address = "127.0.0.1:6666"

EOF

sed -i "s/^[ \t]*name_gulp.*/    name = \"$NODE_NAME\"/" $CONF
sed -i "s/^[ \t]*assemblies_id.*/    assemblies_id = \"$ASSEMBLIES_ID\"/" $CONF
sed -i "s/^[ \t]*assembly_id.*/    assembly_id = \"$ASSEMBLY_ID\"/" $CONF





case "$DISTRO" in
   "Ubuntu") 
   dist=`grep VERSION_ID /etc/*-release | awk -F '="' '{print $2}'`
   v=$(echo $dist | awk -F '"' '{print $1;}')

  case "$v" in
         "14.04")
	  stop verticegulpd
	  start verticegulpd
	  ;;
         "16.04")
          service verticegulpd stop
  	  service verticegulpd start 
     	  service cadvisor stop
          service cadvisor start
	  ;;
  esac
HOSTNAME=`hostname`
sudo cat >> //etc/hosts <<EOF
127.0.0.1  $HOSTNAME localhost
EOF

   ;;
   "Debian") 
systemctl stop verticegulpd.service
systemctl start verticegulpd.service
systemctl stop cadvisor.service
systemctl start cadvisor.service
   ;;
   "Red Hat") 
systemctl stop verticegulpd.service
systemctl start verticegulpd.service
systemctl stop cadvisor.service
systemctl start cadvisor.service
   ;;
   "CoreOS") 
if [ -f /mnt/context.sh ]; then
  . /mnt/context.sh
fi

#IP_ADDRESS=$ETH0_IP  
#GATEWAY=$ETH0_GATEWAY

#echo "ip -" $ETH0_IP "IPADD -" $IP_ADDRESS  "GW- " $GATEWAY >>/root/test
sudo cat >> //home/core/.ssh/authorized_keys <<EOF
$SSH_PUBLIC_KEY
EOF

sudo -s

sudo cat > //etc/hostname <<EOF
$HOSTNAME
EOF

sudo cat > //etc/hosts <<EOF
127.0.0.1 $HOSTNAME localhost

EOF
sudo cat > //etc/systemd/network/static.network <<EOF
[Match]
Name=ens3

[Network]
Address=$ETH0_IP/27
Gateway=$ETH0_GATEWAY
DNS=8.8.8.8
DNS=8.8.4.4

[Address]
Address=$ETH0_IP/27
Peer=$ETH0_GATEWAY

EOF

sudo systemctl restart systemd-networkd


systemctl stop verticegulpd.service
systemctl start verticegulpd.service
systemctl stop cadvisor.service
systemctl start cadvisor.service


   ;;
esac

