# To Run the script
# ./reprepro.sh version=1.0 distro=trusty release=stable
# ./reprepro.sh version=1.0 distro=trusty release=testing
# ./reprepro.sh version=1.5 distro=trusty release=testing
# ./reprepro.sh version=1.5 distro=trusty release=stable
# ./reprepro.sh version=1.5 distro=debian release=testing
# ./reprepro.sh version=1.5 distro=xenial release=stable

for i in "$@"
do
case $i in
	release=*)
	release="${i#*=}"
	;;

	version=*)
	version="${i#*=}"
	;;

	distro=*)
	distro="${i#*=}"
	;;

esac
done

function stable1()
{
mkdir -p /var/repo/1.0/ubuntu/14.04/stable/conf
cd /var/repo/1.0/ubuntu/14.04/stable/conf
cat > distributions <<EOF
Origin: ubuntu
Label: ubuntu
Suite: stable
Codename: trusty
Version: 1.0
Architectures: amd64
Components: stable
Description: vertice
SignWith: 9B46B611
EOF

cd
find ./stable -name \*.deb -exec reprepro --ask-passphrase -Vb /var/repo/1.0/ubuntu/14.04/stable includedeb trusty {} \;
mv $(find /var/www/html/repo/1.0/ubuntu/14.04/stable/pool/stable -name *.deb) /var/www/html/arkave
rm -r /var/www/html/repo/1.0/ubuntu/14.04/stable

mv /var/repo/1.0/ubuntu/14.04/stable /var/www/html/repo/1.0/ubuntu/14.04
chmod -R ugo+rX /var/www/html/repo/1.0/ubuntu/14.04
}

function testing1()
{
mkdir -p /var/repo/1.0/ubuntu/14.04/testing/conf
cd /var/repo/1.0/ubuntu/14.04/testing/conf
cat > distributions <<EOF
Origin: ubuntu
Label: ubuntu
Suite: stable
Codename: trusty
Version: 1.0
Architectures: amd64
Components: testing
Description: vertice
SignWith: 9B46B611
EOF

cd

find ./testing -name \*.deb -exec reprepro --ask-passphrase -Vb /var/repo/1.0/ubuntu/14.04/testing includedeb trusty {} \;
mv $(find /var/www/html/repo/1.0/ubuntu/14.04/testing/pool/testing -name *.deb) /var/www/html/arkave
rm -r /var/www/html/repo/1.0/ubuntu/14.04/testing

mv /var/repo/1.0/ubuntu/14.04/testing /var/www/html/repo/1.0/ubuntu/14.04
chmod -R ugo+rX /var/www/html/repo/1.0/ubuntu/14.04

}


function stable2()
{
mkdir -p /var/repo/1.5/ubuntu/14.04/stable/conf
cd /var/repo/1.5/ubuntu/14.04/stable/conf
cat > distributions <<EOF
Origin: ubuntu
Label: ubuntu
Suite: stable
Codename: trusty
Version: 1.5
Architectures: amd64
Components: stable
Description: vertice
SignWith: 9B46B611
EOF

cd
find ./stable -name \*.deb -exec reprepro --ask-passphrase -Vb /var/repo/1.5/ubuntu/14.04/stable includedeb trusty {} \;
mv $(find /var/www/html/repo/1.5/ubuntu/14.04/stable/pool/stable -name *.deb) /var/www/html/arkave
rm -r /var/www/html/repo/1.5/ubuntu/14.04/stable

mv /var/repo/1.5/ubuntu/14.04/stable /var/www/html/repo/1.5/ubuntu/14.04
chmod -R ugo+rX /var/www/html/repo/1.5/ubuntu/14.04
}
function testing2()
{
mkdir -p /var/repo/1.5/ubuntu/14.04/testing/conf
cd /var/repo/1.5/ubuntu/14.04/testing/conf
cat > distributions << EOF
Origin: ubuntu
Label: ubuntu
Suite: stable
Codename: trusty
Version: 1.5
Architectures: amd64
Components: testing
Description: vertice
SignWith: 9B46B611
EOF

cd
find testing-1.5 -name \*.deb -exec reprepro --ask-passphrase -Vb /var/repo/1.5/ubuntu/14.04/testing includedeb trusty {} \;
mv $(find /var/www/html/repo/1.5/ubuntu/14.04/testing/pool/testing -name *.deb) /var/www/html/arkave
rm -r /var/www/html/repo/1.5/ubuntu/14.04/testing
mv /var/repo/1.5/ubuntu/14.04/testing /var/www/html/repo/1.5/ubuntu/14.04
chmod -R ugo+rX /var/www/html/repo/1.5/ubuntu/14.04
}

function testing3()
{
mkdir -p /var/repo/1.5/ubuntu/16.04/testing/conf
cd /var/repo/1.5/ubuntu/16.04/testing/conf
pwd
cat > distributions <<EOF
Origin: ubuntu
Label: ubuntu
Suite: testing
Codename: xenial
Version: 1.5
Architectures: amd64
Components: testing
Description: vertice
SignWith: 9B46B611
EOF

cd
find ./stable-1.5 -name \*.deb -exec reprepro --ask-passphrase -Vb /var/repo/1.5/ubuntu/16.04/testing includedeb xenial {} \;
mv $(find /var/www/html/repo/1.5/ubuntu/16.04/testing/pool/testing -name *.deb) /var/www/html/arkave
rm -r /var/www/html/repo/1.5/ubuntu/16.04/testing

mv /var/repo/1.5/ubuntu/16.04/testing /var/www/html/repo/1.5/ubuntu/16.04
chmod -R ugo+rX /var/www/html/repo/1.5/ubuntu/16.04
}

function stable3()
{
mkdir -p /var/repo/1.5/ubuntu/16.04/stable/conf
cd /var/repo/1.5/ubuntu/16.04/stable/conf
pwd
cat > distributions <<EOF
Origin: ubuntu
Label: ubuntu
Suite: stable
Codename: xenial
Version: 1.5
Architectures: amd64
Components: stable
Description: vertice
SignWith: 9B46B611
EOF

cd
find ./ubuntu_16.04 -name \*.deb -exec reprepro --ask-passphrase -Vb /var/repo/1.5/ubuntu/16.04/stable includedeb xenial {} \;
mv $(find /var/www/html/repo/1.5/ubuntu/16.04/stable/pool/stable -name *.deb) /var/www/html/arkave
rm -r /var/www/html/repo/1.5/ubuntu/16.04/stable

mv /var/repo/1.5/ubuntu/16.04/stable /var/www/html/repo/1.5/ubuntu/16.04
chmod -R ugo+rX /var/www/html/repo/1.5/ubuntu/16.04
}

function centos1()
{
	createrepo ./centos
	cp -r ./centos/repodata /var/www/html/repo/1.5/centos/7.2/stable/
	cp ./centos/*.rpm /var/www/html/repo/1.5/centos/7.2/stable/
	chmod -R o-w+r /var/www/html/repo/1.5/centos/7.2
}

case $version in
	1.0)
		case $distro in
			trusty)
				case $release in
					stable)
					stable1
					;;
					testing)
					testing1
					;;
				esac
			;;
		esac
	;;
	1.5)
		case $distro in
			trusty)
				case $release in
					testing)
					testing2
					;;
					stable)
					stable2
					;;
				esac
			;;
			xenial)
			      	case $release in
		        		testing)
			        	testing3
			                ;;
					stable)
					stable3
					;;
      				esac
			;;
			centos)
				case $release in
					stable)
					centos1
					;;
				esac
			;;
		esac
	;;
esac


