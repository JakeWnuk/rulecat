`Rulecat` is a multi-tool for working with text streams for password cracking.

Rulecat (`cat` rule) focuses on the conversion of text to rules or specialized
output and features several functions:

- Creates append rules from `stdin`
- Creates prepend rules from `stdin`
- Creates blank lines from `stdin`
- Create the cartesian product of a file and `stdin`
- Creates custom rules per character from `stdin`
- Creates insert rules from `stdin`
- Creates overwrite rules from `stdin`
- Creates toggle rules from `stdin`
- Creates custom rehashing of text from `stdin`
- Creates URL, HTML, & ASCII escape encoded text from `stdin`

Rulecat fits into a small tool ecosystem for password cracking and is designed for lightweight and easy usage with its companion tools:

- [maskcat](https://github.com/JakeWnuk/maskcat)
- [rulecat](https://github.com/JakeWnuk/rulecat)
- [mode](https://github.com/JakeWnuk/mode)

### Getting Started

Usage information and other documentation can be found below:

- Usage documentation:
    - [Append and Prepend Rules](https://github.com/JakeWnuk/rulecat/blob/main/docs/APPEND_AND_PREPEND.md)
    - [Insert and Overwrite Rules](https://github.com/JakeWnuk/rulecat/blob/main/docs/INSERT_AND_OVERWRITE.md)
    - [Toggle and Character to Rules](https://github.com/JakeWnuk/rulecat/blob/main/docs/TOGGLE_AND_CHARACTER.md)
    - [Cartesian and Custom Hashing Expressions](https://github.com/JakeWnuk/rulecat/blob/main/docs/CARTESIAN_AND_CUSTOM.md)
    - [Blank Lines and Encoding Text](https://github.com/JakeWnuk/rulecat/blob/main/docs/BLANK_AND_ENCODING.md)

- For more application examples: 
    - [Rulecat Examples](https://jakewnuk.com/posts/brewing-hash-cracking-resources-w-the-twin-cats/) (external link)
    - [Rulecat Usages](https://jakewnuk.com/posts/how-to-use-rulecat-to-crack-the-perfect-eggs-every-time/) (external link)

### Install from Source
```
git clone https://github.com/JakeWnuk/rulecat && cd rulecat && go build ./cmd/rulecat && mv ./rulecat ~/go/bin/
```

### Current Version 1.2.0:
```
Modes for rulecat (version 1.2.0):

  append        Creates append rules from text
                Example: stdin | rulecat append
                Example: stdin | rulecat append remove
                Example: stdin | rulecat append shift

  prepend       Creates prepend rules from text
                Example: stdin | rulecat prepend
                Example: stdin | rulecat prepend remove
                Example: stdin | rulecat prepend shift

  blank         Creates blank lines from text
                Example: stdin | rulecat blank

  [RULE-FILE]   Create cartesian product of a file and text
                Example: stdin | rulecat [FILE]

  chars         Creates custom rules per character from text
                Example: stdin | rulecat chars [RULE]

  insert        Creates insert rules from from text
                Example: stdin | rulecat insert [START-INDEX]

  overwrite     Creates overwrite rules from from text
                Example: stdin | rulecat overwrite [START-INDEX]

  toggle        Creates toggle rules from from text
                Example: stdin | rulecat toggle [START-INDEX]

  custom        Creates custom text rehashing from an expression
                Example: stdin | rulecat custom [EXPRESSION]
                Example: stdin | rulecat custom 2xmd5(sha1(p))

  dehex         Dehexes $HEX[...] input to standard out
                Example: stdin | rulecat dehex

  encode        URL, HTML, and ASCII escape encodes input and prints new output
                Example: stdin | rulecat encode
```
