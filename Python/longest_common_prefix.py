class Solution:
    def longestCommonPrefix(self, strs) -> str:
        result = ""
        for num in zip(*strs):
            if len(set(num)) == 1:
                result += num[0]
            else:
                return result
            
        return result