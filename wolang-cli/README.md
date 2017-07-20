## The Workshop Task

run `go get github.com/wwgberlin/wwgb-wolang/wolang-cli`
to have the code locally.

Inside the `wolang-cli` directory run `go run main.go` to start the wolang interpreter console (provided to us by Talon.One) 

When the console starts for you, try the following phrases:

```
wolang> (+ 1 2 3 4)
10
wolang> (concat abc defg)
abcdefg
wolang> (> 2 1)
method not defined
```

Inside `main.go` you will find 3 functions that have not been defined yet: largerThan, equals and smallerThan.

Your task is to implement them. You will know you implemented them correctly when you have run go test, and not get any errors.

Then open the interpreter again and run:
```
wolang> (> 2 1)
true
wolang> (== 1 1)
true
wolang> (< 1 2)
true
```

## Good luck!
