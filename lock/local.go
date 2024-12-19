package lock

import (
	"context"
	"sync"
	"time"

	"github.com/air-iot/errors"
)

// KeyInfo 结构体，用于存储每个 key 的锁和过期时间
type KeyInfo struct {
	lock    *sync.Mutex
	expires time.Time
}

// KeyLock 结构体，包含锁表和锁过期时间的管理
type KeyLock struct {
	mu     sync.Mutex
	locks  map[string]*KeyInfo
	cancel context.CancelFunc
}

// NewKeyLock 创建一个 KeyLock 实例，并启动一个清理 Goroutine
func NewKeyLock() *KeyLock {
	ctx, cancel := context.WithCancel(context.Background())
	kl := &KeyLock{
		locks:  make(map[string]*KeyInfo),
		cancel: cancel,
	}
	go kl.cleanupExpiredKeys(ctx)
	return kl
}

// Lock 尝试锁定特定的 key，并在超时后返回错误
func (kl *KeyLock) Lock(ctx context.Context, key, value string, ttl time.Duration) error {
	kl.mu.Lock()
	// 如果 key 不存在，创建一个新的锁并设置过期时间
	if _, exists := kl.locks[key]; !exists {
		kl.locks[key] = &KeyInfo{
			lock:    &sync.Mutex{},
			expires: time.Now().Add(ttl),
		}
	}
	keyInfo := kl.locks[key]
	kl.mu.Unlock()

	// 使用 channel 通知锁定成功
	done := make(chan struct{})
	go func() {
		//fmt.Println(1, key, value)
		keyInfo.lock.Lock()
		//fmt.Println(2, key, value)
		close(done) // 锁定成功，关闭通道
	}()

	select {
	case <-done:
		//fmt.Println(3, key, value)
		return nil
	case <-ctx.Done():
		return errors.New("获取锁超时")
	}
}

// Unlock 解锁特定的 key，并重置其过期时间
func (kl *KeyLock) Unlock(_ context.Context, key, value string, ttl time.Duration) error {
	kl.mu.Lock()
	defer kl.mu.Unlock()
	if keyInfo, exists := kl.locks[key]; exists {
		keyInfo.lock.Unlock()
		// 更新该 key 的过期时间
		keyInfo.expires = time.Now().Add(ttl)
	}
	return nil
}

// cleanupExpiredKeys 定期清理过期的 key，支持通过 context 取消
func (kl *KeyLock) cleanupExpiredKeys(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second) // 每秒检查一次
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			kl.mu.Lock()
			now := time.Now()
			for key, keyInfo := range kl.locks {
				if now.After(keyInfo.expires) {
					delete(kl.locks, key)
				}
			}
			kl.mu.Unlock()
		case <-ctx.Done():
			// 当外部 context 取消时，停止清理 Goroutine
			return
		}
	}
}

// Stop 用于手动停止清理 Goroutine
func (kl *KeyLock) Stop() {
	kl.cancel()
}
