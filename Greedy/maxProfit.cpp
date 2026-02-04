#include <iostream>
#include <stack>
#include <algorithm>
#include <sstream>
#include <vector>

using namespace std;

class Solution {
public:
    int maxProfit(vector<int>& prices) {
        if(prices.size()<1) {
            return 0;
        }
        int min = prices[0];
        int max = 0;
        

        for(int c:prices) {
            if(min>c) {
                min = c;

            }
          else  if(c-min>max) {
                max = c-min;
            }
        }
        return max;
    }

};
int main() {
    if (!freopen("input.txt", "r", stdin))
    {
        cout << "找不到" << endl;
        return 0;
    }
    string line;
    while (getline(cin, line))
    {
        for(char& c:line) {
            if(c == '['||c==']'||c==',') {
                c = ' ';

            }
        }
        stringstream ss(line);
        vector<int> nums;
        int temp;
        while(ss>>temp) {
            nums.push_back(temp);
            
        }

        Solution sol;
        int res = sol.maxProfit(nums);
        cout<<res<<endl;
    }
}