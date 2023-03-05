# validate
A simple, reusable parameter validator

Usage
```
var (
    param1 *int
    param2 []int{1, 2, 3}
    param3 = 3
)
    
err := validate.Run(
    validate.Exists("good1", param1),
    validate.Length("bad2", param2, 2),
    validate.Between("good3", param3, 0, 5),
)

if err != nil {
    fmt.Printf("Validation error: %s", err)
}
```