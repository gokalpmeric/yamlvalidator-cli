# YAML Validator CLI Tool

## Introduction
YAML Validator is a command-line interface (CLI) tool built in Go, using the Cobra library. It's designed to validate the syntax of YAML files, offering a simple and efficient way to ensure YAML files are error-free and adhere to the correct format.

## Features
- Simple and user-friendly command-line interface.
- Fast and accurate validation of YAML files.
- Built using Go and the Cobra library for robust performance.

## Installation

### Prerequisites
- Go (version 1.13 or later)
- Git (for cloning the repository)

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/yamlvalidator.git
2. Navigate to the project directory:
   ```bash
   cd yamlvalidator
3. Run the install script:
   ```bash
    ./install.sh

### Usage
To validate a YAML file, run the following command:
   ```bash
    yamlvalidator -v path/to/yourfile.yaml
```
or
   ```bash
    yamlvalidator validate path/to/yourfile.yaml
```
This command checks the specified YAML file for syntax correctness.

### Building from Source
You can also build the tool from the source code:

   ```bash
    go build
```
This command creates an executable in your current directory.





