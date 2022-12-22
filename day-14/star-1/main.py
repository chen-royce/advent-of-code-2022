from util import parse_file


def main():
    filepath = './problem/sample.txt'
    # filepath = './problem/problem.txt'
    file = open(filepath, 'r')
    paths, max_x, min_x, max_y, min_y = parse_file(file)


if __name__ == '__main__':
    main()
