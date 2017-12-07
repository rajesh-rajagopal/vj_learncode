sudo radosgw-admin user create --uid="megamafrica@gmail.com" --display-name="MegamAfrica"
2016-05-25 15:58:58.815347 7f14945fc900  0 RGWZoneParams::create(): error creating default zone params: (17) File exists
{
    "user_id": "megamafrica@gmail.com",
    "display_name": "MegamAfrica",
    "email": "",
    "suspended": 0,
    "max_buckets": 1000,
    "auid": 0,
    "subusers": [],
    "keys": [
        {
            "user": "megamafrica@gmail.com",
            "access_key": "HW51273AGNR5JS0CBPYI",
            "secret_key": "9kbUcvxZsVzSYDvSteSViq0is1lUK5EnyDdsfvu4"
        }
    ],
    "swift_keys": [],
    "caps": [],
    "op_mask": "read, write, delete",
    "default_placement": "",
    "placement_tags": [],
    "bucket_quota": {
        "enabled": false,
        "max_size_kb": -1,
        "max_objects": -1
    },
    "user_quota": {
        "enabled": false,
        "max_size_kb": -1,
        "max_objects": -1
    },
    "temp_url_keys": []
}


{
App 6223 stdout:     "user_id": "megamafrica@gmail.com",
App 6223 stdout:     "display_name": "Megamafrica",
App 6223 stdout:     "email": "",
App 6223 stdout:     "suspended": 0,
App 6223 stdout:     "max_buckets": 1000,
App 6223 stdout:     "auid": 0,
App 6223 stdout:     "subusers": [],
App 6223 stdout:     "keys": [
App 6223 stdout:         {
App 6223 stdout:             "user": "megamafrica@gmail.com",
App 6223 stdout:             "access_key": "EDKIE9MWA762U3UZDNIL",
App 6223 stdout:             "secret_key": "C2bCdOqMQUBDFpz6SJf4UGCDuFzx2DqAqENx5bHv"
App 6223 stdout:         }
App 6223 stdout:     ],
App 6223 stdout:     "swift_keys": [],
App 6223 stdout:     "caps": [],
App 6223 stdout:     "op_mask": "read, write, delete",
App 6223 stdout:     "default_placement": "",
App 6223 stdout:     "placement_tags": [],
App 6223 stdout:     "bucket_quota": {
App 6223 stdout:         "enabled": false,
App 6223 stdout:         "max_size_kb": -1,
App 6223 stdout:         "max_objects": -1
App 6223 stdout:     },
App 6223 stdout:     "user_quota": {
App 6223 stdout:         "enabled": false,
App 6223 stdout:         "max_size_kb": -1,
App 6223 stdout:         "max_objects": -1
App 6223 stdout:     },
App 6223 stdout:     "temp_url_keys": []
App 6223 stdout: }


  "user": "vijay@megam.io",
  "access_key": "D9Y365NCVCG2UNZZL0MH",
  "secret_key": "SooHgSIzDIxoTQB5vdUmIVww5PO0u3Afc13344di"





go version configuration

version: 0.1
log:
  fields:
    service: registry
#storage:
 # cache:
  #  blobdescriptor: inmemory
  #filesystem:
   # rootdirectory: /var/lib/registry
# Ceph Object Gateway Configuration
# See http://ceph.com/docs/master/radosgw/ for details on installing this service.
storage:
  s3:
    accesskey: D9Y365NCVCG2UNZZL0MH            
    secretkey: SooHgSIzDIxoTQB5vdUmIVww5PO0u3Afc13344di
    region: default
    regionendpoint: http://88.198.139.81
    bucket: vijay-test-bucket1
    encrypt: false
    secure: false
    chunksize: 5242880
    rootdirectory: /s3/docker/registry
http:
  addr: :5000
  headers:
    X-Content-Type-Options: [nosniff]
health:
  storagedriver:
    enabled: true
    interval: 10s
    threshold: 3
    
    
    



$ sudo apt-get update
sudo apt-get install apt-transport-https ca-certificates
sudo apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D
echo 'deb https://apt.dockerproject.org/repo ubuntu-trusty main' >/etc/apt/sources.list.d/docker.list

sudo apt-get update

sudo apt-get install linux-image-extra-$(uname -r)

sudo apt-get install docker-engine

docker run -d -p 6000:6000 --restart=always --name registry \
  -v `pwd`/config.yml:/etc/docker/registry/config.yml \
  registry

docker pull ubuntu 

docker tag ubuntu localhost:6000/ubuntu

docker push localhost:6000/ubuntu

docker pull localhost:6000/ubuntu            docker pull localhost:6000/alpine

docker stop registry && docker rm -v registry


$$$$$$$$$$$$$$$$$$$$$$$$$

docker run -d -p 6000:6000  --name s3registry  -v `pwd`/config.yml:/etc/docker/registry/config.yml registry

docker stop s3registry && docker rm -v s3registry
docker run -d -p 6000:6000 --restart=always --name registry  -v `pwd`/config.yml:/etc/docker/registry/config.yml registry:2

sudo docker tag ubuntu localhost:6000/ubuntu
sudo docker push localhost:6000/ubuntu
sudo docker pull localhost:6000/ubuntu

docker push http://88.198.139.81/vijay-test-bucket1?region=default ubuntu


  http://88.198.139.81/vijay-test-bucket1/Screenshot%20from%202016-05-11%2019-07-21.png?AWSAccessKeyId=D9Y365NCVCG2UNZZL0MH&Expires=1464249280&Signature=7f%2B%2F3FKF%2FRzH9ZObNARmdHIUr1k%3D
  
  http://objects.dreamhost.com/my-bucket-name/secret_plans.txt?Signature=XXXXXXXXXXXXXXXXXXXXXXXXXXX&Expires=1316027075&AWSAccessKeyId=XXXXXXXXXXXXXXXXXXX
  
docker push http://88.198.139.81/vijay-test-bucket1/?region

https://hub.docker.com/r/lorieri/registry-ceph-automated/

sudo apt-get install build-essential python-dev libevent-dev python-pip liblzma-dev
  
https://rominirani.com/docker-tutorial-series-part-6-docker-private-registry-15d1fd899255#.3ll7qhp4d
https://docs.docker.com/registry/deploying/
https://github.com/docker/distribution/blob/master/docs/storage-drivers/s3.md
https://github.com/docker/distribution/blob/master/docs/configuration.md#storage
  
           -e AWS_USER=megdc \
         -e AWS_PASSWORD=megdc \
         
  https://hub.docker.com/r/lorieri/registry-ceph-automated/

docker stop s3registry && docker rm -v s3registry


curl --insecure https://127.0.0.1:5001/debug/health


docker run \
         -e SETTINGS_FLAVOR=s3 \
         -e AWS_BUCKET=vijay-test-bucket1 \
         -e AWS_REGION=default \
         -e STORAGE_PATH=/registry \
         -e AWS_KEY=D9Y365NCVCG2UNZ \
         -e AWS_SECRET=C2bCdOqMQUBDFpz6SJf4UGCDuFzx2DqAqENx5bHv \
         -p 5000:5000 \
         -e AWS_HOST=88.198.139.81 \
         -e AWS_SECURE=false \
         -e AWS_ENCRYPT=false \
         -e AWS_PORT=80 \
         -e AWS_DEBUG=true \
         -e AWS_CALLING_FORMAT=OrdinaryCallingFormat \
          --name s3registry \
         registry

