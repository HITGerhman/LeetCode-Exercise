#include <sstream>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    void setZeroes(vector<vector<int>>& matrix) {
        int m = matrix.size();
        int n = matrix[0].size();
        bool col0 = false;

        for(int i)
    }
};

int main() {
    string line;
    while (getline(cin, line)) {
        if (line.empty()) continue;
        
        for (char &c : line) {
            if (c == '[' || c == ']' || c == ',') {
                c = ' ';
            }
        }
        
        stringstream ss(lxz2cine);
        vector<vector<int>> intervals;
        int start, end;

        while (ss >> start >> end) {
            intervals.push_back({start, end});
        }

        Solution sol;
        sol.setZeroes(intervals);

        cout << "[";
        for (int i = 0; i < intervals.size(); ++i) {
            cout << "[" << intervals[i][0] << "," << intervals[i][1] << "]";
            if (i < intervals.size() - 1) {
                cout << ",";
            }
        }
        cout << "]" << endl;
    }
    return 0;
}
