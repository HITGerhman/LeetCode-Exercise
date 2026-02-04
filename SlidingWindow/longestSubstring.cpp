#include <unordered_map>
#include <iostream>
#include <sstream>
#include <vector>


using namespace std;

class Solution {
public:
    int lengthOfLengestSubstring(string s) {
        unordered_map<char,int> mp;
        int maxLength = 0;
        int left = 0;

        for(int right = 0;right<s.size();right++) {
            char currentChar = s[right];
            if(mp.count(currentChar)&&mp[currentChar]>=left) {
                left = mp[currentChar]+1;
            }
            mp[currentChar] = right;
            maxLength = max(maxLength,right - left+1);
            
        }
        return maxLength;
    }

};

int main() {
    string line;
    if(!freopen("input.txt","r",stdin)) {
        return 0;
    }
    while(getline(cin,line)) {
        if(line.empty()) {
            continue;
        }
        // 修正点 2：如果是带引号的测试用例（如 "abc"），去除首尾引号
        if (line.size() >= 2 && line.front() == '"' && line.back() == '"') {
            line = line.substr(1, line.size() - 2);
        }
        int res;
        Solution sol;
        res = sol.lengthOfLengestSubstring(line);
        cout<<res<<endl;

    }
}