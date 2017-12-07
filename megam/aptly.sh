  GNU nano 2.2.6                                  File: aptly-new.sh                                                                           

apt-get -y remove aptly
apt-get purge aptly
apt-get -y autoremove
rm -r ~/.aptly
apt-get -y  install aptly

aptly repo create -comment="testing" -component="testing" -distribution="trusty" -architectures="amd64, all" testing
aptly repo create -comment="nightly" -component="nightly" -distribution="trusty" -architectures="amd64, all" nightly


cd /root/nightly 
aptly repo add nightly *.deb

cd /root/testing
aptly repo add testing *.deb


aptly publish repo -component="testing","nightly" testing nightly  0.9/ubuntu/14.04


#deb http://get.megam.co/0.9/ubuntu/14.04/ testing megam

cp ~/preseed.cfg ~/.aptly/public/

cp ~/B3E0C1B7_gomegam_pub.key ~/.aptly/public/

cp -r ~/meg ~/.aptly/public/

cp ~/cadvisor.tar.gz ~/.aptly/public/

sudo echo 3 > /proc/sys/vm/drop_caches




