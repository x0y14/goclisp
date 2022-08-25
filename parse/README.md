# parse

```text

program = stmt*
stmt = "(" opIdent stmt* ")"
     |  unary

unary   = ("+" | "-")? primary

primary = ident
        | string
        | number
        | nil
        | true

opIdent = ident
        | "+"
        | "-"
        | "*"
        | "/"
        | "="
        | "/="
        | "<"
        | "<="

// atom
ident  = [a-zA-Z_]+[a-zA-Z0-9_]*
string = "\"" .* "\""

number = float | int
float  = [0-9]+[0-9.]*[0-9]*
int    = [0-9]+

nil    = "nil"
true   = "T"
```
