from day9 import part_one


def test_part_one_with_example():
    assert part_one("""2199943210
3987894921
9856789892
8767896789
9899965678
""") == 15


def _test_part_one_with_simple1():
    assert part_one("""
    12
    """) == 2



def test_part_one_for_middle_of_grid():
    assert part_one("""
    987
    612
    765
    """) == 2

def test_part_one_real_file():
    with open("input.txt") as file:
        input = file.read()
        assert part_one(input) == 0