def example():
    return """3,4,3,1,2"""

def example1():
    return """3,4"""


def read_file():
    with open("input.txt") as file:
        return file.read()

    
def lanternfish(input, afterDays):
    timers = [int(r) for r in input.split(',')]
    days_elapsed = 0
    while days_elapsed < afterDays:
        print("days_elapsed", days_elapsed)
        # print(days_elapsed, timers)
        newtimers = []
        spawned = []
        for t in timers:
            newt = t-1
            if t == 0:
                newt = 6
                spawned.append(8)
                
            newtimers.append(newt)
        newtimers = newtimers + spawned
        days_elapsed += 1
        timers = newtimers

    # print(days_elapsed, timers)
    return timers

def shift_left(arr):
    newarr = [0 for x in arr]
    for i in range(len(arr)-1):
        newarr[i] = arr[i+1]
    newarr[-1] = arr[0]
    return newarr

def lanternfish2(input, days):
    timers = [int(r) for r in input.split(',')]
    days_elapsed = 0
    counts = [0, 0, 0, 0, 0, 0, 0, 0, 0]
    print(0, [0,1,2,3,4,5,6,7,8])
    for t in timers:
        counts[t] = counts[t] + 1
    # print(days_elapsed, counts)
    while days_elapsed < days:
        print(days_elapsed, counts)
    
        newcounts = [0,0,0,0,0,0,0,0,0]
        newcounts[0] = counts[1]
        newcounts[1] = counts[2]
        newcounts[2] = counts[3]
        newcounts[3] = counts[4]
        newcounts[4] = counts[5]
        newcounts[5] = counts[6]
        newcounts[6] = counts[0] + counts[7]
        newcounts[7] = counts[8]
        newcounts[8] = counts[0]

        counts = newcounts
        # print(counts)
        days_elapsed += 1

    print(days_elapsed, counts)
    return counts


def part_one(input):
    return len(lanternfish(input, 80))

def part_two(input):
    return sum(lanternfish2(input, 256))
    # return len(lanternfish2(input, 10))


if __name__ == "__main__":
    # print(part_one(example()))
    # print(part_two(example()))
    print(part_two(read_file()))