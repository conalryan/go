Variables
================================================================================
https://www.youtube.com/watch?v=YS4e4q9oBaU&t=3456s
36:32

- Variable declaration
- Redeclaration and shadowing
  - Inner most scope will override outer scope
- Visibility
  - Global: Uppercase variable name is exposed to global scope
  - Package: Lowercase variable name is scoped to package and is only available to files in the same package
  - Block: Variable declared in func block is scope to that func block
- Naming conventions
  - Short var names represent short lived variables (e.g. counters)
  - Exported variable names should be descriptive and consise
  - Acronyms should be uppercase (e.g. theURL, someHTTP, etc)
  - Pascal case and camel case
- Type convertions
  - `float32(someInt)`
  - `import ( "strconv")`
  - `strconv.Itoa(someInt)`


