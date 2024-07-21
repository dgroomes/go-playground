# hello

A "hello world" Go program.


## Instructions

Follow these instructions to build and run a "Hello world" Go program:

1. Build and run the "hello world" program:
    * ```shell
      go run .
      ```
    * Altogether, it should look something like the following.
    * ```text
      Here is one of my first Go programs. It does the following:
          * It takes some quotes (sayings) from the "rsc.io/quote" package' 
          * Puts them in a slice
          * Marshals (serializes) the slice into a JSON document
          * Unmarshalls (deserializes) it back to a slice and prints it! 
      
      Here are the sayings in an indented JSON document:
      [
        "99 bottles of beer on the wall, 99 bottles of beer, ...",
        "Don't communicate by sharing memory, share memory by communicating.",
        "I can eat glass and it doesn't hurt me.",
        "If a program is too slow, it must have a loop."
      ]
      
      Here are the sayings as a string slice formatted by the 'fmt' package and the '%q' format verb:
      ["99 bottles of beer on the wall, 99 bottles of beer, ..." "Don't communicate by sharing memory, share memory by communicating." "I can eat glass and it doesn't hurt me." "If a program is too slow, it must have a loop."]
      ```
