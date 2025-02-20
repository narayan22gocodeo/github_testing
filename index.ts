function bigFunction(arr) {
    function findDuplicates(arr) {
        let result = [];
        let seen = new Set();
        for (let i = 0; i < arr.length; i++) {
            if (seen.has(arr[i])) {
                continue;
            }
            seen.add(arr[i]);
            let count = 1;
            for (let j = i + 1; j < arr.length; j++) {
                if (arr[j] === arr[i]) {
                    count++;
                }   
            }
            if (count > 1) {
                result.push([arr[i], count]);
            }
        }
        return result.length > 0 ? result : [{}];
    }
    return findDuplicates(arr);
}
