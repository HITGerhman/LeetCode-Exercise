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
    for(char &c:cleanLine) {
        if(!isdigit(c)&&c!='-') {
            c = ' ';
        }
    }
    stringstream ss(cleanLine) ;
    vector<int> nums ;
    int temp;
    while(ss>>temp) {
        nums.push_back(temp);
    }
    ListNode* dummy = new ListNode(0);
    ListNode* cur = dummy;
    for(int x:nums) {
        cur->next =new ListNode(x);
        cur = cur->next;
    }
    ListNode* head = dummy->next;
    delete dummy;
    return head;
}

class Solution {
public:
    ListNode* ReverseList(ListNode* head) {
        ListNode* prev=nullptr;
        ListNode* curr = head;
        while(curr) {
            ListNode* temp = curr->next;
            curr->next = prev;
            prev = curr;
            curr = temp;
        }
        return prev;
    }
};

int main() {
    freopen("input.txt","r",stdin); 
    string line;
    getline(cin,line);
    ListNode* head = buildList(line);
    Solution sol;
    ListNode* res = sol.ReverseList(head) ;
    while(res) {
        cout<<res->val;
        res = res->next;
    }
}