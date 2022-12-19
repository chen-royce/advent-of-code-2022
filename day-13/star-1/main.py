def main():
    filepath = "./problem/sample.txt"
    file = open(filepath, "r")
    pairs = file.read().split("\n\n")
    for pair in pairs:
        print(pair.split("\n"))
    print(pairs)


if __name__ == "__main__":
    main()
