# calming-waters

## Goal
A simulation of a tank in which fish (ASCII) move can have different behaviors: e.g. swimming, looking for food, avoiding each other.

## Feature list
- [X] Drawing a rectangle in the terminal (e.g., 20x10)
- [X] Representing a fish as a struct with position and direction
- [X] Main loop: clear screen, update fish positions, draw again
- [ ] Random movement of fish
- [X] Aquarium boundaries (fish turn)
- [ ] Fish do not "hit" each other

## Architecture goals
- Split display logic from domain logic
- Use only standard library
- Entity-Component-System (ECS)
- tick system
- external configuration
- (extra) multithreading
## Setup
TBD

