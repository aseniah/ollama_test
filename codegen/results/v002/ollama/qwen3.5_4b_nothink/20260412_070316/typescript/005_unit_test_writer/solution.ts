function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases
const testCases = [
    {
        description: "Quantity less than 10 (boundary: 9)",
        price: 10,
        quantity: 9,
        expected: 90
    },
    {
        description: "Quantity between 10 and 49 (boundary: 10)",
        price: 10,
        quantity: 10,
        expected: 90
    },
    {
        description: "Quantity near boundary 50 (boundary: 49)",
        price: 10,
        quantity: 49,
        expected: 441
    },
    {
        description: "Quantity at boundary 50",
        price: 10,
        quantity: 50,
        expected: 375
    }
];

// Run tests
testCases.forEach((tc) => {
    const result = calculateDiscount(tc.price, tc.quantity);
    const passed = result === tc.expected;
    
    if (passed) {
        console.log(`PASS: ${tc.description}`);
    } else {
        console.log(`FAIL: ${tc.description} (expected: ${tc.expected}, got: ${result})`);
    }
});