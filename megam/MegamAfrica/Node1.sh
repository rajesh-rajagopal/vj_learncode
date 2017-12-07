Node1 Setup Scripts


useradd megdc -m -s /bin/bash

echo deb https://download.ceph.com/debian-infernalis/ $(lsb_release -sc) main | tee /etc/apt/sources.list.d/ceph.list
wget -q -O- 'https://download.ceph.com/keys/release.asc' | apt-key add -
apt-get install apt-transport-https  sudo
apt-get install ceph-deploy ceph-common ceph-mds dnsmasq openssh-server ntp sshpass ceph radosgw

sudo mkdir /storage1/osd
sudo mkdir /storage2/osd


####################################In Node0 (ceph mon) 
echo "136.243.49.217 node1.megamafrica.net" >> /etc/hosts

sshpass  scp -o StrictHostKeyChecking=no /home/megdc/.ssh/id_rsa.pub megdc@node1.megamafrica.net:/home/megdc/.ssh/authorized_keys

## while run the above command error : Permission denied, please try again.(because megdc doesn't have password) so manually copied the key

edit .ssh/config

Host master
 Hostname master
 User megdc
Host node1
 Hostname node1.megamafrica.net
 User megdc


ceph-deploy osd prepare node1:/dev/storage1  node1:/dev/storage2
