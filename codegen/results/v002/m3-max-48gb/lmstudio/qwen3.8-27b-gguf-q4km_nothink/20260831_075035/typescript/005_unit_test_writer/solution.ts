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
    { description: "Quantity below 10 (9 units) - no discount", price: 100, quantity: 9, expected: 100 * 9 },
    { description: "Quantity at 10 (boundary) - 10% discount", price: 100, quantity: 10, expected: 100 * 10 * 0.9 },
    { description: "Quantity at 49 (boundary) - 10% discount", price: 100, quantity: 49, expected: 100 * 49 * 0.9 },
    { description: "Quantity at 50 (boundary) - 25% discount", price: 100, quantity: 50, expected: 100 * 50 * 0.75 },
];

let allPassed = true;

for (const tc of testCases) {
    const got = calculateDiscount(tc.price, tc.quantity);
    // Compare with a small tolerance for floating point
    const isClose = Math.abs(got - tc.expected) < 1e-9;
    if (isClose) {
        console.log(`PASS: ${tc.description}`);
    } else {
        console.log(`FAIL: ${tc.description} (expected: ${tc.expected}, got: ${got})`);
        allPassed = false;
    }
}

if (allPassed) {
    console.log("\nAll tests passed!");
} else {
    console.log("\nSome tests failed.");
    process.exit(1);
}