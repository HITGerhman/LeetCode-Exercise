#include <iostream>
#include <vector>
#include <sstream>
#include <cstdio>
#include <queue>

using namespace std;

const int INF = -999999;

struct TreeNode
{
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

vector<int> buildVector(string line)
{
    for (char &c : line)
    {
        if (c == '[' || c == ']' || c == ',')
        {
            c = ' ';
        }
    }
    stringstream ss(line);
    vector<int> nums;
    string temp;

    while (ss >> temp)
    {
        if (temp != "null")
        {
            nums.push_back(stoi(temp));
        }
    }
    return nums;
}

void printTree(TreeNode *root)
{
    if (!root)
    {
        cout << "[]" << endl;
        return;
    }
    vector<string> res;
    queue<TreeNode *> q;
    q.push(root);

    while (!q.empty())
    {
        TreeNode *curr = q.front();
        q.pop();
        if (curr)
        {
            res.push_back(to_string(curr->val));
            q.push(curr->left);
            q.push(curr->right);
        }
        else
        {
            res.push_back("null");
        }
    }

    while (!res.empty() && res.back() == "null")
    {
        res.pop_back();
    }

    cout << "[";
    for (int i = 0; i < res.size(); ++i)
    {
        cout << res[i] << (i == res.size() - 1 ? "" : ",");
    }
    cout << "]" << endl;
}

class Solution
{
public:
    TreeNode *sortedArrayToBST(vector<int> &nums)
    {
        return helper(nums, 0, nums.size() - 1);
    }
    TreeNode *helper(vector<int> &nums, int left, int right)
    {
        if (left > right)
        {
            return nullptr;
        }
        int mid = left + (right - left) / 2;

        TreeNode *root = new TreeNode(nums[mid]);
        root->left = helper(nums, left, mid - 1);
        root->right = helper(nums, mid + 1, right);

        return root;
    }
};
int main()
{
    if (!freopen("input.txt", "r", stdin))
    {
        cout << "无法打开 input.txt" << endl;
        return 0;
    }
    string line;
    while (getline(cin, line))
    {
        if (line.empty())
            continue;

        // 1. 构建数组
        vector<int> nums = buildVector(line);

        // 2. 运行算法
        Solution sol;
        TreeNode *root = sol.sortedArrayToBST(nums);

        // 3. 打印结果
        // 注意：输出结果可能包含末尾的 null，例如 [0,-3,9,-10,null,5,null,null,null]
        // 这是正确的，表示叶子节点的左右孩子为空。
        printTree(root);
    }
    return 0;
}