#!/bin/bash
set -x
set -o errexit
set -o nounset
set -o pipefail

# log to /var/log/startup-script.log
exec > >(tee -a /var/log/startup-script.log)
exec 2>&1

# kill apt processes (E: Unable to lock directory /var/lib/apt/lists/)
kill $(ps aux | grep '[a]pt' | awk '{print $2}') || true



apt-get update -y
apt-get install -y apt-transport-https curl ca-certificates

curl -fSsL https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
echo 'deb http://apt.kubernetes.io/ kubernetes-xenial main' > /etc/apt/sources.list.d/kubernetes.list

add-apt-repository -y ppa:gluster/glusterfs-3.10

apt-get update -y
apt-get install -y cron docker.io ebtables git glusterfs-client haveged kubectl kubelet nfs-common socat kubeadm= ntp || true


curl -Lo pre-k https://cdn.appscode.com/binaries/pre-k/0.1.0-alpha.5/pre-k-linux-amd64 \
	&& chmod +x pre-k \
	&& mv pre-k /usr/bin/

systemctl enable docker
systemctl start docker

cat > /etc/systemd/system/kubelet.service.d/20-pharmer.conf <<EOF
[Service]
Environment="KUBELET_EXTRA_ARGS=--node-labels=cloud.appscode.com/pool=master --cloud-provider= "
EOF

systemctl daemon-reload
systemctl restart kubelet

kubeadm reset


mkdir -p /etc/kubernetes/pki

cat > /etc/kubernetes/pki/ca.key <<EOF
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA8mimkDS5kGfSi52zUiuggGykI6bTt9ngdyTldggupA3ffsr4
bcCtd4aEzVh8qWXjsOhMlDRk/VkrfjtSBYu0XuCfxwiepbbUddqRfLK0Vk463yG+
9bhDsua+0gOL7GPyttMd8vjbNpA+kfyRJQbPl/IlF2vzpfzmEQiLyN2VH/Q2DlDI
XVuxzWtxFjwuSY4XNpz8aB7+yedwsc+Y5egjulOKplgNtY3G8XBOPoDdQtaD0XmD
sVz9FxHpEsToXn86wmBNyBErG1JiUoexXQgPUReOryZakY3HguyhWYaFakW7F+y+
vXkxwxYivanQf/Oh2B3qF0Em/RrZ3e4/YaQIWQIDAQABAoIBAQDkA1XxPPbzAwWD
eCHtf8XoJsi6FDj/rXw69cS0onukqrv4GtlU+zDvxeNy9/HCsWlbxRYLLcF0oQSN
JyU+vCs8N1NzNwNrYq8z8BR7U+jXaZxst+RUYy+ivvLsvpNGFrl+CNUZ78Ta7RE9
5nFsf5yTExyajFtRXbHWgdBibzLPKuVaBE0jkXNKPK2Xit0J95MLtkmjLI2yd9oG
+YrQ0AnwoNnJYTZy4cA4ptLkBul50v0Q7oHakAcxNe7OZ/pUgOKfBWP+YWMsMS+V
pZ6w+EVqwcey45V8NMe2wlBsGtUaNOpku5JdFoBp1djDXmmO0oFR3xzpPPMfY1vZ
A0iKIt2RAoGBAPSx1VCTtqWyiw/0c1y9vnWIcTxbCC5b1jZBwr6cDotxJhEI216Z
spVoETaI+z3k6yUnnf+SoZ0IemPH/4QxenDWt4AbOUV0dd1ypBdJT0cz/d5f8zbP
0DrtR9HHI6jdZDyYi+Dg8+6OD9ybZLYbUxiC4F0yIfaKKz/zxwqUYljtAoGBAP2b
x+dmLrWabiDFUwrcGmUBsNrI9iLhN+J9fZCjDH4kIZvXeQWeqR5VFD6HuLADLsdV
3mDWZQXPsFEpEmgBxCumJkwRIMOfC1jY38aVnAJpJEKG7NbfXp/icyz5n72pW0kH
BqNct+dqiVrdIhqABM4znQMNr+NmhWzjaq/n7JudAoGBANw2fuc5SJLuj8AYGwpH
qPRVirLqqf4uoCXMINsxztUnSz7hXatRXyqesX8G4DTfo8+YS9UJvXB5hMvv3cC2
Vp7eXd5ooa4kFT7qQUrDqxlFbR+H8nZNp/SgPpZIRYtfUypFjv0e9eim6Rd8sbJB
/RJ+FjVJPM+HoUy4s1SOHcQxAoGBAJUjYy+FO7q4EkNbRzZ/sVXs5KAExhpE76RV
v92qxNH5VXpcAGN5pmKcDcen8MgxVvUKXjf8XHRKapmP2a8oOsCgeK0K9YWUpgud
nhseKByS0Cv/RJAn03GarbipI3NZgME1TIYNJGkGkbfJiyCiYDES0SbWibTkuVoL
ctX3QYFRAoGAEt6EFAj8d+X+y8t7p0JTDJBkBtTPZJh59inn16n+g3rUMdPULvKU
OcZhazmm37JDk/e6CefX9Ns0XaS+eiUyq4CMIFT0F++3hAeT1ro9H5rVvGpxyWRr
+n9jN30O2qNirGNblaJcn3V0wF0mXQuP/RcuEAl29ZiRv/iFNcTDKkg=
-----END RSA PRIVATE KEY-----

