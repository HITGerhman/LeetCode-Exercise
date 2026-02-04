#include <iostream>
#include <vector>
#include <sstream>

using namespace std;

class Solution
{
public:
    int maxSubArray(vector<int> &nums)
    {
        int currentSum = 0;
        int maxSum = nums[0];

        for (int num : nums)
        {
            if (currentSum > 0)
            {
                currentSum += num;
            }
            else
            {
                currentSum = num;
            }
            maxSum = max(maxSum, currentSum);
        }
        return maxSum;
    }
};
int main()
{
    if (!freopen("input.txt", "r", stdin))
    {
        return 0;
    }
    string line;
    while (getline(cin, line))
    {
        if (line.empty())
        {
            continue;
        }
        for (char &c : line)
        {
            if (c == '[' || c == ']' || c == ',')
            {
                c = ' ';
            }
        }
        stringstream ss(line);
        vector<int> nums;
        int temp;
        while (ss >> temp)
        {
            nums.push_back(temp);
        }
        Solution sol;
        int res = sol.maxSubArray(nums);
        cout << res << endl;
    }
}