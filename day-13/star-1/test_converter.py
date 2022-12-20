from multiprocessing.sharedctypes import Value
import pytest

from converter import *

def test_find_closing_bracket():
    assert find_closing_bracket("[]") == 1
    assert find_closing_bracket("[ ]") == 2
    assert find_closing_bracket("[[] ]") == 4
    assert find_closing_bracket("[3]") == 2
    with pytest.raises(Exception) as errinfo:
        find_closing_bracket("[[ ]")
    assert "can't find closing bracket" in str(errinfo)

def test_convert_input():
    assert convert_input("[3]") == [3]
    assert convert_input("[3,[2,1]]") == [3, [2, 1]]
    assert convert_input("[[3,2],[1]]") == [[3,2],[1]]
    assert convert_input("[1,[2,[3,[4,[5,6,7]]]],8,9]") == [1,[2,[3,[4,[5,6,7]]]],8,9]

def test_divide_into_values():
    assert divide_into_values("4,[3,2],1,[0,1,2]") == [4, '[3,2]', 1, '[0,1,2]']
    assert divide_into_values("4,3,[3,2],1,[0,1,2]") == [4, 3, '[3,2]', 1, '[0,1,2]']