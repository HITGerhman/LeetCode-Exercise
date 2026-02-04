#include <iostream>
#include <stack>
#include <algorithm>
#include <sstream>

using namespace std;

class Solution
{
public:
    bool isValid(string s)
    {
        stack<char> st;
        
        for (char c : s)
        {
            // 2. 新增：遇到空格、换行、引号，直接跳过
            if (c == ' ' || c == '\r' || c == '\n' || c == '"') {
                continue;
            }
            if (c == '(')
            {
                st.push(')');
            }
            else if (c == '[')
            {
                st.push(']');
            }
            else if (c == '{')
            {
                st.push('}');
            } else {
                if(st.empty()||c!=st.top()) {
                    return false;
                }
                st.pop();
            }
        }
        return st.empty();
    }
};

int main()
{
    if (!freopen("input.txt", "r", stdin))
    {
        cout << "找不到" << endl;
        return 0;
    }
    string line;
    while (getline(cin, line))
    {
        if(!line.empty()&&line.back()=='\r') {
            line.pop_back();
        }
        while(!line.empty()&&isspace(line.back())) {
            line.pop_back();
        }
        Solution sol;
        bool res = sol.isValid(line);
        if (res)
        {
            cout << "true" << endl;
        }
        else
        {
            cout << "false" << endl;
        }
    }
}