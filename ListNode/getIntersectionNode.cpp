#include <iostream>
#include <vector>
#include <sstream>
#include <cstdio>

using namespace std;

struct ListNode {
    int val;
    ListNode* next;
    ListNode(int x):val(x),next(NULL){};

};
class Solution {
public:
    ListNode* getIntersectionNode(ListNode* headA,ListNode* headB) { 
        if(headA==nullptr||headB== nullptr) {
            return nullptr;
        }

        ListNode *pA=headA;
        ListNode *pB=headB;
        while (pA!=pB)
        {
            pA = (pA == nullptr)?headB:pA->next;
            pB = (pB == nullptr)?headA:pB->next;
            
        }
        return pA;
    }
};
ListNode* buildList(const string& line) {
    string cleanLine = line;
    for(char &c:cleanLine) {
        if(!isdigit(c)&&c!='-') {
            c = ' ';
        }
    }
    stringstream ss(cleanLine);
    vector<int> nums;
    int temp;
    while(ss>>temp) {
        nums.push_back(temp);
    }

    ListNode* dummy = new ListNode(0);
    ListNode* cur = dummy;
    for(int x:nums) {
        cur->next = new ListNode(x);
        cur=cur->next;
    }
    ListNode* head=dummy->next; 
    delete dummy;
    return head;
}
int main() {
    freopen("input.txt","r",stdin);
    string lineA,lineB;
    if(getline(cin,lineA)&&getline(cin,lineB)) {
        ListNode* headA  = buildList(lineA);
        ListNode* headB  = buildList(lineB);
        Solution sol;
        ListNode* result=sol.getIntersectionNode(headA,headB);
        if(result) {
            cout<<result->val<<endl;
        } else {
            cout<<"no intersection"<<endl;
        }
    }
}