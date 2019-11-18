# simpleansi

Simple ANSI is a package that contains simple ANSI operations using escape sequences. It was developed to be used in the [Pac Go](https://github.com/danicat/pacgo) project, but feel free to use it anywhere you would like :)

## Introduction

This package contain helper functions for some common terminal operations. One important difference in regards to the ANSI standard is that we consider the top left corner to be zero-based (0,0) instead of 1-based (1,1). The only reason for that is that it simplifies the code for drawing when you are iterating over an array of slices, as it is used in Pac Go.

## Contributing

This project is open source under the MIT license, which means you are basically free to do whatever you want with it, just give me the proper credits. :)

If you want to contribute just raise an issue and/or submit a pull request.

If you are looking for inspiration you may browse the open issues or have a look at the [TODO](TODO.md) list. 

Everything on the TODO list should be planned as a new step on the tutorial unless otherwise noted.

## License

See [LICENSE](LICENSE).

## Contacting the Author

If you have any questions, please reach out to me at [daniela.petruzalek@gmail.com](mailto:daniela.petruzalek@gmail.com). I'm also on Twitter as [@danicat83](https://twitter.com/danicat83).

Finally, you can support me on Ko-Fi if you like my work and want me to continue doing crazy stuff like this project:

[![ko-fi](https://www.ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/O5O716D82)