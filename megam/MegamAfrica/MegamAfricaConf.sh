

##########################
'Master'
88.198.139.80 master.megamafrica.net
user: root
password: MegamAfrica


root@master # lsblk

NAME   MAJ:MIN RM   SIZE RO TYPE MOUNTPOINT
sda      8:0    0 223.6G  0 disk 
├─sda1   8:1    0   512M  0 part /boot
├─sda2   8:2    0    32G  0 part [SWAP]
└─sda3   8:3    0   190G  0 part /
sdb      8:16   0   3.7T  0 disk 
└─sdb1   8:17   0   3.7T  0 part /storage3
sdc      8:32   0   3.7T  0 disk 
└─sdc1   8:33   0   3.7T  0 part /storage4



##########################
'Node1'
136.243.49.217 node1.megamafrica.net
user: root
password: MegamAfrica

root@node1 ~ # cat /etc/os-release 
NAME="Ubuntu"
VERSION="14.04.4 LTS, Trusty Tahr"
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME="Ubuntu 14.04.4 LTS"
VERSION_ID="14.04"
HOME_URL="http://www.ubuntu.com/"
SUPPORT_URL="http://help.ubuntu.com/"
BUG_REPORT_URL="http://bugs.launchpad.net/ubuntu/"
root@node1 ~ # lsblk
NAME    MAJ:MIN RM   SIZE RO TYPE  MOUNTPOINT
sda       8:0    0 223.6G  0 disk  
├─sda1    8:1    0   512M  0 part  
│ └─md0   9:0    0 511.4M  0 raid1 /boot
├─sda2    8:2    0    32G  0 part  
│ └─md1   9:1    0    32G  0 raid1 [SWAP]
└─sda3    8:3    0   190G  0 part  
  └─md2   9:2    0 189.9G  0 raid1 /
sdb       8:16   0 223.6G  0 disk  
├─sdb1    8:17   0   512M  0 part  
│ └─md0   9:0    0 511.4M  0 raid1 /boot
├─sdb2    8:18   0    32G  0 part  
│ └─md1   9:1    0    32G  0 raid1 [SWAP]
└─sdb3    8:19   0   190G  0 part  
  └─md2   9:2    0 189.9G  0 raid1 /
sdc       8:32   0   5.5T  0 disk  
└─sdc1    8:33   0   5.5T  0 part  /storage1
sdd       8:48   0   5.5T  0 disk  
└─sdd1    8:49   0   5.5T  0 part  /storage2


root@slave ~ # cat /etc/os-release 
NAME="Ubuntu"
VERSION="14.04.4 LTS, Trusty Tahr"
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME="Ubuntu 14.04.4 LTS"
VERSION_ID="14.04"
HOME_URL="http://www.ubuntu.com/"
SUPPORT_URL="http://help.ubuntu.com/"
BUG_REPORT_URL="http://bugs.launchpad.net/ubuntu/"

oneadmin@slave:~$ # lsblk
NAME    MAJ:MIN RM   SIZE RO TYPE  MOUNTPOINT
sda       8:0    0 223.6G  0 disk  
├─sda1    8:1    0   512M  0 part  
│ └─md0   9:0    0 511.4M  0 raid1 /boot
├─sda2    8:2    0    32G  0 part  
│ └─md1   9:1    0    32G  0 raid1 [SWAP]
└─sda3    8:3    0   190G  0 part  
  └─md2   9:2    0 189.9G  0 raid1 /
sdb       8:16   0 223.6G  0 disk  
├─sdb1    8:17   0   512M  0 part  
│ └─md0   9:0    0 511.4M  0 raid1 /boot
├─sdb2    8:18   0    32G  0 part  
│ └─md1   9:1    0    32G  0 raid1 [SWAP]
└─sdb3    8:19   0   190G  0 part  
  └─md2   9:2    0 189.9G  0 raid1 /



