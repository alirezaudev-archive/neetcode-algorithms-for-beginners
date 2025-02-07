from typing import *


# https://leetcode.com/problems/sort-an-array/
class InsertionSortingSolution:

    def sortArray(self, nums: List[int]) -> List[int]:
        """
        Insertion Sorting Algorithm
        """
        for i in range(1, len(nums)):
            j = i - 1
            while j >= 0 and nums[j + 1] < nums[j]:
                nums[j], nums[j + 1] = nums[j + 1], nums[j]
                j -= 1

        return nums


class MergeSortingSolution:
    def sortArray(self, nums: List[int]) -> List[int]:
        """
        Merge Sorting Algorithm
        """
        return self.__mergeSort(nums, 0, len(nums))

    def __mergeSort(self, nums: List[int], start: int, end: int) -> List[int]:
        if end - start + 1 <= 1:
            return nums[start:end + 1]

        middle = (start + end) // 2

        leftSorted = self.__mergeSort(nums, start, middle)
        rightSorted = self.__mergeSort(nums, middle + 1, end)

        return self.__merge(leftSorted, rightSorted)

    def __merge(self, left: List[int], right: List[int]) -> List[int]:
        merged = []
        i_left = i_right = 0

        while i_left < len(left) and i_right < len(right):
            if left[i_left] <= right[i_right]:
                merged.append(left[i_left])
                i_left += 1
            else:
                merged.append(right[i_right])
                i_right += 1

        merged.extend(left[i_left:])
        merged.extend(right[i_right:])

        return merged


class QuickSortingSolution:
    def sortArray(self, nums: List[int]) -> List[int]:
        """
        Quick Sorting Algorithm
        """
        return self.__quickSort(nums, 0, len(nums) - 1)

    def __quickSort(self, nums: List[int], start: int, end: int) -> List[int]:
        if end - start + 1 <= 1:
            return nums

        pivot = nums[end]
        pointer = start

        for i in range(start, end + 1):
            if nums[i] < pivot:
                nums[pointer], nums[i] = nums[i], nums[pointer]
                pointer += 1

        nums[end] = nums[pointer]
        nums[pointer] = pivot

        self.__quickSort(nums, start, pointer - 1)
        self.__quickSort(nums, pointer + 1, end)

        return nums


insertionSorted = InsertionSortingSolution().sortArray([2, 3, 4, 1, 6])
mergeSorted = MergeSortingSolution().sortArray([2, 3, 4, 1, 6])
quickSorted = QuickSortingSolution().sortArray([2, 3, 4, 1, 6])

print(insertionSorted)
print(mergeSorted)
print(quickSorted)
