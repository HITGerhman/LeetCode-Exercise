#include <unordered_set>
#include <iostream>
#include <vector>
#include <sstream>



using namespace std;

class Solution {
public:
    int longestConsecutive(vector<int>& nums) {
        unordered_set<int> num_set;
        for(int num:nums) {
            num_set.insert(num);
        }
        int longestStreak = 0;

        for(int num:nums) {
            if(!num_set.count(num-1)) {
                int currentNum = num;
                int currentSteak = 1;
            while (num_set.count(currentNum+1))
            {
                currentNum ++;
                currentSteak ++;

            }
            longestStreak = max(longestStreak,currentSteak);
            
            }
        }
        return longestStreak;

    }
};
int main() {
    if(!freopen("input.txt","r",stdin)) {
        return 0;
    }
    string line;
    while(getline(cin,line)) {
        for(char& c:line) {
            if(c=='['||c==']'||c==',') {
                c = ' ';
            }
        }
        stringstream ss(line) ;
        vector<int> nums;
        int temp;
        while(ss>>temp) {
            nums.push_back(temp);
        }
        Solution sol;
        int res = sol.longestConsecutive(nums); 
        cout<<res<<endl;
        
    }

}