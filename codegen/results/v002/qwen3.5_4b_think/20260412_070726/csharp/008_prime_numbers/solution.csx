int N = int.Parse(Args[0]);

for (int i = 2; i <= N; i++) {
    bool isPrime = true;
    for (int j = 2; j * j <= i; j++) {
        if (i % j == 0) {
            isPrime = false;
            break;
        }
    }
    
    if (isPrime) {
        Console.WriteLine(i);
    }
}