class Solution {
    public int countTriplets(int[] nums) {
        int[] cnt = new int[1 << 16];
        for (int x : nums) {
            for (int y : nums) {
                cnt[x & y]++;
            }
        }
        int ans = 0;
        for (int x : nums) {
            x ^= 0xffff;
            for (int s = x; s != 0; s = (s - 1) & x) {
                ans += cnt[s];
            }
            ans += cnt[0];
        }
        return ans;
    }
}
