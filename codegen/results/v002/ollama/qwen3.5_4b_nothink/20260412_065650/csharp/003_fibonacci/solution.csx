var args = Args;
if (int.TryParse(args[0], out int n)) {
    int a = 1;
    int b = 1;
    
    if (n >= 1) {
        Console.WriteLine(a);
        Console.WriteLine(b);
        
        // Calculate subsequent numbers and check if they exceed n
        while (true) {
            int next = a + b;
            if (next > n) {
                break;
            }
            Console.WriteLine(next);
            a = b;
            b = next;
        }
    }
}