import unittest
from day5 import part_one, get_diag_points

class TestDay5(unittest.TestCase):
    def test_Example1(self):
        self.assertEqual(part_one("9,1 -> 9,5"), 0)

    def test_Example2(self):
        self.assertEqual(part_one("1,1 -> 5,1"), 0)

    def test_Example6(self):
        self.assertEqual(part_one("3,4 -> 1,4"), 0)

    def test_Example4(self):
        self.assertEqual(part_one("""1,1 -> 5,1
        4,1 -> 6,1"""), 0)

    def test_Example5(self):
        self.assertEqual(part_one("""0,1 -> 1,1
        1,1 -> 1,2"""), 0)

    def test_diag_points_1(self):
        self.assertEqual(get_diag_points((0,0), (1,1)), [(0,0), (1,1)])

    def test_diag_points_2(self):
        self.assertEqual(get_diag_points((1,1), (0,0)), [(1,1), (0,0)])

    def test_diag_points_3(self):
        self.assertEqual(get_diag_points((1,0), (0,1)), [(1,0), (0,1)])


    def test_Example3(self):
        input = """0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2"""
        self.assertEqual(part_one(input), 5)

    
    

if __name__ == "__main__":
    unittest.main()