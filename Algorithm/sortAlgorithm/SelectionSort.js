const selectionSort = (arr)=>{
    for(let i = 0;i<arr.length;i++){
        let min = i;
        for(let j = i+1;j<arr.length;j++){
            if(arr[min]>arr[j]){
                min = j;
            }
        }

        if(min != i){
            let temp = arr[i];
            arr[i] = arr[min];
            arr[min] = temp;
        }
    }
    return arr;
}

// arr = [5,3,7,4,6];

// console.log(selectionSort(arr));