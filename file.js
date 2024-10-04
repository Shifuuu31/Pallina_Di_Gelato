function split(str, sep) {
    const sepSize = sep.length;
    if (sepSize == 0) {
        return [str];
    }
    let splitted = [];
    let start = 0; // Initialize start to 0
    for (let i = 0; i <= str.length; i++) {
        if (str.slice(i, i + sepSize) === sep || i == str.length) {
            splitted.push(str.slice(start, i));
            start = i + sepSize;
        }
    }
    if (splitted[splitted.length - 1] === '') {
        splitted.splice(splitted.length - 1, 1); // Remove empty element at the end
    }
    return splitted;
}

// Example usage:
let str = 'rrrr';
console.log(split(str, 'rr'));  // ['','']
console.log(str.split('rr'));   // ['','']

let str2 = 'jh bj kj kn';
console.log(split(str2, ' '));  // ['jh', 'bj', 'kj', 'kn']
console.log(str2.split(' '));   // ['jh', 'bj', 'kj', 'kn']
