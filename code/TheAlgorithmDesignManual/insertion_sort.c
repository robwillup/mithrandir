#include <stdio.h>

void insertion_sort(int arr[], int n)
{
	int i, key, j;

	for (i = 1; i < n; i++) {
		key = arr[i];
		j = i - 1;

		while(j >= 0 && arr[j] > key) {
			arr[j + 1] = arr[j];
			j = j - 1;
		}
		arr[j + 1] = key;
	}
}

int main(void) {
	int hand[] = {5, 10, 2, 9, 3};
	int handSize = sizeof(hand) / sizeof(hand[0]);

	insertion_sort(hand, handSize);

	printf("Sorted list: ");
	for (int i = 0; i < handSize; i++) {
		printf("%d ", hand[i]);
	}
	printf("\n");

	return 0;
}
