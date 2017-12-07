#!/bin/bash
sudo echo deb https://download.ceph.com/debian-jewel/ $(lsb_release -sc) main | tee /etc/apt/sources.list.d/ceph.list
wget -q -O- 'https://download.ceph.com/keys/release.asc' | sudo apt-key add -

sudo apt-get update
sudo apt-get install -y ceph ceph-mds ceph-deploy radosgw dnsmasq openssh-server ntp sshpass

cd /home/megdc/ceph-cluster
sudo echo '136.243.49.217 node1' >>/etc/hosts

ceph-deploy new node1

echo "osd_pool_default_size = 2" >> ceph.conf
echo "osd crush chooseleaf type = 1" >> ceph.conf
echo "mon_pg_warn_max_per_osd = 0" >> ceph.conf

echo "osd max object name len = 256" >> ceph.conf
echo "osd max object namespace len = 64" >>ceph.conf

echo "rbd default features = 1">> ceph.conf
echo "osd pool default pg num = 300" >>ceph.conf
echo "osd pool default pgp num = 300" >>ceph.conf

ceph-deploy install node1 node2 master

ceph-deploy mon create-initial

ceph-deploy osd prepare --fs-type ext4 node1:/storage1/osd node1:/storage2/osd 
ceph-deploy osd prepare --fs-type ext4 node2:/storage3/osd node2:/storage4/osd
ceph-deploy osd prepare --fs-type ext4 master:/storage5/osd master:/storage6/osd

ceph-deploy osd activate node1:/storage1/osd node1:/storage2/osd
ceph-deploy osd activate node2:/storage3/osd node2:/storage4/osd
ceph-deploy osd activate master:/storage5/osd master:/storage6/osd 
ceph-deploy admin node1 node2 master
sudo chmod +r /etc/ceph/ceph.client.admin.keyring
sleep 300
ceph osd pool set rbd pg_num 256
sleep 300
ceph osd pool set rbd pgp_num 256


