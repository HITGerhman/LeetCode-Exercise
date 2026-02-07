#include <iostream>
#include <vector>
#include <cctype>
#include <string>

using namespace std;

class Solution {
public:
    void setZeroes(vector<vector<int>>& matrix) {
        int m = matrix.size();
        if (m == 0) return;
        int n = matrix[0].size();
        if (n == 0) return;

        bool row0 = false;
        bool col0 = false;

        for (int j = 0; j < n; j++) {
            if (matrix[0][j] == 0) {
                row0 = true;
                break;
            }
        }
        for (int i = 0; i < m; i++) {
            if (matrix[i][0] == 0) {
                col0 = true;
                break;
            }
        }

        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                if (matrix[i][j] == 0) {
                    matrix[i][0] = 0;
                    matrix[0][j] = 0;
                }
            }
        }

        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                if (matrix[i][0] == 0 || matrix[0][j] == 0) {
                    matrix[i][j] = 0;
                }
            }
        }

        if (row0) {
            for (int j = 0; j < n; j++) {
                matrix[0][j] = 0;
            }
        }
        if (col0) {
            for (int i = 0; i < m; i++) {
                matrix[i][0] = 0;
            }
        }
    }
};

int main() {
    string line;
    while (getline(cin, line)) {
        if (line.empty()) continue;

        vector<vector<int>> matrix;
        vector<int> row;
        int num = 0, sign = 1;
        bool inNum = false;
        for (char c : line) {
            if (c == '-') sign = -1;
            else if (isdigit(static_cast<unsigned char>(c))) {
                num = num * 10 + (c - '0');
                inNum = true;
            } else {
                if (inNum) {
                    row.push_back(sign * num);
                    num = 0;
                    sign = 1;
                    inNum = false;
                }
                if (c == ']' && !row.empty()) {
                    matrix.push_back(row);
                    row.clear();
                }
            }
        }

        Solution sol;
        sol.setZeroes(matrix);

        cout << "[";
        for (int i = 0; i < (int)matrix.size(); ++i) {
            cout << "[";
            for (int j = 0; j < (int)matrix[i].size(); ++j) {
                cout << matrix[i][j];
                if (j + 1 < (int)matrix[i].size()) cout << ",";
            }
            cout << "]";
            if (i + 1 < (int)matrix.size()) cout << ",";
        }
        cout << "]" << endl;
    }
    return 0;
}
