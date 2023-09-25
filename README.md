# Nixl

Nixl is a versatile command-line utility crafted in Go, designed to streamline a range of developer tasks. It simplifies common operations, making them accessible through an intuitive command-line interface.

## Features

- **Hashing**: Compute hashes using popular algorithms like MD5, SHA-1, SHA-256, and SHA-512.
- **Checksums**: Generate and verify checksums for files with ease.
- **Base64 Encoding/Decoding**: Encode text to base64 or decode base64-encoded data effortlessly.

## Usage
### hash command
The 'hash' command allows you to generate hash values for input text using various algorithms.

Supported algorithms:
1. md5: Calculate an MD5 hash (e.g., 'nixl hash md5 "hello world"')
2. sha256: Calculate a SHA-256 hash (e.g., 'nixl hash [text]'
3. sha512: Calculate a SHA-512 hash (e.g., 'nixl hash sha512 [text]')

## Benefits

- **Simplify Development**: Nixl provides essential tools for developers in one lightweight package, simplifying common tasks.
- **Streamlined Workflow**: Whether you're working on data manipulation, encoding, hashing, or other developer operations, Nixl offers a unified and user-friendly interface to streamline your workflow.
- **Versatility**: Nixl is designed to be a go-to companion for managing and enhancing your development processes, offering a wide range of features to meet your needs.

## Credits

Nixl is built using the following open-source libraries and technologies:

- [Cobra](https://github.com/spf13/cobra): Cobra is a popular Go library for creating powerful command-line applications.
- [Go Programming Language](https://golang.org/): Nixl is crafted in Go, a modern and efficient programming language.

We extend our gratitude to the contributors and maintainers of these projects for their valuable work and contributions to the developer community.

## Note

Nixl is a powerful tool, but nothing is 100% perfect. If you encounter any issues, have feature requests, or want to contribute, please feel free to open a pull request and recommend improvements. Your feedback is valuable in making Nixl even better for the developer community.

