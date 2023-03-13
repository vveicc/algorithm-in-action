import java.io.*;
import java.util.*;

public class Main {

    static BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
    static PrintWriter out = new PrintWriter(new BufferedWriter(new OutputStreamWriter(System.out)));

    static int n;
    static String s;

    public static void main(String[] args) throws Exception {
        int t = Integer.parseInt(in.readLine());
        for (; t > 0; t--) {
            StringTokenizer tokenizer = new StringTokenizer(in.readLine());
            n = Integer.parseInt(tokenizer.nextToken());
            s = tokenizer.nextToken();
            int[] a = new int[n];
            int[] b = new int[n];
            for (int i = 0; i < n; i++) {
                a[i] = n - i; // 数组a的LIS最短，先构造成降序排列
                b[i] = i + 1; // 数组b的LIS最长，先构造成升序排列
            }
            reverse(a, '<'); // 翻转数组a中的连续'<'段
            reverse(b, '>'); // 翻转数组b中的连续'>'段
            for (int x : a) {
                out.printf("%d ", x);
            }
            out.println();
            for (int x : b) {
                out.printf("%d ", x);
            }
            out.println();
        }
        out.flush();
        out.close();
        in.close();
    }

    public static void reverse(int[] arr, char sign) {
        for (int i = 0; i < n - 1; i++) {
            if (s.charAt(i) == sign) {
                int l = i;
                for (i++; i < n - 1 && s.charAt(i) == sign; i++) {
                }
                for (int r = i; l < r; r--) {
                    int x = arr[l];
                    arr[l] = arr[r];
                    arr[r] = x;
                    l++;
                }
            }
        }
    }
}
