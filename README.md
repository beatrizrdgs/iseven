# Is Even

`IsEven` is a GoLang package that provides a simple function to determine whether a given integer is even or not.



## Installation

To install `IsEven`, you can use go get:

```bash
go get -u github.com/beatrizrdgs/iseven
```



## Usage

Simply invoke the `IsEven` function to check if the integer is Even.

```go
package main

import (
    "fmt"
    "github.com/beatrizrdgs/iseven"
)

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    fmt.Println("Checking if numbers are even:")
    for _, number := range numbers {
        if iseven.IsEven(number) {
            fmt.Printf("%d is even\n", number)
        } else {
            fmt.Printf("%d is not even\n", number)
        }
    }
}
```



## Documentation

The package contains only one function:

### func [IsEven](iseven.go)
`func IsEven(n int) bool`

IsEven takes an integer n as input and returns true if the number is even, otherwise it returns false.



## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.