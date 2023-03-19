import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));

    public static void main(String[] args) throws Exception {
        StringTokenizer nk = new StringTokenizer(in.readLine());
        int n = Integer.parseInt(nk.nextToken());
        int k = Integer.parseInt(nk.nextToken());
        long ans = c2(n + 1); // 异或前缀和数对总数
        int s = 0, mask = (1 << k) - 1;
        Map<Integer, Integer> cnt = new HashMap<>();
        cnt.put(s, 1);
        StringTokenizer tokenizer = new StringTokenizer(in.readLine());
        for (; n > 0; n--) {
            s ^= Integer.parseInt(tokenizer.nextToken());
            // s和s^mask的个数一起统计
            cnt.compute(Math.min(s, s ^ mask), (v, c) -> c == null ? 1 : c + 1);
        }
        for (int c : cnt.values()) {
            // 减去相同数对数目，平均分配s和s^mask的个数，使相同异或前缀和数对数目最少
            ans -= c2(c >> 1) + c2((c + 1) >> 1);
        }
        System.out.println(ans);
        in.close();
    }

    public static long c2(int n) {
        return (((long) n) * ((long) (n - 1))) >> 1; // n个数字的不同数对数目
    }
}
