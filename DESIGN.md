## Compile to bytecode
Code should be compiled to bytecode, seperate bytecode interpreter from compiler? This would allow you to add bytecode binaries with the overhead of compiling and preprocessing.

## Preprocessor
Does having a preprocessor make sense? Advantages:
1. Determine stack size for chunks
2. Deduplicate constants
3. Resolve variable names to indices
4. Bytecode verification
2. Macros?
