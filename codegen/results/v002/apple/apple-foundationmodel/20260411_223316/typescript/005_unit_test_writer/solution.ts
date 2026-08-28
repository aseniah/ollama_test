function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

const args = process.argv.slice(2);

if (args.length !== 2) {
    console.log("Usage: node discount_calculator.ts <price> <quantity>");
    return;
}

const price = parseFloat(args[0]);
const quantity = parseInt(args[1]);

try {
    if (isNaN(price) || isNaN(quantity)) {
        console.log("FAIL: Invalid input values (expected: number, expected: number)");
        return;
    }

    const result = calculateDiscount(price, quantity);
    const expected = price * quantity;

    if (quantity === 9) {
        if (result === expected) {
            console.log("PASS: Quantity 9");
        } else {
            console.log("FAIL: Quantity 9 (expected: ", expected, ", got: ", result, ")");
        }
    } else if (quantity === 10) {
        if (result === expected) {
            console.log("PASS: Quantity 10");
        } else {
            console.log("FAIL: Quantity 10 (expected: ", expected, ", got: ", result, ")");
        }
    } else if (quantity === 49) {
        if (result === expected) {
            console.log("PASS: Quantity 49");
        } else {
            console.log("FAIL: Quantity 49 (expected: ", expected, ", got: ", result, ")");
        }
    } else if (quantity === 50) {
        if (result === expected) {
            console.log("PASS: Quantity 50");
        } else {
            console.log("FAIL: Quantity 50 (expected: ", expected, ", got: ", result, ")");
        }
    } else {
        console.log("FAIL: Quantity out of range (expected: 9, 10, 49, 50)");
    }
} catch (error) {
    console.log("FAIL: Error during calculation (expected: number, expected: number)");
}