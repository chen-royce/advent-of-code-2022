from converter import *

INT_TYPE = type(1)
LIST_TYPE = type([])

def compare_lists(l1,l2):
    # if left list and right list run out simultaneously, inconclusive
    if len(l1) == 0 and len(l2) == 0:
        return

    # if left list runs out first, return T
    elif len(l1) == 0:
        return True
    
    # if right list runs out and left list isn't empty, return F
    elif len(l2) == 0:
        return False

    # if first elements both numbers, compare numbers
    if type(l1[0]) == INT_TYPE and type(l2[0]) == INT_TYPE:
        if l1[0] < l2[0]:
            return True
        if l1[0] > l2[0]:
            return False

    # if first element both lists, compare lists
    if type(l1[0]) == LIST_TYPE and type (l2[0]) == LIST_TYPE:
        if compare_lists(l1[0], l2[0]) == True:
            return True
        if compare_lists(l1[0], l2[0]) == False:
            return False
    
    # if one is number and other is list, convert and compare
    if type(l1[0]) == INT_TYPE and type (l2[0]) == LIST_TYPE:
        if compare_lists([l1[0]], l2[0]) == True:
            return True
        if compare_lists([l1[0]], l2[0]) == False:
            return False
    if type(l1[0]) == LIST_TYPE and type (l2[0]) == INT_TYPE:
        if compare_lists(l1[0], [l2[0]]) == True:
            return True
        if compare_lists(l1[0], [l2[0]]) == False:
            return False
    
    # if still no conclusion, continue with next elem
    return compare_lists(l1[1:], l2[1:])
