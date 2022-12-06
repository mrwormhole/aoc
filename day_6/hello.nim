import sequtils

proc solve():void =
  var file = open("input.txt")
  defer:
    file.close()

  var l = readLine(file)
  for index, letter in l:
    if index <= len(l) - 4:
      var temp = newSeq[char]()
      for i in index..index+3:
        temp.add(l[i])

      var newTemp = deduplicate[char](temp)
      if len(newTemp) == len(temp):
        echo "found index: ", index+4
        break

  for index, letter in l:
    if index <= len(l) - 4 and index >= 13:
      var forward = newSeq[char]()
      for i in countup(index, index+3):
        forward.add(l[i])
      var backward = newSeq[char]()
      for i in countdown(index, index-13):
          backward.add(l[i])

      var dedupForward = deduplicate[char](forward)
      var dedupBackward = deduplicate[char](backward)
      if len(dedupForward) == len(forward) and len(dedupBackward) == len(backward) and len(dedupBackward) == 14:
        echo "found index: ", index+1
        break

solve()