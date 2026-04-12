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
        description: "Quantity less than 10 (no discount)",
        price: 10,
        quantity: 9,
        expected: 90
    },
    {
        description: "Boundary: Quantity exactly 10 (10% discount)",
        price: 10,
        quantity: 10,
        expected: 90
    },
    {
        description: "Quantity less than 50 (10% discount)",
        price: 10,
        quantity: 49,
        expected: 441
    },
    {
        description: "Boundary: Quantity exactly 50 (25% discount)",
        price: 10,
        quantity: 50,
        expected: 375
    },
    {
        description: "Quantity greater than 50 (25% discount)",
        price: 20,
        quantity: 100,
        expected: 1500
    }
];

let passed = 0;
let failed = 0;

console.log("Running tests for calculateDiscount function\n");

for (const testCase of testCases) {
    const result = calculateDiscount(testCase.price, testCase.quantity);
    
    // Use strict equality with a small epsilon for floating point comparison
    const match = Math.abs(result - testCase.expected) < 0.001;
    
    if (match) {
        console.log(`PASS: ${testCase.description}`);
        passed++;
    } else {
        console.log(`FAIL: ${testCase.description} (expected: ${testCase.expected}, got: ${result})`);
        failed++;
    }
}

console.log(`\n${passed}/${testCases.length} tests passed`);