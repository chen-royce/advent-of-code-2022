from util import parse
from bfs import *


def main():
    # filepath = "./problem/sample.txt"
    filepath = "./problem/problem.txt"
    file = open(filepath, "r")
    graph = parse(file)
    start_row, start_col = find_start(graph)
    starting_node = Node("S", start_row, start_col, 0)
    min_score = bfs(starting_node, graph)  # use this as a seed score
    starting_coordinates = find_zero_altitudes(graph)
    for c in starting_coordinates:
        row = c[0]
        col = c[1]
        starting_node = Node(graph[row][col], row, col, 0)
        score = bfs(starting_node, graph)
        if score == None:  # this could use some cleaning up
            continue
        if score < min_score:
            min_score = score
    print(min_score)


if __name__ == "__main__":
    main()
