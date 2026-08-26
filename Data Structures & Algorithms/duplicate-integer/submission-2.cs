public class Solution {
    public bool hasDuplicate(int[] nums) {
        HashSet<int> numbers = new HashSet<int>();
        foreach (var n in nums){
            if (!numbers.Add(n)) return true;
        }
        return false;
    }
}