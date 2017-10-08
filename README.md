# apiholdit

[![License][License-Image]][License-Url]
[![ReportCard][ReportCard-Image]][ReportCard-Url]

## Table of contents

* [Introduction](https://github.com/repejota/apiholdit#introduction)
* [User Documentation](https://github.com/repejota/apiholdit#user-documentation)
	* [URL Parameters](https://github.com/repejota/apiholdit#url-parameters)
	* [Examples](https://github.com/repejota/apiholdit#examples)
* [Developer Documentation](https://github.com/repejota/apiholdit#developer-documentation)
* [Continuous Integration](https://github.com/repejota/apiholdit#continuous-integration)
  * [Tests](https://github.com/repejota/apiholdit#license)
  * [Coverage](https://github.com/repejota/apiholdit#coverage)
* [License](https://github.com/repejota/apiholdit#license)

## Introduction

APIholdit is a drop-in and easy to deploy microservice who exposes an HTTP API you can use to generate placeholder images with a desired text, automatically centered and scaled.

For example:

![psh screenshot](https://github.com/repejota/apiholdit/raw/develop/placeholder-example-1.png "placehoder example")

## User Documentation

The main entry point for the REST API is the following URL:

```
http://<host>:<port>/i
```

### URL Parameters

This URL accept the following parameters:

#### Width

`width` of the placeholder image in pixels. 

> By default is `640` pixels.

#### Height

`height` of the placeholder image in pixels. 

> By default is `480` pixels.

#### Text

`text` defines the placeholder text that will appear autoscaled and centered on the placeholdwr image.

This text should be encoded to allow special characters.

> By default is the string `Placeholder`.

#### Background color

`bgcolor` sets the background color to be used to render the placeholder image. 

> By default is the color called *Silver* from the [FlatUIColors](http://flatuicolors.com/) color palette. And its code is `#bdc3c7`.

#### Foreground color

`fgcolor` sets the foreground color, or the color being used to draw the text on the placeholder image. 

> By default is the color called *Midnight Blue* from the [FlatUIColors](http://flatuicolors.com/) color palette. And its code is `#2c3e50`.

### Examples

Here are a few eamples using the previous defined API:


| URL           | Placeholdwer           |
| ------------- |:-------------:|
| http://example.com/i?width=320&height=200      | ![psh screenshot](https://github.com/repejota/apiholdit/raw/develop/placeholder-example-2.png "placehoder example")  |
| http://example.com/i?width=320&height=200&text=placeholder-text      | ![psh screenshot](https://github.com/repejota/apiholdit/raw/develop/placeholder-example-3.png "placehoder example")      |
| http://example.com/i?width=320&height=200&text=placeholder-text&bgcolor=c0392b&fgcolor=2980b9 | ![psh screenshot](https://github.com/repejota/apiholdit/raw/develop/placeholder-example-4.png "placehoder example")      |

## Developer Documentation

### Deployment

You can also deploy *APIHoldit* on Heroku with this one click button:

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

## Continuous integration

### Tests

* Develop: [![CircleCI](https://circleci.com/gh/repejota/apiholdit/tree/develop.svg?style=svg)](https://circleci.com/gh/repejota/apiholdit/tree/develop)
* Master: [![CircleCI](https://circleci.com/gh/repejota/apiholdit/tree/master.svg?style=svg)](https://circleci.com/gh/repejota/apiholdit/tree/master)

### Coverage

* Develop: [![Coverage Status](https://coveralls.io/repos/github/repejota/apiholdit/badge.svg?branch=develop)](https://coveralls.io/github/repejota/apiholdit?branch=develop)
* Master: [![Coverage Status](https://coveralls.io/repos/github/repejota/apiholdit/badge.svg?branch=master)](https://coveralls.io/github/repejota/apiholdit?branch=master)


## License

(The MIT License)

Copyright (c) 2017 apiholdit Authors.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to
deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
sell copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.

[License-Url]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/badge/License-MIT-blue.svg
[ReportCard-Url]: http://goreportcard.com/report/repejota/apiholdit
[ReportCard-Image]: http://goreportcard.com/badge/github.com/repejota/apiholdit