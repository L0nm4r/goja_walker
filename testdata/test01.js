function add(a, b) {
    if (a > 0) {
        return a + b;
    } else {
        return b - a;
    }
}
const result = add(5, 3);
while (result > 0) {
    result--;
}