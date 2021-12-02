def read_input_lines():
    lines = []
    with open("input.txt") as file:
        for line in file:
            lines.append(line.rstrip())
    return lines


def part_one(instructions):
    horizontal_position = 0
    depth = 0

    for i in instructions:
        [command, num] = i.split(" ")
        num = int(num)
        if command == "forward":
            horizontal_position = horizontal_position + num
            continue
        if command == "up":
            depth = depth - num
            continue
        if command == "down":
            depth = depth + num
            continue
    
    return horizontal_position * depth


if __name__ == "__main__":
    
    instructions = read_input_lines()
    print("part one: ", part_one(instructions))
   