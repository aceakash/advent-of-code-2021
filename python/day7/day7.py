def example1():
    return """16,1,2,0,4,2,7,1,2,14"""


def read_file():
    with open("input.txt") as file:
        return file.read()

def part_one(input):
    nums = [int(x) for x in input.split(',')]
    nums.sort()
    median = nums[len(nums)//2]
    f = get_fuels(nums, median)
    return f

def get_fuels2(list, num):
    fs = [abs(x-num) for x in list]
    s = [f*(f+1)/2 for f in fs]
    return sum(s)

def part_two(input):
    nums = [int(x) for x in input.split(',')]
    nums.sort()
    # print(nums)
    avg = round(sum(nums)/len(nums))
    # print(avg)
    i = nums.index(avg)
    for j in range(10):
        k = i - j
        
        print(nums[k], get_fuels2(nums, nums[k]))
    




def get_fuels(list, num):
    # print(list, num)
    return [abs(x-num) for x in list]

if __name__ == "__main__":
    input = read_file()
    # input = example1()
    p1 = part_one(input)
    # print(sum(p1))
    p2 = part_two(input)
    # print(sum(p1))
