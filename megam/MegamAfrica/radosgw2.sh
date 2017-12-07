#/bin/bash

su megdc
cd /home/megdc/ceph-cluster
ceph-deploy install --rgw slave
ceph-deploy admin slave
ceph-deploy rgw create slave



// edit /home/megdc/ceph-cluster/ceph.conf

[client.rgw.slave]
rgw_frontends = "civetweb port=80"

ceph-deploy --overwrite-conf config push slave node1 node2 master

ssh slave 'sudo chmod +x /etc/ceph/ceph.client.admin.keyring '
ssh slave 'sudo chown -R megdc:megdc /etc/ceph/ceph.client.admin.keyring'
ssh slave 'sudo apt-get install radosgw-agent -y'


sudo service radosgw restart id=rgw.slave
//ubuntu 16.o4 to stop start
sudo systemctl start ceph-radosgw@rgw.slave

ceph-deploy admin slave
[ceph_deploy.conf][DEBUG ] found configuration file at: /home/megdc/.cephdeploy.conf
[ceph_deploy.cli][INFO  ] Invoked (1.5.33): /usr/bin/ceph-deploy admin slave
[ceph_deploy.cli][INFO  ] ceph-deploy options:
[ceph_deploy.cli][INFO  ]  username                      : None
[ceph_deploy.cli][INFO  ]  verbose                       : False
[ceph_deploy.cli][INFO  ]  overwrite_conf                : False
[ceph_deploy.cli][INFO  ]  quiet                         : False
[ceph_deploy.cli][INFO  ]  cd_conf                       : <ceph_deploy.conf.cephdeploy.Conf instance at 0x7fb29650bea8>
[ceph_deploy.cli][INFO  ]  cluster                       : ceph
[ceph_deploy.cli][INFO  ]  client                        : ['slave']
[ceph_deploy.cli][INFO  ]  func                          : <function admin at 0x7fb296da37d0>
[ceph_deploy.cli][INFO  ]  ceph_conf                     : None
[ceph_deploy.cli][INFO  ]  default_release               : False
[ceph_deploy.admin][DEBUG ] Pushing admin keys and conf to slave
[slave][DEBUG ] connection detected need for sudo
[slave][DEBUG ] connected to host: slave 
[slave][DEBUG ] detect platform information from remote host
[slave][DEBUG ] detect machine type
[slave][DEBUG ] find the location of an executable
[slave][INFO  ] Running command: sudo /sbin/initctl version
[slave][DEBUG ] write cluster configuration to /etc/ceph/{cluster}.conf


megdc@node1:~/ceph-cluster$ ceph-deploy rgw create slave
[ceph_deploy.conf][DEBUG ] found configuration file at: /home/megdc/.cephdeploy.conf
[ceph_deploy.cli][INFO  ] Invoked (1.5.33): /usr/bin/ceph-deploy rgw create slave
[ceph_deploy.cli][INFO  ] ceph-deploy options:
[ceph_deploy.cli][INFO  ]  username                      : None
[ceph_deploy.cli][INFO  ]  verbose                       : False
[ceph_deploy.cli][INFO  ]  rgw                           : [('slave', 'rgw.slave')]
[ceph_deploy.cli][INFO  ]  overwrite_conf                : False
[ceph_deploy.cli][INFO  ]  subcommand                    : create
[ceph_deploy.cli][INFO  ]  quiet                         : False
[ceph_deploy.cli][INFO  ]  cd_conf                       : <ceph_deploy.conf.cephdeploy.Conf instance at 0x7fcf6bf96998>
[ceph_deploy.cli][INFO  ]  cluster                       : ceph
[ceph_deploy.cli][INFO  ]  func                          : <function rgw at 0x7fcf6c85c6e0>
[ceph_deploy.cli][INFO  ]  ceph_conf                     : None
[ceph_deploy.cli][INFO  ]  default_release               : False
[ceph_deploy.rgw][DEBUG ] Deploying rgw, cluster ceph hosts slave:rgw.slave
[slave][DEBUG ] connection detected need for sudo
[slave][DEBUG ] connected to host: slave 
[slave][DEBUG ] detect platform information from remote host
[slave][DEBUG ] detect machine type
[slave][DEBUG ] find the location of an executable
[slave][INFO  ] Running command: sudo /sbin/initctl version
[ceph_deploy.rgw][INFO  ] Distro info: Ubuntu 14.04 trusty
[ceph_deploy.rgw][DEBUG ] remote host will use upstart
[ceph_deploy.rgw][DEBUG ] deploying rgw bootstrap to slave
[slave][DEBUG ] write cluster configuration to /etc/ceph/{cluster}.conf
[slave][WARNIN] rgw keyring does not exist yet, creating one
[slave][DEBUG ] create a keyring file
[slave][DEBUG ] create path recursively if it doesn't exist
[slave][INFO  ] Running command: sudo ceph --cluster ceph --name client.bootstrap-rgw --keyring /var/lib/ceph/bootstrap-rgw/ceph.keyring auth get-or-create client.rgw.slave osd allow rwx mon allow rw -o /var/lib/ceph/radosgw/ceph-rgw.slave/keyring
[slave][INFO  ] Running command: sudo initctl emit radosgw cluster=ceph id=rgw.slave
[ceph_deploy.rgw][INFO  ] The Ceph Object Gateway (RGW) is now running on host slave and default port 7480



