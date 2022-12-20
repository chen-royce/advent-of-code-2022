STRING_TYPE = type("")

def find_closing_bracket(text):
    l_brackets = 0
    r_brackets = 0
    for i in range(len(text)):
        if text[i] == "[":
            l_brackets += 1
        if text[i] == "]":
            r_brackets += 1
            if l_brackets == r_brackets:
                return i
    raise Exception("can't find closing bracket")

def divide_into_values(text):
    ret = []
    while text.find("[") != -1:
        # find bracketed area
        list_start = text.find("[")
        list_end = find_closing_bracket(text)

        # if there's stuff before the bracket, add it first
        if list_start != 0:
            head = text[:list_start-1].split(",") # 1 to omit comma
            for h in head:
                ret.append(int(h))
        
        # otherwise, just add the bracket
        ret.append(text[list_start:list_end+1])

        # trim string and repeat
        text = text[list_end+2:]

    # if rest of string is not empty, append it
    if len(text) > 0:
        split = text.split(",")
        for item in split:
            ret.append(int(item))
    
    return ret



def convert_input(text):
    # base case: no brackets to be found
    if "[" not in text:
        return int(text)
    
    # if brackets found
    list_start = text.index("[")
    list_end = find_closing_bracket(text)

    text = text[list_start+1:list_end]
    items = divide_into_values(text)

    ret = []

    for item in items:
        if type(item) == STRING_TYPE:
            item = convert_input(item)
        ret.append(item)

    return ret