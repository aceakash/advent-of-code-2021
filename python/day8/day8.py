def example2():
    return """be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce"""

def example1():
    return """acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"""


def real_input():
    with open("input.txt") as file:
        return file.read()


def part_one(input):
    # how many segments of 2,3,4,7 lengths?
    items = parse_input(input)                    # ['line1', 'line2']
    # print(items)
    
    return count_1_4_7_8s(items)

def part_two(input):
    

    items = parse_input(input)
    total = 0
    for item in items:
        nums = extract_nums(item)
        total += sum(nums)

    return total
    
def extract_nums(item):
    lookup = make_lookup(item)
    # lookup = {
    #     'abcdefg': 8,
    #     'bcdef': 5,
    #     'acdfg': 2,
    #     'abcdf': 3,
    #     'abd': 7,
    #     'abcdef': 9,
    #     'bcdefg': 6,
    #     'abef': 4,
    #     'abcdeg': 0,
    #     'ab': 1
    # }
    strs = item['four_segments']
    nums = []
    for s in strs:
        s2 = ''.join(sorted(s))
        if s2 == 'bceg':
            print(item)
        num_str = lookup[s2]
        nums.append(int(num_str))        
    return nums

def make_lookup(item):
    unknowns = []
    lookups = {}
    for s in item['ten_segments']:
        s2 = ''.join(sorted(s))
        if len(s2) == 2:
            lookups[s2] = '1'
            continue
        if len(s2) == 3:
            lookups[s2] = '7'
            continue
        if len(s2) == 7:
            lookups[s2] = '8'
            continue  
        if len(s2) == 4:
            lookups[s2] = '4'
            continue  
        unknowns.append(s2)
    
    print(lookups)
    print(unknowns)
    return lookups

def parse_input(input):
    lines = input.split('\n')
    items = []
    for l in lines:
        parts = l.strip().split('|')
        ten_segments = parts[0].strip().split(' ')
        four_segments = parts[1].strip().split(' ')
        
        item = {
            'raw_line': l,
            'ten_segments': ten_segments,
            'four_segments': four_segments 
        }
        items.append(item)

    return items


def count_1_4_7_8s(items):
    count = 0
    for item in items:
        strs = item['four_segments']
        for s in strs:
            if len(s) == 2 or len(s) == 4 or len(s) == 3 or len(s) == 7:
                count += 1
    
    return count


if __name__ == "__main__":
    # print(part_one(real_input()))

    print(part_two(example1()))
    # print(part_two(example2()))
    # print(part_two(real_input()))