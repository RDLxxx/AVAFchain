package merkle

import (
	"fmt"
	"sync"

	"golang.org/x/crypto/sha3"

	"github.com/near/borsh-go"
)

type HashNode []byte
type ValueNode []byte

type NodeFlag struct {
	hash  HashNode // Хэш узла
	dirty bool     // Флаг, указывающий, был ли узел изменён
}

type Node interface {
	fstring(string) string
	cache() (HashNode, bool)
}

type FullNode struct {
	Children [17]Node // 16 ветвей + значение
	flags    NodeFlag
	mu       sync.RWMutex // Мьютекс для потокобезопасности
}

type ShortNode struct {
	Key   []byte // Ключ (путь)
	Val   Node   // Значение (дочерний узел)
	flags NodeFlag
	mu    sync.RWMutex // Мьютекс для потокобезопасности
}

// Методы для FullNode

func (n *FullNode) FHash() (HashNode, error) {
	n.mu.RLock()
	if n.flags.hash != nil {
		defer n.mu.RUnlock()
		return n.flags.hash, nil
	}
	n.mu.RUnlock()

	n.mu.Lock()
	defer n.mu.Unlock()

	// Повторная проверка, так как хэш мог быть вычислен другой горутиной
	if n.flags.hash != nil {
		return n.flags.hash, nil
	}

	data, err := borsh.Serialize(n)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize FullNode: %w", err)
	}

	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	n.flags.hash = hash.Sum(nil)

	return n.flags.hash, nil
}

func (n ValueNode) Fstring(indent string) string {
	return fmt.Sprintf("ValueNode{%x}", []byte(n))
}

func (n *FullNode) FMarkDirty() {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.flags.dirty = true
}

func (n *FullNode) FCache() (HashNode, bool) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.flags.hash, n.flags.dirty
}

// Методы для ShortNode

func (n *ShortNode) SHash() (HashNode, error) {
	n.mu.RLock()
	if n.flags.hash != nil {
		defer n.mu.RUnlock()
		return n.flags.hash, nil
	}
	n.mu.RUnlock()

	n.mu.Lock()
	defer n.mu.Unlock()

	// Повторная проверка, так как хэш мог быть вычислен другой горутиной
	if n.flags.hash != nil {
		return n.flags.hash, nil
	}

	data, err := borsh.Serialize(n)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize ShortNode: %w", err)
	}

	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	n.flags.hash = hash.Sum(nil)

	return n.flags.hash, nil
}

func (n *ShortNode) SMarkDirty() {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.flags.dirty = true
}

func (n *ShortNode) SCache() (HashNode, bool) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.flags.hash, n.flags.dirty
}
