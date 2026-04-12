using System;

if (Convert.ToInt32(Args[0]) >= 1) {
    int a = 1;
    int b = 1;
    Console.WriteLine(a);
    while (true) {
        int next = a + b;
        if (next > Convert.ToInt32(Args[0])) {
            break;
        }
        a = b;
        b = next;
        Console.WriteLine(b);
    }
}