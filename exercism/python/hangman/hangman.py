# Game status categories
# Change the values as you see fit
STATUS_WIN = "win"
STATUS_LOSE = "lose"
STATUS_ONGOING = "ongoing"


class Hangman:
    def __init__(self, word):
        self.remaining_guesses = 9
        self.status = STATUS_ONGOING
        self.__word = word
        self.__guessed_word = ["_"] * len(word)

    def guess(self, char):
        if self.remaining_guesses >= 0 and "_" in self.__guessed_word:
            if char in self.__word and char not in self.__guessed_word:
                _indexes = self._get_all_the_index_occur(char)
                for indx in _indexes:
                    self.__guessed_word[indx] = char
            else:
                self.remaining_guesses -= 1
        else:
            raise ValueError("The Game Is Over Sir !!")

    def get_masked_word(self):

        return "".join(self.__guessed_word)

    def get_status(self):

        print("".join(self.__guessed_word), self.__word)
        if self.remaining_guesses >= 0 and "".join(
                self.__guessed_word) != self.__word:
            self.status = STATUS_ONGOING
        elif "".join(self.__guessed_word) == self.__word:
            self.status = STATUS_WIN
        else:
            self.status = STATUS_LOSE

        return self.status

    def _get_all_the_index_occur(self, char):

        indexes = []
        for c, word in enumerate(self.__word):
            if char == word:
                indexes.append(c)

        return indexes
