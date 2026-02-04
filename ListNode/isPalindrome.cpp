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

ListNode* buildList(const string& line) {
    string cleanLine = line;
    for(char& c:cleanLine) {
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
        cur = cur->next;
    }
    ListNode* head = dummy->next;
    delete dummy;
    return head;
}

class Solution {
public:
    bool isPalindrome(ListNode* head) {
        if(!head||!head->next) {
            return true;
        }
        ListNode* fast = head;
        ListNode* slow = head;
        while(fast->next&&fast->next->next) {
            slow=slow->next;
            fast=fast->next->next;
        }

        ListNode* prev = nullptr;
        ListNode* curr = slow->next;
        while(curr) {
            ListNode* temp = curr->next;
            curr->next = prev;
            prev=curr;
            curr=temp;
        }
        ListNode* p1 = head;
        ListNode* p2 = prev;
        while(p2) {
            if(p1->val!=p2->val){
                return false;
            }
            p1 = p1->next;
            p2 = p2->next;
        }
        return true;
}
};

int main(){
    freopen("input.txt","r",stdin);
    string line;
    getline(cin,line);
    ListNode* head = buildList(line);
    Solution sol;
    bool res = sol.isPalindrome(head) ;
    if(res) {
        cout<<"true"<<endl;
    } else {
        cout<<"false"<<endl;
    }
}