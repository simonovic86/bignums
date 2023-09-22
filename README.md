# :abacus: Bignums Library

The `bignums` library is a highly sophisticated toolkit, meticulously designed for developers to perform arithmetic computations on extensive numerical values with ease and precision. Developed in the versatile Go programming language, it specializes in handling big numeric types, namely `*big.Int` and `*big.Float`. It presents a fluent and cohesive API, enabling the chaining of operations seamlessly and allowing for sophisticated calculations with enhanced readability and maintainability.

## :star2: Key Features

- **:link: Fluent API**: Enables seamless chaining of multiple arithmetic operations, significantly enhancing code readability and maintainability and reducing cognitive load, offering a clean and efficient coding experience.

- **:1234: Support for Various Types**: Expertly manages `*big.Int` and `*big.Float`, along with standard numeric types and their string representations, ensuring comprehensive adaptability and versatility in dealing with numeric representations.

- **:shield: Advanced Error Handling**: Provides robust error management for scenarios such as division by zero and mismatched brackets, with detailed error messages, ensuring the resilience and reliability of the computation processes.

- **:fast_forward: Operation Prioritization**: Utilizes brackets to effectively prioritize and accurately execute operations, allowing for the execution of more intricate calculations and offering a high level of control over operation order.

## :hammer_and_wrench: Installation

```shell
go get github.com/simonovic86/bignums
```

## :computer: Usage

### :building_construction: Creating a Chain
Initialize new chains using `NewBigIntChain` or `NewBigFloatChain`:

```go
intChain := bignums.NewBigIntChain(10)
floatChain := bignums.NewBigFloatChain(10.5)
```

### :gear: Performing Operations
Leverage the fluent API to execute a series of operations:

```go
resultIntChain := intChain.Add(5).Subtract(3).Multiply(2).Divide(3)
resultFloatChain := floatChain.Add(5.5).Subtract(3.5).Multiply(2.5).Divide(3.5)
```

### :arrow_forward: Prioritization using Brackets
Employ `Begin()` and `End()` to assign operation priority:

```go
resultChain := intChain.Add(10).Begin().Add(20).End().Multiply(2) // ((10+20)*2)
```

### :mag: Fetching Results
Retrieve the final results post-operation:

```go
value, err := resultChain.Value()
if err != nil {
log.Fatalf("Error occurred: %v", err)
}
fmt.Println("Result: ", value)
```

## :book: Examples

For a better understanding, here are some practical examples demonstrating the versatile applications of `bignums`:

```go
package main

import (
"fmt"
"github.com/simonovic86/bignums"
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

## :page_facing_up: Documentation
Refer to the [GoDoc](https://pkg.go.dev/github/simonovic86/bignums) for an in-depth overview and detailed insights into each function and feature.

## :handshake: Contributing
We warmly welcome contributions! Feel free to open a pull request for minor changes. For discussions on major changes or feature requests, kindly open an issue first.

## :balance_scale: License
This project is licensed under the MIT License. For more details, please see the [LICENSE](LICENSE) file.
