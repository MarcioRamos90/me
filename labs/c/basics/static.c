#include<stdio.h>

int runner_non_static(){
    int count = 0;
    count++;
    return count;
 }

int runner_static()
{
    static int count = 0;
    count++;
    return count;
}

int main()
{
    printf("static %d\n", runner_static());
    printf("static %d\n", runner_static());
	
    printf("non static %d\n", runner_non_static());
    printf("non static %d\n", runner_non_static());
    return 0;
}
