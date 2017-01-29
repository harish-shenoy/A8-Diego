// Copyright 2016 IBM Corporation
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package database

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"

	"github.com/amalgam8/amalgam8/registry/utils/logging"
)

type redisDB struct {
	conn     redis.Conn
	pool     *redis.Pool
	logger   *log.Entry
	address  string
	password string
}

// NewRedisDB returns an instance of a Redis database
func NewRedisDB(address string, password string) Database {
	return NewRedisDBWithConn(nil, address, password)
}

// NewRedisDBWithConn returns an instance of a Redis database using an existing connection
func NewRedisDBWithConn(conn redis.Conn, address string, password string) Database {
	db := &redisDB{
		conn:     conn,
		pool:     nil,
		address:  address,
		password: password,
		logger:   logging.GetLogger(module),
	}

	return db

}

// NewRedisDBWithPool returns an instance of a Redis database using an existing connection pool
func NewRedisDBWithPool(pool *redis.Pool) Database {
	db := &redisDB{
		conn:     nil,
		pool:     pool,
		address:  "",
		password: "",
		logger:   logging.GetLogger(module),
	}

	return db

}

func (rdb *redisDB) connect() (redis.Conn, error) {
	var conn redis.Conn
	var err error
	if rdb.pool == nil {
		// Connect to Redis
		conn, err = redis.Dial("tcp", rdb.address)
		if err != nil {
			return nil, err
		}
		_, err = conn.Do("AUTH", rdb.password)
		if err != nil {
			conn.Close()
			return nil, err
		}
	} else {
		conn = rdb.pool.Get()
	}
	return conn, nil
}

func (rdb *redisDB) ReadKeys(match string) ([]string, error) {
	var err error
	conn := rdb.conn
	if rdb.conn == nil {
		conn, err = rdb.connect()
		if err != nil {
			return nil, err
		}
		defer conn.Close()
	}

	matches, err := rdb.scan(conn, match)

	return matches, err
}

func (rdb *redisDB) ReadEntry(key string) ([]byte, error) {
	var err error
	conn := rdb.conn
	if rdb.conn == nil {
		conn, err = rdb.connect()
		if err != nil {
			return nil, err
		}
		defer conn.Close()
	}

	value, err := conn.Do("GET", key)
	if value == nil || err != nil {
		return nil, err
	}

	return redis.Bytes(value, err)
}

func (rdb *redisDB) ReadAllEntries(match string) (map[string]string, error) {
	var err error
	conn := rdb.conn
	if rdb.conn == nil {
		conn, err = rdb.connect()
		if err != nil {
			return nil, err
		}
		defer conn.Close()
	}

	entries := make(map[string]string)

	matches, err := rdb.scan(conn, match)

	var args []interface{}
	for _, k := range matches {
		args = append(args, k)
	}
	values, err := redis.Strings(conn.Do("MGET", args...))
	if err != nil {
		return nil, err
	}

	for i, value := range values {
		entries[matches[i]] = value
	}

	return entries, nil
}

func (rdb *redisDB) InsertEntry(key string, entry []byte) error {
	var err error
	conn := rdb.conn
	if rdb.conn == nil {
		conn, err = rdb.connect()
		if err != nil {
			return err
		}
		defer conn.Close()
	}

	_, err = conn.Do("SET", key, entry)
	return err
}

func (rdb *redisDB) DeleteEntry(key string) (int, error) {
	var err error
	conn := rdb.conn
	if rdb.conn == nil {
		conn, err = rdb.connect()
		if err != nil {
			return 0, err
		}
		defer conn.Close()
	}

	return redis.Int(conn.Do("DEL", key))
}

func (rdb *redisDB) Expire(key string, ttl time.Duration) error {
	var err error
	conn := rdb.conn
	if rdb.conn == nil {
		conn, err = rdb.connect()
		if err != nil {
			return err
		}
		defer conn.Close()
	}

	_, err = conn.Do("EXPIRE", key, ttl.Seconds())
	return err
}

func (rdb *redisDB) scan(conn redis.Conn, match string) ([]string, error) {
	var (
		cursor int64
		keys   []string
	)

	var matches []string
	for {
		items, err := redis.Values(conn.Do("SCAN", cursor, "MATCH", match))
		if err != nil || items == nil || len(items) == 0 {
			return matches, err
		}

		items, err = redis.Scan(items, &cursor, &keys)
		if err != nil {
			return matches, err
		}

		matches = append(matches, keys...)

		if cursor == 0 {
			break
		}
	}
	return matches, nil
}
