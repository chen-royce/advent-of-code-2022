def generate_2d_array(width, height):
    a = []
    for i in range(height):
        a.append(['.'] * width)
    return a


def print_cave(cave):
    for row in cave:
        to_print = ""
        for char in row:
            to_print += char
        print(to_print)


class Cave:
    def __init__(self, min_x, max_x, max_y):
        self.map = generate_2d_array(max_x - min_x + 1, max_y+1)
        self.map_offset = min_x


def add_to_cave(cave, x, y, to_add):
    print(x)
    print(y)
    print(cave.map_offset)
    cave.map[y][x - cave.map_offset] = to_add
    print_cave(cave.map)


def remove_from_cave(cave, x, y):
    cave.map[y][x - cave.map_offset] = '.'


def paths_to_cave(paths, min_x, max_x, max_y):
    cave = Cave(min_x, max_x, max_y)
    for point in paths:
        add_to_cave(cave, point[0], point[1], '#')
    return cave
