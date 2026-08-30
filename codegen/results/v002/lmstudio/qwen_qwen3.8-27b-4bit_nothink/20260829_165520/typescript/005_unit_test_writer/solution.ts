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

const testCases: TestCase[] = [
    {
        description: "quantity 9 (below discount threshold)",
        price: 10,
        quantity: 9,
        expected: 90,
    },
    {
        description: "quantity 10 (at 10% discount threshold)",
        price: 10,
        quantity: 10,
        expected: 90,
    },
    {
        description: "quantity 49 (just below 25% discount threshold)",
        price: 10,
        quantity: 49,
        expected: 441,
    },
    {
        description: "quantity 50 (at 25% discount threshold)",
        price: 10,
        quantity: 50,
        expected: 375,
    },
    {
        description: "quantity 1 (minimal purchase, no discount)",
        price: 25.5,
        quantity: 1,
        expected: 25.5,
    },
    {
        description: "quantity 25 (middle of 10% discount range)",
        price: 20,
        quantity: 25,
        expected: 450,
    },
];

function runTests(testCases: TestCase[]): void {
    console.log("Running tests...\n");
    let passCount = 0;
    let failCount = 0;

    for (const test of testCases) {
        const result = calculateDiscount(test.price, test.quantity);
        // Use a small epsilon for floating-point comparison
        const isPass = Math.abs(result - test.expected) < 1e-9;

        if (isPass) {
            console.log(`PASS: ${test.description}`);
            passCount++;
        } else {
            console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
            failCount++;
        }
    }

    console.log(`\n${passCount} passed, ${failCount} failed, ${passCount + failCount} total.`);
}

runTests(testCases);