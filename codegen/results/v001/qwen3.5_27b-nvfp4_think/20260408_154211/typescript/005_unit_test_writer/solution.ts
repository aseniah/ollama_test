function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

type TestCase = {
    description: string;
    price: number;
    quantity: number;
    expected: number;
};

function runTest(testCase: TestCase): void {
    const result = calculateDiscount(testCase.price, testCase.quantity);
    const isPass = Math.abs(result - testCase.expected) < 0.01; // float comparison tolerance
    
    if (isPass) {
        console.log(`PASS: ${testCase.description}`);
    } else {
        console.log(`FAIL: ${testCase.description} (expected: ${testCase.expected}, got: ${result})`);
    }
}

function main(): void {
    console.log('Running calculateDiscount tests...\n');

    const testCases: TestCase[] = [
        {
            description: 'Quantity 9 (no discount)',
            price: 100,
            quantity: 9,
            expected: 100 * 9 // 900
        },
        {
            description: 'Quantity 10 (boundary - 10% discount starts)',
            price: 100,
            quantity: 10,
            expected: 100 * 10 * 0.9 // 900
        },
        {
            description: 'Quantity 49 (boundary - 10% discount ends)',
            price: 100,
            quantity: 49,
            expected: 100 * 49 * 0.9 // 4410
        },
        {
            description: 'Quantity 50 (boundary - 25% discount starts)',
            price: 100,
            quantity: 50,
            expected: 100 * 50 * 0.75 // 3750
        },
        {
            description: 'Quantity 1 (minimum - no discount)',
            price: 50,
            quantity: 1,
            expected: 50 * 1 // 50
        },
        {
            description: 'Quantity 75 (well into 25% discount tier)',
            price: 200,
            quantity: 75,
            expected: 200 * 75 * 0.75 // 11250
        }
    ];

    let passCount = 0;
    let failCount = 0;

    testCases.forEach((testCase) => {
        runTest(testCase);
        if (Math.abs(calculateDiscount(testCase.price, testCase.quantity) - testCase.expected) < 0.01) {
            passCount++;
        } else {
            failCount++;
        }
    });

    console.log(`\nResults: ${passCount} passed, ${failCount} failed`);
}

main();