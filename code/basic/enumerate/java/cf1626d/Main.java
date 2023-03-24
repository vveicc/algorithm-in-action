import java.io.*;
import java.util.*;
import java.util.function.Function;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(System.out);

    public static void main(String[] args) throws Exception {
        int T = Integer.parseInt(in.readLine());
        for (; T > 0; T--) {
            int n = Integer.parseInt(in.readLine());
            StringTokenizer at = new StringTokenizer(in.readLine());
            Integer[] a = new Integer[n];
            for (int i = 0; i < n; i++) {
                a[i] = Integer.parseInt(at.nextToken());
            }
            Arrays.sort(a);

            int ans = cost(n) + 2;
            // 枚举与第一组的大小最接近的2的幂次
            for (int i = 1; ans != 0 && i < n; i <<= 1) {
                int cx = search(a, a[i]); // < x = a[i]
                ans = Math.min(ans, i - cx + cost(n - cx) + 1);
                // 枚举与第二组的大小最接近的2的幂次
                for (int j = 1; ans != 0 && cx + j < n; j <<= 1) {
                    int cy = search(a, a[cx + j]); // < y = a[cx+j]
                    ans = Math.min(ans, i + j - cy + cost(n - cy));
                }
            }
            out.println(ans);
        }
        out.flush();
        out.close();
        in.close();
    }

    public static int bits(int x) {
        int c = 0;
        for (; x != 0; x >>= 1) {
            c++;
        }
        return c;
    }

    public static int cost(int x) {
        int l = bits(x);
        if (x == 1 << l >> 1) {
            return 0;
        } else {
            return (1 << l) - x;
        }
    }

    public static <T extends Comparable<T>> int search(T[] arr, T x) {
        return search(arr.length, i -> arr[i].compareTo(x) != -1);
    }

    public static int search(int n, Function<Integer, Boolean> f) {
        int i = 0, j = n;
        while (i < j) {
            int h = i + ((j - i) >> 1);
            if (!f.apply(h)) {
                i = h + 1; // f(i-1) = false
            } else {
                j = h; // f(j) = true
            }
        }
        return i;
    }
}
