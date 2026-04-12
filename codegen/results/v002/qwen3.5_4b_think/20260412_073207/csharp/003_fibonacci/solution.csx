int n = int.Parse(Args[0]);

if (n >= 1) {
    int a = 1;
    int b = 1;
    
    Console.WriteLine(a);
    
    while (b <= n) {
        Console.WriteLine(b);
        int next = a + b;
        a = b;
        b = next;
    }
}