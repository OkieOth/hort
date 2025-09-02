![WIP](https://img.shields.io/badge/work%20in%20progress-red)

# json.schema.parser
small library to parse json schema into custom types

# Additions supported by the parser
* x-ref: marks a foreign key relation to another object. I follows the same syntax as '$ref'
* x-tag: list of tags/annotations you can put as strings on attributes and types.

# What's not supported yet
* $ref links to other files
* allOf for complex types
* anyOf, oneOf terms
