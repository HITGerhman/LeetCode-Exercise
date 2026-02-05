#include <vector>
#include <iostream>
#include <sstream>

using namespace std;

class Solution {
public:
    vector<int> productExceptSelf(vector<int>& nums) {
        int n = nums.size();
        vector<int> res(n);
        res[0] = 1;
        for(int i = 1;i<n;i++) {
            res[i] = res[i-1]*nums[i-1];
        }
        int R = 1;
        for(int i = n-1;i>=0;i--) {
            res[i] = res[i]*R;
            R *= nums[i];
        }
        return res;
    }
};
int main() {
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
        vector<int> res = sol.productExceptSelf(nums);
        
        cout<<'[';
        for(int i = 0;i<res.size();i++) {
            cout<<res[i];
            if(i!=res.size()-1) {
                cout<<',';
            }
        }
        cout<<']';

    }
}