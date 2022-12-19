LIST_TYPE = type([])
INT_TYPE = type(1)


# def convert_list_items(text):
#     return text.split(",")


def convert_input(text):
    print("CALLING CONVERT INPUT ON:")
    print(text)
    # base case: no brackets to be found
    if "[" not in text:
        strings = text.split(",")
        if len(strings) > 1:
            ints = []
            for str in strings:
                ints.append(int(str))
            return ints
        else:
            return int(strings[0])

    # if brackets found
    list_start = text.index("[")
    list_end = (len(text) - 1) - text[::-1].index("]")

    if list_start == 0 and list_end == len(text) - 1:
        print("HERE 0")
        return convert_input(text[list_start+1:list_end])

    ret = []

    # example: input is 0,[1,2,3]
    # we want to handle just 0
    if list_start != 0:
        print("HERE 1")
        ret.append(convert_input(text[0:list_start-1]))

    print("HERE 2")
    ret.append(convert_input(text[list_start+1:list_end]))

    # example: input is [1,2,3],0
    # we want to handle just 0
    if list_end != len(text)-1:
        print("HERE 3")
        ret.append(convert_input(text[list_end+2:]))

    return ret
