function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases definition
const testCases = [
    { desc: "Single item purchase at $10", expected: 10, price: 10, quantity: 1 },
    { desc: "Boundary 9 items at $100 (discount tier 1)", expected: 900, price: 100, quantity: 9 },
    { desc: "Boundary 10 items at $100 (first discount tier applies)", expected: 900, price: 100, quantity: 10 },
    { desc: "Boundary 49 items at $200 (last discount tier 1)", expected: 8820, price: 200, quantity: 49 },
    { desc: "Boundary 50 items at $200 (second discount tier applies)", expected: 7500, price: 200, quantity: 50 },
];

// Run tests
let passedCount = 0;
const results: string[] = [];

for (const tc of testCases) {
    const actual = calculateDiscount(tc.price, tc.quantity);
    
    if (actual === tc.expected) {
        console.log(`PASS: ${tc.desc}`);
        passedCount++;
    } else {
        console.log(`FAIL: ${tc.desc} (expected: ${tc.expected}, got: ${actual})`);
    }
}

console.log(`\n${passedCount}/${testCases.length} tests passed.`);