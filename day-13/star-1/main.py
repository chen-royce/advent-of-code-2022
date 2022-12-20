from compare import *
from converter import *

def sum_array(arr):
    acc = 0
    for num in arr:
        acc+=num
    return acc

def main():
    # filepath = "./problem/sample.txt"
    filepath = "./problem/problem.txt"
    file = open(filepath, "r")
    pairs = file.read().split("\n\n")
    correct_pairs = []
    for i in range(len(pairs)):
        pair = pairs[i].split("\n")
        print(pair[0])
        print(pair[1])
        left = convert_input(pair[0])
        right = convert_input(pair[1])
        result = compare_lists(left,right)
        print(result)
        if result == True:
            correct_pairs.append(i+1)
    print("CORRECT PAIRS")
    print(correct_pairs)
    print("SUMS")
    print(sum_array(correct_pairs))

if __name__ == "__main__":
    main()
