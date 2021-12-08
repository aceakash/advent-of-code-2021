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
    print(items)
    
    digit_strings = extract_digit_strings(items)  # [['abcde', 'bcdefr'], ['abcde', 'bcdefr'], ['abcde', 'bcdefr']]

    return count_1_4_7_8s(digit_strings)

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

def extract_digit_strings(inputs):
    return []

def count_1_4_7_8s(digit_strings):
    return 0


if __name__ == "__main__":
    # print(part_one(example1()))
    print(part_one(example2()))
    # print(part_one(real_input()))