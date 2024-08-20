# Pokedex CLI

A simple CLI-based Pokedex built with Go that interacts with the [PokeAPI](https://pokeapi.co/) to fetch data about Pokémon. This project implements a REPL (Read-Eval-Print Loop) to create an interactive environment where users can explore locations, catch Pokémon, and manage their Pokedex.

## Features

- **Interactive CLI (REPL)**: An interactive command-line interface that allows users to input commands and receive responses in real time.
- **Fetch Pokémon Data**: Uses the PokeAPI to fetch and display data about Pokémon.
- **Explore Locations**: Allows users to explore different locations in the Pokémon world.
- **Catch Pokémon**: Users can catch Pokémon and add them to their Pokedex.
- **Cache Implementation**: Efficient caching system to store responses and reduce repeated API calls.
- **Command Set**: Includes commands for viewing the map, exploring locations, catching Pokémon, listing caught Pokémon, and inspecting individual Pokémon.

## Requirements

- Go 1.16 or later
- Internet connection to access the PokeAPI

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/konfucius1/pokedexcli.git
   cd pokedexcli
   ```

2. **Install dependencies**:

   Since this is a Go project, dependencies are managed via Go modules. Ensure your Go environment is set up correctly, and then run:

   ```bash
   go mod tidy
   ```

3. **Build the project**:

   ```bash
   go build -o pokedexcli
   ```

4. **Run the CLI**:

   After building, you can run the CLI tool with:

   ```bash
   ./pokedexcli
   ```

## Usage

Once you run the application, you will enter the REPL, where you can execute various commands to interact with the Pokedex.

### Commands

- `help`: Displays a list of available commands.
- `exit`: Exit the Pokedex CLI.
- `map`: Get the next page of locations.
- `mapb`: Get the previous page of locations.
- `explore <location_name>`: Explore a specific location in the Pokémon world.
- `catch <pokemon_name>`: Catch a Pokémon by its name.
- `inspect <pokemon_name>`: Inspect details of a caught Pokémon.
- `pokedex`: List all the Pokémon you've caught.

### Example

```bash
Pokedex > help
Available commands:
help - Displays a help message
exit - Exit the Pokedex
map - Get the next page of locations
mapb - Get the previous page of locations
explore <location_name> - Explore a location
catch <pokemon_name> - Catch a Pokémon
inspect <pokemon_name> - Inspect a Pokémon
pokedex - List all your Pokémon

Pokedex > explore pallet-town
Exploring Pallet Town...
```

## Project Structure

- `cmd`: Entry point of the application. Initializes the REPL and sets up the configuration.
- `internal/pokeapi`: Contains the client responsible for interacting with the PokeAPI.
- `internal/pokecache`: Implements a caching mechanism to store and retrieve API responses efficiently.
- `internal/repl`: Handles the REPL logic, including command processing and input/output operations.

## Caching

The project includes a custom caching mechanism (`pokecache`) that stores API responses for a specified duration to minimize redundant API calls. The cache is periodically cleaned up using a reaping function.

## Configuration

The configuration (`config` struct) in this project stores the state of the REPL, including:
- The PokeAPI client
- Caught Pokémon
- Pagination for location navigation

## Contributing

Contributions are welcome! If you have ideas to improve the project, feel free to submit a pull request or open an issue.
