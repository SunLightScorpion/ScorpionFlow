# ScorpionFlow

A unified application for managing your Linux workflow with an elegant GUI built on top of Go and Fyne.

## Features

- **Cross-Platform GUI**: Modern graphical user interface powered by Fyne
- **Linux Automation**: Simplify your Linux system workflow
- **Easy Installation**: Available as Debian package (.deb)
- **Lightweight**: Built with Go for fast performance and minimal resource usage

## Requirements

- **Go**: 1.26 or higher (for building from source)
- **Linux**: Ubuntu, Debian, or other Linux distributions
- **System Libraries**: GTK 3+ (for Fyne GUI rendering)

### For Building Debian Packages

- `dpkg-deb` (usually included in `dpkg`)
- Standard build tools

## Installation

### From Binary (Debian/Ubuntu)

```bash
# Build the Debian package
./build-deb.sh

# Install the generated package
sudo dpkg -i dist/scorpionflow_x.x.x_amd64.deb

# Run ScorpionFlow
scorpionflow
```

### From Source

```bash
# Clone the repository
git clone https://github.com/yourusername/scorpionflow.git
cd scorpionflow

# Install dependencies
go mod download

# Build the binary
go build -o scorpionflow ./cmd/scorpionflow

# Run
./scorpionflow
```

## Project Structure

```
ScorpionFlow/
├── cmd/
│   └── scorpionflow/
│       ├── core.go          # Main application logic
│       └── assets/          # Application icons and resources
├── go.mod                    # Go module definition
├── build-deb.sh             # Debian package build script
└── README.md               # This file
```

## Development

### Building

```bash
# Compile the Go binary
go build -o scorpionflow ./cmd/scorpionflow
```

### Building a Debian Package

```bash
./build-deb.sh
```

This script will:
1. Build the Go binary
2. Create the package structure
3. Generate the control file
4. Create a desktop entry for the application menu
5. Package everything as a `.deb` file

### Dependencies

Main dependencies are managed via `go.mod`:

- **fyne.io/fyne/v2** - GUI framework for the graphical interface
- Other transitive dependencies (see `go.mod` for complete list)

Install all dependencies with:

```bash
go mod download
```

## Usage

After installation, you can launch ScorpionFlow:

```bash
# From command line
scorpionflow

# Or from your application menu (if installed via .deb)
# Look for "ScorpionFlow" in your desktop environment's application menu
```

## Building and Packaging

The `build-deb.sh` script automates the creation of a Debian package:

- **Current Version**: 0.1.5
- **Architecture**: amd64
- **Output**: `dist/scorpionflow_0.1.5_amd64.deb`

To update the version, edit the `VERSION` variable in `build-deb.sh`.

## Troubleshooting

### Missing GTK Libraries

If you get GTK-related errors when running:

```bash
# Ubuntu/Debian
sudo apt-get install libgtk-3-0 libglib2.0-0
```

### Icon Not Loading

ScorpionFlow looks for icons in multiple locations:
- `cmd/scorpionflow/assets/icon.png`
- `assets/icon.png`
- `/usr/share/icons/hicolor/256x256/apps/scorpionflow.png`

Ensure your icon file is placed in one of these locations.

## License

Please add your license information here (e.g., MIT, GPL-3.0, Apache-2.0)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Maintainer

- **NightDev701**

## Support

For issues, questions, or suggestions, please open an issue on GitHub.
