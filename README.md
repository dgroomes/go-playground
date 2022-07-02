# go-playground

📚 Learning and experimenting with the [Go programming language](https://golang.org/).


## Instructions

Follow these instructions to build and run a "Hello world" Go program:

1. Build and run the program:
   * ```shell
     go run .
     ```


## Learning Strategy

These are the components I'm using to guide my learning:

* [x] Set up instructions
    * I have my own installation instructions, including Bash autocompletion, in the section named "Install Go" in my
      personal [my-config](https://github.com/dgroomes/my-config/blob/main/mac-os/MACOS_SETUP.md).
* [x] [GoLand IDE](https://www.jetbrains.com/go/)
    * Make sure to enable Go modules integration in GoLand. See the below screenshot for the setting:
      ![goland-enable-go-modules.png](goland-enable-go-modules.png)
* [ ] [Official Go docs: *Tutorial: Get started with Go*](https://golang.org/doc/tutorial/getting-started)
    * IN PROGRESS
* [ ] [Official Go docs: *
  Packages The Project Help Blog Play How to Write Go Code (with GOPATH)*](https://golang.org/doc/gopath_code)


## Observations about Go

These are my observations about Go the programming language, the surrounding toolchain, the ecosystem of Go libraries,
and the Go community.

* It's really cool that you can add an import statement and then do `go mod tidy` to add the module as a requirement.
  That's an interesting inversion. Normally (in environments like Java) it's up to you to find the dependency and add it
  to the project model file (`build.gradle` for Gradle, `pom.xml` for Maven).
* I like the omission of commas in the import statement. When the compiler can just as easily make sense of a grammar
  without commas, in a certain context, then why require commas? Bash is comma-less, for example! I bet it's hard on the
  compiler and the IDE and will result in harder to read compiler error messages.
* GoLand doesn't have an [intention action](https://www.jetbrains.com/help/idea/intention-actions.html) for converting a
  string to a raw string. I'm not sure why. Do raw strings work much differently than raw/multiline strings in other
  languages? I use this all the time in Intellij on Kotlin and JavaScript code.
* I really like that I can name a local variable `_` so that the Go compiler won't complain about an unused variable.
* I'm not a huge fan of having a `GOPATH`. This makes me put all my Go projects in a specific directory (on
  the `GOPATH`). In my case, it's at `~/repos/go` but what I prefer is to put my repos in `~/repos/personal`
  or `~/repos/opensource`. This is a convention that I use to great effect and I take with me as I use different
  computers. The Go toolchain however is forcing its own directory constraint on me. So now I have to invent new
  conventions. Should I make `~/repos/go/personal`
  and `~/repos/go/opensource`? Or should I symlink repos from `~/repos/go` into my conventional `~/repos/personal`
  and `~/repos/opensource`?
* I don't have to declare variables with a `var` keyword? That's pretty cool. Note that you need `var` I think only whe
  you don't assign the variable in the same statement.
* The code samples in the docs at <https://pkg.go.dev> is amazing. Wow! You can even edit and run the code samples right there!


## Wish list

General clean ups, TODOs and things I wish to implement for this project:

* [x] DONE Print the JSON document. I'd like to print the JSON string in the program output. This gives
  better contrast and proves that the program actually does something with JSON.
* [ ] Figure out how to modularize Go code. I often struggle with the GOPATH.


## Reference

* <https://pkg.go.dev/>
    * What do I call this website, the "Go package site"?
* [Official blog post: *Tidying up the Go web experience*](https://go.dev/blog/tidy-web)
    * The marketing for Go is great! I wish Java had something similar.
* [The `fmt` package docs](https://pkg.go.dev/fmt)
    * Read this to understand the format verbs like `%s` and `%q`.  