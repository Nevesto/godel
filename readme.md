# Godel (2.0.0)

![PREVIEW](/public/PREVIEW.png)

## About

Godel is a command line tool to easily clear Discord accounts with built-in security features to avoid account bans.

This project is under active development. For any bugs suport or bugs contact open a new issue.


## Features

- **Clean Discord account messages and data** - Delete your messages from DMs and guilds
- **Enhanced DM clearing** - Clear all DMs including closed conversations
- **Multiple security profiles** - Choose between conservative, default, or aggressive rate limiting
- **Advanced rate limiting** - Configurable delays and exponential backoff to avoid bans
- **Random delays** - Randomized timing to appear more human-like
- **Retry logic** - Automatic retry with exponential backoff on rate limit errors
- **Multiple tokens support** - Manage and switch between multiple Discord accounts
- **Simple command line interface** - Easy to use commands
- **Fast and efficient operation** - Optimized message deletion
- **Cross-platform compatibility** - Works on Windows, macOS, and Linux
- **Modular architecture** - Clean, maintainable code structure

## Security Features

Godel includes several safety mechanisms to reduce the risk of account bans:

- **Rate Limiting**: Configurable requests per second to avoid hitting Discord's rate limits
- **Random Delays**: Adds randomization between operations to appear more natural
- **Exponential Backoff**: Automatically backs off when rate limited
- **Batch Processing**: Processes messages in configurable batches with delays
- **Security Profiles**: Three pre-configured profiles for different risk tolerances

### Security Profiles

- **Conservative** (recommended): Slowest but safest, 1 request every 4 seconds
- **Default**: Balanced approach, 1 request every 2 seconds
- **Aggressive**: Faster but higher risk, 1 request per second

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

# Switch to a token
godel token-switch [name]

# Clear all DMs (including closed ones) with default security
godel clear-all-dms

# Clear all DMs with conservative security (slower, safer)
godel clear-all-dms --security conservative

# Clear all DMs with aggressive security (faster, riskier)
godel clear-all-dms --security aggressive

# Clear only specific direct message channel
godel clear-dm [channel_id]

# Clear all messages on a guild
godel clear-guild [guild_id] 

# Get help
godel --help
```

## Architecture

The project follows a modular architecture:

- **`pkg/config`**: Security configuration and profiles
- **`pkg/ratelimit`**: Rate limiting and backoff logic
- **`pkg/client`**: Enhanced Discord client with retry logic
- **`pkg/cleaner`**: Message deletion operations
- **`auth`**: Authentication and token management
- **`cmd`**: CLI commands
- **`scripts`**: Legacy script wrappers

## Requirements

- Go version: go1.23.1 or higher
- Discord account token

## Contributing

1. Fork the repository
2. Create a new branch (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -m 'Add a feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a Pull Request

## Disclaimer

Discord does not allow account automation. This tool is developed for educational purposes. Use at your own risk. The security features included in this tool are designed to reduce risk but cannot guarantee that your account will not be banned.

## License

Distributed under the MIT License. See `LICENSE` file for more information.