/*
344. Reverse String

Write a function that takes a string as input and returns the string reversed.
Example:
Given s = "hello", return "olleh".

===========================================================================================================================================
*/

class Solution {
public:
    string reverseString(string s) {
        if(s.length() == 0)
            return "";
            
        int i = 0, j = s.size() - 1;
        while(i < j){
            swap(s[i++], s[j--]);
        }
        
        return s;
    }
};
