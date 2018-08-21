/*
1. Two Sum

Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution.

Example:
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

===========================================================================================================================================
*/

class Solution {
public:
	vector<int> twoSum(vector<int>& nums, int target) {
		vector<int> viRet;
		vector<int>::iterator tempIterator;
		map<int, int> miiMap;

		for (int i = 0; i < nums.size(); ++i) {
			if (miiMap.find(target - nums[i]) != miiMap.end()) {
				viRet.push_back(miiMap[target - nums[i]]);
				viRet.push_back(i);
			}
			miiMap[nums[i]] = i;
		}

		return viRet;

	}
};
