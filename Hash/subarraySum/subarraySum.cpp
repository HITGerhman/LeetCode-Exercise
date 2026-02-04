#include <iostream>
#include <unordered_map>
#include <sstream>
#include <vector>

using namespace std;

class Solution
{
public:
    int subarraySum(vector<int> &nums, int k)
    {
        unordered_map<int, int> mp;
        mp[0] = 1;
        int count = 0;
        int currentSum = 0;

        for (int num : nums)
        {
            currentSum += num;
            if (mp.find(currentSum - k) != mp.end())
            {
                count += mp[currentSum - k];
            }
            mp[currentSum]++;
        }
        return count;
    }
};

int main()
{
    if (!freopen("input.txt", "r", stdin))
    {
        return 0;
    }
    string line;
    int k;
    while (getline(cin, line) && cin >> k)
    {
        string dummy;
        getline(cin, dummy);
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
        int res = sol.subarraySum(nums, k);
        cout << res << endl;
    }
}