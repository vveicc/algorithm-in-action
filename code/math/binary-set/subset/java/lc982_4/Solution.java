class Solution {
    public int countTriplets(int[] nums) {
        int m = 1;
        for (int x : nums) {
            for (; m <= x; m <<= 1) {
            }
        }
        int[] cnt = new int[m];
        cnt[0] = nums.length;
        int mask = m - 1;
        for (int x : nums) {
            x ^= mask;
            for (int s = x; s != 0; s = (s - 1) & x) {
                cnt[s]++;
            }
        }
        int ans = 0;
        for (int x : nums) {
            for (int y : nums) {
                ans += cnt[x & y];
            }
        }
        return ans;
    }
}
