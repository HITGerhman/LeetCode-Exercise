#include <iostream>
#include <vector>
#include <sstream>

using namespace std;


class Solution {
public:
    vector<int> findAnagrams(string s,string p) {
        int s_len = s.size();
        int p_len = p.size();

        if(s_len<p_len) {
            return;
        }

        vector<int> res;
        vector<int> pcount(26,0);
        vector<int> windowCount(26,0);

        for(int i = 0;i<p_len;i++) {
        
        }
    }
};

int main() {
    if(!freopen("input.txt","r",stdin)) {
        return 0;
    }
    string line;
    string target;
    while(getline(cin,line)&&getline(cin,target)) {
        if(line.empty())     {
            continue;
        }
        if(line.size()>=2&&line.front()=='"'&&line.back()=='"') {
            line = line.substr(1,line.size()-2);
        }
        
        if(target.size()>=2&&target.front()=='"'&&target.back()=='"') {
            target = target.substr(1,target.size()-2);
        }
        vector<int> res;
        Solution sol;
        res = sol.findAnagrams(line,target);
        cout<<'"';
        for(int i;i<res.size();i++) {
            cout<<res[i];
            if(i!=res.size()-1) {
                cout<<",";
            }
        }
         cout<<'"';

    }
}