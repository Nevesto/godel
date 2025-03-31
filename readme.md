# Godel (1.0.0)

![PREVIEW](/public/PREVIEW.png)

## About

Godel is a command line tool to easily clear Discord accounts.

This project is under active development. For any bugs suport or bugs contact open a new issue.


## Features

- Clean Discord account messages and data
- Multiple tokens suport
- Simple command line interface
- Fast and efficient operation
- Cross-platform compatibility

## Installation

```bash
# Clone the repository
git clone https://github.com/Nevesto/godel.git

# Navigate to the project directory
cd godel

# Install dependencies
go get

# Build the project
go build
```

## Basic usage

```bash
# Register a new token
godel set-token [name] [token]

# Clear only direct messages
godel clear-dm [id]

# Clear all messages on a guild
godel clear-guild [guild_id] 

# Get help
godel --help
```

## Requirements

- Go version: go1.23.1
- Discord account token

## Contributing

1. Fork the repository
2. Create a new branch (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -m 'Add a feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a Pull Request

## Disclaimer

Discord does not allow account automation. This tool is developed for educational purposes. Use at your own risk.

## License

Distributed under the MIT License. See `LICENSE` file for more information.