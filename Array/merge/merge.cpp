#include <algorithm>
#include <iostream>
#include <vector>
#include <sstream>

using namespace std;

class Solution {
public:
    vector<vector<int>> merge(vector<vector<int>>& intervals) {
        if(intervals.empty()) {
            return{};
        }
        sort(intervals.begin(),intervals.end());
        vector<vector<int>> merged;
        merged.push_back(intervals[0]);
        for(int i = 1;i<intervals.size();i++) {
            vector<int>& last = merged.back();
            vector<int>& curr = intervals[i];
            if(curr[0]<=last[1]) {
                last[1] = max(last[1],curr[1]);
            }else {
                merged.push_back(curr);
            }
        }
    return merged;
    }
};
int main() {
    if (!freopen("input.txt", "r", stdin))
    {
        return 0;
    }
    string line;
    while(getline(cin,line)) {
        if(line.empty()) {
            continue;
        }
        for (char &c : line) {
        if (c == '[' || c == ']' || c == ',') {
                c = ' ';
            }
        }
        stringstream ss(line);
        vector<vector<int>> intervals;
        int start,end;
        while(ss>>start>>end) {
            intervals.push_back({start,end});
        }
        Solution sol;
        vector<vector<int>> res = sol.merge(intervals);
        cout<<"[";
        for(int i = 0;i<res.size();++i) {
            cout<<"["<<res[i][0]<<","<<res[i][1]<<"]";
            if(i<res.size()-1) {
                cout<<",";
            }
        }
        cout<<"]"<<endl;
    }
}