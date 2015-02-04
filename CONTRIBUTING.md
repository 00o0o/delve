# Contributing to Delve

Want to help contribute to Delve? Great! Any and all help is certainly appreciated, whether it's code, documentation, or spelling corrections.

## Filing issues

When filing an issue, make sure to answer these five questions:

1. What version of Delve are you using (`dlv -v`)?
2. What operating system and processor architecture are you using?
3. What did you do?
4. What did you expect to see?
5. What did you see instead?

## Contributing code

Fork this repo and create your own feature branch. Install all dependencies as documented in the README.

### Guidelines

Consider the following guidelines when preparing to submit a patch:

* Follow standard Go conventions (document any new exported types, funcs, etc.., ensuring proper punctuation).
* Ensure that you test your code. Any patches sent in for new / fixed functionality must include tests in order to be merged into master.
* If you plan on making any major changes, create an issue before sending a patch. This will allow for proper discussion beforehand.
* Keep any os / arch specific code contained to os / arch specific files. Delve leverages Go's filename based conditional compilation, i.e do not put Linux specific functionality in a non Linux specific file.


