# workspaces

NOT YET FULLY IMPLEMENTED and ABANDONED. I don't quite get the point of Go workspaces, especially because they don't
encapsulate modules from each other. It seems that all modules in the workspace are visible to each other in the
"module path". For example, I found that the Grafana codebase uses modules (i.e. a `go.work` file) but if I comment out
[this line](https://github.com/grafana/grafana/blob/f1b4964b24c1339935cda99498b89980934f1f50/pkg/apiserver/go.mod#L7), I
can still compile and test the `apiserver` code... It is able to resolve code in the `apimachinery` module. This is the
behavior I first found when writing my own example here and it was confusing so I went to see a real world example, like
in Grafana. It seems like Go workspaces are useful for the `replace` directive (neat, and this is implemented in other
toolchain's like Gradle's dependency substitution feature), but overall I'm not seeing the modularization. 

Modularizing code with Go workspaces.


## Overview

In other programming languages and toolchains, I'm used to authoring multi-module projects where some modules are
library-style modules are consumed from other application-style modules. The separation is useful because it helps you
design the library-style modules with a clear API and without accidentally importing some code (like a collections
utility) and creating a coupling to another library. This design also helps you eject the library at some future date
if you've wound up with something rock solid that can get more usage on its own (or even just copy/pasted).

Go 1.18 added support for "workspaces" which lets us design a project in this way.


## Instructions

Follow these instructions to build and run some modules that express inter-module dependencies on each other.

1. TODO Build the library module. Can you build it if other code in the workspace doesn't compile? That's kind of an
   important advantage in multi-module projects.
2. TODO Build and run the CLI/app module.
3. TODO analyze the dependencies. The lib module should have none, while the app has many.
