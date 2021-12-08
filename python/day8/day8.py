from typing import List


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


def example3():
    return """cf bcdf acf abcdefg abdfg acdfg acdeg abcdfg abcefg abdefg | acdeg acdfg abcdfg abdefg"""

def real_input():
    with open("input.txt") as file:
        return file.read()


def part_one(input):
    # how many segments of 2,3,4,7 lengths?
    items = parse_input(input)
    return count_1_4_7_8s(items)

def part_two(input):
    items = parse_input(input)
    nums = [extract_num(item) for item in items]
    return sum(nums)
    
def extract_num(item):
    lookup = make_lookup(item)
    strs = item['four_segments']
    num_strs = [lookup[''.join(sorted(s))] for s in strs]
    return int(''.join(num_strs))

def make_lookup(item):
    unknowns, lookups = [], {}
    
    find_one_four_seven_eight(lookups, unknowns, item)
    two_code = find_two(lookups, unknowns, item)
    three_code = find_three(two_code, lookups, unknowns, item)
    five_code = find_five(lookups, unknowns, item)
    find_zero(five_code, lookups, unknowns, item)
    find_nine(three_code, lookups, unknowns, item)
    find_six(lookups, unknowns, item)

    return lookups

def find_one_four_seven_eight(lookups, unknowns, item):
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
    

def find_six(lookups, unknowns, item):
    lookups[unknowns[0]] = '6'
    unknowns.remove(unknowns[0])

def is_subset_of(a, b):
    return len(set(a).difference(set(b))) == 0

def find_nine(three_code, lookups, unknowns, item):
    # nine is the only one containing three_code as a subset
    nine_code = [x for x in unknowns if len(x) == 6 and is_subset_of(three_code, x)][0]
    lookups[nine_code] = '9'
    unknowns.remove(nine_code)

def find_zero(five_code, lookups, unknowns, item):
    # zero is the only one which doesnt contain five_code as a subset
    zero_code = [x for x in unknowns if len(x) == 6 and not is_subset_of(five_code, x)][0]
    lookups[zero_code] = '0'
    unknowns.remove(zero_code)


def find_five(lookups, unknowns, item):
    # five is the only five-letter code remaining
    five_code = [x for x in unknowns if len(x) == 5][0]
    lookups[five_code] = '5'
    unknowns.remove(five_code)
    return five_code

def find_two(lookups, unknowns, item):
    # the code for two is not a subset of any other unknowns
    five_letters = [x for x in unknowns if len(x) == 5]
    six_letters = [x for x in unknowns if len(x) == 6]

    two_code = None
    for fl in five_letters:
        found_subset = False
        for sl in six_letters:
            if is_subset_of(fl, sl):
                found_subset = True 
                break
        if not found_subset:
            two_code = fl
            break

    lookups[two_code] = '2'
    unknowns.remove(two_code)
    return two_code

def find_three(two_code, lookups, unknowns, item):
    # the code for three is one letter off from the code for two
    three_code = [x for x in unknowns if len(x) == 5 and has_one_difference(x, two_code)][0]
    lookups[three_code] = '3'
    unknowns.remove(three_code)
    return three_code

def has_one_difference(a, b):
    return len(set(a).difference(set(b))) == 1

def parse_input(input):
    items = []
    for l in input.split('\n'):
        parts = l.strip().split('|')
        items.append({
            'ten_segments': parts[0].strip().split(' '),
            'four_segments': parts[1].strip().split(' ') 
        })
    return items


def count_1_4_7_8s(items):
    count = 0
    for item in items:
        for s in item['four_segments']:
            if len(s) in [2,4,3,7]:
                count += 1
    
    return count


if __name__ == "__main__":
    assert part_one(real_input()) == 539 
    assert part_two(real_input()) == 1084606 
    