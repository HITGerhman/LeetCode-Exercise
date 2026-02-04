#include <iostream>

using namespace std;

class Solution {
public:
    int climbStairs(int n) {
        if(n == 1) {
            return 1;
        }
        if(n == 2) {
            return 2;
        }
        int a = 1;
        int b = 2;
        int sum = 0;

        for(int i = 3;i<=n;i++) {
            sum = a+b;
            a = b;
            b = sum;
        }
        return b;
    }
};
int main() {
    freopen("input.txt","r",stdin);
    int stairs;
    cin>>stairs;
    Solution sol;
    int res = sol.climbStairs(stairs) ;
    cout<<res<<endl;
}