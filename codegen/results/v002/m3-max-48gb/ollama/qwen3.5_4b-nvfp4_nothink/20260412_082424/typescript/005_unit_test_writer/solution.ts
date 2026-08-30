function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function main() {
    console.log("Running discount tests...\n");

    // Test Case 1: Quantity < 10 (No discount)
    const tc1Price = 10;
    const tc1Qty = 9;
    const tc1Expected = calculateDiscount(tc1Price, tc1Qty);
    const tc1Actual = tc1Price * tc1Qty; // Should be 90
    if (tc1Expected === tc1Actual) {
        console.log(`PASS: No discount for quantity ${tc1Qty} at price ${tc1Price} (result: ${tc1Expected})`);
    } else {
        console.log(`FAIL: No discount for quantity ${tc1Qty} at price ${tc1Price}`);
        console.log(`  Expected: ${tc1Actual}, Got: ${tc1Expected}`);
    }

    // Test Case 2: Quantity = 10 (9% discount)
    const tc2Price = 20;
    const tc2Qty = 10;
    const tc2Expected = calculateDiscount(tc2Price, tc2Qty);
    const tc2Actual = tc2Price * tc2Qty * 0.9; // Should be 180
    if (tc2Expected === tc2Actual) {
        console.log(`PASS: 9% discount for quantity ${tc2Qty} at price ${tc2Price} (result: ${tc2Expected})`);
    } else {
        console.log(`FAIL: 9% discount for quantity ${tc2Qty} at price ${tc2Price}`);
        console.log(`  Expected: ${tc2Actual}, Got: ${tc2Expected}`);
    }

    // Test Case 3: Quantity = 49 (Still 9% discount) - Boundary condition just below 50
    const tc3Price = 15;
    const tc3Qty = 49;
    const tc3Expected = calculateDiscount(tc3Price, tc3Qty);
    const tc3Actual = tc3Price * tc3Qty * 0.9; // Should be 643.5
    if (tc3Expected === tc3Actual) {
        console.log(`PASS: 9% discount for quantity ${tc3Qty} at price ${tc3Price} (result: ${tc3Expected})`);
    } else {
        console.log(`FAIL: 9% discount for quantity ${tc3Qty} at price ${tc3Price}`);
        console.log(`  Expected: ${tc3Actual}, Got: ${tc3Expected}`);
    }

    // Test Case 4: Quantity = 50 (7.5% discount) - Boundary condition at threshold
    const tc4Price = 25;
    const tc4Qty = 50;
    const tc4Expected = calculateDiscount(tc4Price, tc4Qty);
    const tc4Actual = tc4Price * tc4Qty * 0.75; // Should be 937.5
    if (tc4Expected === tc4Actual) {
        console.log(`PASS: 7.5% discount for quantity ${tc4Qty} at price ${tc4Price} (result: ${tc4Expected})`);
    } else {
        console.log(`FAIL: 7.5% discount for quantity ${tc4Qty} at price ${tc4Price}`);
        console.log(`  Expected: ${tc4Actual}, Got: ${tc4Expected}`);
    }

    // Test Case 5: Quantity > 50 (7.5% discount) - Normal case above threshold
    const tc5Price = 12;
    const tc5Qty = 60;
    const tc5Expected = calculateDiscount(tc5Price, tc5Qty);
    const tc5Actual = tc5Price * tc5Qty * 0.75; // Should be 540
    if (tc5Expected === tc5Actual) {
        console.log(`PASS: 7.5% discount for quantity ${tc5Qty} at price ${tc5Price} (result: ${tc5Expected})`);
    } else {
        console.log(`FAIL: 7.5% discount for quantity ${tc5Qty} at price ${tc5Price}`);
        console.log(`  Expected: ${tc5Actual}, Got: ${tc5Expected}`);
    }

    // Test Case 6: Quantity = 9 (Boundary condition just below 10)
    const tc6Price = 8;
    const tc6Qty = 5;
    const tc6Expected = calculateDiscount(tc6Price, tc6Qty);
    const tc6Actual = tc6Price * tc6Qty; // Should be 40
    if (tc6Expected === tc6Actual) {
        console.log(`PASS: No discount for quantity ${tc6Qty} at price ${tc6Price} (result: ${tc6Expected})`);
    } else {
        console.log(`FAIL: No discount for quantity ${tc6Qty} at price ${tc6Price}`);
        console.log(`  Expected: ${tc6Actual}, Got: ${tc6Expected}`);
    }

    // Test Case 7: Quantity = 49 vs 50 comparison logic check
    const tc7Price = 100;
    const tc7Qty = 49;
    const result49 = calculateDiscount(tc7Price, tc7Qty);
    const tc8Price = 100;
    const tc8Qty = 50;
    const result50 = calculateDiscount(tc8Price, tc8Qty);

    // Check if 49 falls into 9% and 50 falls into 7.5%
    const expected49 = tc7Price * tc7Qty * 0.9;
    const expected50 = tc8Price * tc8Qty * 0.75;

    console.log("\n--- Logic Check ---");
    if (result49 === expected49 && result50 === expected50) {
        console.log("PASS: Correct discount rate applied at boundary (49 uses 9%, 50 uses 7.5%)");
    } else {
        console.log(`FAIL: Discount logic mismatch at boundary`);
        console.log(`  Quantity 49 - Expected: ${expected49}, Got: ${result49}`);
        console.log(`  Quantity 50 - Expected: ${expected50}, Got: ${result50}`);
    }

    console.log("\nTests complete.");
}

main();