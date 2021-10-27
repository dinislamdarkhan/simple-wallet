package conns

import (
	"fmt"
	"time"

	"github.com/dinislamdarkhan/simple-wallet/src/app/config"

	"github.com/gocql/gocql"
)

type Connections struct {
	Cassandra *gocql.Session
}

func (c Connections) Close() {
	c.Cassandra.Close()
}

func New(cass config.CassandraConfig) (c Connections, err error) {
	if c.Cassandra, err = cassandraConn(cass); err != nil {
		return c, fmt.Errorf("cassandraConn: %v", err)
	}

	return c, nil
}

func cassandraConn(cfg config.CassandraConfig) (*gocql.Session, error) {
	cluster := gocql.NewCluster(cfg.ConnectionIP...)
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = 10 * time.Second
	cluster.Timeout = 10 * time.Second
	cluster.SocketKeepalive = 10 * time.Second

	cluster.Keyspace = cfg.Keyspace
	cluster.Consistency = gocql.Quorum

	client, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return client, nil
}
