REGISTRY=quay.io/coreos/etcd:v3.5.9
DATA_DIR=/Users/c/Developer/gin/Examples/EtcdData
NODE1=node1

docker run -d -p 2379:2379 \
  -p 2380:2380 \
  --volume=${DATA_DIR}:/etcd-data \
  --name etcd ${REGISTRY} \
  /usr/local/bin/etcd \
  --data-dir=/etcd-data --name node1 \
  --initial-advertise-peer-urls http://${NODE1}:2380 --listen-peer-urls http://0.0.0.0:2380 \
  --advertise-client-urls http://${NODE1}:2379 --listen-client-urls http://0.0.0.0:2379 \
  --initial-cluster node1=http://${NODE1}:2380


# 使用
# 因为没有了sh bash 只能直接使用
# docker exec -it 443b44ff2390 /usr/local/bin/etcdctl get a
#(base)   docker % docker exec -it 443b44ff2390 /usr/local/bin/etcdctl put a b
#OK
#(base)   docker % docker exec -it 443b44ff2390 /usr/local/bin/etcdctl get a
#a
#b