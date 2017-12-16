# rectpuzz-solver
*Library &amp; command line tool for solving 'the rectangle puzzle'. Written in Go!*

The `rectpuzz` library contains [Golang](https://golang.org) representations of `Cell`s 
and `Rect`s found in '[the rectangle puzzle](http://www.chiark.greenend.org.uk/~sgtatham/puzzles/js/rect.html)' 
introduced in CS-135.

The library contains a `SearchGraph` function that takes in a grid of unsigned integers (`[][]uint`), and 
solves the puzzle using an efficient linear graph search with backtracing. It is significantly faster than 
the original [Racket](https://racket-lang.org) implementation (<1 ms for a 20 x 20 grid).

The `main` package implements the function as a command line tool (WIP). 

**Check out the [video demo](https://youtu.be/l_vsUksOWCE)!**

---

## Usage
### Using the library
1. Download this repository as a zip, and place it in the `vendor` folder under your project root.
2. Use `import "github.com/steven-xie/rectpuzz-solver/rectpuzz` as needed.

### Running the executable binary
Head over to [the releases](https://github.com/steven-xie/rectpuzz-solver/releases) and grab the latest version. Make sure you get the one that matches your computer's architecture!

#### A Note for Linux/Mac users

If you're running on Mac/Linux, you may need to give the binary executable permissions using the terminal, like this:
``` bash
chmod +x solver-mac-x64
```

Then, run the executable either by double clicking it, or by typing `./solver-mac-x64` in your terminal.

*Make sure that you're in the same directory as the binary when you do this!*

