# rectpuzz-solver
*Library &amp; command line tool for solving 'the rectangle puzzle'. Written in Go!*

The `rectpuzz` library contains [Golang](https://golang.org) representations of `Cell`s 
and `Rect`s found in '[the rectangle puzzle](http://www.chiark.greenend.org.uk/~sgtatham/puzzles/js/rect.html)' 
introduced in CS-135.

The library contains a `SearchGraph` function that takes in a grid of unsigned integers (`[][]uint`), and 
solves the puzzle using an efficient linear graph search with backtracing. It is significantly faster than 
the original [Racket](https://racket-lang.org) implementation (<1 ms for a 20 x 20 grid).

The `main` package implements the function as a command line tool (WIP). 

---

### Usage
1. Download this repository as a zip, and place it in the `vendor` folder under your project root.
2. Use `import "github.com/steven-xie/rectpuzz-solver/rectpuzz` as needed.
