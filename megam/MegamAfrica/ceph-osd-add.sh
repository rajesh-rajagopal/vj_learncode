sshpass -p "megdc" scp -o StrictHostKeyChecking=no /home/megdc/.ssh/id_rsa.pub megdc@node2:/home/megdc/.ssh/authorized_keys

ceph-deploy install node2 master

ceph-deploy disk list node2

[ERROR ] ConfigError: Cannot load config: [Errno 2] No such file or directory: 'ceph.conf'; has `ceph-deploy new` been run in this directory?

sshpass -p "megdc" scp -o StrictHostKeyChecking=no /etc/ceph/ceph.conf megdc@node2:/home/megdc/ceph.conf
ssh node2 'sudo mv ceph.conf /etc/ceph/ceph.conf'

ssh node2 'sudo umount /storage3'
ssh node2 'sudo umount /storage4'

ceph-deploy disk list node2

ceph-deploy disk zap node2:sdc  
ceph-deploy disk zap node2:sdd 

ceph-deploy osd prepare node2:sdc 
ceph-deploy osd prepare node2:sdd 

sudo lsblk    //we cannot view mount point as megdc user 

ceph-deploy disk list node2

ceph-deploy osd activate node2:/dev/sdc1:/dev/sdc2
ceph-deploy osd activate node2:/dev/sdd1:/dev/sdd2

ceph-deploy admin node2

sudo chmod +r /etc/ceph/ceph.client.admin.keyring
sudo chown -R megdc:megdc /etc/ceph/ceph.client.admin.keyring


sleep 150
ceph osd pool set rbd pg_num 128
sleep 150
ceph osd pool set rbd pgp_num 128
