def parse(file):
    rows = []
    for row in file.readlines():
        toAdd = []
        for char in row:
            if char != '\n':
                toAdd += char
        rows.append(toAdd)
    return rows

def find_start(input):
    for i in range(len(input)):
        for j in range(len(input[0])):



def main():
    filepath = "./problem/sample.txt"
    file = open(filepath, "r")
    input = parse(file)
    my_list = [1,2,3]


if __name__ == "__main__":
    main()
