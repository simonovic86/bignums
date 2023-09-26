### bignums Package

The `bignums` package is a Go library designed to provide utility functions for performing arithmetic operations on big numbers of various types, including `*big.Int` and `*big.Float`. It facilitates the chaining of arithmetic operations and ensures graceful error handling, allowing for cleaner and more concise code.

### Installation

To install the `bignums` package, run the following command:

```sh
go get github.com/simonovic86/bignums
```

Once installed, import it in your project:

```go
import "github.com/simonovic86/bignums"
```

### Usage

#### BigIntChain

```go
bic := bignums.NewBigIntChain(5) // Initialize with an integer, *big.Int, or a string.
bic.Add(10).Subtract(3).Multiply(2) // Chain arithmetic operations.

result, err := bic.Value() // Retrieve the result.
if err != nil {
    log.Fatal(err) // Handle errors gracefully.
}
fmt.Println(result) // Output: 24
```

#### BigFloatChain

```go
bfc := bignums.NewBigFloatChain(3.5) // Initialize with a float, *big.Float, or a string.
bfc.Add("1.5").Multiply(2) // Chain arithmetic operations.

result, err := bfc.Value() // Retrieve the result.
if err != nil {
    log.Fatal(err) // Handle errors gracefully.
}
fmt.Println(result) // Output: 10
```

### Brackets and Prioritization

When using the chaining functionality, be mindful of operation precedence. Operations are executed in the order they are chained, and there are no built-in brackets for prioritization. For instance, `bic.Add(10).Multiply(2)` will first add 10 and then multiply the result by 2, not multiply 10 by 2 and then add it.

To prioritize operations, you may need to chain them in the correct order, or separate them and use temporary variables.

### Supported Operations

- `Add`, `Subtract`, `Multiply`, `Divide`, `Mod` (Only in `BigIntChain`), `Pow`, `Abs`: Perform arithmetic operations.
- `Value()`: Retrieve the result of the chain.

### Error Handling

Errors occurring during an operation prevent the execution of subsequent operations in the chain. Always check for errors using `Value()` before using the result.

### Converting to Big Numbers

The package supports conversion from various types, including integers, unsigned integers, floating-point numbers, and their string representations, to `*big.Int` and `*big.Float`.

### Notes
- Ensure non-zero divisors for division and modulo.
- Be mindful of non-integer and negative exponents for power operations.
- Handle errors gracefully.

### License

This project is licensed under the MIT License. Refer to the LICENSE file for more details.

### Contribution

We welcome your contributions! Feel free to submit a Pull Request.

### Authors
- [Janko Simonovic](mailto:simonovic86@gmail.com)