EOF
pre-k get cacert --common-name=ca < /etc/kubernetes/pki/ca.key > /etc/kubernetes/pki/ca.crt

cat > /etc/kubernetes/pki/front-proxy-ca.key <<EOF
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA5Fl/SvO7F4v83R/MXf5XnBrbq1C0rv9OmxGTwa0VYhvmHMhr
mzjJ2ZQ6pNgTH/o6qfFTXtk9Su2HcONWI5Nl+zeqvgEEXQLDHaxlU7WRFRQf7Ca5
A4/GtHCILlt0YQrov0cUNucP1h/mE7k+4DA8kLS2uoSqE7sww6M7DyuJGhHtJmte
DC1zB78hUCaX4P7A+oHzFOZsqq9rtXlbbZDES+gVvkKRz8iBtsnnmOp7OXgtQzbd
Bn7vyg26YFZKD7sHgqDQ2P4QsrIlN7pZsW1ZGTjMt7QfRz1aRVM2jTkUGuuma+X7
tO3z5MKnbNqqFf7lNrNrju2XyIK+LyOMKS93mwIDAQABAoIBAFejAXeCkws2bpVu
3sA0HOuMY0IcdzuyQDjDRjHOe6ZrjqT9Rk8MPGuEt/0L8EsVfe468AQ2248mGhT5
KtX8e+cuNhk1yrbnmHqK71hxFn1Ae//l3JvBUkFt7tXbw4564ViHrvpjM3C6TVLE
hZ0E5jmjIX3z7FImRVDQFJp92uBlxyjkOdCL2B5M+7AFT3s87jweZtcInDMUJacU
vK9qkoQhTNgwDKRJ/PB97W0hmIcesdwEdSuH8DBS1KXpno9BSPK4r2PT/8oaLLuP
7qJgD66UJrfWJt850w+Gilge+3BnK/j0oP/LvSquonuz/QauYgIXOtEzsmqL8+4C
9dAdiAECgYEA+sIa+Qpe0AfozTDwkdQW/2Z8GwP9ZYIN+cn1qaeM8kS8/kUEcCaN
XOiwiyb50JeLi2RjATNSOuDbUkRq7wqcMOm56ahRx2fTMBMetmOhh3E9X95aBLo6
SBh4BFjMcSbNeexhEFQzmkvsRJBUHoRLHXCHEwSh7tqxTupb90o6UlECgYEA6R95
uxsndu04bOfU+F2NcOAbosJCZk6SNgov84XHOHlJiOImTmaO1gq7QUyoTbgpSK64
+fPGJ6LQpWoa3LiVJYnnNRFfDwLz5UnFjKTvIu6eta8K+MlCZXWNapCAoX6DgNpi
IY/TIlNX4XQSq4q/F2cjQfsqZtRuHJo0pup8ZCsCgYAlpGnEJMhi1i0eUFBss7fY
ExxiOdEaIdHTQ3v7QcmovKTaTqYB5+ekoNGPsgQiSf2KqUErdvbLG+IQkXCz/ZXA
yfU8nCXalz1epCAiJQakfF1IPOqqdwI0vGeXz3SxL6/8s31GpD0y1l7bVAcspZNZ
EgFU24crgezH+epgTuY+YQKBgQCAzaVCngQ2D2entIfVGkt13D0otlBdmnfqZ8KK
uJWUPF9GxD0KAFxXIuLHd4VeO6yDvhXx9KC9cZX6bx+caFfvv7wBdDGjZ6w+OX5u
8Qo3sMP7+lUDlJ+c01BpWqHXmMErxn2FsaOkSYw6wxk1splWYuP2JsIMwr58Cy+g
Mdq8tQKBgQCDgGwdTiJlkaUS++uyPbF2sHeBV5cx8xgvx/8cXtZSA6/IXUXeG4vM
Ptgn45wjmk2sqC70fg6mfbwX9bkJdQS5bGVEh7eYGqSDMrwzVZuu4o/YFF9gKXfF
o/6fE3uWyQI3BlWYLcjh/HgpWS/iHW8//H2N35lZbp9Cf0e90lDTbQ==
-----END RSA PRIVATE KEY-----

