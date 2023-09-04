### Quick Start
Create toggle rules
```
$ cat test.tmp | rulecat toggle
T0
T0 T3
T0
```
Create character to rules
```
$ cat test.tmp | rulecat chars @
@T @h @i @s
@I @s @  @A
@T @e @s @t @1 @2 @3
```

### Creating Toggle Rules
Rulecat can be used to create toggle rules from `stdin`. This will convert
input into valid `Hashcat` rules and identify where toggles are.
```
Example: stdin | rulecat toggle [START-INDEX]
```

When the `toggle` option is used with a valid `START-INDEX` value the starting
index of the toggle rule can be changed.
```
$ cat test.tmp | rulecat insert 6
i6T i7h i8i i9s
i6I i7s i8  i9A
i6T i7e i8s i9t iA1 iB2 iC3
```

### Creating Character to Rules
Rulecat can be used to create custom rules per character from `stdin`. This
will take input and insert them before each character. This option has
multibyte character support.

This option can be used to create unique combinations like `@`, `!`, and `/`.
```
Example: stdin | rulecat chars [RULE]
```

When the `chars` option is used with a `RULE` input the text will be inserted
in front of each character. This can support any text.
```
$ cat test.tmp | rulecat chars @
@T @h @i @s
@I @s @  @A
@T @e @s @t @1 @2 @3
```
