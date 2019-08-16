const bubbleSort = (arr)=>{
    for(let i = 0;i<arr.length;i++){
        for(let j = 0;j<arr.length;j++){
            if(arr[j]>arr[j+1]){
                let temp = arr[j];
                arr[j] = arr[j+1];
                arr[j+1] = temp;
            }
        }
    }

    return arr;
}

// arr = [5,3,7,4,6];

// bubbleSort(arr);
// console.log(arr);
// Because the code will run until i == length,so it may run on an already sorted
//array more than once ,we can set a  boolean swap to decide whether  
const bubbleSort2 = (arr)=>{
    let swap ;
    do{
        swap = false;
        for(let i =0 ;i<arr.length;i++){
            if(arr[i]>arr[i+1]){
                let temp = arr[i];
                arr[i] = arr[i+1];
                arr[i+1] = temp;
                swap = true;
            }
        }
    }while(swap);
    return arr;
}

// arr2 = [5,3,7,4,6];
// console.log(bubbleSort2(arr2)); 