EOF
pre-k get cacert --common-name=front-proxy-ca < /etc/kubernetes/pki/front-proxy-ca.key > /etc/kubernetes/pki/front-proxy-ca.crt

chmod 600 /etc/kubernetes/pki/ca.key /etc/kubernetes/pki/front-proxy-ca.key




mkdir -p /etc/kubernetes/kubeadm


cat > /etc/kubernetes/kubeadm/config.yaml <<EOF
api:
  advertiseAddress: ""
  bindPort: 6443
apiVersion: kubeadm.k8s.io/v1alpha1
certificatesDir: ""
cloudProvider: external
etcd:
  caFile: ""
  certFile: ""
  dataDir: ""
  endpoints: null
  image: ""
  keyFile: ""
imageRepository: ""
kind: MasterConfiguration
kubernetesVersion: v1.8.0
networking:
  dnsDomain: ""
  podSubnet: ""
  serviceSubnet: ""
nodeName: ""
token: ""
tokenTTL: 0s
unifiedControlPlaneImage: ""

EOF


pre-k merge master-config \
	--config=/etc/kubernetes/kubeadm/config.yaml \
	--apiserver-advertise-address=$(pre-k get public-ips --all=false) \
	--apiserver-cert-extra-sans=$(pre-k get public-ips --routable) \
	--apiserver-cert-extra-sans=$(pre-k get private-ips) \
	--apiserver-cert-extra-sans= \
	> /etc/kubernetes/kubeadm/config.yaml
kubeadm init --config=/etc/kubernetes/kubeadm/config.yaml --skip-token-print



kubectl apply \
  -f http://docs.projectcalico.org/v2.3/getting-started/kubernetes/installation/hosted/kubeadm/1.6/calico.yaml \
  --kubeconfig /etc/kubernetes/admin.conf



kubectl apply \
  -f https://raw.githubusercontent.com/appscode/pharmer/master/addons/kubeadm-probe/ds.yaml \
  --kubeconfig /etc/kubernetes/admin.conf

mkdir -p ~/.kube
sudo cp -i /etc/kubernetes/admin.conf ~/.kube/config
sudo chown $(id -u):$(id -g) ~/.kube/config



until [ $(kubectl get pods -n kube-system -l k8s-app=kube-dns -o jsonpath='{.items[0].status.phase}' --kubeconfig /etc/kubernetes/admin.conf) == "Running" ]
do
   echo '.'
   sleep 5
done

kubectl apply -f "https://raw.githubusercontent.com/appscode/pharmer/master/cloud/providers/digitalocean/cloud-control-manager.yaml" --kubeconfig /etc/kubernetes/admin.conf

until [ $(kubectl get pods -n kube-system -l app=cloud-controller-manager -o jsonpath='{.items[0].status.phase}' --kubeconfig /etc/kubernetes/admin.conf) == "Running" ]
do
   echo '.'
   sleep 5
done

cat > /etc/systemd/system/kubelet.service.d/20-pharmer.conf <<EOF
[Service]
Environment="KUBELET_EXTRA_ARGS=--cloud-provider=external --node-labels=cloud.appscode.com/pool=master "
EOF

NODE_NAME=$(uname -n)
kubectl taint nodes ${NODE_NAME} node.cloudprovider.kubernetes.io/uninitialized=true:NoSchedule --kubeconfig /etc/kubernetes/admin.conf

systemctl daemon-reload
systemctl restart kubelet

# sleep 10
# reboot

