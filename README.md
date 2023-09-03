# dlock Locking Library

dlock provides a unified way to acquire and manage distributed locks using various 3rd party backends.


## Introduction

Distributed systems often require coordination and synchronization to avoid conflicts and ensure smooth operation. dlock simplifies this process by offering an easy-to-use interface for acquiring and managing distributed locks.

Here's what you can do with this library:

- Acquire locks with a customizable time-to-live (TTL).
- Refresh locks to prevent automatic expiration.
- Retry lock acquisition with a customizable interval and limit.

## Supported Backends
- SoonTM

## Installation

```shell
go get github.com/charconstpointer/dlock
```

# Usage

## Lock Configuration

Before acquiring a lock, you can configure its behavior using the LockConfig struct and various LockOpt functions.

```go
type LockConfig struct {
    // TTL is the time-to-live of the lock. If not refreshed within this time, it will be released.
    TTL time.Duration `json:"ttl"`
    // RefreshInterval is the interval at which the lock is refreshed.
    RefreshInterval time.Duration `json:"refresh_interval"`
    // RetryInterval is the interval at which the lock is retried if held by another client.
    RetryInterval time.Duration `json:"retry_interval"`
    // RetryLimit is the maximum number of retry attempts before giving up.
    RetryLimit int `json:"retry_limit"`
}
```

Here are some configuration options:

- WithTTL: Set the time-to-live (TTL) of the lock.
- WithRefreshInterval: Configure the refresh interval of the lock.
- WithRetryInterval: Customize the retry interval for acquiring the lock.
- WithRetryLimit: Set the maximum retry attempts before giving up.

## Acquiring a Lock

To acquire a lock, you can use the Locker interface:

```go
type Locker interface {
    Lock(key, value string, opts ...LockOpt) (*Lock, error)
}
```

Here's an example of acquiring a lock:

```go
locker := NewLocker() // Initialize your locker instance
lock, err := locker.Lock("resource-key", "unique-value", WithTTL(time.Minute), WithRetryLimit(3))
if err != nil {
    // Handle lock acquisition failure
} else {
    // You've acquired the lock!
    // Access your critical resource here.
    defer lock.Release() // Release the lock when done.
}
```