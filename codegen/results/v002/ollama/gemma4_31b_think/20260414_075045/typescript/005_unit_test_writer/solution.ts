/**
 * Calculates a bulk discount based on the quantity of items.
 * 
 * @param price - The unit price of the item.
 * @param quantity - The number of items being purchased.
 * @returns The total price after applying the applicable discount.
 */
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

/**
 * Interface to define a test case structure
 */
interface TestCase {
    price: number;
    quantity: number;
    expected: number;
    description: string;
}

/**
 * Main function to run the test suite
 */
function runTests() {
    const testCases: TestCase[] = [
        { 
            price: 100, 
            quantity: 9, 
            expected: 900, 
            description: "No discount for quantities under 10 (Boundary: 9)" 
        },
        { 
            price: 100, 
            quantity: 10, 
            expected: 900, 
            description: "10% discount for quantities starting at 10 (Boundary: 10)" 
        },
        { 
            price: 100, 
            quantity: 49, 
            expected: 4410, 
            description: "10% discount for quantities up to 49 (Boundary: 49)" 
        },
        { 
            price: 100, 
            quantity: 50, 
            expected: 3750, 
            description: "25% discount for quantities 50 and above (Boundary: 50)" 
        },
        { 
            price: 10, 
            quantity: 100, 
            expected: 750, 
            description: "High volume discount application" 
        }
    ];

    testCases.forEach(({ price, quantity, expected, description }) => {
        const result = calculateDiscount(price, quantity);
        
        // Using a small epsilon to handle floating point precision if necessary,
        // though these specific numbers typically align well in IEEE 754.
        const isPass = Math.abs(result - expected) < 0.0001;

        if (isPass) {
            console.log(`PASS: ${description}`);
        } else {
            console.log(`FAIL: ${description} (expected: ${expected}, got: ${result})`);
        }
    });
}

// Execute the test suite
runTests();