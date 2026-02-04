#include <iostream>
#include <sstream>
#include <vector>
#include <cstdio>

using namespace std;


struct ListNode {
    int val;
    ListNode* next;
    ListNode():val(0),next(nullptr){}
    ListNode(int x):val(x),next(nullptr){}
    ListNode(int x,ListNode* next):val(x),next(next){}
};

ListNode* buildList(const string& line,int pos) {
    string cleanString = line;
    for(char& c:cleanString) {
        if(!isdigit(c)&&c!='-') {
            c = ' ';
        }
    }
    stringstream ss(cleanString);
    vector<int> nums;
    int temp;
    while (ss>>temp)
    {
        nums.push_back(temp);
    }
    if(nums.empty()) {
        return nullptr;
    }
    //用vector记录每个节点的地址
    vector<ListNode*> nodes;
    ListNode* dummy = new ListNode(0);
    ListNode* cur = dummy->next;
    for(int x:nums) {
        cur = new ListNode(x);
        cur = cur->next;
        nodes.push_back(cur);
    }
    if(pos>=-1&&pos<(int)nodes.size()) {
        cur->next = nodes[pos];
    }
    ListNode* head = dummy->next;
    delete dummy;
    return head;
}

class Solution {
public:
    bool hasCycle(ListNode* head) {
        if(head == nullptr||head->next == nullptr) {
            return false;
        }
        ListNode* fast = head;
        ListNode* slow = head;
        while(fast!=nullptr&&fast->next!=nullptr) {
            slow = slow->next;
            fast = fast->next->next;
            if(slow = fast) {
                return true;
            }
        }
        return false;
    }
};

int main() {
    freopen("input.txt","r",stdin);
    string line;
    int pos;
    if (getline(cin, line) && cin >> pos) {
        ListNode* head = buildList(line, pos);
        Solution sol;
        bool hasCycle = sol.hasCycle(head);
        cout << (hasCycle ? "true" : "false") << endl;
    }
    return 0;
}