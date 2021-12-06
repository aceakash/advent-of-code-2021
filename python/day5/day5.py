import os

def example():
    # return "0,9 -> 5,9"
    return """0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2"""


def get_diag_points(start, end):
    (startx, starty) = start
    (endx, endy) = end
    (x, y) = (startx, starty)
    diag_points = []
    (xdir, ydir) = (1, 1)
    if startx > endx:
        xdir = -1
    if starty > endy:
        ydir = -1
    while x != endx:
        diag_points.append((x, y))
        x += xdir
        y += ydir
    
    diag_points.append(end)

    return diag_points 

def read_file():
    with open("input.txt") as file:
        return file.read()


def read_input_lines():
    # return read_file().splitlines()
    return example().splitlines()
    
    
def parse_point(str):
    parts = str.split(',')
    return (int(parts[0]), int(parts[1]))

def parse(input):
    rawlines = input.splitlines()
    lines = []
    for rl in rawlines:
        rawpoints = rl.split(' -> ')
        lines.append([parse_point(rawpoints[0]), parse_point(rawpoints[1])])
    return lines

def print_points(lines, covered_points):
    for y in range (0, 10):
        print("\n")
        for x in range (0, 10):
            if (x,y) in covered_points:
                print(covered_points[(x,y)], end=" ")
            else:
                print("-", end=" ")

def part_one(input):
    lines = parse(input)

    covered_points = {}

    for l in lines:
        print("line", l)
        (startx, starty) = l[0]
        (endx, endy) = l[1]
        # print(startx, starty, endx, endy)
        if startx == endx:
            print("startx == endx", l)
            for i in range(min(starty, endy), max(starty, endy)+1):
                covered_points[(startx, i)] = covered_points.setdefault((startx, i), 0) + 1
                # print("added point", covered_points)

        elif starty == endy:
            print("starty == endy", l)
            for i in range(min(startx, endx), max(startx, endx)+1):
                covered_points[(i, starty)] = covered_points.setdefault((i, starty), 0) + 1
                # print("added point", covered_points)

        
        elif abs(startx-endx) == abs(starty-endy):
            print("diag line", l)
            points = get_diag_points((startx, starty), (endx, endy))
            for p in points:
                covered_points[p] = covered_points.setdefault(p, 0) + 1

        else:
            print("will not use line", l)      


            
    # print("covered_points", covered_points)
    # print_points(lines, covered_points)        
    count = 0
    for v in covered_points.values():
        # print(v)
        if v > 1:
            count = count + 1

    return count


if __name__ == "__main__":
    print(part_one(read_file()))