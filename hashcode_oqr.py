import math


def main():
    f = open("demo.txt", "r")

    firstLineArray = f.readline().split(" ")
    numBooks = firstLineArray[0]
    numLibraries = firstLineArray[1]
    scanningDays = int(firstLineArray[2])
    scoreOfBooks = f.readline().split(" ")
    scoreBooksMap = []
    for x in scoreOfBooks:
        scoreBooksMap.append(
            {
                "id": int(scoreOfBooks.index(x)),
                "score": int(x),
            }
        )
    libraryInfo = []
    for i in range(0, int(numLibraries)):
        library = f.readline().strip('\n').split(" ")
        libraryMore = f.readline().strip('\n').split(" ")
        libraryMoreMap = []
        for x in libraryMore:
            libraryMoreMap.append(
                {
                    "id": x,
                    "score": scoreOfBooks[int(x)]
                }
            )

        libraryMoreMap.sort(key=lambda l: l["score"], reverse=True)
        if int(library[1]) < scanningDays:
            libraryInfo.append(
                {
                    "id": int(i),
                    "numBooks": int(library[0]),
                    "signup": int(library[1]),
                    "shipBooks": int(library[2]),
                    "books": libraryMoreMap
                }
            )

    libraryInfo.sort(key=lambda k: k["signup"], reverse=True)

    print(len(libraryInfo))
    for u in libraryInfo:
        leftDays = scanningDays - u["signup"]
        # print(u["shipBooks"])
        # print(leftDays)
        canSend = math.ceil(leftDays - 1 / u["shipBooks"])
        print(u["id"], canSend)

        for y in range(canSend):
            if y == (canSend - 1):
                print(u["books"][y]['id'])
            else:
                print(u["books"][y]['id'], end=" ")

main()
