// Original function to test
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test runner
interface TestCase {
    description: string;
    price: number;
    quantity: number;
    expected: number;
}

const testCases: TestCase[] = [
    {
        description: "No discount for quantity 9",
        price: 100,
        quantity: 9,
        expected: 900
    },
    {
        description: "10% discount for quantity 10 (boundary)",
        price: 100,
        quantity: 10,
        expected: 900
    },
    {
        description: "10% discount for quantity 49 (boundary)",
        price: 100,
        quantity: 49,
        expected: 4410
    },
    {
        description: "25% discount for quantity 50 (boundary)",
        price: 100,
        quantity: 50,
        expected: 3750
    },
    {
        description: "Large quantity with 25% discount",
        price: 200,
        quantity: 100,
        expected: 15000
    }
];

// Run tests
let passCount = 0;
let failCount = 0;

for (const test of testCases) {
    const result = calculateDiscount(test.price, test.quantity);
    // Use a small epsilon for floating point comparison
    const epsilon = 0.0001;
    
    if (Math.abs(result - test.expected) < epsilon) {
        console.log(`PASS: ${test.description}`);
        passCount++;
    } else {
        console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
        failCount++;
    }
}

console.log(`\nResults: ${passCount} passed, ${failCount} failed`);