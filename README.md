```markdown
# Bignums Library

The `bignums` library provides utility functions for performing arithmetic operations on big numbers of various types, including `*big.Int` and `*big.Float`, in Go programming language. It provides a fluent API to chain operations and perform complex calculations with big numbers seamlessly.

## Features

- **Fluent API**: Allows chaining of multiple arithmetic operations, enhancing readability and maintainability.
- **Support for Different Types**: Handles `*big.Int` and `*big.Float`, along with standard numeric types, and their string representations.
- **Error Handling**: Robust error handling to deal with division by zero and mismatched brackets.
- **Operation Prioritization**: Offers the use of brackets to prioritize operations.

## Installation

```shell
go get <your-library-path>/bignums
```

## Usage

### Creating a Chain
Create a new chain using `NewBigIntChain` or `NewBigFloatChain`:

```go
intChain := bignums.NewBigIntChain(10)
floatChain := bignums.NewBigFloatChain(10.5)
```

### Performing Operations
Use the chainable API to perform operations:

```go
resultIntChain := intChain.Add(5).Subtract(3).Multiply(2).Divide(3)
resultFloatChain := floatChain.Add(5.5).Subtract(3.5).Multiply(2.5).Divide(3.5)
```

### Using Brackets for Prioritization
To prioritize operations, use `Begin()` and `End()`:

```go
resultChain := intChain.Add(10).Begin().Add(20).End().Multiply(2) // ((10+20)*2)
```

### Fetching Results
After performing operations, fetch the result:

```go
value, err := resultChain.Value()
if err != nil {
    log.Fatalf("Error occurred: %v", err)
}
fmt.Println("Result: ", value)
```

## Examples
Here are some examples to illustrate the usage of `bignums`:

```go
package main

import (
    "fmt"
    "<your-library-path>/bignums"
)

func main() {
    chain := bignums.NewBigIntChain(10)
    resultChain := chain.Add(5).Subtract(2).Multiply(3).Divide(2) // ((10 + 5 - 2) * 3) / 2

    value, err := resultChain.Value()
    if err != nil {
        fmt.Println("Error occurred:", err)
        return
    }

    fmt.Println("Result:", value)
}
```

## Documentation
For detailed information about each function and feature, please refer to the [GoDoc](https://pkg.go.dev/<your-library-path>/bignums).

## Contributing
If you want to contribute, feel free to open a pull request. For major changes, please open an issue first to discuss what you would like to change.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```
