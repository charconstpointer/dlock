package dlock

import "time"

type LockConfig struct {
	// TTL is the time-to-live of the lock. If the lock is not refreshed within
	// this time, it will be released.
	TTL time.Duration `json:"ttl"`
	// RefreshInterval is the interval at which the lock is refreshed.
	RefreshInterval time.Duration `json:"refresh_interval"`
	// RetryInterval is the interval at which the lock is retried if it is
	// already held by another client.
	RetryInterval time.Duration `json:"retry_interval"`
	// RetryLimit is the maximum number of times the lock is retried before
	// giving up.
	RetryLimit int `json:"retry_limit"`
}

// Lock represents a distributed lock acquired by a client.
type Lock struct {
	// Key is the name of the lock.
	Key string `json:"key"`
	// Value is the value of the key that was set when the lock was acquired.
	Value string `json:"value"`
	// Identity is the unique identifier of the client that acquired the lock.
	Identity string `json:"identity"`
	// Timestamp is the time at which the lock was acquired.
	Timestamp time.Time `json:"timestamp"`
}

// Locker represents a service that can be used to acquire distributed locks.
type Locker interface {
	Lock(key, value string, opts ...LockOpt) (*Lock, error)
}

// LockOpt is a function that can be used to configure a lock.
type LockOpt func(*LockConfig)

// WithTTL configures the TTL of the lock.
func WithTTL(ttl time.Duration) LockOpt {
	return func(c *LockConfig) {
		c.TTL = ttl
	}
}

// WithRefreshInterval configures the refresh interval of the lock.
func WithRefreshInterval(interval time.Duration) LockOpt {
	return func(c *LockConfig) {
		c.RefreshInterval = interval
	}
}

// WithRetryInterval configures the retry interval of the lock.
func WithRetryInterval(interval time.Duration) LockOpt {
	return func(c *LockConfig) {
		c.RetryInterval = interval
	}
}

// WithRetryLimit configures the retry limit of the lock.
func WithRetryLimit(limit int) LockOpt {
	return func(c *LockConfig) {
		c.RetryLimit = limit
	}
}
