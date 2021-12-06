def example():
    return """3,4,3,1,2"""


def read_file():
    with open("input.txt") as file:
        return file.read()

    
def lanternfish(input, afterDays):
    timers = [int(r) for r in input.split(',')]
    days_elapsed = 0
    while days_elapsed < afterDays:
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


def part_one(input):
    return len(lanternfish(input, 80))

if __name__ == "__main__":
    # print(part_one(example()))
    print(part_one(read_file()))