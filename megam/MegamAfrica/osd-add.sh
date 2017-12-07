
#  In Node 2

echo deb https://download.ceph.com/debian-infernalis/ $(lsb_release -sc) main | tee /etc/apt/sources.list.d/ceph.list
wget -q -O- 'https://download.ceph.com/keys/release.asc' | apt-key add -

sudo apt-get -y install ceph-deploy ceph-common ceph-mds

sudo mkdir /storage3/osd
sudo mkdir /storage4/osd

adduser megdc

echo 'megdc ALL = (root) NOPASSWD:ALL' | sudo tee /etc/sudoers.d/megdc



#    In  NODE 1

cat /etc/hosts
### Hetzner Online GmbH installimage
# nameserver config
# IPv4
127.0.0.1 localhost.localdomain localhost
136.243.49.217 node1.megamafrica.net node1

2a01:4f8:212:11e8::2 node1.megamafrica.net node1
127.0.0.1  node1
136.243.49.217 node1 
138.201.21.215 node2
 node1

cat /home/megdc/.ssh/config

Host node1
 Hostname node1
 User megdc

Host node2
 Hostname node2
 User megdc

echo "public network = 138.201.98.0/27" >> ceph.conf

cat /etc/ceph/ceph.conf
[global]
fsid = 6773ed61-b6c1-48a4-9be2-c6840e0de9e7
mon_initial_members = node1
mon_host = 136.243.49.217
auth_cluster_required = cephx
auth_service_required = cephx
auth_client_required = cephx
filestore_xattr_use_omap = true
osd_crush_chooseleaf_type = 0
osd_pool_default_size = 2
mon_pg_warn_max_per_osd = 0
#public network = 136.243.49.0/27

[mon]
keyring = /etc/ceph/keyring.$name

[mds]
keyring = /etc/ceph/keyring.$name

[mon.a]
host = node1
mon addr = 136.243.49.217:6789

[mds.a]
host = node1

[osd.0]
host = node1

[osd.1]
host = node1

[osd.2]
host = node2

[osd.3]
host = node2


sshpass -p "megdc" scp -o StrictHostKeyChecking=no /home/megdc/.ssh/id_rsa.pub megdc@node2:/home/megdc/.ssh/authorized_keys
sudo scp /etc/ceph/ceph.conf megdc@node2:/home/megdc/ceph.conf
ssh node2 'sudo mv ceph.conf /etc/ceph/ceph.conf'

sudo scp /home/megdc/ceph-cluster/ceph.bootstrap-osd.keyring megdc@node2:/home/megdc/ceph.keyring

su megdc
# ceph-deploy osd create node2:/storage3/osd node2:/storage4/osd
ceph-deploy --overwrite-conf osd prepare node2:/storage3/osd node2:/storage4/osd
ceph-deploy osd activate node2:/storage3/osd node2:/storage4/osd

sudo scp /etc/ceph/*.keyring megdc@node2:/home/megdc/

sudo scp /etc/ceph/*.conf megdc@node2:/home/megdc/

ssh megdc@node2 'sudo mv  /home/megdc/*.keyring /etc/ceph/'

ssh megdc@node2 'sudo mv  /home/megdc/*.conf /etc/ceph/'

#In Node 2 




cat /etc/hosts
### Hetzner Online GmbH installimage
# nameserver config
# IPv4
127.0.0.1 localhost.localdomain localhost
138.201.21.215 node2.megamafrica.net node2

127.0.0.1  node2
136.243.49.217 node1


for down_osd in $(ceph osd tree | awk '/down/ {print $1}')
do
 host=$(ceph osd find $down_osd | awk -F\" '$2 ~ /host/ {print $4}')
 ssh $host restart ceph-osd id=$down_osd
done


ID WEIGHT   TYPE NAME      UP/DOWN REWEIGHT PRIMARY-AFFINITY 
-1 21.64975 root default                                     
-2 10.82977     host node1                                   
 0  5.41489         osd.0       up  1.00000          1.00000 
 1  5.41489         osd.1       up  1.00000          1.00000 
-3 10.81998     host node2                                   
 2  5.40999         osd.2     down        0          1.00000 
 3  5.40999         osd.3     down        0          1.00000 


megdc@node1:~/ceph-cluster$ ceph osd create
2
megdc@node1:~/ceph-cluster$ ceph osd create
3
megdc@node1:~/ceph-cluster$ ceph osd crush set 2 5.40999 pool=default rack=unknownrack host=node2   /[138.201.21.215]
set item id 2 name 'osd.2' weight 2 at location {host=138.201.21.215,pool=default,rack=unknownrack} to crush map
megdc@node1:~/ceph-cluster$ ceph osd crush set 3 5.40999 pool=default rack=unknownrack host=node2  /[138.201.21.215]
set item id 3 name 'osd.3' weight 3 at location {host=138.201.21.215,pool=default,rack=unknownrack} to crush map


ID WEIGHT   TYPE NAME               UP/DOWN REWEIGHT PRIMARY-AFFINITY 
-5  5.00000 rack unknownrack                                          
-4  5.00000     host 138.201.21.215                                   
 2  2.00000         osd.2              down        0          1.00000 
 3  3.00000         osd.3              down        0          1.00000 
-1 10.82977 root default                                              
-2 10.82977     host node1                                            
 0  5.41489         osd.0                up  1.00000          1.00000 
 1  5.41489         osd.1                up  1.00000          1.00000 
-3        0     host node2 


ceph osd rm 2
ceph osd rm 3
ln -s /storage1/osd /var/lib/ceph/osd/ceph-2 
ln -s /storage2/osd /var/lib/ceph/osd/ceph-3 
sudo ceph auth add osd.2 osd 'allow *' mon 'allow rwx' -i /var/lib/ceph/osd/ceph-2/keyring
sudo ceph auth add osd.3 osd 'allow *' mon 'allow rwx' -i /var/lib/ceph/osd/ceph-3/keyring





ceph-deploy uninstall node1 node2 master
ceph-deploy purgedata node1 node2 master
ceph-deploy purge node1 node2 master


ssh megdc@node2 'sudo apt-get -y autoremove'

ssh megdc@node2 'sudo rm -r /run/ceph'
ssh megdc@node2 'sudo rm -r /var/lib/ceph'
ssh megdc@node2 'sudo rm /var/log/upstart/ceph*'
ssh megdc@node2 'sudo rm -rf /home/megdc/ceph-cluster/*'
