# validate
A simple, reusable parameter validator

Usage
```
var (
    age *int
    ids []int{1, 2, 3}
    count = 3
)
    
err := errors.Join(
    validate.Exists("age", age),
    validate.Length("ids", ids, 2),
    validate.Between("count", count, 0, 5),
)

if err != nil {
    fmt.Printf("Validation error: %s", err)
}
```
