import java.util.*;
import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Deque;
import java.util.Set;
// import java.util.regex.Pattern;
import java.util.HashSet;
import java.util.HashMap;
import java.util.Map;
// import java.util.Locale;


public class Solution {
    public boolean idAnagram(String s, String t){
        char[] stringToCharArrayS = s.toCharArray();
        char[] stringToCharArrayT = t.toCharArray();

        Arrays.sort(stringToCharArrayS);
        Arrays.sort(stringToCharArrayT);

        String sortS = new String(stringToCharArrayS);
        String sortT = new String(stringToCharArrayT);

        if(sortS.equals(sortT)){
            return true;
        } else {
            return false;
        }
    }

    public int majorityElement(int[] num){
        int result = 0;
        int ans = 0;
        List<Integer> list = new ArrayList<Integer>(num.length);
        for (int i = 0; i < num.length; i++) {
            list.add(num[i]);
        }

        Set<Integer> set = new HashSet<>(list);
        for (Integer i: set){
            if(Collections.frequency(list, i) > result){
                result = Collections.frequency(list, i);
                ans = i;
            }
        }
        return ans;
    }

    public boolean isPalindrome(String s) {
        if(s == ""){
            return true;
        }
        String filtering = s.replaceAll("[^A-Za-z0-9]", "").toLowerCase();
        
        StringBuilder sb = new StringBuilder();
        sb.append(filtering);

        return sb.toString().equals(sb.reverse().toString());
    }


    public int[] twosum(int[] nums, int target) {
        int[] intResult = new int[2];
        System.out.println();
        int pos = 0;
        int length = 1;
        
        while(true){
            if(nums[pos] + nums[length] == target){          
               intResult[0] = pos;
               intResult[1] = length;
               return intResult;
            }
            if(length == nums.length){
                length = 0;
                pos += 1;
                length = pos + 1;
            }else{
                length += 1;
            }
        }
    }

    public int maximum69Number(int num){
        List<Integer> numList = new ArrayList<>();
        while(num != 0){
            numList.add(num % 10);
            num /= 10;
        }
        for(int i = 0; numList.size() > i; i++){
            if(numList.get(i) == 6){
                numList.set(i, 9);
            }else if(numList.get(i) == 9){
                numList.set(i, 6);
            }
        }
        return num;
    }


    public int removeDuplicate(int[] nums){
        if(nums.length == 0){
            return 0;
        }
        int pos = 0;
        for(int i = 0; nums.length > i; i++){
            if(nums[i] != nums[pos]){
                pos++;
                nums[pos] = nums[i];
            }
        }
        System.out.println(nums);
        return pos + 1;
    }

    public int[] plusOne(int[] digits){
        int len = digits.length - 1;
        for(int i = len; i >= 0; i--){ 
            if(digits[i] <= 9){
                digits[i]++;
                return digits;
            }
            digits[i] = 0;
        }
        int[] newDigits = new int[len + 1];
        newDigits[0] = 1;
        return newDigits;
    }

    public boolean isValid(String s){        
        Map<Character, Character> map = new HashMap<>();
        map.put(']', '[');
        map.put('}', '{');
        map.put(')', '(');
    
        Deque<Character> deque = new ArrayDeque<>();
        for(char c: s.toCharArray()){
            if(map.values().contains(c)){
                deque.push(c);
            }else if(map.keySet().contains(c)){
                if(deque.isEmpty() || map.get(c) != deque.pop()){
                    return false;
                }
            }else{
                return false;
            }
        }
        return deque.isEmpty();
    }

    public int lengthOfLastWord(String s){
        String[] splitS = s.split(" ");
        String test = splitS[splitS.length - 1];
        return test.length();
    }

    public boolean wordPattern(String pattern, String s) {
        Set<String> hashSet = new HashSet<String>();
        String arr[] = pattern.split(",");
        
        System.out.println(arr);

        for(int i = 0; arr.length < i;i++){
            hashSet.add(arr[i]);
        }

        System.out.println(hashSet);
        return true;
    }

    public int[] distributeCandies(int candies, int numPepople){
        int[] ans = {numPepople};
        
        int handOverCandies = 1;
        int pos = 0;

        while(candies > 0){
            if(pos == numPepople){
                pos = 0;
            }

            if(handOverCandies < candies){
                ans[pos] += candies;
                candies  -= handOverCandies;
                pos += 1;
                handOverCandies += 1;
                continue;
            }

            ans[pos] += handOverCandies;
            candies  -= handOverCandies;

            pos += 1;
            handOverCandies += 1;
        }
        return ans;
    }

    public int maxProduct(int[] nums){
        int maxNum = 0;
        int tmp2 = 0;
        int ans = 0;

        for(int i = 0; nums.length - 1 > i; i++){
            maxNum  = (nums[i] > maxNum) ? nums[i] : maxNum;
            tmp2 = (maxNum - 1) * (nums[i + 1] - 1);
            ans  = (tmp2 > ans) ? tmp2 : ans;
        }

        return ans;
    }

    public int busyStudent(int[] startTime, int[] endTime, int queryTime){
        int ans = 0;
        for (int i = 0; i < startTime.length; i++) {
            for (int l = startTime[i]; l <= endTime[i]; l++) {
                if (l == queryTime) {
                    ans += 1;
                    break;
                }
            }
        }
    
        return ans;
    }

    public String sortSentence(String s) {
        String[] sentenceLists = s.split(" ");
        ArrayList<String> ans = new ArrayList<String>();
            
        int i = 0;
        int roopBreakNum = 0;
        int isMatchNumber = 1;
    
        while (sentenceLists.length > roopBreakNum) {
            int wordLength = sentenceLists[i].length();
            int lastLength = wordLength - 1;
            String lastChar = sentenceLists[i].substring(lastLength);
    
            if (lastChar.equals(String.valueOf(isMatchNumber))) {
                ans.add(sentenceLists[i].substring(0, lastLength));
                isMatchNumber += 1;
                roopBreakNum += 1;
            }
    
            i = (i + 1) % sentenceLists.length;
        }
    
        return String.join(" ", ans);
    }

    public int isPrefexOfWord(String sentence, String searchWord){
        String[] arr = sentence.split(" ");
        char[] matchWordSplit = searchWord.toCharArray();
        int ans = 0;

        for(int i = 0; arr.length > i; i++){
            char[] oneSentenceSplit = arr[i].toCharArray();

            if(matchWordSplit[0] == oneSentenceSplit[0]){
                for(int x = 0; matchWordSplit.length > x; x++){
                    if(matchWordSplit[i] != oneSentenceSplit[i]){
                        break;
                    }
                }
                ans = i + 1;
            }
        }

        return ans;
    }

    public int sumOfMultiples(int n){
        List<Integer> testNums = new ArrayList<Integer>();
        for(int i = 1; i <= n ; i++){
            testNums.add(i);
        }
        int[] numArray = testNums.stream().mapToInt(i->i).toArray();
        int ans = 0;
        for(int x = 0; x < numArray.length ; x++){
            if(numArray[x] % 3 == 0 || numArray[x] % 5 == 0 || numArray[x] % 7 == 0){
                ans += numArray[x];
            }
        }
        return ans;
    }

}