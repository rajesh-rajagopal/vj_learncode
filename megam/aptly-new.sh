apt-get -y remove aptly
apt-get purge aptly
apt-get -y autoremove
rm -r ~/.aptly
apt-get   install aptly

aptly repo create -comment="stable" -component="stable" -distribution="trusty" -architectures="amd64, all" stable
aptly repo create -comment="nightly" -component="nightly" -distribution="trusty" -architectures="amd64, all" nightly
aptly repo create -comment="jessie-testing" -component="testing" -distribution="jessie" -architectures="amd64, all" jessie
aptly repo create -comment="testing" -component="testing" -distribution="trusty" -architectures="amd64, all" testing

#aptly repo create -comment="testing-1.5" -component="testing" -distribution="trusty" -architectures="amd64, all" testing

cd /root/nightly 
aptly repo add nightly *.deb

cd /root/testing
aptly repo add testing *.deb

cd /root/stable
aptly repo add stable *.deb

cd /root/testing-jessie
aptly repo add jessie *.deb

#cd /root/testing-1.5
#aptly repo add testing *.deb


aptly publish repo -component="stable","testing","nightly" stable testing nightly 1.0/ubuntu/14.04

aptly publish repo -component="testing" jessie 1.0/debian/8

#aptly publish repo  -component="testing" testing 1.5/ubuntu/14.04 
#deb http://get.megam.co/0.9/ubuntu/14.04/ trusty testing
#deb http://get.megam.io/0.9/debian/8/ jessie testing

cp ~/preseed.cfg ~/.aptly/public/

cp ~/B3E0C1B7_gomegam_pub.key ~/.aptly/public/

cp -r ~/meg ~/.aptly/public/

cp ~/cadvisor.tar.gz ~/.aptly/public/

sudo echo 3 > /proc/sys/vm/drop_caches


