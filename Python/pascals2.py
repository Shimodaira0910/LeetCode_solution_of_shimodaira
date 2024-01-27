import string

def excel(columnNumber):
    alphabets = list(string.ascii_uppercase)
    alphabets_max_num = 26
    
    if columnNumber <= 26:
        return alphabets[columnNumber - 1]
    
    ans = []
    if columnNumber >= alphabets_max_num:
        ans.append("A")
        second_str_num = columnNumber - alphabets_max_num
        ans.append(alphabets[second_str_num - 1])
        return " ".join(ans)

columnNumber = 52
print(excel(columnNumber))