from util import *
from cave import *


def main():
    filepath = './problem/sample.txt'
    # filepath = './problem/problem.txt'
    file = open(filepath, 'r')
    paths, min_x, max_x, max_y = file_to_paths(file)
    cave = paths_to_cave(paths, min_x, max_x, max_y)


if __name__ == '__main__':
    main()
