# Go Method Receivers and Map Semantics

This document summarizes the architectural decisions regarding the `SetCopies` implementation in the `books` package, focusing on method receivers and Go map behavior.

## 1. Multiple Receivers with the Same Method Name
In Go, methods are identified by their **receiver type**. It is perfectly valid (and often preferred) to have multiple types implement a method with the same name.

### The `Book` Receiver
The `Book` struct implements `SetCopies` to handle internal logic and validation for a single entity.
```go
func (b Book) SetCopies(copies int) error {
    // Validation logic (e.g., no negative numbers)
    // ...
}
```

### The `Catalog` Receiver
The `Catalog` type (a `map[string]Book`) implements `SetCopies` to handle data orchestration: finding the book by ID and persisting changes back to the map.
```go
func (catalog Catalog) SetCopies(ID string, copies int) error {
    // Orchestration logic: Find, Delegate, Update
    // ...
}
```

## 2. Delegation Pattern (DRY)
Inside `Catalog.SetCopies`, we call `book.SetCopies(copies)`. This is **delegation**. 
- The `Catalog` is responsible for **finding** the record.
- The `Book` is responsible for **validating** the new state.

This adheres to the **DRY (Don't Repeat Yourself)** principle. If validation rules for "copies" change, they only need to be updated in the `Book` method.

## 3. Go Map Semantics
A critical reason for the implementation pattern in `Catalog.SetCopies` is how Go handles maps:

1. **Maps Return Copies**: When you retrieve a struct from a map (`book := catalog[ID]`), Go returns a copy of that struct.
2. **Immutability in Maps**: You cannot modify a struct field directly inside a map (e.g., `catalog[ID].Copies = 5` is a compiler error).
3. **The Update Cycle**: To update a value in a map of structs, you must:
   - Extract the copy.
   - Modify the copy (using the `Book` receiver method).
   - Re-assign the modified copy back to the map key.

## 4. TDD Benefits
By separating these concerns:
- We can write a unit test for `Book` to verify validation logic.
- We can write a unit test for `Catalog` to verify that it correctly identifies and updates the right record without re-testing the validation logic.