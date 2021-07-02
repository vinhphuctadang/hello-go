

int sum_array(int* array, int n) {
    int sum = 0;
    for(int i = 0; i<n; ++i) { 
        sum += array[i];
    }
    return sum;
}