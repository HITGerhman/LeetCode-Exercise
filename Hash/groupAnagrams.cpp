#include <iostream>
#include <vector>
#include <algorithm>
#include <sstream>
#include <unordered_map>

using namespace std;

class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        unordered_map<string,vector<string>> mp;
        for(string s:strs) {
            string key = s;

            sort(key.begin(),key.end());
            mp[key].push_back(s);
        }
        vector<vector<string>> result;
        for(auto it:mp) {
            result.push_back(it.second);

        }
        return result;
    }
};
int main() {
    if(!freopen("input.txt","r",stdin))  {
        return 0;
    }
    string line;
    while(getline(cin,line)) {
        if(line.size()<2) {
            continue;
        }
        vector<string> strs;

        size_t start = line.find('[');
        size_t end = line.find(']');

        if(start != string::npos && end!=string::npos && end>start) {
            string content = line.substr(start+1,end-start-1);
            stringstream ss(content);
            string segment;

            while(getline(ss,segment,',')) {
                
            }
        }
    }
}