#include <sstream>
#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

class Solution
{
public:
    int maxArea(vector<int> &height)
    {
        int left = 0;
        int right = height.size() - 1;
        int max_val = 0;
        while (left < right)
        {
            int h = min(height[left], height[right]);
            int w = right - left;
            int currentArea = w * h;
            max_val = max(currentArea, max_val);
            if (height[left] < height[right])
            {
                left++;
            }
            else
            {
                right--;
            }
        }
        return max_val;
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
        int res = sol.maxArea(nums);
        cout << res << endl;
    }
}