ceph auth list
installed auth entries:

osd.0
	key: AQCl7UJXDsGXDBAAgPqoapJQUlXKFoaOHkdU7Q==
	caps: [mon] allow profile osd
	caps: [osd] allow *
osd.1
	key: AQDG7UJXynoJOBAAu3Vf0FpHSmKmX/Kxx1Uh+w==
	caps: [mon] allow profile osd
	caps: [osd] allow *
osd.2
	key: AQB1A0NXiUGbFxAAn9kg6dPKxGReLCXcXA37Qw==
	caps: [mon] allow profile osd
	caps: [osd] allow *
osd.3
	key: AQChA0NXJxDvMxAAIJax+kjeEF7/4YzTYI+Rgg==
	caps: [mon] allow profile osd
	caps: [osd] allow *
osd.4
	key: AQANBkNXiE1iExAAQZTVXB800g8rl023D3EZrA==
	caps: [mon] allow profile osd
	caps: [osd] allow *
osd.5
	key: AQAnBkNXVB2RKxAAzw+VLcXXtNoTfjkf6x8zpg==
	caps: [mon] allow profile osd
	caps: [osd] allow *
client.admin
	key: AQA47UJX3jiuMhAAv7hxaJtZj4sWgTj+TDSX/Q==
	caps: [mds] allow *
	caps: [mon] allow *
	caps: [osd] allow *
client.bootstrap-mds
	key: AQA57UJXsMccHBAAQg/2Smf4Pw+7wNReGRPmEw==
	caps: [mon] allow profile bootstrap-mds
client.bootstrap-osd
	key: AQA57UJXjy6DAhAAz0OC/r35OfKzp1gCACrn8w==
	caps: [mon] allow profile bootstrap-osd
client.bootstrap-rgw
	key: AQA57UJXbepADxAAzCCWC6aKN5JiNxtvD6gMBw==
	caps: [mon] allow profile bootstrap-rgw
client.libvirt
	key: AQDtB0NX444QNhAAawmxU+jczyy6txdhxbRlwA==
	caps: [mon] allow r
	caps: [osd] allow class-read object_prefix rbd_children, allow rwx pool=one
client.radosgw.gateway
	key: AQDo3UNXrdtsDRAAnVudGX+iWEQNziSp2unpnQ==
	caps: [mon] allow rwx
	caps: [osd] allow rwx
client.rgw.slave
	key: AQBSE0NX1QfpHxAANkihSDFJm5MS1RD8FA5puw==
	caps: [mon] allow rw
	caps: [osd] allow rwx


sudo ceph-authtool --create-keyring /etc/ceph/ceph.client.radosgw.keyring
sudo chmod +r /etc/ceph/ceph.client.radosgw.keyring
sudo ceph -k /etc/ceph/ceph.client.admin.keyring auth add client.rgw.slave -i /etc/ceph/ceph.client.radosgw.keyring




radosgw-admin bucket rm --purge-data --uid=megamafrica@gmail.com


radosgw-admin user rm --uid=megamafrica@gmail.com

radosgw-admin user info --uid=megamafrica@gmail.com
