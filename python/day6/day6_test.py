import unittest
from day6 import shift_left

class TestDay6(unittest.TestCase):
    def test_shift_left_Example1(self):
        self.assertEqual(shift_left([1]), [1])
        
    def test_shift_left_Example2(self):
        self.assertEqual(shift_left([0, 1]), [1, 0])

    def test_shift_left_Example3(self):
        self.assertEqual(shift_left([2, 1, 1]), [1, 1, 2])
       
if __name__ == "__main__":
    unittest.main()