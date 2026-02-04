#include <sstream>
#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        vector<vector<int>> res;
        sort(nums.begin(),nums.end());
        int n = nums.size();
        if(n<3) {
            return res;
        }
        for(int i = 0;i<nums.size()-2;i++) {
            if(nums[i]>0) {
                break;
            }
            if(i>0&&nums[i]== nums[i-1]) {
                continue;
            }
            int left = i+1;
            int right =  n-1 ;
            while(left<right) {
                int sum = nums[i]+nums[left]+nums[right];
                if(sum==0) {
                    res.push_back({nums[i],nums[left],nums[right]});
                    while (left<right&&nums[left]==nums[left+1])
                    {
                       left++; 
                    }
                    while(left<right&&nums[right]==nums[right-1]) {
                        right--;
                    }
                    right--;
                    left++;
                    
                }
                else if(sum<0) {
                    left++;
                }else {
                    right--;
                }
            }

        }
        return res;
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
        vector<vector<int>> res = sol.threeSum(nums);
        cout<<"[";
        for(int i = 0;i<res.size();i++) {
            cout<<"[";
            for(int j = 0;j<res[i].size();j++) {
                cout<<res[i][j];
                if(j<res[i].size()-1) {
                    cout<<",";
                }
            }
            cout<<"]";
            if(i<res.size()-1) {
                cout<<",";
            }
            
        }
        cout<<"]"<<endl;
    }
    return 0;
}