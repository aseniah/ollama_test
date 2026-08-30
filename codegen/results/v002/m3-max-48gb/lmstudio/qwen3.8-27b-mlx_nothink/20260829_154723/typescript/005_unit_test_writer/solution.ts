function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

interface TestCase {
    description: string;
    price: number;
    quantity: number;
    expected: number;
}

function runTest(testCase: TestCase): void {
    const got = calculateDiscount(testCase.price, testCase.quantity);
    // Use a small epsilon to account for floating point issues
    const epsilon = 1e-10;
    if (Math.abs(got - testCase.expected) < epsilon) {
        console.log(`PASS: ${testCase.description}`);
    } else {
        console.log(`FAIL: ${testCase.description} (expected: ${testCase.expected}, got: ${got})`);
    }
}

function main(): void {
    const testCases: TestCase[] = [
        {
            description: "quantity below discount threshold (quantity 9)",
            price: 10,
            quantity: 9,
            expected: 90, // 10 * 9 = 90
        },
        {
            description: "quantity at first discount threshold (quantity 10)",
            price: 10,
            quantity: 10,
            expected: 90, // 10 * 10 * 0.9 = 90
        },
        {
            description: "quantity just below second discount threshold (quantity 49)",
            price: 10,
            quantity: 49,
            expected: 441, // 10 * 49 * 0.9 = 441
        },
        {
            description: "quantity at second discount threshold (quantity 50)",
            price: 10,
            quantity: 50,
            expected: 375, // 10 * 50 * 0.75 = 375
        },
        {
            description: "quantity well above second discount threshold (quantity 100)",
            price: 20,
            quantity: 100,
            expected: 1500, // 20 * 100 * 0.75 = 1500
        },
        {
            description: "quantity zero (no purchase)",
            price: 5,
            quantity: 0,
            expected: 0, // 5 * 0 = 0
        },
    ];

    testCases.forEach(runTest);
}

main();