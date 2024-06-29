# Code-Concat

Code-Concat is a Go application that recursively processes files in a directory, copies their content to the clipboard, and provides options for handling specific file types like SRT files.

## Repository Structure

```
directory-copier/
├── .gitignore
├── LICENSE
├── README.md
├── go.mod
├── go.sum
├── main.go
├── cmd/
│   └── codeconcat/
│       └── main.go
├── internal/
│   ├── processor/
│   │   └── processor.go
│   ├── filehandler/
│   │   └── filehandler.go
│   └── utils/
│       └── utils.go
└── scripts/
    ├── install.sh
    └── install.ps1
```

## Features

- Recursively process files in a specified directory
- Copy processed content to clipboard
- Handle different file types (text, Python, HTML, SRT)
- Exclude specified directories
- Skip files larger than a specified size
- Clean and group SRT files by speaker
- Respect .gitignore rules

## Installation

### Using Go

1. Ensure you have Go installed on your system (version 1.16 or later).
2. Clone this repository:
   ```
   git clone https://github.com/yourusername/directory-copier.git
   ```
3. Navigate to the project directory:
   ```
   cd directory-copier
   ```
4. Build the application:
   ```
   go build -o codeconcat cmd/codeconcat/main.go
   ```
5. Move the built binary to a location in your PATH.

### Using Installation Scripts

#### For Linux/macOS:

```bash
curl -sSL https://raw.githubusercontent.com/timsamart/code-concat/main/scripts/install.sh | bash
```

#### For Windows (PowerShell):

```powershell
Invoke-Expression (Invoke-WebRequest -Uri https://raw.githubusercontent.com/timsamart/code-concat/main/scripts/install.ps1 -UseBasicParsing).Content
```

## Usage

```
codeconcat [flags] <directoryPath>
```

### Flags:

- `--srt`: Process SRT files to clean timestamps and group by speaker
- `--size` or `-s`: Maximum file size in KB to process (default: 1024)
- `--exclude` or `-e`: Directories to exclude (comma-separated)

### Example:

```
codeconcat --srt --size 2048 --exclude node_modules,vendor /path/to/directory
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.# Code-Concat

Code-Concat is a Go application that recursively processes files in a directory, copies their content to the clipboard, and provides options for handling specific file types like SRT files.

## Repository Structure

```
directory-copier/
├── .gitignore
├── LICENSE
├── README.md
├── go.mod
├── go.sum
├── main.go
├── cmd/
│   └── codeconcat/
│       └── main.go
├── internal/
│   ├── processor/
│   │   └── processor.go
│   ├── filehandler/
│   │   └── filehandler.go
│   └── utils/
│       └── utils.go
└── scripts/
    ├── install.sh
    └── install.ps1
```

## Features

- Recursively process files in a specified directory
- Copy processed content to clipboard
- Handle different file types (text, Python, HTML, SRT)
- Exclude specified directories
- Skip files larger than a specified size
- Clean and group SRT files by speaker
- Respect .gitignore rules

## Installation

### Using Go

1. Ensure you have Go installed on your system (version 1.16 or later).
2. Clone this repository:
   ```
   git clone https://github.com/yourusername/directory-copier.git
   ```
3. Navigate to the project directory:
   ```
   cd directory-copier
   ```
4. Build the application:
   ```
   go build -o codeconcat cmd/codeconcat/main.go
   ```
5. Move the built binary to a location in your PATH.

### Using Installation Scripts

#### For Linux/macOS:

```bash
curl -sSL https://raw.githubusercontent.com/timsamart/code-concat/main/scripts/install.sh | bash
```

#### For Windows (PowerShell):

```powershell
Invoke-Expression (Invoke-WebRequest -Uri https://raw.githubusercontent.com/timsamart/code-concat/main/scripts/install.ps1 -UseBasicParsing).Content
```

## Usage

```
codeconcat [flags] <directoryPath>
```

### Flags:

- `--srt`: Process SRT files to clean timestamps and group by speaker
- `--size` or `-s`: Maximum file size in KB to process (default: 1024)
- `--exclude` or `-e`: Directories to exclude (comma-separated)

### Example:

```
codeconcat --srt --size 2048 --exclude node_modules,vendor /path/to/directory
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.