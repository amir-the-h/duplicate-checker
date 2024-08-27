# Duplicate Checker

## Description

The Go Duplicate Checker is a command-line application that helps you identify and manage duplicate files in a given directory. It uses the MD5 checksum algorithm to compare the contents of files and determine if they are duplicates. 

With the Go Duplicate Checker, you have the option to remove duplicate files based on your preference. You can choose to keep either the oldest or newest file, depending on your needs. 

This tool provides a convenient way to declutter your file system and free up storage space by eliminating unnecessary duplicate files. 


## Installation

To use the latest release, you can download it from the [latest release link](https://github.com/amir-the-h/duplicate-checker/releases/latest).

To clone the repository and build the app locally, follow these steps:

1. Clone the repository:
  ```shell
  git clone https://github.com/amir-the-h/duplicate-checker.git
  ```

2. Navigate to the project directory:
  ```shell
  cd duplicate-checker
  ```

3. Install the dependencies:
  ```shell
  go mod tidy
  ```

4. Build the app:
  ```shell
  go build -o dup-checker cmd/main.go
  ```

5. Move the built binary to the desired location (optional):
  ```shell
  sudo mv dup-checker /usr/local/bin
  ```

## Usage

To run the app locally, use the following command:

```shell
dup-checker [-d=<directory | default: PWD>] [-r] [-o]
```

- `-d`: The path to the directory to scan for duplicates. If not provided, the current working directory will be used.
- `-r`: Whether to remove the duplicates. If provided, the duplicates will be removed.
- `-o`: Whether to keep the oldest file. If provided, the oldest file will be kept instead of the newest one.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.