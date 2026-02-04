#include <iostream>
#include <vector>
#include <sstream>
#include <cstdio>
#include <queue>

using namespace std;

class Solution {
public:
    int searchInsert(vector<int>& nums,int target) {
        int left = 0;
        int right = nums.size()-1;

        while(left<=right) {
            int mid = left + (right - left)/2 ;

            if(nums[mid]==target) {
                return mid;
            } else if(nums[mid]<target) {
                left = mid+1;
            } else {
                right = mid-1;
            }
        }
        return left;
    }

};

int main() {
    if (!freopen("input.txt", "r", stdin))
    {
        cout << "无法打开 input.txt" << endl;
        return 0;
    }
    string line;
    string targetStr;
    int target;

    while (getline(cin, line)&&getline(cin,targetStr))
    {
        if (line.empty()) {
            continue;
        }
        target = stoi(targetStr);
        for(char &c:line) {
            if(!isdigit(c)&&c!='-') {
                c = ' ';
            }
        }
        stringstream ss(line);
        vector<int> nums;
        int temp;
        while(ss>>temp) {
            nums.push_back(temp);
        }
        Solution Sol;
        int res  =  Sol.searchInsert(nums,target);
        cout<<res<<endl;
    }
    return 0;
}