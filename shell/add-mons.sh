add Mons  and osd 

#ceph-deploy new node2  [node3...]   //dont run when add mon to existing cluster

ceph-deploy install node2                                               // ceph-deploy uninstall node2
ceph-deploy mon add node2                                             // ceph-deploy mon destroy node2

sudo scp /etc/ceph/ceph.conf  megdc@node2:/home/megdc/
ssh megdc@node2 'sudo mkdir /etc/ceph'
ssh megdc@node2 'sudo mv ceph.conf /etc/ceph/ceph.conf'



ON NODE2
ceph-deploy mon create-initial






ceph-deploy purge node2
ceph-deploy purgedata node2
