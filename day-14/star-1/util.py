def file_to_paths(file):
    paths_text = file.read().split('\n')
    paths_text = paths_text[:len(paths_text)]
    paths = []

    min_x, max_x, max_y = -1, -1, -1

    for path in paths_text:
        path = path.split(' -> ')  # skip last line
        for point in path:
            point = point.split(',')
            coordinates = []
            for coordinate in point:
                coordinate = int(coordinate)
                coordinates.append(coordinate)
            if len(coordinates) != 2:
                raise Exception("failed to parse")

            # Update min/max vals
            if coordinates[0] < min_x or min_x == -1:
                min_x = coordinates[0]
            if coordinates[0] > max_x or max_x == -1:
                max_x = coordinates[0]
            if coordinates[1] > max_y or max_y == -1:
                max_y = coordinates[1]

            # Append to paths
            paths.append((coordinates[0], coordinates[1]))

    return paths, min_x, max_x, max_y
