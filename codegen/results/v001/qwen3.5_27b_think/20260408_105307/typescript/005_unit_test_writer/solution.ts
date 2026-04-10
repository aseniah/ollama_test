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

function runTest(test: TestCase): void {
    const result = calculateDiscount(test.price, test.quantity);
    // Allow for small floating point differences
    const isClose = Math.abs(result - test.expected) < 0.01;
    
    if (isClose) {
        console.log(`PASS: ${test.description}`);
    } else {
        console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
    }
}

// Test cases covering boundary conditions
const testCases: TestCase[] = [
    {
        description: "Quantity 9 - no discount (below 10)",
        price: 100,
        quantity: 9,
        expected: 900 // 100 * 9 * 1.0
    },
    {
        description: "Quantity 10 - 10% discount (boundary)",
        price: 100,
        quantity: 10,
        expected: 900 // 100 * 10 * 0.9
    },
    {
        description: "Quantity 49 - 10% discount (below 50)",
        price: 100,
        quantity: 49,
        expected: 4410 // 100 * 49 * 0.9
    },
    {
        description: "Quantity 50 - 25% discount (boundary)",
        price: 100,
        quantity: 50,
        expected: 3750 // 100 * 50 * 0.75
    }
];

console.log("Running calculateDiscount tests...\n");

testCases.forEach(test => runTest(test));