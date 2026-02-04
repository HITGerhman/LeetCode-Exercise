#include <iostream>
#include <sstream>
#include <cstdio>
#include <vector>
#include <algorithm>

using namespace std;

class Solution {
public:
    void MoveZero(vector<int>& nums) {
        int slow = 0;
        int fast = nums.size();
        for(int i = 0;i<fast;i++) {
            if(nums[i]!=0) {
                nums[slow] = nums[i];
                slow++;
            }
        }
        for(int i = slow;i<fast;i++) {
            nums[i] = 0;
        }
    }
};

int main() {
    freopen("input.txt","r",stdin);
    string line; 
    while(getline(cin,line)) {
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
        Sol.MoveZero(nums);

        for(int i;i<nums.size();i++) {
            cout << nums[i] << (i == nums.size() - 1 ? "" : " ");
        }
        cout<<endl;
    }
    return 0;
}