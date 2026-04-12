function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Read command-line arguments
const price = parseInt(process.argv[2] || "100", 10);

// Test cases covering boundary conditions
const testCases = [
    { qty: 9, desc: "Quantity < 10 (boundary)" },
    { qty: 10, desc: "Quantity >= 10 (boundary)" },
    { qty: 49, desc: "Quantity < 50 (boundary)" },
    { qty: 50, desc: "Quantity >= 50 (boundary)" },
];

// Run tests
testCases.forEach(test => {
    const result = calculateDiscount(price, test.qty);
    
    // Calculate expected value manually based on the function's logic
    let expected: number;
    if (test.qty < 10) {
        expected = price * test.qty;
    } else if (test.qty < 50) {
        expected = price * test.qty * 0.9;
    } else {
        expected = price * test.qty * 0.75;
    }

    // Compare results with a small epsilon for floating point safety
    const isPass = Math.abs(result - expected) < 0.00001;

    if (isPass) {
        console.log(`PASS: ${test.desc}`);
    } else {
        console.log(`FAIL: ${test.desc} (expected: ${expected}, got: ${result})`);
    }
});