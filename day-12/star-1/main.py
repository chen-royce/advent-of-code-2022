def open_file(filepath):
    file = open(filepath, "r")
    return file


def read_input(file):

    rows = []

    for row in file.readlines():
        toAdd = []
        for char in row:
            if char != '\n':
                toAdd += char
        rows.append(toAdd)

    return rows


def main():
    read_input(open_file("./problem/sample.txt"))


if __name__ == "__main__":
    main()
