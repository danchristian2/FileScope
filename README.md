# FileScope

**FileScope** is a fast, lightweight command-line tool for analyzing software projects.

It scans a project directory and gives you useful information about the files inside it, including the number of files, directories, total storage size, and lines of code.

The goal is to make it easy to understand the structure and size of a project directly from the terminal.

## 🚀 Current Features

FileScope currently supports:

* 📁 Recursive directory scanning
* 📄 File counting
* 📂 Directory counting
* 💾 Total project size
* 📝 Total number of lines
* 📊 Simple terminal-based project statistics

Example:

```text
==== FileScope ====

Directory: .
Files: 42
Directories: 8
Total size: 2.31 MB
Total lines: 4,827
```

## 🛠️ Built With

* **Go**
* Go standard library
* Command-line interface

FileScope is designed to have minimal dependencies and remain fast and lightweight.

## 📦 Installation

### From source

Make sure you have Go installed.

Clone the repository:

```bash
git clone https://github.com/danchristian2/filescope.git
```

Enter the project:

```bash
cd filescope
```

Run FileScope:

```bash
go run . .
```

## 💻 Usage

FileScope accepts a directory as an argument.

### Analyze the current directory

```bash
filescope .
```

### Analyze another directory

```bash
filescope ./my-project
```

### Analyze a project using the development version

If you're running directly from the source:

```bash
go run . ./my-project
```

## 🧠 How It Works

FileScope recursively walks through the selected directory.

For every file it finds, it can collect information such as:

```text
Project
   │
   ├── Files
   │
   ├── Directories
   │
   ├── File sizes
   │
   └── Lines
```

The collected information is then summarized and displayed in the terminal.

## 🗺️ Roadmap

FileScope is actively being developed.

Planned features include:

* [ ] File extension/language statistics
* [ ] Code line counting
* [ ] Blank line counting
* [ ] Comment line counting
* [ ] Largest files
* [ ] Largest directories
* [ ] File-type breakdown
* [ ] Duplicate file detection
* [ ] Git repository analysis
* [ ] Exclude directories such as `node_modules` and `.git`
* [ ] JSON output
* [ ] CSV output
* [ ] Custom analysis options
* [ ] Improved terminal interface
* [ ] Windows package distribution
* [ ] Linux package distribution

## 🎯 Why FileScope?

Software projects can become difficult to understand as they grow.

FileScope aims to provide a quick answer to questions such as:

* How many files are in this project?
* How many directories does it contain?
* How much storage does it use?
* How many lines does it contain?
* Which parts of the project are taking the most space?
* What types of files make up the project?

Instead of manually checking folders, FileScope brings this information together in one command.

## 🔒 Privacy

FileScope is designed with a **local-first approach**.

Your project is analyzed locally on your computer. FileScope does not need to upload your source code to a remote server to perform its analysis.

## 📚 Project Status

FileScope is currently under active development.

The project is also being used as a practical way to learn Go, software architecture, file systems, and command-line application development.

## 🤝 Contributing

Contributions, ideas, bug reports, and improvements are welcome.

If you find a problem or have an idea for a feature, feel free to open an issue or submit a pull request.

## 📄 License

License information will be added as the project approaches its first public release.
