import pytest

from compare import compare_lists

def test_compare_lists():
    assert compare_lists([],[]) == True
    assert compare_lists([1],[2]) == True
    assert compare_lists([2],[1]) == False
    assert compare_lists([2],[2]) == True
    assert compare_lists([2,2],[2]) == False
    assert compare_lists([9],[[8,7,6]]) == False
    assert compare_lists([],[2]) == True
    assert compare_lists([[[]]],[[]]) == False
    assert compare_lists([1,[2,[3,[4,[5,6,7]]]],8,9],[1,[2,[3,[4,[5,6,0]]]],8,9]) == False