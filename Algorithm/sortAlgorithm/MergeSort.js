// compare the

const merge = (leftArr,rightArr)=>{
    let result = [];
    let indexLeft = 0;
    let indexRight = 0;

    while(indexLeft < leftArr.length && indexRight< rightArr.length){
        if(leftArr[indexLeft]<rightArr[indexRight]){
            result.push(leftArr[indexLeft]);
            indexLeft++;
        }else{
            result.push(rightArr[indexRight]);
            indexRight++;
        }
    }

    return result.concat(leftArr.slice(indexLeft)).concat(rightArr.slice(indexRight));
}

const mergeSort = (arr)=>{
    // return arr with one item
    if(arr.length === 1){
        return arr;
    }

    let mid = Math.floor(arr.length/2); //get the middle
    let leftArr = arr.slice(0,mid); //get left array 
    let rightArr = arr.slice(mid);//get Right array
    return merge(mergeSort(leftArr),mergeSort(rightArr));
}

// arr = [5,3,7,4,6];
// console.log(mergeSort(arr));