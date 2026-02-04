#include <iostream>
#include <vector>
#include <sstream>
#include <cstdio>

using namespace std;

struct ListNode
{
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

ListNode *buildList(const string &line)
{
    string cleanLine = line;
    for (char &c : cleanLine)
    {
        if (!isdigit(c) && c != '-')
        {
            c = ' ';
        }
    }
    stringstream ss(cleanLine);
    vector<int> nums;
    int temp;
    while (ss >> temp)
    {
        nums.push_back(temp);
    }

    ListNode *dummy = new ListNode(0);
    ListNode *cur = dummy;
    for (int x : nums)
    {
        cur->next = new ListNode(x);
        cur = cur->next;
    }
    ListNode *head = dummy->next;
    delete dummy;
    return head;
}

void printList(ListNode *head)
{
    cout << "[";
    ListNode *cur = head;
    while (cur)
    {
        cout << cur->val;
        if (cur->next)
        {
            cout << ",";
        }
        cur = cur->next;
    }
    cout << "]" << endl;
}

class Solution
{
public:
    ListNode *mergeTwoLists(ListNode *list1, ListNode *list2)
    {
        ListNode dummy(0);
        ListNode *tail = &dummy;
        while (list1 != nullptr && list2 != nullptr)
        {
            if (list1->val > list2->val)
            {
                tail->next = list2;
                list2 = list2->next;
            }
            else
            {
                tail->next = list1;
                list1 = list1->next;
            }
            tail = tail->next;
        }
        // 只要 list1 不为空，就把剩下的全部接上；否则接上 list2
        tail->next = (list1 != nullptr) ? list1 : list2;

        return dummy.next;
    }
};

int main()
{
    freopen("input.txt", "r", stdin);
    string line1, line2;
    if (getline(cin, line1), getline(cin, line2))
    {
        ListNode *head1 = buildList(line1);
        ListNode *head2 = buildList(line2);
        Solution sol;
        ListNode *res = sol.mergeTwoLists(head1, head2);
        printList(res);
    }
}