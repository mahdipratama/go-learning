# Debugging Deadlocks: Mutex Reentrancy in Go

## The Problem
In the `books.go` implementation, a deadlock occurred when calling `SetCopies`. 

### Deadlock Trace
1. `catalog.SetCopies(id, count)` is called.
2. It acquires a write lock: `catalog.mu.Lock()`.
3. It then calls `catalog.GetBook(id)`.
4. `GetBook` attempts to acquire a read lock: `catalog.mu.RLock()`.
5. **Deadlock**: The goroutine is now waiting for itself to release the Write Lock so it can take the Read Lock. Since it's waiting, it never reaches the `Unlock()` call.

## Why it happens
In Go, `sync.Mutex` and `sync.RWMutex` are **not reentrant** (also known as recursive locks). 
- **Reentrant**: A lock that can be acquired multiple times by the same goroutine without deadlocking.
- **Non-Reentrant (Go)**: If a goroutine tries to lock a mutex it already holds, it will block forever.

## The Solution: Internal vs. External Methods
To follow Go best practices and keep the code DRY (Don't Repeat Yourself) without deadlocking, we use a pattern of exported and unexported methods.

1. **Exported Methods**: (e.g., `GetBook`, `SetCopies`) Handle the locking logic and provide the public API.
2. **Internal Helpers**: (e.g., `getBook`) Perform the actual data manipulation/retrieval *without* touching the mutex.

### Example Refactor
Instead of having `SetCopies` call a locked method, it should either perform the logic directly or call a private helper:

```go
// Exported - Handles locking
func (catalog *Catalog) GetBook(ID string) (Book, bool) {
	catalog.mu.RLock()
	defer catalog.mu.RUnlock()
	return catalog.getBook(ID)
}

// Internal - No locking, safe to call from other locked methods
func (catalog *Catalog) getBook(ID string) (Book, bool) {
	book, ok := catalog.data[ID]
	return book, ok
}

// Exported - Handles locking
func (catalog *Catalog) SetCopies(ID string, copies int) error {
	catalog.mu.Lock()
	defer catalog.mu.Unlock()

	// Call the internal helper that doesn't attempt to Lock()
	book, ok := catalog.getBook(ID)
    // ... logic ...
}
```

## TDD Takeaway
When tests hang or report `all goroutines are asleep`, it is a signal to check for:
1. Circular dependencies in function calls that both require locks.
2. Forgetting to call `Unlock()` (always use `defer` to ensure safety).
3. Logic that tries to upgrade a Read Lock to a Write Lock (or vice versa) on the same goroutine.