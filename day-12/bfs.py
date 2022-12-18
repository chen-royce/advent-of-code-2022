from collections import deque


def calculate_altitude(char):
    if char == "S":
        return 0
    elif char == "E":
        return 25

    altitudes = "abcdefghijklmnopqrstuvwxyz"
    return altitudes.index(char)


class Node:
    def __init__(self, value, row, col, distance_traveled):
        self.value = value
        self.altitude = calculate_altitude(value)
        self.row = row
        self.col = col
        self.distance_traveled = distance_traveled


def stringify_coordinates(node):
    return (f'{node.row}-{node.col}')


def get_neighbors(node, graph):
    neighbors = []
    if node.row != 0:
        top_neighbor = Node(
            graph[node.row-1][node.col],
            node.row-1,
            node.col,
            node.distance_traveled+1
        )
        neighbors.append(top_neighbor)
    if node.row != len(graph) - 1:
        bot_neighbor = Node(
            graph[node.row+1][node.col],
            node.row + 1,
            node.col,
            node.distance_traveled + 1
        )
        neighbors.append(bot_neighbor)
    if node.col != 0:
        left_neighbor = Node(
            graph[node.row][node.col-1],
            node.row,
            node.col - 1,
            node.distance_traveled + 1
        )
        neighbors.append(left_neighbor)
    if node.col != len(graph[0]) - 1:
        right_neighbor = Node(
            graph[node.row][node.col+1],
            node.row,
            node.col + 1,
            node.distance_traveled + 1
        )
        neighbors.append(right_neighbor)
    return neighbors


def find_start(input):
    for i in range(len(input)):
        for j in range(len(input[0])):
            if input[i][j] == "S":
                return i, j
    raise ValueError("Could not find starting node 'S'")


def find_end(input):
    for i in range(len(input)):
        for j in range(len(input[0])):
            if input[i][j] == "E":
                return i, j
    raise ValueError("Could not find ending node 'E'")


def find_zero_altitudes(input):
    zero_altitudes = []
    for i in range(len(input)):
        for j in range(len(input[0])):
            if calculate_altitude(input[i][j]) == 0:
                zero_altitudes.append((i, j))
    return zero_altitudes


def bfs(start_node, graph):
    # Initialize queue of nodes to visit
    to_visit = deque([start_node])

    # Initialize dictionary of visited nodes' locations
    visited = {stringify_coordinates(start_node): True}

    # Find the ending target
    end = find_end(graph)

    while len(to_visit) != 0:
        # Get next node to visit and mark as visited
        curr_node = to_visit.popleft()

        # If the node is what we want, return its distance value
        if (curr_node.row, curr_node.col) == end:
            return curr_node.distance_traveled

        # Otherwise, check its neighbors
        neighbors = get_neighbors(curr_node, graph)
        for neighbor in neighbors:
            if neighbor.altitude - curr_node.altitude > 1:
                continue
            if stringify_coordinates(neighbor) not in visited:
                to_visit.append(neighbor)
                visited[stringify_coordinates(neighbor)] = True
