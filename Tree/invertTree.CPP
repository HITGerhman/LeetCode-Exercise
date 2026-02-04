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

TreeNode *buildTree(const string &line)
{
    string cleanLine = line;
    for (char &c : cleanLine)
    {
        if (c == '[' || c == ']' || c == ',')
        {
            c = ' ';
        }
    }
    stringstream ss(cleanLine);
    vector<string> nodes;
    string temp;
    while (ss >> temp)
    {
        nodes.push_back(temp);
    }
    if (nodes.empty() || nodes[0] == "null")
    {
        return nullptr;
    }
    TreeNode *root = new TreeNode(stoi(nodes[0]));
    queue<TreeNode *> q;
    q.push(root);

    int i = 1;
    while (!q.empty() && i < nodes.size())
    {
        TreeNode *curr = q.front();
        q.pop();
        if (i < nodes.size() && nodes[i] != "null")
        {
            curr->left = new TreeNode(stoi(nodes[i]));
            q.push(curr->left);
        }
        i++;
        if (i < nodes.size() && nodes[i] != "null")
        {
            curr->right = new TreeNode(stoi(nodes[i]));
            q.push(curr->right);
        }
        i++;
    }
    return root;
}

void printLevelOrder(TreeNode *root)
{
    if (!root)
    {
        cout << "";
        return;
    }
    queue<TreeNode *> q;
    q.push(root);
    cout << "[";

    bool first = true;
    while (!q.empty())
    {
        TreeNode *curr = q.front();
        q.pop();
        if (!first)
        {
            cout << ",";
        }
        cout << curr->val;
        first = false;
        if (curr->left)
        {
            q.push(curr->left);
        }
        if (curr->right)
        {
            q.push(curr->right);
        }
    }
    cout << "]";
}

class Solution
{
public:
    TreeNode *invertTree(TreeNode *root)
    {
        if (root == nullptr)
        {
            return nullptr;
        }

        swap(root->left, root->right);

        invertTree(root->left);
        invertTree(root->right);
        return root;
    }
};

int main()
{
    if (!freopen("input.txt", "r", stdin))
    {
        return 0;
    }
    string line;
    while (getline(cin, line))
    {
        if (line.empty())
        {
            continue;
        }
        TreeNode *root = buildTree(line);
        Solution sol;
        TreeNode *res = sol.invertTree(root);
        printLevelOrder(res);
    }
}
