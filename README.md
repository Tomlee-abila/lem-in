# Lem-in

A Go implementation of an ant colony pathfinding program that finds the most efficient way to move ants through a colony from start to end room.

## Description

Lem-in is a pathfinding program that simulates ant colony movement through a network of rooms and tunnels. The program reads a colony description from a file and calculates the optimal paths to move all ants from the start room to the end room while adhering to specific constraints and rules.

### Key Features

- Finds the quickest path(s) for ants to traverse the colony
- Handles multiple possible paths and traffic optimization
- Validates input data format and colony structure
- Provides clear visualization of ant movements
- Error handling for various invalid input scenarios

## Technical Requirements

- Go 1.21 or higher
- Standard Go packages only (no external dependencies)

## Installation

```bash
# Clone the repository
git clone https://learn.zone01kisumu.ke/git/tabila/lem-in

# Navigate to project directory
cd lem-in
```

## Usage

Basic usage:
```bash
go run . [input_file]
```

Example:
```bash
go run . test0.txt
```

### Input File Format Specifications

The input file must contain the following elements in order:

1. **Number of ants** (first line)
   - Must be a positive integer

2. **Room definitions**
   - Format: `name coord_x coord_y`
   - Example: `room1 23 3`
   - Special rooms marked with:
     - `##start` (starting room)
     - `##end` (ending room)

3. **Room connections**
   - Format: `name1-name2`
   - Example: `room1-room2`

Example of a valid input file:
```
3
##start
0 1 0
##end
1 5 0
2 9 0
3 13 0
0-2
2-3
3-1
```

### Rules and Constraints

1. **Room Names**
   - Cannot start with 'L' or '#'
   - Must not contain spaces
   - Must be unique

2. **Tunnels**
   - Connect exactly two rooms
   - No duplicate tunnels between same rooms
   - A room can connect to multiple other rooms

3. **Ant Movement**
   - Only one ant per room (except start/end rooms)
   - Each ant moves only once per turn
   - Multiple ants can't use the same tunnel in the same turn

### Output Format

The program outputs:

1. The complete input data (colony description)
2. Ant movements in the format: `Lx-y Lz-w Lr-o ...`
   - x, z, r: ant numbers (1 to number_of_ants)
   - y, w, o: destination room names

Example output:
```
3
##start
0 1 0
##end
1 5 0
2 9 0
3 13 0
0-2
2-3
3-1

L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3
L3-1
```

### Error Handling

The program handles various error cases and displays appropriate messages:
- Invalid number of ants
- Missing start/end rooms
- Invalid room coordinates
- Invalid connections
- Unreachable end room
- Other formatting errors

## Contributors
- [Shayo Victor](https://github.com/makebelief/shayo-victor)
- [Kevin Wesonga](https://github.com/kevwasonga)
- [Tomlee Abila](https://github.com/Tomlee-abila)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
