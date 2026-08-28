if (Args.Count > 0) {
    int n = int.Parse(Args[0]);
    if (n < 1) {
        return;
    }
    long a = 1;
    long b = 1;
    long c = 0;
    do {
        if (a <= n) {
            Console.WriteLine(a);
        }
        c = a + b;
        a = b;
        b = c;
    } while (a <= n);
}