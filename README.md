BS-JVM
======

About
-----

The BS ("Blinding Speed") Java Virtual Machine is a work-in-progress JVM
implementation in the Go programming language.

I'm just occasionally working on this for fun; don't expect anything big
any time soon.

Usage
-----

This still barely runs anything, due to some unimplemented instructions,
little testing, and basically no standard library.

For the most part, nothing really "works" yet, but it is capable of running a
simple test. To use the JVM:
```bash
# First, build the JVM executable
cd jvm/
go build .
cd ../class_file/test_data

# These two commands should produce identical output.
../../jvm/jvm RandomDotsSimple.class
java RandomDotsSimple
```

(The extremely basic `RandomDotsSimple` test doesn't rely on any standard
library functionality apart from `System.out.print(char)`. See the source code
in the same directory.)

There is also a basic disassembler:
```bash
cd disassemble/
go build .

# This path can be replaced with a path to any valid class file.
./disassemble -filename ../class_file/test_data/RandomDotsSimple.class
```

At the moment, the JVM doesn't try to load classes outside of whichever
standard classes are built in; trying to disassemble or run files depending on
separate class files will encounter errors.

