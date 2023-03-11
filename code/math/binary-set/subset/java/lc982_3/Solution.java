class Solution {
    public int countTriplets(int[] nums) {
        int[] cnt = new int[1 << 16];
        cnt[0] = nums.length;
        for (int x : nums) {
            x ^= 0xffff;
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
