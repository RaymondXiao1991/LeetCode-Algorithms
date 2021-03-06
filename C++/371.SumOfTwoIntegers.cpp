/*
371. Sum of Two Integers

Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -.
Example:
Given a = 1 and b = 2, return 3.

===========================================================================================================================================
*/

class Solution {
public:
    int getSum(int a, int b) {
		if (b == 0)
		{
			return a;
		}
    
    // a^b: 异或运算（两位相同的话则为0，否则为1）
    // a&b: 位与运算（两个全为1，则为1，否则为0）
    // (a&b)<<1: 对于0加0、0加1、1加0而言，都不会产生进位，只有1加1时，会向前产生一个进位。因此两个数先做位与运算，然后再向左移动一位
		return getSum(a^b, (a&b)<<1);
	}
};
