def part_one(input):
    points = parse_input(input)
    low_points = []
    for pl in points:
        for p in pl:
            adj = find_adjacent_points(p, points)
            if p[2] < min_adj_val(adj):
                low_points.append(p)    
        
    return sum([h[2]+1 for h in low_points])


def min_adj_val(adj):
    min = 9999999999
    if adj['n'] != None and adj['n'][2][2] < min:
        min = adj['n'][2][2]
    if adj['s'] != None and adj['s'][2][2] < min:
        min = adj['s'][2][2]
    if adj['e'] != None and adj['e'][2][2] < min:
        min = adj['e'][2][2]
    if adj['w'] != None and adj['w'][2][2] < min:
        min = adj['w'][2][2]
    
    return min

def parse_input(input):
    lines = [l.strip() for l in input.strip().split("\n")]

    points = []
    i = 0
    for i in range(len(lines)):
        l = lines[i]
        point_line = []
        for j in range(len(l)):
            point_line.append((i, j, int(l[j])))
        points.append(point_line)
        
    return points

def find_adjacent_points(p, points):
    (i, j, v) = p
    
    # corners
    if i == 0 and j == 0:
        return {
            'n': None,
            'w': None,
            'e': (i, j+1, points[i][j+1]),
            's': (i+1, j, points[i+1][j])
        }
    
    if i == 0 and j == len(points[0])-1:
        return {
            'n': None,
            'w': (i, j-1, points[i][j-1]),
            'e': None,
            's': (i+1, j, points[i+1][j])
        }

    if i == len(points)-1 and j == 0:
        return {
            'n': (i-1, j, points[i-1][j]),
            'w': None,
            'e': (i, j+1, points[i][j+1]),
            's': None
        }

    if i == len(points)-1 and j == len(points[0])-1:
        return {
            'n': (i-1, j, points[i-1][j]),
            'w': (i, j-1, points[i][j-1]),
            'e': None,
            's': None
        }
    
    # top edge
    if i == 0:
        return {
            'n': None,
            'w': (i, j-1, points[i][j-1]),
            'e': (i, j+1, points[i][j+1]),
            's': (i+1, j, points[i+1][j])
        }

    if i == len(points) - 1: # bottom edge
        return {
            'n': (i-1, j, points[i-1][j]),
            'w': (i, j-1, points[i][j-1]),
            'e': (i, j+1, points[i][j+1]),
            's': None
        }

    if j == 0: # left edge
        return {
            'n': (i-1, j, points[i-1][j]),
            'w': None,
            'e': (i, j+1, points[i][j+1]),
            's': (i+1, j, points[i+1][j])
        }

    if j == len(points[0]) - 1: # right edge
        return {
            'n': (i-1, j, points[i-1][j]),
            'w': (i, j-1, points[i][j-1]),
            'e': None,
            's': (i+1, j, points[i+1][j])
        }

    # has all neighbours
    return {
            'n': (i-1, j, points[i-1][j]),
            'w': (i, j-1, points[i][j-1]),
            'e': (i, j+1, points[i][j+1]),
            's': (i+1, j, points[i+1][j])
        }


    