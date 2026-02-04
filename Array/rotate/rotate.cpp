#include <iostream>
#include <sstream>
#include <vector>
#include <algorithm>

using namespace std;

class Solution {
public:
    void rotate(vector<int>& nums,int k) {
        int n = nums.size();
        k = k%n;
        if(k == 0) {
            return;
        }
        reverse(nums.begin(),nums.end());
        reverse(nums.begin(),nums.begin()+k);
        reverse(nums.begin()+k,nums.end());

    }
};
int main() {
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
        sol.rotate(nums, k);
        cout<<'[';
        for(int i = 0;i<nums.size();i++) {
            cout<<nums[i];
            if(i!=nums.size()-1) {
                cout<<',';
            }
        }
        cout<<']';
    }
}