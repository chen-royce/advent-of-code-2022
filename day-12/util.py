def parse(file):
    rows = []
    for row in file.readlines():
        toAdd = []
        for char in row:
            if char != '\n':
                toAdd += char
        rows.append(toAdd)
    return rows