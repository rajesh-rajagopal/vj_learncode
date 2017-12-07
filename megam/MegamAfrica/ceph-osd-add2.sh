sshpass -p "megdc" scp -o StrictHostKeyChecking=no /home/megdc/.ssh/id_rsa.pub megdc@master:/home/megdc/.ssh/authorized_keys

ceph-deploy install master

ceph-deploy disk list master

i didnot meet this error for master only even there wasn't exist /etc/ceph/ceph.conf file{
#[ERROR ] ConfigError: Cannot load config: [Errno 2] No such file or directory: 'ceph.conf'; has `ceph-deploy new` been run in this directory?

#sshpass -p "megdc" scp -o StrictHostKeyChecking=no /etc/ceph/ceph.conf megdc@master:/home/megdc/ceph.conf
#ssh master 'sudo mv ceph.conf /etc/ceph/ceph.conf'
}


ssh master 'sudo umount /storage4'
ssh master 'sudo umount /storage5'

ceph-deploy disk list master

ceph-deploy disk zap master:sdb  
ceph-deploy disk zap master:sdc 

ceph-deploy osd prepare master:sdb 
ceph-deploy osd prepare master:sdc 

ssh master 'sudo lsblk'    //we cannot view mount point as megdc user 

ceph-deploy disk list master

ceph-deploy osd activate master:/dev/sdb1:/dev/sdb2
ceph-deploy osd activate master:/dev/sdc1:/dev/sdc2

ceph-deploy admin master

sleep 150
ceph osd pool set rbd pg_num 256
sleep 150
ceph osd pool set rbd pgp_num 256

sshpass -p "megdc" scp -o StrictHostKeyChecking=no /etc/ceph/*.keyring megdc@master:/home/megdc/
ssh master 'sudo mv *.keyring /etc/ceph/

ssh master 'sudo chmod +r /etc/ceph/ceph.client.admin.keyring'
ssh master 'sudo chown -R megdc:megdc /etc/ceph/ceph.client.admin.keyring'



