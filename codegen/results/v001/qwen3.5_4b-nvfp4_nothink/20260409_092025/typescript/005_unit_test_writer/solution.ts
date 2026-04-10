function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases covering boundary conditions and general logic
const testCases = [
    { desc: "Single item (boundary below 10)", expected: 20, price: 20, quantity: 1 },
    { desc: "Just under threshold for 10% discount (quantity 9)", expected: 18, price: 2, quantity: 9 },
    { desc: "At threshold for 10% discount (quantity 10)", expected: 18, price: 1.8, quantity: 10 },
    { desc: "Just under threshold for 25% discount (quantity 49)", expected: 46.68, price: 1.8, quantity: 49 },
    { desc: "At threshold for 25% discount (quantity 50)", expected: 36, price: 1.8, quantity: 50 }
];

// Main execution block
let successCount = 0;
const totalTests = testCases.length;

for (const tc of testCases) {
    const result = calculateDiscount(tc.price, tc.quantity);
    
    if (result === tc.expected) {
        console.log(`PASS: ${tc.desc}`);
        successCount++;
    } else {
        console.log(`FAIL: ${tc.desc} (expected: ${tc.expected}, got: ${result})`);
    }
}

console.log(`\nResults: ${successCount}/${totalTests} tests passed.`);