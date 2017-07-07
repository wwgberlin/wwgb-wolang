# wwgb-wolang
A demo project for the Women Who Go Berlin presentation @Talon.One HQ.

Steps:

Step 1
 - Create plus and strconcat functions
 - Create eval and evalFCall functions
 - Create a main() function to run examples for eval

Step 2
- Write extfuncdef struct type with name and implementation
- Write callable type interface and functions Name and Call
- Create a simple extfuncdef and edit evalFCall to include extfuncdef's

Step 3
- Split the functions into different files. Somthing like:
    - builtInFunctions: plus, strconcat
    - extension: extfuncdef, callable, name, call
    - wolang: eval, evalFCall
- Be sure all the files have `package wolang` on the first line
- Make the necessary functions public by capitalizing the func name, such as eval -> Eval
- Create another package wolang-cli and move main (into a main.go file) there with wolang imported
- `go run main.go` to check that everything works

Step 4
- Make it such that extfuncdef's can be added with Eval
    - Make necessary functions public
    - Write RegExtFunc

Step 5
- Change Eval and/or EvalFCall such that a nested expression can be evaluated
    - example: `nestedExpr := []interface{}{"+", 3, 4, []interface{}{"+", 5, 6}}`