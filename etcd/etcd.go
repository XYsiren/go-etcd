package etcd

import clientv3 "go.etcd.io/etcd/client/v3"

func getEtcdEndpoints() []string {
	return []string{"0.0.0.0:2379"}
}

func GetEtcdClient() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: getEtcdEndpoints(),
	})
	return cli, err
}
