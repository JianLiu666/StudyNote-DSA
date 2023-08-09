#include <vector>
#include <set>

using namespace std;

class Solution {
public:
    vector<int> maxSlidingWindow(vector<int>& nums, int k) {
        // BST 維護 sliding window
        multiset<int>Set;
        vector<int>result;

        for (int i=0; i < nums.size(); i++)
        {
            Set.insert(nums[i]);
            if (Set.size() > k)
            {
                Set.erase(Set.lower_bound(nums[i-k]));
            }
            if (Set.size() == k)
            {
                result.push_back(*Set.rbegin());
            }
        }

        return result;
    }
};