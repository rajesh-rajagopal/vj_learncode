#!/bin/bash


ceph-deploy uninstall node1 
ceph-deploy purgedata node1 
ceph-deploy forgetkeys
ceph-deploy purge node1 
sudo apt-get -y autoremove

sudo apt-get -y remove ceph-deploy ceph-common ceph-mds ceph
sudo apt-get -y purge ceph-deploy ceph-common ceph-mds ceph

sudo rm -r /run/ceph
sudo rm -r /var/lib/ceph
sudo rm /var/log/upstart/ceph*
sudo rm -rf /home/megdc/ceph-cluster/*


sudo echo deb https://download.ceph.com/debian-jewel/ $(lsb_release -sc) main | tee /etc/apt/sources.list.d/ceph.list
wget -q -O- 'https://download.ceph.com/keys/release.asc' | sudo apt-key add -

sudo apt-get update
sudo apt-get install -y ceph ceph-mds ceph-deploy radosgw dnsmasq openssh-server ntp sshpass

ceph-deploy new node1

// change IPV6 to IPV4 ip   ms_bind_ipv6 = true, mon_host = [2a01:4f8:212:11e8::2] 

removed:     mon_host = [2a01:4f8:212:11e8::2] 
add ==>     mon_host = 136.243.49.217

echo "filestore_xattr_use_omap = true" >>ceph.conf
echo "osd_pool_default_size = 2" >> ceph.conf
echo "osd crush chooseleaf type = 0" >> ceph.conf
echo "mon_pg_warn_max_per_osd = 0" >> ceph.conf
 
 cat ceph.conf 
[global]
fsid = 911c2c04-91f2-409c-a021-b6d49388d6ec
mon_initial_members = node1
mon_host = 136.243.49.217
auth_cluster_required = cephx
auth_service_required = cephx
auth_client_required = cephx
filestore_xattr_use_omap = true
osd_pool_default_size = 2
osd crush chooseleaf type = 0
mon_pg_warn_max_per_osd = 0


ceph-deploy disk zap node1:sdc  
ceph-deploy disk zap node1:sdd 

ceph-deploy mon create-initial

ceph-deploy disk list node1

sudo chown -R megdc:megdc /etc/ceph/ceph.client.admin.keyring

ceph-deploy osd prepare node1:sdc 
ceph-deploy osd prepare node1:sdd 

sudo lsblk    //we cannot view mount point as megdc user 

ceph-deploy disk list node1

ceph-deploy osd activate node1:/dev/sdc1:/dev/sdc2

ceph-deploy osd activate node1:/dev/sdd1:/dev/sdd2

ceph-deploy admin node1

sudo chmod +r /etc/ceph/ceph.client.admin.keyring
sudo chown -R megdc:megdc /etc/ceph/ceph.client.admin.keyring


sleep 150
ceph osd pool set rbd pg_num 128
sleep 150
ceph osd pool set rbd pgp_num 128

