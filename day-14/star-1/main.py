from util import file_to_paths


def main():
    filepath = './problem/sample.txt'
    # filepath = './problem/problem.txt'
    file = open(filepath, 'r')
    paths, min_x, max_x, max_y = file_to_paths(file)


if __name__ == '__main__':
    main()
