from compare import *
from converter import *
from functools import cmp_to_key

def sum_array(arr):
    acc = 0
    for num in arr:
        acc+=num
    return acc

def make_comparator(compare_func):
    def compare(x,y):
        if compare_func(x,y):
            return -1
        elif compare_func(y,x):
            return 1
        else:
            return 0
    return compare

def main():
    filepath = "./problem/problem.txt"
    file = open(filepath, "r")
    pairs = file.read().split("\n\n")
    items = []
    for i in range(len(pairs)):
        pair = pairs[i].split("\n")
        left = convert_input(pair[0])
        right = convert_input(pair[1])
        items.append(left)
        items.append(right)
    items.append([[2]])
    items.append([[6]])
    custom_key = cmp_to_key(make_comparator(compare_lists))
    sorted_list = sorted(items, key=custom_key)
    for i in range(len(sorted_list)):
        if sorted_list[i] == [[2]]:
            print("2 at", i+1)
        if sorted_list[i] == [[6]]:
            print("6 at", i+1)

if __name__ == "__main__":
    main()
