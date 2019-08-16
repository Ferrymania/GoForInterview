// swap item in first to last 
const swap = (arr,firstIndex,lastIndex)=>{
    let temp = arr[firstIndex];
    arr[firstIndex] = arr[lastIndex];
    arr[lastIndex] = temp;
}
// Create the Max Heap
const buildMaxHeap = (arr)=>{
    let i = Math.floor(arr.length/2 - 1);
    while(i>=0){
        heapify(arr,i,arr.length);
        i -=1;
    }
}
// moving one element at a time down to its correct location in the heap
const heapify = (heap,i,max)=>{
    let index,leftChild,rightChild;
    while(i < max){
        index = i;
        leftChild = 2*i + 1;
        rightChild = leftChild +1;
        if(leftChild < max && heap[leftChild] >heap[index]){
            index = leftChild;
        }
        if(rightChild < max && heap[rightChild]>heap[index]){
            index = rightChild;
        }

        if(index === i){
            return;
        }
        swap(heap,i,index);
        i = index;
    }
}

const heapSort=(arr)=>{
    //create max heap
    buildMaxHeap(arr);

    //Find last element
    lastElement = arr.length -1;

    //Continue heap sort until we have only one element left in the array.
    while(lastElement > 0){
        swap(arr,0,lastElement);
        heapify(arr,0,lastElement);
        lastElement -=1;
    }



}

arr = [5,3,7,4,6];
heapSort(arr);
console.log(arr);