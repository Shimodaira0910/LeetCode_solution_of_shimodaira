class Solution:
    def isPalindrome(self, x: int) -> bool:
        str_num = str(x)
        if str_num == str_num[::-1]:
            return True
        else:
            return False