import pytest

from converter import *


def test_convert_input():
    assert convert_input("[3]") == [3]
    assert convert_input("[3,[2,1]]") == [3, [2, 1]]
