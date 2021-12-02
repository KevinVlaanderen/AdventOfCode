# AIVD Kerstpuzzel 2017

## Prerequisites

- Python 3.x
- Optional: Virtualenv

## Installation

- Check out the repo
- Create virtual environment, e.g. `virtualenv aivd2017`
- Activate virtual environment, e.g. `workon aivd2017`
- Install requirements: `pip install requirements.txt`
- Install package: `pip install --editable .`

## Running

- Activate virtual environment, e.g. `workon aivd2017`
- Run the app, e.g. `aivd solve`

### Available command-line options

```
Usage: aivd solve [OPTIONS] [PROBLEM] [STRATEGY]

Options:
  -c, --console    Output to the console
  -f, --file PATH  Output to a file
  --help           Show this message and exit.`
```

## How-to

### Define a new problem

New problems are defined by creating a new class derived from `Problem`, within a new sub-package of the `problems` package.

For each problem, you should set a title field, and implement a property `data` that loads the initial data passed to each strategy.

### Define a new strategy

New strategies are defined by creating a new class derived from `strategy`, within the problem directory.

For each strategy, you should set a title field, and implement the execute method.

## Utilìties

### Logging

Both problems and strategies have a logger field available to them (e.g. `self.logger`).
This is a basic Python logger instance (https://docs.python.org/3.5/library/logging.html).
You can (and should) use it to write data to the console and/or file, depending on the command line flags passed.