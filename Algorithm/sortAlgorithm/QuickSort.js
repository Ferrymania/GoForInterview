//swap two elements
const swap = (arr,i,j)=>{
    let temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;
};
// take last element as pivot,places the pivot element  at its correct position in
//sorted array,and places all smaller to left of pivot and all bigger elements
// to right of pivot
const partition = (arr,low,high)=>{
    let pivot = arr[high];
    let j = low,i;
    for(i = low;i<high;i++){
        if(arr[i] <pivot){
            swap(arr,i,j);
            j++;
        }
    }
    swap(arr,i,j);
    return j;
}
//low--->starting index
//high--->ending index
const quickSort = (arr,low,high) =>{
    if(low < high){
        let pivot = partition(arr,low,high);
        quickSort(arr,low,pivot-1);
        quickSort(arr,pivot+1,high);
        return arr;
    }
}

// arr = [5,3,7,4,6];
// console.log(quickSort(arr,0,4));
// console.log(quickSort([10,7,8,9,1,5],0,5));