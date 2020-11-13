import random


def generate_random_list(size_of_list):

    rl = []
    for _ in range(1, size_of_list):
        n = random.randint(1, 100)
        rl.append(n)
    return rl


def main():
    unsorted_list = generate_random_list(40)
    print("unsorted list ", unsorted_list)
    sorted_array = slicer(unsorted_list)
    print(sorted_array)


def slicer(unsorted_list):

    len_of_unsorted_list = len(unsorted_list)
    if len_of_unsorted_list <= 1:
        return unsorted_list

    middle = int(len_of_unsorted_list / 2)

    left_slice = unsorted_list[:middle]
    right_slice = unsorted_list[middle:]

    # print(left_slice)
    # print(right_slice)

    return sort_merge(slicer(left_slice), slicer(right_slice))


def sort_merge(left_slice, right_slice):

    len_of_left_slice = len(left_slice)
    len_of_right_slice = len(right_slice)
    left_nobe, right_nobe = 0, 0
    size_of_the_full_array = len_of_left_slice + len_of_right_slice
    sorted_array = [0] * size_of_the_full_array
    for k in range(0, size_of_the_full_array):
        if left_nobe > (len_of_left_slice - 1) and right_nobe <= (len_of_right_slice - 1):

            sorted_array[k] = right_slice[right_nobe]
            right_nobe += 1

        elif right_nobe > (len_of_right_slice - 1) and left_nobe <= (len_of_left_slice - 1):

            sorted_array[k] = left_slice[left_nobe]
            left_nobe += 1

        elif left_slice[left_nobe] < right_slice[right_nobe]:
            sorted_array[k] = left_slice[left_nobe]
            left_nobe += 1
        else:
            sorted_array[k] = right_slice[right_nobe]
            right_nobe += 1

    return sorted_array


if __name__ == "__main__":
    print("welcome to python merge sort")
    main()