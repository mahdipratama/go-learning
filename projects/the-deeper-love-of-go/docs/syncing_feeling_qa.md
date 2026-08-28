# Implementation and Testing Patterns: "That Syncing Feeling"

This document captures the logic behind the implementation of the `books` package as discussed during the development of the Catalog API.

## 1. Method Delegation and Map Semantics

### The Question
In `Catalog.SetCopies`, why do we call `book.SetCopies(copies)` instead of updating the value directly?

### The Logic
1. **Receiver Types**: Go allows multiple types to have methods with the same name. `Catalog.SetCopies` handles the map lookup, while `Book.SetCopies` handles the internal validation.
2. **Map Semantics**: In Go, retrieving a struct from a map (`book := catalog[ID]`) returns a **copy**, not a reference. 
3. **DRY (Don't Repeat Yourself)**: By calling the `Book` method, we reuse existing validation logic (e.g., preventing negative copy counts).
4. **The Pattern**: 
   - Extract the copy.
   - Delegate the update/validation to the `Book` type.
   - Re-insert the modified copy back into the map.

## 2. Test Sanity Checks (Pre-conditions)

### The Question
In `TestAddBook_ReturnsErrorIfIDExists`, why do we check if the book is present *before* performing the action?

```go
_, ok := catalog.GetBook("ABC04")
if !ok {
    t.Fatal("Book not present")
}
```

### The Logic
This is a **Sanity Check** (or Pre-condition). Its purpose is to ensure the test environment is set up correctly before the "real" test begins.
1. **Avoiding False Positives**: If the book isn't in the catalog, `AddBook` might fail for the wrong reason (or succeed when it should have failed).
2. **Distinguishing Errors**: 
   - If the test fails at the check: The **test data** is wrong.
   - If the test fails at the action: The **code logic** is wrong.
3. **Self-Documenting Tests**: It tells future developers exactly what state the catalog must be in for this specific test case to be valid.

## 3. Concurrency and Safety

### The Question
Is the duplicate check in `AddBook` related to concurrency?

### The Logic
Yes. In a multi-user (concurrent) environment, a duplicate check is the first line of defense:
1. **Preventing Overwrites**: Without this check, User B could call `AddBook` and silently overwrite work just completed by User A.
2. **Atomic Rules**: By making `AddBook` return an error for existing IDs, we force developers to use the correct API for updates (`SetCopies`), preventing accidental data loss.
3. **Identifying Race Conditions**: This highlights the "Check-then-Act" pattern. Even with the check, if two users call the function at the exact same time, they might both pass the check. This creates the necessity for **Mutexes** (locks) to ensure only one person can "Check-then-Act" at a time.

## 4. TDD AAA Pattern
Our tests follow the **Arrange, Act, Assert** pattern:
- **Arrange**: Set up the catalog and verify its state (Sanity Check).
- **Act**: Call the function being tested (`AddBook` or `SetCopies`).
- **Assert**: Verify the result is exactly what we expected.