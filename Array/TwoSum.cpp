#include <iostream>
#include <cstdio> // 必须包含这个头文件
#include <sstream>
#include <algorithm>
#include <vector>
#include <map>
#include <unordered_map>

using namespace std;
class Solution {
public:
    vector<int> twosum(vector<int>& nums,int target ){
        unordered_map<int,int>hashtable;
        for(int i=0;i<nums.size();++i) {
            int complement = target - nums[i];
            auto it = hashtable.find(complement);
            if(it != hashtable.end()) {
                return {it->second,i}; 
            }
            hashtable[nums[i]]=i;
        }
        return {};

    }
};

int main() {
    // 只要这行代码在，程序就会自动从目录下的 input.txt 读取数据
    // 而不需要你手动输入。本地测试完删掉这行即可提交。
    freopen("input.txt", "r", stdin); 
    string line;
    if(getline(cin,line)) {
        for(char &c:line) {
            if(!isdigit(c)&& c!='-') {
                c = ' ';
            }
        }
    }
    stringstream ss(line);
    vector<int> nums;
    int temp;
    while(ss>>temp) {
        nums.push_back(temp);
    }
    int target;
    cin>>target;

    Solution sol;
    vector<int> result= sol.twosum(nums,target);

    if(!result.empty()) {
        cout<<"["<<result[0]<<","<<result[1]<<"]"<<endl;
    }else {
        cout<<"No solution"<<endl;
    }
    
    return 0;
}