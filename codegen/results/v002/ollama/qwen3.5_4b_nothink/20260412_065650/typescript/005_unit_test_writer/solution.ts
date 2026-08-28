// Copy the discount calculation function
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases covering boundary conditions and typical cases
const testCases = [
    {
        price: 10,
        quantity: 5,
        expected: 50,
        description: "Less than 10 items with no discount"
    },
    {
        price: 10,
        quantity: 9,
        expected: 90,
        description: "Boundary: 9 items (no discount)"
    },
    {
        price: 10,
        quantity: 10,
        expected: 90,
        description: "Boundary: 10 items (90% discount)"
    },
    {
        price: 10,
        quantity: 49,
        expected: 441,
        description: "Boundary: 49 items (90% discount)"
    },
    {
        price: 10,
        quantity: 50,
        expected: 375,
        description: "Boundary: 50 items (75% discount)"
    },
    {
        price: 10,
        quantity: 100,
        expected: 750,
        description: "Large quantity with maximum discount"
    }
];

// Run tests and print results
for (const testCase of testCases) {
    const result = calculateDiscount(testCase.price, testCase.quantity);
    const isPass = result === testCase.expected;
    
    if (isPass) {
        console.log(`PASS: ${testCase.description}`);
    } else {
        console.log(`FAIL: ${testCase.description} (expected: ${testCase.expected}, got: ${result})`);
    }
}