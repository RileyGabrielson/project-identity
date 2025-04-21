![banner](./identity-banner.png)

# The Identity

A large, unnecessary identity function written in go.

## Why?

Because it's fun to do something unecessary in a creative way :) You've written enough code that is "useful".

Inspired by [The Art of Code](https://www.youtube.com/watch?v=6avJHaC3C2U).

## Constraints

- Input range: -1000 to 1000 (inclusive)

## Contributing

1. Create a new module in the `modules` folder:
   ```go
   // modules/yourmodule/yourmodule.go
   package yourmodule

   func Identity(num int) int {
       // Your unnecessarily complex code here
       return num
   }
   ```

2. Add your module to the identity chain in `main.go`:
   ```go
   func Identity(num int) int {
       // ...
       result = yourmodule.Identity(result) // Add your module here
       return result
   }
   ```

3. Keep your module self-contained

4. Test your changes:
   - Run `go test -v` to ensure your module maintains the identity property
   - Add new test cases if your module has specific edge cases
