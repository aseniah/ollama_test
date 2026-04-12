function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function runTest(description: string, expected: number, actual: number): void {
    const output = 
        actual === expected 
            ? `PASS: ${description}` 
            : `FAIL: ${description} (expected: ${expected}, got: ${actual})`;
    console.log(output);
}

function main() {
    // Define a test price and run test cases
    const testPrice = 100;
    
    // Test case 1: quantity = 9, no discount
    const result1 = calculateDiscount(testPrice, 9);
    const expected1 = 900;
    runTest("quantity=9 (no discount)", expected1, result1);

    // Test case 2: quantity = 10, 10% discount applies
    const result2 = calculateDiscount(testPrice, 10);
    const expected2 = 900;
    runTest("quantity=10 (10% discount)", expected2, result2);

    // Test case 3: quantity = 49, still 10% discount applies
    const result3 = calculateDiscount(testPrice, 49);
    const expected3 = 441;
    runTest("quantity=49 (10% discount)", expected3, result3);

    // Test case 4: quantity = 50, 75% discount applies
    const result4 = calculateDiscount(testPrice, 50);
    const expected4 = 3750;
    runTest("quantity=50 (25% off)", expected4, result4);

    // Additional test case: quantity < 10 (edge)
    const result5 = calculateDiscount(testPrice, 1);
    const expected5 = 100;
    runTest("quantity=1 (discount level 1)", expected5, result5);

    // Additional test case: quantity > 50
    const result6 = calculateDiscount(testPrice, 100);
    const expected6 = 7500;
    runTest("quantity=100 (25% off)", expected6, result6);
}

main();