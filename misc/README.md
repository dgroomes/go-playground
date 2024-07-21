# misc

Exploring miscellaneous features and concepts of Go.


## Instructions

Follow these instructions to build and run the exploratory programs:

1. Build and run the "meta" program:
    * ```shell
      go run cmd/meta/meta.go
       ```
2. Build and run the "file" program:
    * ```shell
      go run cmd/file/file.go
      ```
3. Build and run the "exec-subprocess" program:
    * ```shell
      go run cmd/exec-subprocess/exec-subprocess.go
      ```
    * It should output something like the following.
    * ```text
      The 'ls' command found the following files in the current working directory:
      README.md
      cmd
      
      The 'docker' command reports the following version:
      27.0.3
      ```
4. Build and run the "exec-process-image-replacement" program:
    * ```shell
      go run cmd/exec-process-image-replacement/exec-process-image-replacement.go
      ```
    * It should output something like the following.
    * ```text
      Hello from a Go program. This process will be 'morphed' from a Go program into an execution of the 'ls' program. The process stays the same, but we can consider that 'image' (memory) is replaced.
      total 8
      -rw-r--r--  1 dave  staff   1.5K Jul 21 12:33 README.md
      drwxr-xr-x@ 6 dave  staff   192B Jun  3 19:56 cmd      
      ```
