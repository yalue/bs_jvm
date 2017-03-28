Another Go-based JVM Implementation
===================================

About
-----

This project aims to implement a function JVM in the Go programming language.
I am aware that such a project exists already, so my own effort here is only
for my personal benefit for now.

The JVM is not actually implemented yet. All that has been completed is parsing
class files and disassembly.

Usage
-----

For now, only disassembly works. Usage:

```bash
cd disassemble/
go build .

# The path to RandomDots.class can be replaced with a path to any valid class
# file.
./disassemble -filename ../class_file/test_data/RandomDots.class
```
